package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuDeleteLogic {
	return MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDeleteLogic) MenuDelete(req types.DeleteMenuReq) (*types.DeleteMenuResp, error) {
	_, err := l.svcCtx.Sys.MenuDelete(l.ctx, &sysclient.MenuDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.DeleteMenuResp{}, nil
}
