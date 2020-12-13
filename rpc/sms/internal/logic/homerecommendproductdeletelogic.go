package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductDeleteLogic {
	return &HomeRecommendProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendProductDeleteLogic) HomeRecommendProductDelete(in *sms.HomeRecommendProductDeleteReq) (*sms.HomeRecommendProductDeleteResp, error) {
	err := l.svcCtx.SmsHomeRecommendProductModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.HomeRecommendProductDeleteResp{}, nil
}
