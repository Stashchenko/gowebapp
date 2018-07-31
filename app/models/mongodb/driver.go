package mongodb

import (
	"fmt"
)

var MaxPool int
var PATH string
var DBNAME string

func CheckAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = PATH
		fmt.Printf("=CheckAndInitServiceConnection: Path = %v", PATH)
		err := service.New()
		if err != nil {
			panic(err)
		}
	}
}
