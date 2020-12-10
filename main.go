package main

import (
	"fmt"
	"ginessential/common"
	"github.com/gin-gonic/gin"
)


func main()  {
	db := common.InitDB()
	fmt.Println(db)
	r := gin.Default()
	r = CollectRoute(r)
	//r.POST("/api/auth/register", controller.Register)
	//r.Run() // listen and serve on 0.0.0.0:8080
	panic(r.Run())

}



