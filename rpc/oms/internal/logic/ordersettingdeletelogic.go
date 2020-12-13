package logic

import (
	"context"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingDeleteLogic {
	return &OrderSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderSettingDeleteLogic) OrderSettingDelete(in *oms.OrderSettingDeleteReq) (*oms.OrderSettingDeleteResp, error) {
	err := l.svcCtx.OmsOrderSettingModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &oms.OrderSettingDeleteResp{}, nil
}
