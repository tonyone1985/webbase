// userlist
package control

import (
	"context"
	"errors"
	"strings"
	"webbase/db"
)

type Userlist struct {
}

func (this *Userlist) Post(c context.Context, req Requester) (interface{}, error) {
	m := req.FormValue("method")
	var e error
	result := &BaseJson{}
	switch m {
	case "editinfo":
		e = db.Editinfo(c, req.FormValue("username"), req.FormValue("realname"))
		result.Msg = "信息修改成功"
		break
	case "resetpwd":
		e = db.ResetPwd(c, req.FormValue("username"), req.FormValue("pwd"))
		result.Msg = "密码重置成功"
		break
	case "adduser":
		u := &db.UserBean{}
		u.Username = strings.ToLower(req.FormValue("username"))
		u.Pwd = req.FormValue("pwd")
		u.Real_name = req.FormValue("realname")
		u.Nick_name = req.FormValue("realname")
		u.Role_id = 101
		e = db.AddUser(c, u)
		result.Msg = "用户添加成功"
		break
	default:
		e = errors.New("未知错误")
		break
	}
	if e == nil {
		result.Suss = true

	} else {
		result.Suss = false
		result.Msg = e.Error()
	}
	return result, nil
}
func (this *Userlist) Get(c context.Context) (map[string]interface{}, error) {
	uss := db.GetAllUsers(c)
	r := make(map[string]interface{})
	r["users"] = uss
	return r, nil
}
