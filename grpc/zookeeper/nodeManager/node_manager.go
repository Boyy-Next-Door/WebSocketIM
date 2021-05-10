/**
node注册中心
*/
package nodeManager

import (
	"errors"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/wonderivan/logger"
	"math/rand"
	"sync"
	"time"
)

const MaxNode = 5
const HeartBeatInterval = 10 // 节点死亡时间：距离上一次心跳超过10秒
const RecoverTime = 20       // 节点复活时间：距离上一次心跳20秒以内
const MaxNodeLoad = 1000     // 每个节点最大负载用户数量
const LoadThreshold = 0.9    // 当节点负载达到百分之90时，开始将新的连接请求分配给剩余空闲节点  当所有节点都达到负载阈值 开始平均分配

// node定义
type Node struct {
	NodeName     string     `json:"nodeName,omitempty"`
	GrpcAddr     string     `json:"grpcAddr,omitempty"`
	HttpAddr     string     `json:"httpAddr,omitempty"`
	LastHeatBeat int64      `json:"-"` // 上一次心跳时间戳
	Status       int        `json:"-"` // 节点状态  1-正常 0-死亡
	UserSet      mapset.Set `json:"-"` // 当前登录用户set
}

type Manager struct {
	// 维护node map
	NodeMap map[string]Node
	// 用来记录node注册顺序的slice
	NodeSlice []string
}

// 全局单例
var ins *Manager
var mu sync.Mutex

func GetIns() *Manager {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			// 初始化注册中心
			ins = &Manager{}
			ins.NodeMap = make(map[string]Node, 0)
			ins.NodeSlice = make([]string, 0)
			// 定时任务 每一秒钟检查一次节点活性
			go ins.MonitorNodes()
		}
	}
	return ins
}

/**
添加一个节点
*/
func (manager *Manager) AddNode(node Node) error {
	val, exist := manager.NodeMap[node.NodeName]
	if exist {
		// 节点复活
		if val.Status == 0 {
			val.Status = 1
			manager.NodeMap[node.NodeName] = val
			return nil
		} else {
			// 同名节点存活  添加失败
			return errors.New("同名节点已经注册，注册失败")
		}
	} else {
		// 节点不存在 添加成功
		node.LastHeatBeat = time.Now().Unix()
		node.Status = 1
		node.UserSet = mapset.NewSet()
		manager.NodeMap[node.NodeName] = node
		manager.NodeSlice = append(manager.NodeSlice, node.NodeName)
		return nil
	}
}

/**
接受节点心跳
*/
func (manager *Manager) HeartBeat(nodeName string) error {
	val, exist := manager.NodeMap[nodeName]
	if exist {
		if val.Status == 1 || val.Status == 0 && time.Now().Unix()-val.LastHeatBeat < RecoverTime {
			val.Status = 1
			val.LastHeatBeat = time.Now().Unix()
			manager.NodeMap[nodeName] = val
			return nil
		} else {
			// 节点已经死亡且超过了复活时间 删除中之
			delete(manager.NodeMap, nodeName)
			for i := 0; i < len(manager.NodeSlice); i++ {
				if manager.NodeSlice[i] == nodeName {
					manager.NodeSlice = append(manager.NodeSlice[:i], manager.NodeSlice[i+1:]...)
					break
				}
			}
			return errors.New("该节点已经死亡，请重新注册")
		}
	} else {
		// 节点不存在
		return errors.New("该节点不存在，请注册")
	}
}

/**
监听注册中心的节点活性
*/
func (manager *Manager) MonitorNodes() {
	//定时任务
	ticker := time.NewTicker(time.Second * 10)
	for range ticker.C {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " --- MonitorNodes new tick.")
		// 检查nodeMap中死亡的节点
		//遍历map
		logger.Info("------ checking nodeManager status ------")
		for key, val := range manager.NodeMap {
			logger.Info(key, val)
			dif := time.Now().Unix() - val.LastHeatBeat
			if dif > HeartBeatInterval {
				if dif < RecoverTime {
					val.Status = 0
				} else {
					// 节点死亡
					delete(manager.NodeMap, key)
					for i := 0; i < len(manager.NodeSlice); i++ {
						if manager.NodeSlice[i] == key {
							manager.NodeSlice = append(manager.NodeSlice[:i], manager.NodeSlice[i+1:]...)
							break
						}
					}
				}
			}
		}
		logger.Info("------ check finished ------")
	}
}

