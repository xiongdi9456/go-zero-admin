package logic

import (
	"context"
	"go-zero-admin/rpc/model/omsmodel"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderSettingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderSettingUpdateLogic {
	return &OrderSettingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderSettingUpdateLogic) OrderSettingUpdate(in *oms.OrderSettingUpdateReq) (*oms.OrderSettingUpdateResp, error) {
	err := l.svcCtx.OmsOrderSettingModel.Update(omsmodel.OmsOrderSetting{
		Id:                  in.Id,
		FlashOrderOvertime:  in.FinishOvertime,
		NormalOrderOvertime: in.NormalOrderOvertime,
		ConfirmOvertime:     in.ConfirmOvertime,
		FinishOvertime:      in.FinishOvertime,
		CommentOvertime:     in.CommentOvertime,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderSettingUpdateResp{}, nil
}
