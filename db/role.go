// role
package db

import (
	"context"
)

type Role struct {
}

func (this *Role) convert(rst []interface{}) []*RoleBean {
	l := len(rst)
	result := make([]*RoleBean, l)
	for i := 0; i < l; i++ {
		result[i] = rst[i].(*RoleBean)

	}
	return result
}

func (this *Role) GetRoles(ctx context.Context) ([]*RoleBean, error) {
	rst, err := ssql.Select(ctx, T_Role)
	if err != nil {
		return nil, err
	}
	return this.convert(rst), nil
}
func (this *Role) TxGetRoles(tx *SqlTx) ([]*RoleBean, error) {

	rst, err := tx.Select(T_Role)

	if err != nil {
		return nil, err
	}

	return this.convert(rst), nil
}

func (this *Role) TxInsertRole(tx *SqlTx, items ...interface{}) error {

	return tx.Insert(items...)
}
func (this *Role) InsertRole(ctx context.Context, items ...interface{}) error {
	return ssql.Insert(ctx, items...)
}

func GetAuths(ctx context.Context) ([]*AuthBean, error) {
	aus, e := ssql.SelectWhere(ctx, T_Auth, "1=1 order by ordered")
	if e != nil {
		return nil, e
	}
	al := len(aus)
	result := make([]*AuthBean, al)
	for k, i := range aus {
		result[k] = i.(*AuthBean)
	}
	return result, nil

}

func GetRoles(ctx context.Context) ([]*RoleBean, error) {
	aus, e := ssql.SelectWhere(ctx, T_Role, "1=1")
	if e != nil {
		return nil, e
	}
	al := len(aus)
	result := make([]*RoleBean, al)
	for k, i := range aus {
		result[k] = i.(*RoleBean)
	}
	return result, nil
}
