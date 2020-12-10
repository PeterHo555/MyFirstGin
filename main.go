package main

import (
	"fmt"
	"ginessential/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)


func main()  {
	InitConfig()
	db := common.InitDB()
	fmt.Println(db)

	r := gin.Default()
	r = CollectRoute(r)

	//r.Run() // listen and serve on 0.0.0.0:8080
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":"+port))
	}

	panic(r.Run())

}

func InitConfig()  {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

	}

}