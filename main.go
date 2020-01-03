package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tbox-homework/common"
	"tbox-homework/middleware"
	"tbox-homework/modules/user/usrstorage"
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

	// authenticated required endpoints
	storage := usrstorage.NewUserMongoStorage(sc.Mongo)
	authMiddleware := middleware.NewAuthMiddleware(storage)
	authRequired := api.Group("/")
	authRequired.Use(authMiddleware.AuthRequired)
	{
		// Define required endpoint here

	}

	if err := r.Run(":8081"); err != nil {
		fmt.Println(err.Error())
	}

}
