package account

import (
	"crud/dto"
	"crud/utils/auth"
)

type ControllerAdmin interface {
	CreateAdmin(req AdminParam) (any, error)
	GetAdminById(id uint) (FindAdmin, error)
	GetAdmins(username string, page uint) (FindAllAdmins, error)
	UpdateAdminById(req UpdateAdmin, id uint) (any, error)
	DeleteAdminById(id uint) (FindAdmin, error)
	Login(req LoginParam) (any, error)
	Register(req RegisterParam) (any, error)
}

type controllerAdmin struct {
	adminUseCase UseCaseAdmin
}

func (uc controllerAdmin) CreateAdmin(req AdminParam) (any, error) {

	admin, err := uc.adminUseCase.CreateAdmin(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create user",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username:   admin.Username,
			RoleID:     admin.RoleID,
			Password:   admin.Password,
			IsVerified: admin.IsVerified,
			IsActive:   admin.IsActive,
		},
	}
	return res, nil
}

func (uc controllerAdmin) GetAdminById(id uint) (FindAdmin, error) {
	var res FindAdmin
	admin, err := uc.adminUseCase.GetAdminById(id)
	if err != nil {
		return FindAdmin{}, err
	}
	res.Data = admin
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get user",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (uc controllerAdmin) GetAdmins(username string, page uint) (FindAllAdmins, error) {
	var res FindAllAdmins
	admin, err := uc.adminUseCase.GetAdmins(username, page)
	if err != nil {
		return FindAllAdmins{}, err
	}

	res.Data = admin
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get user",
		Message:      "Success",
		ResponseTime: "",
	}

	return res, nil
}

func (uc controllerAdmin) UpdateAdminById(req UpdateAdmin, id uint) (any, error) {

	admin, err := uc.adminUseCase.UpdateAdminById(req, id)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create user",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username:   admin.Username,
			RoleID:     admin.RoleID,
			Password:   admin.Password,
			IsVerified: admin.IsVerified,
			IsActive:   admin.IsActive,
		},
	}
	return res, nil
}

func (uc controllerAdmin) DeleteAdminById(id uint) (FindAdmin, error) {
	var res FindAdmin
	admin, err := uc.adminUseCase.DeleteAdminById(id)
	if err != nil {
		return FindAdmin{}, err
	}
	res.Data = admin
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success Delete Admin",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (uc controllerAdmin) Login(req LoginParam) (any, error) {

	admin, err := uc.adminUseCase.Login(req)
	if err != nil {
		return SuccessCreate{}, err
	}

	accessToken := auth.GenerateTokenJwt(admin)

	res := SuccessLogin{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success Login",
			Message:      "Success Login",
			ResponseTime: "",
		},
		Data: SuccessLoginParam{
			Username:    admin.Username,
			RoleID:      admin.RoleID,
			IsVerified:  admin.IsVerified,
			IsActive:    admin.IsActive,
			AccessToken: accessToken,
		},
	}
	return res, nil
}

func (uc controllerAdmin) Register(req RegisterParam) (any, error) {

	admin, err := uc.adminUseCase.Register(req)
	if err != nil {
		return SuccessCreate{}, err
	}

	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success Login",
			Message:      "Success Login",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username:   admin.Username,
			RoleID:     admin.RoleID,
			Password:   admin.Password,
			IsVerified: admin.IsVerified,
			IsActive:   admin.IsActive,
		},
	}
	return res, nil
}
