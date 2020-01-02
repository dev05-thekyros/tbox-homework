package ginusr

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tbox-homework/common"
	"tbox-homework/modules/user/usrhandler"
	"tbox-homework/modules/user/usrmodel"
	"tbox-homework/modules/user/usrrepo"
	"tbox-homework/modules/user/usrstorage"
)

func VerifyOtp(sc *common.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rq usrmodel.VerifyOTPRequest
		if err := c.ShouldBindJSON(&rq); err != nil {
			c.JSON(http.StatusOK, usrmodel.VerifyOTPResponse{Message: err.Error(), IsError: true})
			return
		}

		if err := rq.Validate(); err != nil {
			c.JSON(http.StatusOK, usrmodel.VerifyOTPResponse{Message: err.Error(), IsError: true})
			return
		}

		ctx := c.Request.Context()
		redis := common.NewRedisStorageProvider(ctx, sc.RedisClient, sc.Logger)
		storage := usrstorage.NewUserMongoStorage(sc.Mongo)
		smsService := common.NewTwilloSMSService()
		rp := usrrepo.NewUserRepo(storage, redis, smsService)
		hdl := usrhandler.NewVerifyUserOTPrHandler(rp)

		accessToken, err := hdl.Response(c.Request.Context(), rq.PhoneNumber, rq.OTP)
		if err != nil {
			c.JSON(http.StatusOK, usrmodel.VerifyOTPResponse{Message: err.Error(), IsError: true})
			return
		}
		c.JSON(http.StatusOK, usrmodel.VerifyOTPResponse{Message: "Welcome to system!", IsError: false, Data: usrmodel.User{AccessToken: accessToken}})
	}
}
