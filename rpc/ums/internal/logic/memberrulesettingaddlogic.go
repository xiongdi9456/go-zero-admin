package logic

import (
	"context"
	"go-zero-admin/rpc/model/umsmodel"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingAddLogic {
	return &MemberRuleSettingAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingAddLogic) MemberRuleSettingAdd(in *ums.MemberRuleSettingAddReq) (*ums.MemberRuleSettingAddResp, error) {
	_, err := l.svcCtx.UmsMemberRuleSettingModel.Insert(umsmodel.UmsMemberRuleSetting{
		ContinueSignDay:   in.ContinueSignDay,
		ContinueSignPoint: in.ContinueSignPoint,
		ConsumePerPoint:   float64(in.ConsumePerPoint),
		LowOrderAmount:    float64(in.LowOrderAmount),
		MaxPointPerOrder:  in.MaxPointPerOrder,
		Type:              in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberRuleSettingAddResp{}, nil
}
