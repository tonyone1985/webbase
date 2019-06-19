// mapper
package main

import (
	"webbase/control"
)

func GetMapper() map[string]control.HttpExecutor {
	mapper := make(map[string]control.HttpExecutor)
	mapper["userlist"] = &control.Userlist{}
	mapper["mydetails"] = &control.Mydetails{}

	return mapper
}
