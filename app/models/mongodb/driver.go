package mongodb

import "log"

var MaxPool int
var PATH string
var DBNAME string

func CheckAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = PATH
		log.Printf("=CheckAndInitServiceConnection: Path = %v", PATH)
		err := service.New()
		if err != nil {
			panic(err)
		}
	}
}
