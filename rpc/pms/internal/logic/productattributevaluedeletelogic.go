package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeValueDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueDeleteLogic {
	return &ProductAttributeValueDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueDeleteLogic) ProductAttributeValueDelete(in *pms.ProductAttributeValueDeleteReq) (*pms.ProductAttributeValueDeleteResp, error) {
	err := l.svcCtx.PmsProductAttributeValueModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeValueDeleteResp{}, nil
}
