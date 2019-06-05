// mydetails
package control

import (
	"context"
	"webbase/db"
	"webbase/models"
)

type Mydetails struct {
}

func (this *Mydetails) Post(c context.Context, req Requester) (interface{}, error) {
	rest := &BaseJson{}
	m := req.FormValue("method")
	if m == "resetpwd" {

		u := c.Value(S_USER_KEY).(*models.User)
		p := req.FormValue("pwd")
		if e := db.ResetPwd(c, u.Username, p); e == nil {
			rest.Msg = "密码重置成功"
			rest.Suss = true
		} else {
			rest.Msg = "密码重置失败" + e.Error()
			rest.Suss = false
		}

	} else {
		rest.Msg = "参数错误"
		rest.Suss = false
	}
	return rest, nil

}
func (this *Mydetails) Get(c context.Context) (map[string]interface{}, error) {
	return nil, nil
}
