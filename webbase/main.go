package main

import (
	"context"
	"control"
	"db"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/muesli/cache2go"
	"github.com/satori/go.uuid"
)

const cockey = "ljxu"

const sauths = "auths"
const STimeOut = 60 * 60

func GetAutus(c *gin.Context) []*db.AuthBean {
	return c.Value(sauths).([]*db.AuthBean)
}
func GetCUser(c *gin.Context) *db.UserBean {
	return c.Value(control.S_USER_KEY).(*db.UserBean)
}
func GetMenuHtml(c *gin.Context) template.HTML {
	auths := GetAutus(c)
	menu := ""
	for k, v2 := range auths {
		if v2.Pid == 0 {
			if k != 0 {
				menu += `</ul></li>`
			}
			menu += `<li class="slide" >
			<a class="side-menu__item" data-toggle="slide" href="#"><i class="side-menu__icon fa fa-table"></i><span class="side-menu__label">` +
				v2.Auth_name.String + `</span><i class="angle fa fa-angle-right"></i></a><ul class="slide-menu">`
		} else {
			menu += `<li><a href="` + v2.Auth_val.String + `" class="slide-item">` + v2.Auth_name.String + `</a></li>`
		}
	}
	if len(menu) > 0 {
		menu += "</ul></li>"
	}
	return template.HTML(menu)
}

func plogin(c *gin.Context) {
	if c.Request.Method == "GET" {
		ses := sessions.Default(c)
		ses.Set(cockey, "")
		ses.Save()
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "",
		})
	} else if c.Request.Method == "POST" {

		//control.Login(c, c.Request.FormValue("mail"), c.Request.FormValue("password"))
		//u := db.Login2(c, c.Request.FormValue("mail"), c.Request.FormValue("password"))
		u := control.Login(c, c.Request.FormValue("mail"), c.Request.FormValue("password"))
		if u != nil {
			ses := sessions.Default(c)
			uid, _ := uuid.NewV4()
			uids := uid.String()
			cache.Add(uids, STimeOut*time.Second, u)
			ses.Set(cockey, uids)
			ses.Save()

			c.Redirect(http.StatusFound, "/mydetails")
		} else {
			c.Redirect(http.StatusFound, "/login")
		}
		//c.JSON(http.StatusOK, "[]")

	}

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

func initAuthMap(c context.Context) {
	roles, _ := db.GetRoles(c)
	auths, _ := db.GetAuths(c)
	g_authmap = make(map[int][]*db.AuthBean)

	for _, r := range roles {
		authids := getauthids(r)
		roleauths := getauths(authids, auths)
		roleauthsorderd := orderauths(roleauths)
		g_authmap[r.Role_id] = roleauthsorderd
	}

}
func islogin(c *gin.Context) bool {
	ses := sessions.Default(c)
	k := ses.Get(cockey)

	if k != "" {
		v, e := cache.Value(k)
		if e != nil || v == nil {
			return false
		}
		c.Set(control.S_USER_KEY, v.Data())

		initAuthMap(c)
		u := v.Data().(*db.UserBean)

		c.Set(sauths, g_authmap[u.Role_id])
		return true
	}
	return false
}

var cache *cache2go.CacheTable
var g_authmap map[int][]*db.AuthBean

func checkAuth(page string, c *gin.Context) bool {
	if page == "login" {
		return true
	}
	if !islogin(c) {
		return false
	}
	if page == "mydetails" {
		return true
	}
	aus := GetAutus(c)
	page2au := "/" + page
	for _, a := range aus {
		if !a.Auth_val.Valid {
			continue
		}
		if a.Auth_val.String == page2au {
			return true
		}
	}

	return false

}

type Config struct {
	Bind   string `json:"bind"`
	Db     string `json:"bd"`
	Maildb string `json:""maildb`
}

