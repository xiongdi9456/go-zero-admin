package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderReturnApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyListLogic {
	return &OrderReturnApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnApplyListLogic) OrderReturnApplyList(in *oms.OrderReturnApplyListReq) (*oms.OrderReturnApplyListResp, error) {
	all, _ := l.svcCtx.OmsOrderReturnApplyModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*oms.OrderReturnApplyListData
	for _, item := range *all {

		list = append(list, &oms.OrderReturnApplyListData{
			Id:               item.Id,
			OrderId:          item.OrderId,
			CompanyAddressId: item.CompanyAddressId,
			ProductId:        item.ProductId,
			OrderSn:          item.OrderSn,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
			MemberUsername:   item.MemberUsername,
			ReturnAmount:     int64(item.ReturnAmount),
			ReturnName:       item.ReturnName,
			ReturnPhone:      item.ReturnPhone,
			Status:           item.Status,
			HandleTime:       item.HandleTime.Format("2006-01-02 15:04:05"),
			ProductPic:       item.ProductPic,
			ProductName:      item.ProductName,
			ProductBrand:     item.ProductBrand,
			ProductAttr:      item.ProductAttr,
			ProductCount:     item.ProductCount,
			ProductPrice:     int64(item.ProductPrice),
			ProductRealPrice: int64(item.ProductRealPrice),
			Reason:           item.Reason,
			Description:      item.Description,
			ProofPics:        item.ProofPics,
			HandleNote:       item.HandleNote,
			HandleMan:        item.HandleMan,
			ReceiveMan:       item.ReceiveMan,
			ReceiveTime:      item.ReceiveTime.Format("2006-01-02 15:04:05"),
			ReceiveNote:      item.ReceiveNote,
		})
	}

	fmt.Println(list)
	return &oms.OrderReturnApplyListResp{
		Total: 10,
		List:  list,
	}, nil
}
