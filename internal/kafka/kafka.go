package kafka

import (
    "context"
    "time"

    "github.com/Shopify/sarama"
)

type KafkaClient struct {
    producer sarama.SyncProducer
    topic    string
}

func NewKafkaClient(brokers []string, topic string) (*KafkaClient, error) {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll        
    config.Producer.Retry.Max = 3                           
    config.Producer.Return.Successes = true           
    config.Producer.Timeout = 5 * time.Second

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }
    return &KafkaClient{
        producer: producer,
        topic:    topic,
    }, nil
}

func (kc *KafkaClient) SendMessage(value string) (partition int32, offset int64, err error) {
    msg := &sarama.ProducerMessage{
        Topic: kc.topic,
        Value: sarama.StringEncoder(value),
    }
    partition, offset, err = kc.producer.SendMessage(msg)
    return
}

func (kc *KafkaClient) SendMessageWithKey(key string, value string) (partition int32, offset int64, err error) {
    msg := &sarama.ProducerMessage{
        Topic: kc.topic,
        Key:   sarama.StringEncoder(key),
        Value: sarama.StringEncoder(value),
    }
    partition, offset, err = kc.producer.SendMessage(msg)
    return
}

func (kc *KafkaClient) Close() error {
    return kc.producer.Close()
}

type KafkaConsumer struct {
    ctx      context.Context
    consumer sarama.Consumer
    topic    string
}

func NewKafkaConsumer(brokers []string, topic string) (*KafkaConsumer, error) {
    consumer, err := sarama.NewConsumer(brokers, nil)
    if err != nil {
        return nil, err
    }
    ctx := context.Background()
    return &KafkaConsumer{
        ctx:      ctx,
        consumer: consumer,
        topic:    topic,
    }, nil
}

func (kc *KafkaConsumer) ConsumePartition(partition int32, offset int64, handleFunc func(msg *sarama.ConsumerMessage)) error {
    pc, err := kc.consumer.ConsumePartition(kc.topic, partition, offset)
    if err != nil {
        return err
    }
    defer pc.Close()
    for msg := range pc.Messages() {
        handleFunc(msg)
    }
    return nil
}

func (kc *KafkaConsumer) ListPartitions() ([]int32, error) {
    return kc.consumer.Partitions(kc.topic)
}

func (kc *KafkaConsumer) Close() error {
    return kc.consumer.Close()
}