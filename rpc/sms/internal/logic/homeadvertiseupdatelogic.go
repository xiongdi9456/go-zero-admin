package logic

import (
	"context"
	"go-zero-admin/rpc/model/smsmodel"
	"time"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeAdvertiseUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseUpdateLogic {
	return &HomeAdvertiseUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeAdvertiseUpdateLogic) HomeAdvertiseUpdate(in *sms.HomeAdvertiseUpdateReq) (*sms.HomeAdvertiseUpdateResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	err := l.svcCtx.SmsHomeAdvertiseModel.Update(smsmodel.SmsHomeAdvertise{
		Id:         in.Id,
		Name:       in.Name,
		Type:       in.Type,
		Pic:        in.Pic,
		StartTime:  StartTime,
		EndTime:    EndTime,
		Status:     in.Status,
		ClickCount: in.ClickCount,
		OrderCount: in.OrderCount,
		Url:        in.Url,
		Note:       in.Note,
		Sort:       in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeAdvertiseUpdateResp{}, nil
}
