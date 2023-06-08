package unit_test

import (
	mocks "crud/mock"
	"crud/utils/auth"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"crud/entity" // Import the entity package
	"crud/modules/account"
)

func TestCreateAdmin(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Prepare the input data
	admin := account.AdminParam{
		Username:   "admin1",
		RoleID:     1,
		Password:   "password123",
		IsVerified: "false",
		IsActive:   "true",
	}

	// Prepare the expected output
	password, salt := auth.GenerateHash(admin.Password)

	expectedAdmin := entity.Actor{
		Username:   admin.Username,
		RoleID:     admin.RoleID,
		Password:   password,
		IsVerified: admin.IsVerified,
		IsActive:   admin.IsActive,
		Salt:       salt,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Set up expectations
	mockRepo.On("CreateActor", mock.AnythingOfType("*entity.Actor")).Return(&expectedAdmin, nil)

	// Call the method under test
	result, err := useCase.CreateAdmin(admin)

	isSame := bcrypt.CompareHashAndPassword([]byte(password), []byte(admin.Password+salt))
	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin.Username, result.Username)
	assert.Equal(t, expectedAdmin.RoleID, result.RoleID)
	assert.Equal(t, expectedAdmin.IsVerified, result.IsVerified)
	assert.Equal(t, expectedAdmin.IsActive, result.IsActive)
	assert.WithinDuration(t, expectedAdmin.CreatedAt, result.CreatedAt, time.Second)
	assert.WithinDuration(t, expectedAdmin.UpdatedAt, result.UpdatedAt, time.Second)
	// Add password hashing comparison
	assert.True(t, isSame == nil)
}

func TestGetAdminById(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Prepare the input data
	adminID := uint(1)

	// Prepare the expected output
	expectedAdmin := entity.Actor{
		ID:        adminID,
		Username:  "admin1",
		RoleID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set up expectations
	mockRepo.On("GetActorById", adminID).Return(expectedAdmin, nil)

	// Call the method under test
	admin, err := useCase.GetAdminById(adminID)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin.ID, admin.ID)
	assert.Equal(t, expectedAdmin.Username, admin.Username)
	assert.Equal(t, expectedAdmin.RoleID, admin.RoleID)
	assert.WithinDuration(t, expectedAdmin.CreatedAt, admin.CreatedAt, time.Second)
	assert.WithinDuration(t, expectedAdmin.UpdatedAt, admin.UpdatedAt, time.Second)
}

func TestGetAdmins(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Prepare the input data
	username := ""
	page := uint(1)

	// Prepare the expected output
	expectedAdmins := []entity.Actor{
		{
			ID:        1,
			Username:  "admin1",
			RoleID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Username:  "admin2",
			RoleID:    2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Set up expectations
	mockRepo.On("GetActors", username, page).Return(expectedAdmins, nil)

	// Call the method under test
	admins, err := useCase.GetAdmins(username, page)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, len(expectedAdmins), len(admins))
	for i := 0; i < len(expectedAdmins); i++ {
		assert.Equal(t, expectedAdmins[i].ID, admins[i].ID)
		assert.Equal(t, expectedAdmins[i].Username, admins[i].Username)
		assert.Equal(t, expectedAdmins[i].RoleID, admins[i].RoleID)
		assert.WithinDuration(t, expectedAdmins[i].CreatedAt, admins[i].CreatedAt, time.Second)
		assert.WithinDuration(t, expectedAdmins[i].UpdatedAt, admins[i].UpdatedAt, time.Second)
	}
}

func TestUpdateAdminById(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Prepare the input data
	admin := account.UpdateAdmin{
		Username:   "admin1",
		Password:   "password123",
		IsVerified: "false",
		IsActive:   "true",
	}
	adminId := uint(1)

	// Prepare the expected output
	password, salt := auth.GenerateHash(admin.Password)

	expectedAdmin := entity.Actor{
		Username:   admin.Username,
		Password:   password,
		IsVerified: admin.IsVerified,
		IsActive:   admin.IsActive,
		Salt:       salt,
	}

	// Set up expectations
	mockRepo.On("UpdateActorById", mock.AnythingOfType("*entity.Actor"), adminId).Return(&expectedAdmin, nil)

	// Call the method under test
	result, err := useCase.UpdateAdminById(admin, adminId)

	isSame := bcrypt.CompareHashAndPassword([]byte(password), []byte(admin.Password+salt))
	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin.Username, result.Username)
	assert.Equal(t, expectedAdmin.RoleID, result.RoleID)
	assert.Equal(t, expectedAdmin.IsVerified, result.IsVerified)
	assert.Equal(t, expectedAdmin.IsActive, result.IsActive)
	assert.WithinDuration(t, expectedAdmin.UpdatedAt, result.UpdatedAt, time.Second)
	// Add password hashing comparison
	assert.True(t, isSame == nil)
}

