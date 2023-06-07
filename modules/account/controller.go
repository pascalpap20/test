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
	GetRegisterApproval() (FindAllRegisterApproval, error)
	UpdateRegisterApprovalStatusById(req UpdateRegisterApproval, id uint, adminInfo map[string]uint) (any, error)
	SetActivateAdminById(id uint) (any, error)
	SetDeactivateAdminById(id uint) (any, error)
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
			MessageTitle: "Success create admin",
			Message:      "Success create",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username: admin.Username,
			RoleID:   admin.RoleID,
			//Password:   admin.Password,
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
		MessageTitle: "Success get admin",
		Message:      "Success retrieve an admin data",
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
		MessageTitle: "Success get all admins",
		Message:      "Success retrieve all admins",
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
			MessageTitle: "Success update admin",
			Message:      "Success update",
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
			MessageTitle: "Success register new admin",
			Message:      "Success register",
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

func (uc controllerAdmin) GetRegisterApproval() (FindAllRegisterApproval, error) {
	var res FindAllRegisterApproval
	registerApproval, err := uc.adminUseCase.GetRegisterApproval()
	if err != nil {
		return FindAllRegisterApproval{}, err
	}

	res.Data = registerApproval
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get register approval",
		Message:      "Success",
		ResponseTime: "",
	}

	return res, nil
}

func (uc controllerAdmin) UpdateRegisterApprovalStatusById(req UpdateRegisterApproval, id uint, adminInfo map[string]uint) (any, error) {

	registerApproval, err := uc.adminUseCase.UpdateRegisterApprovalStatusById(req, id, adminInfo)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessUpdateRegisterApproval{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success update status",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: RegisterApprovalParam{
			AdminID:      registerApproval.ID,
			SuperAdminID: registerApproval.SuperAdminID,
			Status:       registerApproval.Status,
		},
	}
	return res, nil
}

func (uc controllerAdmin) SetActivateAdminById(id uint) (any, error) {

	admin, err := uc.adminUseCase.SetActivateAdminById(id)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success activate an admin",
			Message:      "Success activate",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username:   admin.Username,
			RoleID:     admin.RoleID,
			IsVerified: admin.IsVerified,
			IsActive:   admin.IsActive,
		},
	}
	return res, nil
}

func (uc controllerAdmin) SetDeactivateAdminById(id uint) (any, error) {

	admin, err := uc.adminUseCase.SetDeactivateAdminById(id)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success deactivate an admin",
			Message:      "Success deactivate",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username:   admin.Username,
			RoleID:     admin.RoleID,
			IsVerified: admin.IsVerified,
			IsActive:   admin.IsActive,
		},
	}
	return res, nil
}
