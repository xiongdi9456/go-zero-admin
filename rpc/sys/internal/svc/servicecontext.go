package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-admin/rpc/model/sysmodel"
	"go-zero-admin/rpc/sys/internal/config"
)

type ServiceContext struct {
	c config.Config

	UserModel     *sysmodel.SysUserModel
	RoleModel     *sysmodel.SysRoleModel
	MenuModel     *sysmodel.SysMenuModel
	DictModel     *sysmodel.SysDictModel
	DeptModel     *sysmodel.SysDeptModel
	LoginLogModel *sysmodel.SysLoginLogModel
	SysLogModel   *sysmodel.SysLogModel
	ConfigModel   *sysmodel.SysConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		c:             c,
		UserModel:     sysmodel.NewSysUserModel(sqlConn),
		RoleModel:     sysmodel.NewSysRoleModel(sqlConn),
		MenuModel:     sysmodel.NewSysMenuModel(sqlConn),
		DictModel:     sysmodel.NewSysDictModel(sqlConn),
		DeptModel:     sysmodel.NewSysDeptModel(sqlConn),
		LoginLogModel: sysmodel.NewSysLoginLogModel(sqlConn),
		SysLogModel:   sysmodel.NewSysLogModel(sqlConn),
		ConfigModel:   sysmodel.NewSysConfigModel(sqlConn),
	}
}
