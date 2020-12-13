package logic

import (
	"context"
	"go-zero-admin/rpc/model/smsmodel"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductUpdateLogic {
	return &HomeRecommendProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendProductUpdateLogic) HomeRecommendProductUpdate(in *sms.HomeRecommendProductUpdateReq) (*sms.HomeRecommendProductUpdateResp, error) {
	err := l.svcCtx.SmsHomeRecommendProductModel.Update(smsmodel.SmsHomeRecommendProduct{
		Id:              in.Id,
		ProductId:       in.ProductId,
		ProductName:     in.ProductName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeRecommendProductUpdateResp{}, nil
}
