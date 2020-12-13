package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleDeleteLogic {
	return RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req types.DeleteRoleReq) (*types.DeleteRoleResp, error) {
	_, err := l.svcCtx.Sys.RoleDelete(l.ctx, &sysclient.RoleDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.DeleteRoleResp{}, nil
}
