package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRegisterLogic {
	return MemberRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRegisterLogic) MemberRegister(req types.MemberRegisterReq) (*types.MemberRegisterResp, error) {
	resp, _ := l.svcCtx.Ums.MemberList(l.ctx, &umsclient.MemberListReq{
		Current:  1,
		PageSize: 2,
	})

	for _, data := range resp.List {
		fmt.Println(data)
	}

	return &types.MemberRegisterResp{}, nil
}
