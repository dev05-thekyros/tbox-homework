package ginusr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tbox-homework/common"
	"tbox-homework/modules/user/usrhandler"
	"tbox-homework/modules/user/usrmodel"
	"tbox-homework/modules/user/usrrepo"
	"tbox-homework/modules/user/usrstorage"
)

func SignUp(sc *common.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var signUpUserModel usrmodel.SignUpUser
		if err := c.ShouldBindJSON(&signUpUserModel); err != nil {
			c.JSON(http.StatusOK, usrmodel.SignUpUsrResponse{Message: err.Error(), IsError: true})
			return
		}

		if err := signUpUserModel.Validate(); err != nil {
			c.JSON(http.StatusOK, usrmodel.SignUpUsrResponse{Message: err.Error(), IsError: true})
			return
		}

		ctx := c.Request.Context()
		redis := common.NewRedisStorageProvider(ctx, sc.RedisClient, sc.Logger)
		storage := usrstorage.NewUserMongoStorage(sc.Mongo)
		smsService := common.NewTwilloSMSService()
		rp := usrrepo.NewUserRepo(storage, redis, smsService)
		hdl := usrhandler.NewSignUpUserHandler(rp)

		accessToken, err := hdl.Response(c.Request.Context(), signUpUserModel.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusOK, usrmodel.SignUpUsrResponse{Message: err.Error(), IsError: true})
			return
		}
		if accessToken != "" {
			c.JSON(http.StatusOK, usrmodel.SignUpUsrResponse{Message: "Welcome to system!", IsError: false, Data: usrmodel.User{AccessToken: accessToken}})
		} else {
			c.JSON(http.StatusOK, usrmodel.SignUpUsrResponse{
				Message: fmt.Sprintf("New OTP code is sent to your phone number %s!", signUpUserModel.PhoneNumber),
				IsError: false,
				Data:    usrmodel.User{AccessToken: accessToken}})
		}
	}
}
