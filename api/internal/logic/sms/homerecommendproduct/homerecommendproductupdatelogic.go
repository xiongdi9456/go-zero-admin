package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductUpdateLogic {
	return HomeRecommendProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductUpdateLogic) HomeRecommendProductUpdate(req types.UpdateHomeRecommendProductReq) (*types.UpdateHomeRecommendProductResp, error) {
	_, err := l.svcCtx.Sms.HomeRecommendProductUpdate(l.ctx, &smsclient.HomeRecommendProductUpdateReq{
		Id:              req.Id,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateHomeRecommendProductResp{}, nil
}
