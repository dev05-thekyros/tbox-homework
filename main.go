package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tbox-homework/common"
	"tbox-homework/modules/user/usrtransport/ginusr"
)

func main() {
	//Loading configure file
	sc := common.InitServiceContext()

	// Setting REST API
	r := gin.Default()

	// routing
	// basic authentication endpoints
	api := r.Group("/api/v1")
	{
		api.POST("/sign-up", ginusr.SignUp(sc))
		api.POST("/verify-otp", ginusr.VerifyOtp(sc))
	}

	if err := r.Run(":8081"); err != nil {
		fmt.Println(err.Error())
	}

}
