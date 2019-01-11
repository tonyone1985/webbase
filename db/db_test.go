package db

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"log"
	"testing"
)

func deftes(t *testing.T) {
	ctx := context.Background()
	uu := GetUser(ctx, "cxx", "cxx")
	if uu == nil {
		log.Println(" uu == nil")
		t.Error()
		return
	}

	log.Println(uu.Nick_name)
	u := GetRole(ctx, uu.Role_id)
	if u == nil {
		log.Println("err")
	}
	log.Println(u.Remark)
}

func Test_fff(t *testing.T) {
	sr := sha1.Sum([]byte("111"))

	log.Println(len(base64.StdEncoding.EncodeToString(sr[:])))
	log.Printf("%x\n", sha1.Sum([]byte("111")))

}
