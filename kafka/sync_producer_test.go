package kafka

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/require"
)

var producerConfigForTest = ProducerConf{
	Brokers: []string{"localhost:9092"},
}

var val = "testing 123"

var msg = &sarama.ProducerMessage{
	Topic: "iris",
	Value: sarama.StringEncoder(val),
}

var msgs = []*sarama.ProducerMessage{
	&sarama.ProducerMessage{
		Topic: "iris",
		Value: sarama.StringEncoder(val),
	},
	&sarama.ProducerMessage{
		Topic: "iris",
		Value: sarama.StringEncoder(val),
	},
}

func TestSyncProducer(t *testing.T) {
	syncProducer, err := NewSyncProducer(producerConfigForTest)
	require.NotNil(t, syncProducer, "创建syncProducer失败：%+v", err)

	partition, offset, err := syncProducer.SendMessage(msg)
	require.Nil(t, err, "发送消息失败")
	t.Logf("partition: %d, offset: %d", partition, offset)

	err = syncProducer.SendMessages(msgs)
	require.Nil(t, err, "发送多条消息失败")

	err = syncProducer.Close()
	require.Nil(t, err, "关闭链接失败")
}
