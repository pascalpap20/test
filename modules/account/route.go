package account

import (
	"crud/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterAdmin struct {
	AdminRequestHandler RequestHandlerAdmin
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterAdmin {
	return RouterAdmin{AdminRequestHandler: NewAdminRequestHandler(
		dbCrud,
	)}
}

func (r RouterAdmin) Handle(router *gin.Engine) {
	basepath := "/admin"
	admin := router.Group(basepath)

	admin.POST("/",
		r.AdminRequestHandler.CreateAdmin,
	)

	admin.GET("/:id",
		r.AdminRequestHandler.GetAdminById,
	)

	admin.GET("/", middleware.IsAdmin(),
		r.AdminRequestHandler.GetAdmins,
	)

	admin.PUT("/:id",
		r.AdminRequestHandler.UpdateAdminById,
	)

	admin.DELETE("/:id",
		r.AdminRequestHandler.DeleteAdminById,
	)

	admin.POST("/login",
		r.AdminRequestHandler.Login,
	)

	admin.POST("/register",
		r.AdminRequestHandler.Register,
	)
}
