package main

import (

	//"github.com/gin-gonic/contrib/sessions"
	"net/http"

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

func pformeditor(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "formeditor.html", gin.H{
			"title": "",
		})
	} else if c.Request.Method == "POST" {

		//c.JSON(http.StatusOK, "[]")

	}
}

// func pmydetails(c *gin.Context) {
// 	if c.Request.Method == "GET" {

// 		c.HTML(http.StatusOK, "mydetails.html", gin.H{
// 			"user": GetCUser(c),
// 			"menu": GetMenuHtml(c),
// 		})
// 	} else if c.Request.Method == "POST" {

// 		rest := &excutStr{}
// 		m := c.Request.FormValue("method")
// 		if m == "resetpwd" {

// 			u := c.Value(susrname).(*db.UserBean)
// 			p := c.Request.FormValue("pwd")
// 			if e := db.ResetPwd(c, u.Username, p); e == nil {
// 				rest.Msg = "密码重置成功"
// 				rest.Suss = true
// 			} else {
// 				rest.Msg = "密码重置失败" + e.Error()
// 				rest.Suss = false
// 			}

// 		} else {
// 			rest.Msg = "参数错误"
// 			rest.Suss = false
// 		}

// 		c.JSON(http.StatusOK, rest)

// 	}
// }
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

// func puserlist(c *gin.Context) {
// 	if c.Request.Method == "GET" {

// 		//uss := db.GetAllJUsers(c)
// 		uss := db.GetAllUsers(c)

// 		c.HTML(http.StatusOK, "userlist.html", gin.H{
// 			"users": uss,
// 			"user":  GetCUser(c),
// 			"menu":  GetMenuHtml(c),
// 		})
// 	} else if c.Request.Method == "POST" {
// 		m := c.Request.FormValue("method")
// 		var e error
// 		result := &excutStr{}
// 		switch m {
// 		case "editinfo":
// 			e = db.Editinfo(c, c.Request.FormValue("username"), c.Request.FormValue("realname"))
// 			result.Msg = "信息修改成功"
// 			break
// 		case "resetpwd":
// 			e = db.ResetPwd(c, c.Request.FormValue("username"), c.Request.FormValue("pwd"))
// 			result.Msg = "密码重置成功"
// 			break
// 		case "adduser":
// 			u := &db.UserBean{}
// 			u.Username = strings.ToLower(c.Request.FormValue("username"))
// 			u.Pwd = c.Request.FormValue("pwd")
// 			u.Real_name = c.Request.FormValue("realname")
// 			u.Nick_name = c.Request.FormValue("realname")
// 			u.Role_id = 101
// 			e = db.AddUser(c, u)
// 			result.Msg = "用户添加成功"
// 			break
// 		default:
// 			e = errors.New("未知错误")
// 			break
// 		}
// 		if e == nil {
// 			result.Suss = true

// 		} else {
// 			result.Suss = false
// 			result.Msg = e.Error()
// 		}

// 		c.JSON(http.StatusOK, result)

// 	}
// }

func ptable(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "table.html", gin.H{
			"title": "",
		})
	} else if c.Request.Method == "POST" {

		//c.JSON(http.StatusOK, "[]")

	}
}
