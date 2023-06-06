package account

import (
	"crud/entity"
	"crud/repository"
	"crud/utils/auth"
	"time"
)

type UseCaseAdmin interface {
	CreateAdmin(user AdminParam) (entity.Actor, error)
	GetAdminById(id uint) (entity.Actor, error)
	GetAdmins(username string, page uint) ([]entity.Actor, error)
	UpdateAdminById(admin UpdateAdmin, id uint) (entity.Actor, error)
	DeleteAdminById(id uint) (entity.Actor, error)
	Login(admin LoginParam) (*entity.Actor, error)
	Register(admin RegisterParam) (*entity.Actor, error)
}

type useCaseAdmin struct {
	adminRepo repository.ActorInterfaceRepo
}

func (uc useCaseAdmin) CreateAdmin(admin AdminParam) (entity.Actor, error) {
	var newAdmin *entity.Actor
	var salt string

	admin.Password, salt = auth.GenerateHash(admin.Password)

	newAdmin = &entity.Actor{
		Username:   admin.Username,
		RoleID:     admin.RoleID,
		Password:   admin.Password,
		IsVerified: admin.IsVerified,
		IsActive:   admin.IsActive,
		Salt:       salt,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err := uc.adminRepo.CreateActor(newAdmin)
	if err != nil {
		return *newAdmin, err
	}
	return *newAdmin, nil
}

func (uc useCaseAdmin) GetAdminById(id uint) (entity.Actor, error) {
	var admin entity.Actor
	admin, err := uc.adminRepo.GetActorById(id)
	return admin, err
}

func (uc useCaseAdmin) GetAdmins(username string, page uint) ([]entity.Actor, error) {
	var admin []entity.Actor
	admin, err := uc.adminRepo.GetActors(username, page)
	return admin, err
}

func (uc useCaseAdmin) UpdateAdminById(admin UpdateAdmin, id uint) (entity.Actor, error) {
	var updateAdmin *entity.Actor
	var salt string

	admin.Password, salt = auth.GenerateHash(admin.Password)

	updateAdmin = &entity.Actor{
		Username:   admin.Username,
		Password:   admin.Password,
		IsVerified: admin.IsVerified,
		IsActive:   admin.IsActive,
		Salt:       salt,
	}

	_, err := uc.adminRepo.UpdateActorById(updateAdmin, id)
	if err != nil {
		return *updateAdmin, err
	}
	return *updateAdmin, nil
}

func (uc useCaseAdmin) DeleteAdminById(id uint) (entity.Actor, error) {
	var admin entity.Actor
	admin, err := uc.adminRepo.DeleteActorById(id)
	return admin, err
}

func (uc useCaseAdmin) Login(admin LoginParam) (*entity.Actor, error) {
	var newAdmin *entity.Actor

	newAdmin = &entity.Actor{
		Username: admin.Username,
		Password: admin.Password,
	}

	res, err := uc.adminRepo.Login(newAdmin)
	if err != nil {
		return newAdmin, err
	}

	return res, nil
}

func (uc useCaseAdmin) Register(admin RegisterParam) (*entity.Actor, error) {
	var newAdmin *entity.Actor

	newAdmin = &entity.Actor{
		RoleID:     2,
		Username:   admin.Username,
		Password:   admin.Password,
		IsVerified: "false",
		IsActive:   "false",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	res, err := uc.adminRepo.Register(newAdmin)
	if err != nil {
		return newAdmin, err
	}

	return res, nil
}
