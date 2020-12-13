package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingDeleteLogic {
	return &MemberRuleSettingDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingDeleteLogic) MemberRuleSettingDelete(in *ums.MemberRuleSettingDeleteReq) (*ums.MemberRuleSettingDeleteResp, error) {
	err := l.svcCtx.UmsMemberRuleSettingModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberRuleSettingDeleteResp{}, nil
}
