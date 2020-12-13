package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeCategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryDeleteLogic {
	return &ProductAttributeCategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryDeleteLogic) ProductAttributeCategoryDelete(in *pms.ProductAttributeCategoryDeleteReq) (*pms.ProductAttributeCategoryDeleteResp, error) {
	err := l.svcCtx.PmsProductAttributeCategoryModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeCategoryDeleteResp{}, nil
}
