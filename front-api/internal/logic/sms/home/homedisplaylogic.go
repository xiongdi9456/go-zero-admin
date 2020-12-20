package logic

import (
	"context"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeDisplayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeDisplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeDisplayLogic {
	return HomeDisplayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeDisplayLogic) HomeDisplay(req types.ListHomeDisplayReq) (*types.ListHomeDisplayResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListHomeDisplayResp{}, nil
}
