package kafka

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var consumerConfigForTest = ConsumerGroupConf{
	Brokers: []string{"localhost:9092"},
	GroupId: "group2",
	Ext:     nil,
}

func TestNewConsumerGroup(t *testing.T) {
	consumer, err := NewConsumerGroup(consumerConfigForTest)
	require.NotNil(t, consumer, "创建consumerGroup失败：%+v", err)


	// config.Consumer.Return.Errors = true
	// Track errors
	/*go func() {
		for err := range consumer.Errors() {
			fmt.Println("ERROR", err)
		}
	}()*/

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{"iris"}
		handler := ExampleHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := consumer.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}


}

func TestConsumerGroup_Close(t *testing.T) {
	consumer, err := NewConsumerGroup(consumerConfigForTest)
	require.NotNil(t, consumer, "创建consumerGroup失败：%+v", err)

	err = consumer.Close()
	require.Nil(t, err, "关闭链接失败")
}
