// control
package control

import (
	"context"
	"webbase/db"
	"webbase/models"
)

const S_USER_KEY = "user"

type Requester interface {
	FormValue(string) string
}

type HttpExecutor interface {
	Post(c context.Context, req Requester) (interface{}, error)
	Get(c context.Context) (map[string]interface{}, error)
}

type BaseJson struct {
	Suss bool
	Msg  string
}

func Login(ctx context.Context, uname string, pwd string) *models.User {
	usb := db.Login2(ctx, uname, pwd)
	if usb == nil {
		return nil
	}
	return &models.User{
		Pwd:       usb.Pwd,
		Username:  usb.Pwd,
		Nick_name: usb.Nick_name,
		Real_name: usb.Real_name,
		Role_id:   usb.Role_id,
	}

}
func AddUser(ctx context.Context, u *models.User) error {
	ub := &db.UserBean{
		Pwd:       u.Pwd,
		Username:  u.Pwd,
		Nick_name: u.Nick_name,
		Real_name: u.Real_name,
		Role_id:   u.Role_id,
	}

	return db.AddUser(ctx, ub)

}

func ResetPwd(ctx context.Context, uname string, pwd string) error {
	return db.ResetPwd(ctx, uname, pwd)
}
