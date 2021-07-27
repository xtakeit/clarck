// Author: Steve Zhang
// Date: 2020/9/9 3:17 下午

package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// SyncProducer 定义同步生产者类型
type SyncProducer struct {
	sarama.SyncProducer
}

// NewSyncProducer 以指定生产者配置创建同步生产者实例并返回实例地址
// 创建生产者失败时将返回错误
func NewSyncProducer(cf ProducerConf) (pdr *SyncProducer, err error) {
	if cf.Ext == nil {
		cf.Ext = sarama.NewConfig()
	}

	// 同步生产者必须配置开启
	cf.Ext.Producer.Return.Successes = true

	spdr, err := sarama.NewSyncProducer(cf.Brokers, cf.Ext)
	if err != nil {
		err = fmt.Errorf("sarama.NewSyncProducer: %w", err)
		return
	}

	pdr = &SyncProducer{
		SyncProducer: spdr,
	}

	return
}

// Close 实现io.Closer接口, 关闭生产者, 清理打开的系统资源
func (prd *SyncProducer) Close() (err error) {
	err = prd.SyncProducer.Close()
	return
}
