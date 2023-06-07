package account

import (
	"crud/dto"
	"crud/repository"
	"fmt"
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
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
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
	//request := AdminParam{}
	//err := c.BindQuery(&request)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	//	return
	//}

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
	//request := AdminParam{}
	//err := c.BindQuery(&request)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	//	return
	//}

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
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
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
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
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
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	res, err := h.ctr.Register(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) GetRegisterApproval(c *gin.Context) {
	request := RegisterApprovalParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetRegisterApproval()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) UpdateRegisterApprovalStatusById(c *gin.Context) {
	request := UpdateRegisterApproval{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	registerApprovalId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	adminId := c.GetFloat64("id")
	roleId := c.GetFloat64("role_id")
	fmt.Println(fmt.Sprintf("%T %T", uint(adminId), roleId))
	adminInfo := map[string]uint{
		"id":      uint(adminId),
		"role_id": uint(roleId),
	}

	res, err := h.ctr.UpdateRegisterApprovalStatusById(request, uint(registerApprovalId), adminInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) SetActivateAdminById(c *gin.Context) {
	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.SetActivateAdminById(uint(adminId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) SetDeactivateAdminById(c *gin.Context) {
	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.SetDeactivateAdminById(uint(adminId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}
