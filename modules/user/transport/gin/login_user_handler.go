package ginuser

import (
	"food-client/common"
	"food-client/component"
	"food-client/component/hasher"
	"food-client/component/tokenprovider/jwt"
	"food-client/modules/user/biz"
	"food-client/modules/user/model"
	"food-client/modules/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := storage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := biz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
