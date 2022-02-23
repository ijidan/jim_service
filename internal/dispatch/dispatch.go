package dispatch

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/fatih/color"
	"jim_service/pkg"
)

const (
	TopicSendMessage = "topic.send.message"
)
const (
	GroupSendMessage="group.send.message"
)

func PublishSendMessage(gatewayId string,message []byte) error {
	kafka := pkg.NewKafkaS(pkg.Conf.PubSub.Brokers)
	topic:=fmt.Sprintf("%s.%s",TopicSendMessage,gatewayId)
	return kafka.PublishMessage(topic, string(message))
}

func SubscribeSendMessage(gId string) error {
	f := func(message *sarama.ConsumerMessage) error {
		color.Yellow("kafka message received:%s", string(message.Value))
		return nil
	}
	kafka := pkg.NewKafkaS(pkg.Conf.PubSub.Brokers)
	topic:=fmt.Sprintf("%s.%s",TopicSendMessage,gId)
	return kafka.Subscribe(topic, GroupSendMessage, f)
}