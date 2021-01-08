package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-order/pkg/code"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/common/queue"
	"gitee.com/kelvins-io/kelvins"
	"github.com/RichardKnop/machinery/v1/tasks"
)

type PushNoticeService struct {
	server *queue.MachineryQueue
	tag    PushMsgTag
}

func NewPushNoticeService(server *queue.MachineryQueue, tag PushMsgTag) *PushNoticeService {
	if tag.RetryCount <= 0 {
		tag.RetryCount = 3
	}
	if tag.RetryTimeout <= 0 {
		tag.RetryTimeout = 10
	}
	return &PushNoticeService{
		server: server,
		tag:    tag,
	}
}

type PushMsgTag struct {
	DeliveryTag    string
	DeliveryErrTag string
	RetryCount     int
	RetryTimeout   int
}

func (p *PushNoticeService) PushMessage(ctx context.Context, args interface{}) (string, int) {

	taskSign, retCode := p.buildQueueData(ctx, args)
	if retCode != code.Success {
		return "", retCode
	}

	taskId, retCode := p.sendTaskToQueue(ctx, taskSign)
	if retCode != code.Success {
		return "", retCode
	}

	return taskId, code.Success
}

// 构建队列数据
func (p *PushNoticeService) buildQueueData(ctx context.Context, args interface{}) (*tasks.Signature, int) {

	sign := p.buildTaskSignature(args)

	errSign, err := tasks.NewSignature(
		p.tag.DeliveryErrTag, []tasks.Arg{
			{
				Name:  "data",
				Type:  "string",
				Value: json.MarshalToStringNoError(args),
			},
		})

	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "build errSign err: %v, taskSign: %v", err, sign)
		return nil, code.ErrorServer
	}

	errCallback := make([]*tasks.Signature, 0)
	errCallback = append(errCallback, errSign)
	sign.OnError = errCallback

	return sign, code.Success
}

// 构建任务签名
func (p *PushNoticeService) buildTaskSignature(args interface{}) *tasks.Signature {

	taskSignature := &tasks.Signature{
		Name:         p.tag.DeliveryTag,
		RetryCount:   p.tag.RetryCount,
		RetryTimeout: p.tag.RetryTimeout,
		Args: []tasks.Arg{
			{
				Name:  "data",
				Type:  "string",
				Value: json.MarshalToStringNoError(args),
			},
		},
	}

	return taskSignature
}

// 将任务发送到队列
func (p *PushNoticeService) sendTaskToQueue(ctx context.Context, taskSign *tasks.Signature) (string, int) {

	result, err := p.server.TaskServer.SendTaskWithContext(ctx, taskSign)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "pushMessage err:%v, data:%v", err, taskSign)
		return "", code.ErrorServer
	}

	return result.Signature.UUID, code.Success
}

func (p *PushNoticeService) GetTaskState(taskId string) (*tasks.TaskState, error) {

	return p.server.TaskServer.GetBackend().GetState(taskId)
}
