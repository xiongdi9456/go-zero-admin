package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeAdvertiseDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseDeleteLogic {
	return &HomeAdvertiseDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(in *sms.HomeAdvertiseDeleteReq) (*sms.HomeAdvertiseDeleteResp, error) {
	err := l.svcCtx.SmsHomeAdvertiseModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.HomeAdvertiseDeleteResp{}, nil
}
