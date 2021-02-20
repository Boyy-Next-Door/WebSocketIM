package mq

const (
	MaxMessageCapacity = 10000
	Topic              = "message"
)

var (
	MyClient *Client
)

func Init() {
	MyClient = NewClient()
	MyClient.SetConditions(MaxMessageCapacity)
}
