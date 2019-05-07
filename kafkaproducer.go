package main

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaConfigInterface interface {
	Host() string
	Key() string
	Topic() string
	GroupId() string
}

type KafkaProducer struct {
	cfg KafkaConfigInterface
}

func (k *KafkaProducer) Cfg() KafkaConfigInterface {
	return k.cfg
}

func (k *KafkaProducer) SetCfg(cfg KafkaConfigInterface) {
	k.cfg = cfg
}

func NewKafkaProducer(cfg KafkaConfigInterface) *KafkaProducer {
	return &KafkaProducer{
		cfg: cfg,
	}
}

func (k *KafkaProducer) sendMessage(message []byte) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{k.Cfg().Host()},
		Topic:    k.Cfg().Topic(),
		Balancer: &kafka.LeastBytes{},
	})
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(k.Cfg().Key()),
			Value: message,
		},
	)
	if err != nil {
		return err
	}
	return writer.Close()
}
