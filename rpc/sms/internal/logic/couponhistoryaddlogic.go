package logic

import (
	"context"
	"go-zero-admin/rpc/model/smsmodel"
	"time"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryAddLogic {
	return &CouponHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponHistoryAddLogic) CouponHistoryAdd(in *sms.CouponHistoryAddReq) (*sms.CouponHistoryAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	UseTime, _ := time.Parse("2006-01-02 15:04:05", in.UseTime)
	_, err := l.svcCtx.SmsCouponHistoryModel.Insert(smsmodel.SmsCouponHistory{
		CouponId:       in.CouponId,
		MemberId:       in.MemberId,
		CouponCode:     in.CouponCode,
		MemberNickname: in.MemberNickname,
		GetType:        in.GetType,
		CreateTime:     CreateTime,
		UseStatus:      in.UseStatus,
		UseTime:        UseTime,
		OrderId:        in.OrderId,
		OrderSn:        in.OrderSn,
	})
	if err != nil {
		return nil, err
	}

	return &sms.CouponHistoryAddResp{}, nil
}
