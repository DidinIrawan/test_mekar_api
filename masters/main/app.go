package main

import (
	"testAPI/configs"
	"testAPI/masters/apis"
	"testAPI/utils"
)

func main()  {
	conf := configs.NewAppConfig()
	db, err := configs.InitDB(conf)

	utils.ErrorCheck(err, "Print")
	myRoute := configs.CreateRouter()
	apis.Init(myRoute,db)
	configs.RunServer(myRoute)
	//fmt.Sprintf("database succes",db)
}