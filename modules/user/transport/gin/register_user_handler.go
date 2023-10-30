package ginuser

import (
	"food-client/common"
	"food-client/component"
	"food-client/component/hasher"
	"food-client/modules/user/biz"
	"food-client/modules/user/model"
	"food-client/modules/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data model.UserCreation

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := biz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
