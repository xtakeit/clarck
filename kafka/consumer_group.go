// Author: Steve Zhang
// Date: 2020/9/14 2:10 下午

package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type ConsumerGroup struct {
	sarama.ConsumerGroup

	messages chan *sarama.ConsumerMessage
}

func NewConsumerGroup(cf ConsumerGroupConf) (gc *ConsumerGroup, err error) {
	if cf.Ext == nil {
		cf.Ext = sarama.NewConfig()
	}

	ogc, err := sarama.NewConsumerGroup(cf.Brokers, cf.GroupId, cf.Ext)
	if err != nil {
		err = fmt.Errorf("sarama.NewConsumerGroup: %w", err)
		return
	}

	gc = &ConsumerGroup{
		ConsumerGroup: ogc,
		messages:      make(chan *sarama.ConsumerMessage, 1),
	}

	return
}

func (gc *ConsumerGroup) Close() (err error) {
	_ = gc.ConsumerGroup.Close()
	return
}

type ExampleHandler struct {
}

func (mh ExampleHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (mh ExampleHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (mh ExampleHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		session.MarkMessage(msg, "")
	}
	return nil
}

