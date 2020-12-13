package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnReasonDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnReasonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnReasonDeleteLogic {
	return &OrderReturnReasonDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnReasonDeleteLogic) OrderReturnReasonDelete(in *oms.OrderReturnReasonDeleteReq) (*oms.OrderReturnReasonDeleteResp, error) {
	err := l.svcCtx.OmsOrderReturnReasonModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.OrderReturnReasonDeleteResp{}, nil
}
