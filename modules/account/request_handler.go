package account

import (
	"crud/dto"
	"crud/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerAdmin struct {
	ctr ControllerAdmin
}

func NewAdminRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerAdmin {
	return RequestHandlerAdmin{
		ctr: controllerAdmin{
			adminUseCase: useCaseAdmin{
				adminRepo: repository.NewActor(dbCrud),
			},
		}}
}

func (h RequestHandlerAdmin) CreateAdmin(c *gin.Context) {
	request := AdminParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateAdmin(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) GetAdminById(c *gin.Context) {
	request := AdminParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetAdminById(uint(adminId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) GetAdmins(c *gin.Context) {
	request := AdminParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	username := c.Query("username")
	//page := c.Query("page")
	page, err := strconv.ParseUint(c.Query("page"), 10, 64)
	res, err := h.ctr.GetAdmins(username, uint(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) UpdateAdminById(c *gin.Context) {
	request := UpdateAdmin{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.UpdateAdminById(request, uint(adminId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) DeleteAdminById(c *gin.Context) {
	request := AdminParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.DeleteAdminById(uint(adminId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) Login(c *gin.Context) {
	request := LoginParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) Register(c *gin.Context) {
	request := RegisterParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.Register(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}
