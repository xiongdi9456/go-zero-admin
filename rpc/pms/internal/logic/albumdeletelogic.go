package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlbumDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumDeleteLogic {
	return &AlbumDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumDeleteLogic) AlbumDelete(in *pms.AlbumDeleteReq) (*pms.AlbumDeleteResp, error) {
	err := l.svcCtx.PmsAlbumModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.AlbumDeleteResp{}, nil
}