/**
供sdk使用 请求分配一个服务节点
*/
func (manager *Manager) GetNode(userId string) (Node, error) {
	// 分配一个节点
	busyLoadNodes := []string{}
	idleLoadNodes := []string{}
	bestIdleNode := ""
	bestIdleLoad := -1
	for _, val := range manager.NodeSlice {
		// 根据注册的时间顺序进行遍历
		currNodeLoad := manager.NodeMap[val].UserSet.Cardinality()

		// 忙负载 完全满负载的节点不参与分配
		if currNodeLoad >= MaxNodeLoad*LoadThreshold && currNodeLoad < MaxNodeLoad {
			busyLoadNodes = append(busyLoadNodes, val)
		} else {
			// 空闲负载
			idleLoadNodes = append(idleLoadNodes, val)
			if currNodeLoad > bestIdleLoad {
				bestIdleLoad = currNodeLoad
				bestIdleNode = val
			}
		}
	}

	// 此时有两种情况
	if bestIdleNode == "" {
		// 1. 所有节点均满负载, 则bestIdleIdx为"", 此情况下在fullLoadNodes里选取负载最小的一个进行连接 （当最小负载节点有多个时 随机选取）
		minLoadBusyNodeIdx := []int{}
		minLoad := MaxNodeLoad
		for idx, val := range busyLoadNodes {
			currNodeLoad := manager.NodeMap[val].UserSet.Cardinality()
			if currNodeLoad <= minLoad {
				minLoad = currNodeLoad
				minLoadBusyNodeIdx = append(minLoadBusyNodeIdx, idx)
			}
		}

		// 在minLoadBusyNodeIdx中随机选取一个忙负载节点进行分配
		if len(minLoadBusyNodeIdx) == 0 {
			return Node{}, errors.New("无可用节点")
		} else {
			randIdx := rand.Intn(len(minLoadBusyNodeIdx))
			return manager.NodeMap[busyLoadNodes[randIdx]], nil
		}

	} else {
		// 2. 有空闲节点, bestIdleIdx不为"", 此情况下分配给bestIdleIdx这个节点, 尽可能早地使其满载
		return manager.NodeMap[bestIdleNode], nil
	}
}

/**
供node调用 登记用户登入信息
*/
func (manager *Manager) UserCheckIn(nodeName string, userId string) error {
	// 支持挤下线
	prevNode := ""

	for key, val := range manager.NodeMap {
		if val.UserSet.Contains(userId) {
			prevNode = key
			break
		}
	}
	// 用户已登录
	if prevNode != "" {
		manager.NodeMap[prevNode].UserSet.Remove(userId)

		// todo 远程调用该节点的 forceLogout 方法 强制将用户挤下线 并且处理掉先前那个节点的缓存消息
	}

	// 用户加入节点
	newNode, exist := manager.NodeMap[nodeName]

	// 节点不存在
	if !exist {
		return errors.New("目标节点不存在")
	} else {
		newNode.UserSet.Add(userId)
		return nil
	}
}

/**
供node调用 登记用户登出信息
*/
func (manager *Manager) UserCheckOut(nodeName string, userId string) error {
	// 获取目标节点
	newNode, exist := manager.NodeMap[nodeName]

	// 节点不存在
	if !exist {
		return errors.New("目标节点不存在")
	} else {
		newNode.UserSet.Remove(userId)
		return nil
	}
}

/**
供node调用 查找某个用户所在的node
*/
func (manager *Manager) FindUser(userId string) (Node, error) {
	// 目标用户所在的node
	targetNodeName := ""

	for key, val := range manager.NodeMap {
		if val.UserSet.Contains(userId) {
			targetNodeName = key
			break
		}
	}

	// 寻找到了目标用户
	if targetNodeName != "" {
		return manager.NodeMap[targetNodeName], nil
	} else {
		return Node{}, errors.New("目标用户未登录")
	}
}

func (manager *Manager) GetNodeAddress(nodeName string) (string, error) {
	node, exist := manager.NodeMap[nodeName]
	if exist {
		return node.GrpcAddr, nil
	} else {
		return "", errors.New("目标node不存在")
	}
}
