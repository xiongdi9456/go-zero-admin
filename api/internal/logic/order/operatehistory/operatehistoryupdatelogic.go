package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryUpdateLogic {
	return OperateHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryUpdateLogic) OperateHistoryUpdate(req types.UpdateOperateHistoryReq) (*types.UpdateOperateHistoryResp, error) {
	_, err := l.svcCtx.Oms.OrderOperateHistoryUpdate(l.ctx, &omsclient.OrderOperateHistoryUpdateReq{})

	if err != nil {
		return nil, err
	}

	return &types.UpdateOperateHistoryResp{}, nil
}