package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductRelationDeleteLogic {
	return &CouponProductRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductRelationDeleteLogic) CouponProductRelationDelete(in *sms.CouponProductRelationDeleteReq) (*sms.CouponProductRelationDeleteResp, error) {
	err := l.svcCtx.SmsCouponProductRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.CouponProductRelationDeleteResp{}, nil
}
