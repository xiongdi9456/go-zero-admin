package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionProductRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationDeleteLogic {
	return &FlashPromotionProductRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionProductRelationDeleteLogic) FlashPromotionProductRelationDelete(in *sms.FlashPromotionProductRelationDeleteReq) (*sms.FlashPromotionProductRelationDeleteResp, error) {
	err := l.svcCtx.SmsFlashPromotionProductRelationModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.FlashPromotionProductRelationDeleteResp{}, nil
}
