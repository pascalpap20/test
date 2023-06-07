package customer

import (
	"crud/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterCustomer struct {
	CustomerRequestHandler RequestHandlerCustomer
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterCustomer {
	return RouterCustomer{CustomerRequestHandler: NewCustomerRequestHandler(
		dbCrud,
	)}
}

func (r RouterCustomer) Handle(router *gin.Engine) {
	basepath := "/customer"
	admin := router.Group(basepath)

	admin.POST("/", middleware.IsSuperAdminOrAdmin(),
		r.CustomerRequestHandler.CreateCustomer,
	)

	admin.GET("/:id", middleware.IsSuperAdminOrAdmin(),
		r.CustomerRequestHandler.GetCustomerById,
	)

	admin.GET("/", middleware.IsSuperAdminOrAdmin(),
		r.CustomerRequestHandler.GetCustomers,
	)

	admin.PUT("/:id", middleware.IsSuperAdminOrAdmin(),
		r.CustomerRequestHandler.UpdateCustomerById,
	)

	admin.DELETE("/:id", middleware.IsSuperAdminOrAdmin(),
		r.CustomerRequestHandler.DeleteCustomerById,
	)
}
