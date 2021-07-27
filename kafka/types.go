package kafka

import "github.com/Shopify/sarama"

// ProducerConf 定义同步生产者配置类型
type ProducerConf struct {
	Brokers []string

	// 扩展配置, 需要覆盖sarama默认配置时使用
	Ext *sarama.Config
}

type ConsumerGroupConf struct {
	Brokers []string
	GroupId string
	Ext     *sarama.Config
}

type Config struct {
	Kafka struct {
		Consumer ConsumerGroupConf
		Producer ProducerConf
	}
}
