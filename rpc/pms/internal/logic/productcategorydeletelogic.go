package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryDeleteLogic {
	return &ProductCategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryDeleteLogic) ProductCategoryDelete(in *pms.ProductCategoryDeleteReq) (*pms.ProductCategoryDeleteResp, error) {
	err := l.svcCtx.PmsProductCategoryModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductCategoryDeleteResp{}, nil
}