func SaveDefCfg(c *Config) {
	dir := "../conf"
	file := "../conf/config.json"
	_, err := os.Stat(file)
	if err == nil || os.IsExist(err) {
		return
	}

	_, err = os.Stat(dir)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	d, _ := json.MarshalIndent(c, "", "\t")
	ioutil.WriteFile(file, d, 0666)
}
func ReadConfg() *Config {
	filePth := "../conf/config.json"
	defcfg := &Config{
		Bind:   ":8080",
		Db:     "root:@tcp(localhost:3306)/ljx",
		Maildb: "root:@tcp(localhost:3306)/mail",
	}
	f, err := os.Open(filePth)
	if err != nil {
		SaveDefCfg(defcfg)
		return defcfg
	}
	d, e := ioutil.ReadAll(f)
	if e != nil {
		SaveDefCfg(defcfg)
		return defcfg
	}
	rcfg := &Config{}
	if json.Unmarshal(d, &rcfg) != nil {
		SaveDefCfg(defcfg)
		return defcfg
	}
	return rcfg

}
func main() {

	config := ReadConfg()
	db.Init(config.Db, config.Maildb)

	router := gin.Default()
	//router.Static("/assets", "../static/assets")
	hs := http.FileServer(http.Dir("../static"))

	store := sessions.NewCookieStore([]byte("secret"))
	cache = cache2go.Cache("myCache")

	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("../static/pages/*")

	//router.GET("/login", func(c *gin.Context) {
	//
	//	c.HTML(http.StatusOK, "login.html", gin.H{
	//		"title": "",
	//	})
	//})
	//router.Static("/assets", "../static/assets")

	// pageroutes := make(map[string]gin.HandlerFunc)
	// pageroutes["index"] = pindex
	// pageroutes["login"] = plogin
	// pageroutes["emaildetails"] = pemaildetails
	// pageroutes["formeditor"] = pformeditor
	// pageroutes["table"] = ptable
	// pageroutes["userlist"] = puserlist
	// pageroutes["profile"] = pprofile
	// pageroutes["mydetails"] = pmydetails
	router.GET("/", plogin)

	pageroutes2 := make(map[string]control.HttpExecutor)
	postroutes2 := make(map[string]control.HttpExecutor)

	//_ = pageroutes2
	pageroutes2["userlist"] = &control.Userlist{}
	postroutes2["userlist"] = &control.Userlist{}
	//pageroutes2["profile"] = pprofile
	pageroutes2["mydetails"] = &control.Mydetails{}
	postroutes2["mydetails"] = &control.Mydetails{}

	router.POST("/:p1", func(c *gin.Context) {
		page := c.Param("p1")

		if page == "login" {
			plogin(c)
			return
		}

		if islogin(c) {
			h := postroutes2[page]
			if h == nil {
				c.HTML(http.StatusNotFound, "404.html", gin.H{
					"title": "",
				})
				//c.HTML()
				return
			}

			if checkAuth(page, c) {
				pdata, e := h.Post(c, c.Request)
				if e == nil {
					c.HTML(http.StatusNotFound, "404.html", nil)
				} else {
					if pdata != nil {
						c.JSON(http.StatusOK, pdata)
					} else {
						c.JSON(http.StatusExpectationFailed, nil)
					}

				}

			} else {
				c.HTML(http.StatusNotFound, "505.html", nil)
			}

		} else {
			c.HTML(http.StatusNotFound, "404.html", nil)
		}
	})
	router.GET("/:p1/*p2", func(c *gin.Context) {
		p1 := c.Param("p1")
		if p1 == "assets" {
			hs.ServeHTTP(c.Writer, c.Request)
			return
		}
	})
	router.GET("/:p1", func(c *gin.Context) {

		page := c.Param("p1")

		if page == "login" {
			plogin(c)
			return
		}

		if islogin(c) {
			h := pageroutes2[page]

			if h == nil {
				c.HTML(http.StatusNotFound, "404.html", nil)
				//c.HTML()
				return
			}
			if checkAuth(page, c) {
				pdata, e := h.Get(c)
				if e != nil {
					c.HTML(http.StatusNotFound, "404.html", nil)
				} else {
					var ldata gin.H = nil
					if pdata != nil {
						ldata = pdata
					} else {
						ldata = gin.H{}

					}
					ldata["user"] = GetCUser(c)
					ldata["menu"] = GetMenuHtml(c)
					c.HTML(http.StatusOK, page+".html", ldata)
				}
			} else {
				c.HTML(http.StatusNotFound, "505.html", nil)
			}
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "",
			})
		}

		//_ := db.GetAuths("1")

	})
	//router.Static("/assets", "../static/assets")

	//router.GET("/index.html", func(c *gin.Context) {
	//
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "",
	//	})
	//})

	router.Run(config.Bind)

}
