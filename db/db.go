package db

import (
	"context"

	"log"

	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"strings"
	"webbase/simplesql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var ssql simplesql.Sql

//var ssqlm simplesql.Sql

func Init(db1 string, db2 string) {
	var err error
	//ssql, err = simplesql.New("postgres", "postgres://cxx:123456@localhost/testdb?sslmode=disable")
	ssql, err = simplesql.New("mysql", db1)
	//ssqlm, err = simplesql.New("mysql", db2)

	if err != nil {
		log.Println(err)
		return
	}

	ssql.RegistTable(&UserBean{}, T_User)
	ssql.RegistTable(&AuthBean{}, T_Auth)
	ssql.RegistTable(&RoleBean{}, T_Role)
	//ssqlm.RegistTable(&JameUserBean{}, T_JUser)
}

func init() {

}

const T_JUser = "users"
const T_User = "wuser"
const T_Role = "wrole"
const T_Auth = "wauth"

type JameUserBean struct {
	Username      string `key:"def"`
	PwdHash       string
	PwdAlgorithm  string
	UseForwarding int
}

type UserBean struct {
	Username  string `key:"def"`
	Pwd       string
	Nick_name string
	Real_name string
	Role_id   int
}
type RoleBean struct {
	Role_id   int `key:"auto"`
	Role_name sql.NullString
	Auths     sql.NullString
	Remark    sql.NullString
}
type AuthBean struct {
	Auth_id   int `key:"auto"`
	Auth_name sql.NullString
	Auth_val  sql.NullString
	Remark    sql.NullString
	Pid       int
}

type SqlTx struct {
	simplesql.Tx
}

func Tx(ctx context.Context) (*SqlTx, error) {
	tx, e := ssql.Tx(ctx)
	if e != nil {
		return nil, e
	}
	r := &SqlTx{tx}
	return r, nil
}
func TxEnd(sqltx *SqlTx) {
	sqltx.End()
}

func GetRole(ctx context.Context, rid int) *RoleBean {
	r, e := ssql.SelectOne(ctx, T_Role, rid)
	if e != nil {
		return nil
	}
	if r == nil {
		return nil
	}
	return r.(*RoleBean)
}

func GetUser(ctx context.Context, username string, pwd string) *UserBean {

	r, e := ssql.SelectOne(ctx, T_User, username)
	if e != nil {
		return nil
	}
	if r == nil {
		return nil
	}
	usr := r.(*UserBean)

	if usr.Pwd == pwd {
		return usr
	}
	return nil
}

//func GetAllJUsers(ctx context.Context) []*JameUserBean {
//	all, e := ssqlm.SelectWhere(ctx, T_JUser, "1=1")
//	if e != nil {
//		return nil
//	}
//	ulen := len(all)

//	res := make([]*JameUserBean, ulen)
//	for k, v := range all {
//		res[k] = v.(*JameUserBean)
//	}
//	return res
//}
func AddUser(ctx context.Context, u *UserBean) error {
	u.Pwd = Password(u.Pwd)

	e := ssql.Insert(ctx, u)
	if e != nil {
		return e
	}
	//	ju := &JameUserBean{}
	//	ju.Username = u.Username
	//	ju.PwdHash = Ptojpwd(u.Pwd)
	//	ju.PwdAlgorithm = "SHA"
	//	e = ssqlm.Insert(ctx, ju)
	return e
}
func Editinfo(ctx context.Context, uname string, realname string) error {
	u, e := ssql.SelectOne(ctx, T_User, strings.ToLower(uname))
	if e != nil {
		return e
	}
	uu := u.(*UserBean)
	uu.Real_name = realname
	return ssql.Update(ctx, uu)
}
func GetAllUsers(ctx context.Context) []*UserBean {
	all, e := ssql.SelectWhere(ctx, T_User, "1=1")
	if e != nil {
		return nil
	}
	ulen := len(all)

	res := make([]*UserBean, ulen)
	for k, v := range all {
		res[k] = v.(*UserBean)
	}
	return res
}

func Password(pwd string) string {
	sr := sha1.Sum([]byte(pwd))
	pwdhas := base64.StdEncoding.EncodeToString(sr[:])
	return pwdhas
}

func Ptojpwd(p string) string {
	return p[0:24]
}
func JPassword(pwd string) string {
	sr := sha1.Sum([]byte(pwd))
	pwdhas := base64.StdEncoding.EncodeToString(sr[:])
	return pwdhas[0:24]
}
func AddJUser(ctx context.Context, uname string, pwd string) error {

	ju := &JameUserBean{
		Username:      strings.ToLower(uname),
		PwdHash:       JPassword(pwd),
		PwdAlgorithm:  "SHA",
		UseForwarding: 0,
	}
	return ssql.Insert(ctx, ju)

}

func ResetPwd(ctx context.Context, uname string, pwd string) error {
	upwd := Password(pwd)

	//	ju := &JameUserBean{
	//		Username:      strings.ToLower(uname),
	//		PwdHash:       Ptojpwd(upwd),
	//		PwdAlgorithm:  "SHA",
	//		UseForwarding: 0,
	//	}
	if e := ssql.Execute(ctx, "update wuser set pwd='"+upwd+"' where username='"+strings.ToLower(uname)+"'", nil); e != nil {
		return e
	}
	return nil
	//return ssqlm.Update(ctx, ju)
}

func Login2(ctx context.Context, uname string, pwd string) *UserBean {
	u, er := ssql.SelectOne(ctx, "wuser", strings.ToLower(uname))
	if er != nil || u == nil {
		log.Println(er)
		return nil
	}

	pwdhas := Password(pwd)
	if u.(*UserBean).Pwd != pwdhas {
		return nil
	}
	uu := u.(*UserBean)
	return uu
}

func Login(ctx context.Context, uname string, pwd string) *UserBean {
	u, er := ssql.SelectOne(ctx, "wuser", strings.ToLower(uname))
	if er != nil || u == nil || u.(*UserBean).Pwd != pwd {
		log.Println(er)
		return nil
	}
	sr := sha1.Sum([]byte(pwd))
	pwdhas := base64.StdEncoding.EncodeToString(sr[:])
	if u.(*UserBean).Pwd != pwdhas {
		return nil
	}
	uu := u.(*UserBean)

	return uu
}
