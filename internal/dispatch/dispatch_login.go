package dispatch

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/fatih/color"
	"jim_service/pkg"
)

func SubscribeCmdLogin() error {
	f := func(msg *sarama.ConsumerMessage) error {
		color.Yellow("message received:%s", string(msg.Value))
		return errors.New("test trigger")
	}
	kafka := pkg.NewKafkaS(pkg.Conf.PubSub.Brokers)
	return kafka.Subscribe(TopicCmdLogin, GroupCmdLogin, f)
}