package setup

import (
	"gitee.com/cristiane/go-common/queue"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"log"
)

// SetUpRedisQueue returns *queue.MachineryQueue instance of redis queue.
func SetUpRedisQueue(redisQueueSetting *setting.QueueRedisSettingS, namedTaskFuncs map[string]interface{}) *queue.MachineryQueue {
	if redisQueueSetting == nil {
		log.Fatal("[err] redisQueueSetting is nil")
	}
	if redisQueueSetting.Broker == "" {
		log.Fatal("[err] Lack of redisQueueSetting.Broker")
	}
	if redisQueueSetting.DefaultQueue == "" {
		log.Fatal("[err] Lack of redisQueueSetting.DefaultQueue")
	}
	if redisQueueSetting.ResultBackend == "" {
		log.Fatal("[err] Lack of redisQueueSetting.ResultBackend")
	}
	if redisQueueSetting.ResultsExpireIn < 0 {
		log.Fatal("[err] redisQueueSetting.ResultsExpireIn must >= 0")
	}

	redisQueue, err := queue.NewRedisQueue(
		redisQueueSetting.Broker,
		redisQueueSetting.DefaultQueue,
		redisQueueSetting.ResultBackend,
		redisQueueSetting.ResultsExpireIn,
		namedTaskFuncs,
	)
	if err != nil {
		log.Fatalf("[err] Err queue.NewRedisQueue:%v", err)
	}

	return redisQueue
}

// SetUpAliAMQPQueue returns *queue.MachineryQueue instance of aliyun AMQP queue.
func SetUpAliAMQPQueue(aliAMQPQueueSetting *setting.QueueAliAMQPSettingS, namedTaskFuncs map[string]interface{}) *queue.MachineryQueue {
	if aliAMQPQueueSetting == nil {
		log.Fatal("[err] aliAMQPQueueSetting is nil")
	}
	if aliAMQPQueueSetting.AccessKey == "" {
		log.Fatal("[err] Lack of aliAMQPQueueSetting.AccessKey")
	}
	if aliAMQPQueueSetting.SecretKey == "" {
		log.Fatal("[err] Lack of aliAMQPQueueSetting.SecretKey")
	}
	if aliAMQPQueueSetting.AliUid < 0 {
		log.Fatal("[err] aliAMQPQueueSetting.AliUid must >= 0")
	}
	if aliAMQPQueueSetting.EndPoint == "" {
		log.Fatal("[err] Lack of aliAMQPQueueSetting.EndPoint")
	}
	if aliAMQPQueueSetting.VHost == "" {
		log.Fatal("[err] Lack of aliAMQPQueueSetting.VHost")
	}
	if aliAMQPQueueSetting.DefaultQueue == "" {
		log.Fatal("[err] Lack of aliAMQPQueueSetting.DefaultQueue")
	}
	if aliAMQPQueueSetting.ResultBackend == "" {
		log.Fatal("[err] Lack of aliAMQPQueueSetting.ResultBackend")
	}
	if aliAMQPQueueSetting.ResultsExpireIn < 0 {
		log.Fatal("[err] aliAMQPQueueSetting.ResultsExpireIn must >= 0")
	}

	var aliAMQPConfig = &queue.AliAMQPConfig{
		AccessKey:        aliAMQPQueueSetting.AccessKey,
		SecretKey:        aliAMQPQueueSetting.SecretKey,
		AliUid:           aliAMQPQueueSetting.AliUid,
		EndPoint:         aliAMQPQueueSetting.EndPoint,
		VHost:            aliAMQPQueueSetting.VHost,
		DefaultQueue:     aliAMQPQueueSetting.DefaultQueue,
		ResultBackend:    aliAMQPQueueSetting.ResultBackend,
		ResultsExpireIn:  aliAMQPQueueSetting.ResultsExpireIn,
		Exchange:         aliAMQPQueueSetting.Exchange,
		ExchangeType:     aliAMQPQueueSetting.ExchangeType,
		BindingKey:       aliAMQPQueueSetting.BindingKey,
		PrefetchCount:    aliAMQPQueueSetting.PrefetchCount,
		QueueBindingArgs: nil,
		NamedTaskFuncs:   namedTaskFuncs,
	}

	var aliAMQPQueue, err = queue.NewAliAMQPMqQueue(aliAMQPConfig)
	if err != nil {
		log.Fatalf("[err] Err queue.NewAliAMQPMqQueue:%v", err)
	}

	return aliAMQPQueue
}

// SetUpAMQPQueue returns *queue.MachineryQueue instance of  AMQP queue.
func SetUpAMQPQueue(amqpQueueSetting *setting.QueueAMQPSettingS, namedTaskFuncs map[string]interface{}) *queue.MachineryQueue {
	if amqpQueueSetting == nil {
		log.Fatal("[err] amqpQueueSetting is nil")
	}
	if amqpQueueSetting.Broker == "" {
		log.Fatal("[err] Lack of amqpQueueSetting.Broker")
	}
	if amqpQueueSetting.DefaultQueue == "" {
		log.Fatal("[err] Lack of amqpQueueSetting.DefaultQueue")
	}
	if amqpQueueSetting.ResultBackend == "" {
		log.Fatal("[err] Lack of amqpQueueSetting.ResultBackend")
	}
	if amqpQueueSetting.ResultsExpireIn < 0 {
		log.Fatal("[err] amqpQueueSetting.ResultsExpireIn must >= 0")
	}

	var amqpQueue, err = queue.NewRabbitMqQueue(
		amqpQueueSetting.Broker,
		amqpQueueSetting.DefaultQueue,
		amqpQueueSetting.ResultBackend,
		amqpQueueSetting.ResultsExpireIn,
		amqpQueueSetting.Exchange,
		amqpQueueSetting.ExchangeType,
		amqpQueueSetting.BindingKey,
		amqpQueueSetting.PrefetchCount,
		nil,
		namedTaskFuncs)
	if err != nil {
		log.Fatalf("[err] Err queue.NewRabbitMqQueue:%v", err)
	}

	return amqpQueue
}
