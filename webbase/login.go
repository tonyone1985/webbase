package main

import (
	"errors"
	"strings"

	//"github.com/gin-gonic/contrib/sessions"
	"net/http"

	"db"

	"github.com/gin-gonic/gin"
)

func pindex(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "",
		})
	} else if c.Request.Method == "POST" {

		//c.JSON(http.StatusOK, "[]")

	}
}

func pemaildetails(c *gin.Context) {
	if c.Request.Method == "GET" {
		uarr := c.QueryArray("user")
		usr := ""
		if len(uarr) > 0 {
			usr = uarr[0]
		}
		readonly := ""
		username := ""
		btnval := "添加"
		method := "add"
		if usr == "" { //新增

		} else { //修改
			method = "edit"
			readonly = "readonly"
			username = usr
			btnval = "重置密码"
		}

		c.HTML(http.StatusOK, "emaildetails.html", gin.H{
			"readonly": readonly,
			"username": username,
			"btnval":   btnval,
			"method":   method,
		})
	} else if c.Request.Method == "POST" {

		u := c.Request.FormValue("username")
		p := c.Request.FormValue("password")
		m := c.Request.FormValue("method")

		if m == "add" {
			db.AddJUser(c, u, p)
		} else {
			db.ResetPwd(c, u, p)
		}

		c.Redirect(http.StatusFound, "/emaildetails")
	}
}
func pformeditor(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "formeditor.html", gin.H{
			"title": "",
		})
	} else if c.Request.Method == "POST" {

		//c.JSON(http.StatusOK, "[]")

	}
}

type excutStr struct {
	Suss bool
	Msg  string
}

func pmydetails(c *gin.Context) {
	if c.Request.Method == "GET" {

		c.HTML(http.StatusOK, "mydetails.html", gin.H{
			"user": GetCUser(c),
			"menu": GetMenuHtml(c),
		})
	} else if c.Request.Method == "POST" {

		rest := &excutStr{}
		m := c.Request.FormValue("method")
		if m == "resetpwd" {

			u := c.Value(susrname).(*db.UserBean)
			p := c.Request.FormValue("pwd")
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

		c.JSON(http.StatusOK, rest)

	}
}
func pprofile(c *gin.Context) {
	if c.Request.Method == "GET" {

		auths := GetAutus(c)
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"auths": auths,
		})
	} else if c.Request.Method == "POST" {

		//c.JSON(http.StatusOK, "[]")

	}
}
func puserlist(c *gin.Context) {
	if c.Request.Method == "GET" {

		//uss := db.GetAllJUsers(c)
		uss := db.GetAllUsers(c)

		c.HTML(http.StatusOK, "userlist.html", gin.H{
			"users": uss,
			"user":  GetCUser(c),
			"menu":  GetMenuHtml(c),
		})
	} else if c.Request.Method == "POST" {
		m := c.Request.FormValue("method")
		var e error
		result := &excutStr{}
		switch m {
		case "editinfo":
			e = db.Editinfo(c, c.Request.FormValue("username"), c.Request.FormValue("realname"))
			result.Msg = "信息修改成功"
			break
		case "resetpwd":
			e = db.ResetPwd(c, c.Request.FormValue("username"), c.Request.FormValue("pwd"))
			result.Msg = "密码重置成功"
			break
		case "adduser":
			u := &db.UserBean{}
			u.Username = strings.ToLower(c.Request.FormValue("username"))
			u.Pwd = c.Request.FormValue("pwd")
			u.Real_name = c.Request.FormValue("realname")
			u.Nick_name = c.Request.FormValue("realname")
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

		c.JSON(http.StatusOK, result)

	}
}
func pdatatables(c *gin.Context) {
	if c.Request.Method == "GET" {

		uss := db.GetAllJUsers(c)
		trs := make([]string, 0)
		if uss != nil {
			for _, v := range uss {
				trs = append(trs, v.Username)
			}
		}
		c.HTML(http.StatusOK, "datatables.html", gin.H{
			"trs": trs,
		})
	} else if c.Request.Method == "POST" {

		//c.JSON(http.StatusOK, "[]")

	}
}
func ptable(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "table.html", gin.H{
			"title": "",
		})
	} else if c.Request.Method == "POST" {

		//c.JSON(http.StatusOK, "[]")

	}
}
