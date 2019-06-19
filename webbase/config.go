// Config
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/sipt/GoJsoner"
)

var DEFConfig []byte = []byte(`
{
	//httplisten 
	"bind": ":8080",
	//db connection string
	"bd": "root:123456@tcp(localhost:3306)/ljx",	
	"Maildb": "root:123456@tcp(localhost:3306)/mail"
}
`)

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
func SaveDef() {
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

	ioutil.WriteFile(file, DEFConfig, 0666)
}
func ReadConfg() *Config {
	filePth := "../conf/config.json"
	var d []byte = nil
	f, err := os.Open(filePth)
	defer f.Close()
	if err != nil {
		d = DEFConfig
		SaveDef()

		goto R
	}
	d, err = ioutil.ReadAll(f)
	if err != nil {
		d = DEFConfig
		SaveDef()
	}
R:
	dstr, err := GoJsoner.Discard(string(d))

	if err != nil {
		log.Println("error config file")
		return nil
	}
	d2 := []byte(dstr)
	rcfg := &Config{}
	if json.Unmarshal(d2, &rcfg) != nil {
		log.Println("error config file")
		return nil
	}
	return rcfg

}
