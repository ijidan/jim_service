package dispatch

import (
	"encoding/json"
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

		var data json.RawMessage
		msg:=ClientMessage{
			Data: &data,
		}

		if err3 := json.Unmarshal(message.Value, &msg); err3 != nil {
			color.Red("kafka parse message err:%s",err3.Error())
		}
		switch msg.Cmd {
		case "chat.c2c.txt":
			content:=TextMessage{}
			if err4 := json.Unmarshal(data, &content); err4 != nil {
				color.Red("kafka parse message err:%s",err4.Error())
			}
			toReceiverId:=content.ToReceiverId
			gatewayId,ok1:=ClientIdGatewayIdMap.Load(toReceiverId)
			if !ok1{
				color.Red("kafka compute Gateway Id err")
			}
			err:=PublishSendMessage(gatewayId.(string),message.Value)
			if err!=nil{
				color.Red("kafka trans message err:%s",err.Error())
			}
		}

		return nil
	}
	kafka := pkg.NewKafkaS(pkg.Conf.PubSub.Brokers)
	topic:=fmt.Sprintf("%s.%s",TopicSendMessage,gId)
	return kafka.Subscribe(topic, GroupSendMessage, f)
}