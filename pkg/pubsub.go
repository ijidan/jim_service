package pkg

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type PubSub struct {
	brokers []string
}

func (w *PubSub) PublishMessage(topic string, content []byte) error {
	publisher, err := kafka.NewPublisher(kafka.PublisherConfig{
		Brokers:   w.brokers,
		Marshaler: kafka.DefaultMarshaler{},
	}, watermill.NewStdLogger(false, false))
	if err != nil {
		return err
	}
	m := message.NewMessage(watermill.NewUUID(), content)
	return publisher.Publish(topic, m)
}

func (w *PubSub) Subscribe(topic string) (<-chan *message.Message, error) {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	subscriber, err := kafka.NewSubscriber(kafka.SubscriberConfig{
		Brokers:                w.brokers,
		Unmarshaler:            kafka.DefaultMarshaler{},
		OverwriteSaramaConfig:  saramaSubscriberConfig,
		ConsumerGroup:          "default",
		NackResendSleep:        0,
		ReconnectRetrySleep:    0,
		InitializeTopicDetails: nil,
	}, watermill.NewStdLogger(false, false))
	if err != nil {
		return nil, err
	}
	return subscriber.Subscribe(context.Background(), topic)
}

func NewPubSub(brokers []string) *PubSub {
	pubSub:=&PubSub{brokers: brokers}
	return pubSub
}