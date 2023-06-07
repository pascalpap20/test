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

	admin.POST("/", middleware.IsSuperAdmin(),
		r.AdminRequestHandler.CreateAdmin,
	)

	admin.GET("/:id", middleware.IsSuperAdminOrAdmin(),
		r.AdminRequestHandler.GetAdminById,
	)

	admin.GET("/", middleware.IsSuperAdminOrAdmin(),
		r.AdminRequestHandler.GetAdmins,
	)

	admin.PUT("/:id", middleware.IsSuperAdminOrAdmin(),
		r.AdminRequestHandler.UpdateAdminById,
	)

	admin.DELETE("/:id", middleware.IsSuperAdmin(),
		r.AdminRequestHandler.DeleteAdminById,
	)

	admin.POST("/login",
		r.AdminRequestHandler.Login,
	)

	admin.POST("/register", middleware.IsSuperAdminOrAdmin(),
		r.AdminRequestHandler.Register,
	)

	admin.GET("/register-approval", middleware.IsSuperAdmin(),
		r.AdminRequestHandler.GetRegisterApproval,
	)

	admin.PUT("/register-approval/:id", middleware.IsSuperAdmin(),
		r.AdminRequestHandler.UpdateRegisterApprovalStatusById,
	)

	admin.PUT("/set-activate/:id", middleware.IsSuperAdmin(),
		r.AdminRequestHandler.SetActivateAdminById,
	)

	admin.PUT("/set-deactivate/:id", middleware.IsSuperAdmin(),
		r.AdminRequestHandler.SetDeactivateAdminById,
	)

}
