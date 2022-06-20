package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	bootstrap.SetupRoute(router)

	//运行服务
	err := router.Run(":8000")

	if err != nil {
		//错误处理，端口占用了或者其他错误
		fmt.Println(err.Error())
	}
}
