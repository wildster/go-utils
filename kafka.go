package goutils

import (
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
	"github.com/twmb/franz-go/plugin/kzap"
)

var (
	seeds      = []string{GetStringEnv("KAFKA_BROKER", DefaultString("localhost:9092"))}
	Mq         *kgo.Client
	ErrorTopic = "kafka_application_error"
)

type KafkaOption struct {
	consumerGroup string
	topics        []string
}

func CreateKafkaClient(opt *KafkaOption) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup(opt.consumerGroup),
		kgo.HeartbeatInterval(time.Second*1),
		kgo.ConsumeTopics(opt.topics...),
		kgo.SASL(plain.Auth{
			User: GetStringEnv("KAFKA_USER_NAME", nil),
			Pass: GetStringEnv("KAFKA_PASSWORD", nil),
		}.AsMechanism()),
		kgo.WithLogger(kzap.New(Logger)),
		kgo.BlockRebalanceOnPoll(),
		kgo.AllowAutoTopicCreation(),
	)

	if err != nil {
		panic(err)
	}

	Mq = cl
}
