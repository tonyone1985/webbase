// auths
package control

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"webbase/db"
)

type AuthsControl struct {
	authmap map[int][]*db.AuthBean
}

func (this *AuthsControl) GetRoleAuths(roleid int) ([]*db.AuthBean, error) {
	ats, ok := this.authmap[roleid]
	if !ok {
		return nil, errors.New(fmt.Sprintf("role %d not find auth", roleid))
	}
	return ats, nil
}

func (this *AuthsControl) InitAuthMap(c context.Context) error {
	roles, err := db.GetRoles(c)
	if err != nil {
		return err
	}
	auths, err := db.GetAuths(c)
	if err != nil {
		return err
	}
	this.authmap = make(map[int][]*db.AuthBean)

	for _, r := range roles {
		authids := getauthids(r)
		roleauths := getauths(authids, auths)
		roleauthsorderd := orderauths(roleauths)
		this.authmap[r.Role_id] = roleauthsorderd
	}
	return nil
}

func getauthids(r *db.RoleBean) []int {
	if !r.Auths.Valid {
		return nil
	}
	aus := strings.Split(r.Auths.String, ",")
	l := len(aus)
	rst := make([]int, l)
	for k, v := range aus {
		rst[k], _ = strconv.Atoi(v)
	}
	return rst
}
func getauths(ids []int, objs []*db.AuthBean) []*db.AuthBean {
	var r = make([]*db.AuthBean, 0)
	for _, v := range ids {
		for _, vv := range objs {
			if v == vv.Auth_id {
				r = append(r, vv)
				break
			}
		}
	}
	return r
}

func orderauths(objs []*db.AuthBean) []*db.AuthBean {
	r := make([]*db.AuthBean, 0)
	for _, v := range objs {
		if v.Pid == 0 {
			r = append(r, v)
			for _, vv := range objs {
				if vv.Pid == v.Auth_id {
					r = append(r, vv)
				}
			}

		}
	}
	return r
}