func TestDeleteAdminById(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Prepare the input data
	id := uint(1)

	// Prepare the expected output
	expectedAdmin := entity.Actor{
		ID: 1,
	}

	// Set up expectations
	mockRepo.On("DeleteActorById", id).Return(expectedAdmin, nil)

	// Call the method under test
	deletedAdmin, err := useCase.DeleteAdminById(id)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin.ID, deletedAdmin.ID)
}

func TestLogin(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Prepare the input data
	admin := account.LoginParam{
		Username: "testuser",
		Password: "testpassword",
	}

	// Prepare the expected output
	expectedActor := &entity.Actor{
		Username: admin.Username,
		Password: admin.Password,
	}

	// Set up expectations
	mockRepo.On("Login", expectedActor).Return(expectedActor, nil)

	// Call the method under test
	actor, err := useCase.Login(admin)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)
}

func TestRegister(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Prepare the input data
	admin := account.RegisterParam{
		Username: "testuser",
		Password: "testpassword",
	}

	// Prepare the expected output
	password, salt := auth.GenerateHash(admin.Password)

	expectedActor := &entity.Actor{
		RoleID:     2,
		Username:   admin.Username,
		Password:   admin.Password,
		IsVerified: "false",
		IsActive:   "false",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Set up expectations
	mockRepo.On("Register", mock.AnythingOfType("*entity.Actor")).Return(expectedActor, nil)

	// Call the method under test
	result, err := useCase.Register(admin)

	isSame := bcrypt.CompareHashAndPassword([]byte(password), []byte(admin.Password+salt))
	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedActor.Username, result.Username)
	assert.Equal(t, expectedActor.RoleID, result.RoleID)
	assert.Equal(t, expectedActor.IsVerified, result.IsVerified)
	assert.Equal(t, expectedActor.IsActive, result.IsActive)
	assert.WithinDuration(t, expectedActor.UpdatedAt, result.UpdatedAt, time.Second)
	assert.True(t, isSame == nil)
}

func TestGetRegisterApproval(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create the use case instance with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Create the expected result
	expectedApproval := []entity.RegisterApproval{
		{ID: 1, AdminID: 2, SuperAdminID: 1, Status: "pending"},
		{ID: 2, AdminID: 3, SuperAdminID: 1, Status: "approved"},
		{ID: 3, AdminID: 4, SuperAdminID: 1, Status: "rejected"},
	}

	// Set up the mock repository expectation
	mockRepo.On("GetRegisterApproval").Return(expectedApproval, nil)

	// Call the method on the use case
	approval, err := useCase.GetRegisterApproval()

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedApproval, approval)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestUpdateRegisterApprovalStatusById(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create the use case instance with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Create the input parameters
	reg := account.UpdateRegisterApproval{
		Status: "approved",
	}

	id := uint(1)
	adminInfo := map[string]uint{
		"id": 1,
	}

	// Create the expected result
	expectedApproval := entity.RegisterApproval{
		SuperAdminID: adminInfo["id"],
		Status:       reg.Status,
	}

	// Set up the mock repository expectation
	mockRepo.On("UpdateRegisterApprovalStatusById", mock.AnythingOfType("*entity.RegisterApproval"), id).Return(&expectedApproval, nil)

	// Call the method on the use case
	approval, err := useCase.UpdateRegisterApprovalStatusById(reg, id, adminInfo)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedApproval, approval)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestSetActivateAdminById(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create the use case instance with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Create the input parameter
	id := uint(1)

	// Create the expected result
	expectedAdmin := entity.Actor{
		ID:       id,
		IsActive: "true",
	}

	// Set up the mock repository expectation
	mockRepo.On("SetActivateAdminById", id).Return(expectedAdmin, nil)

	// Call the method on the use case
	admin, err := useCase.SetActivateAdminById(id)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestSetDeactivateAdminById(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create the use case instance with the mock repository
	useCase := account.NewUseCaseAdmin(mockRepo)

	// Create the input parameter
	id := uint(1)

	// Create the expected result
	expectedAdmin := entity.Actor{
		ID:       id,
		IsActive: "false",
	}

	// Set up the mock repository expectation
	mockRepo.On("SetDeactivateAdminById", id).Return(expectedAdmin, nil)

	// Call the method on the use case
	admin, err := useCase.SetDeactivateAdminById(id)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
