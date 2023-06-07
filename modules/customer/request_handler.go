package customer

import (
	"crud/dto"
	"crud/repository"
	"crud/utils/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerCustomer struct {
	ctr ControllerCustomer
}

func NewCustomerRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerCustomer {
	return RequestHandlerCustomer{
		ctr: controllerCustomer{
			customerUseCase: useCaseCustomer{
				customerRepo: repository.NewCustomer(dbCrud),
			},
		}}
}

func (h RequestHandlerCustomer) CreateCustomer(c *gin.Context) {
	request := CustomerParam{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	_, err = helper.ValidateRequestCreateOrUpdateCustomer(request.FirstName, request.LastName, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	res, err := h.ctr.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) GetCustomerById(c *gin.Context) {
	request := CustomerParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetCustomerById(uint(customerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) GetCustomers(c *gin.Context) {
	request := CustomerParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	name := c.Query("name")
	email := c.Query("email")
	page, err := strconv.ParseUint(c.Query("page"), 10, 64)
	res, err := h.ctr.GetCustomers(name, email, uint(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) UpdateCustomerById(c *gin.Context) {
	request := UpdateCustomer{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	_, err = helper.ValidateRequestCreateOrUpdateCustomer(request.FirstName, request.LastName, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.UpdateCustomerById(request, uint(customerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) DeleteCustomerById(c *gin.Context) {
	request := CustomerParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.DeleteCustomerById(uint(customerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}
