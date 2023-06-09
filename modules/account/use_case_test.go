package account

import (
	"crud/entity"
	mocks "crud/mock"
	"crud/repository"
	"crud/utils/auth"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"testing"
	"time"
)

func TestNewUseCaseAdmin(t *testing.T) {
	type args struct {
		adminRepo repository.ActorInterfaceRepo
	}

	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := NewUseCaseAdmin(mockRepo)

	tests := []struct {
		name string
		args args
		want UseCaseAdmin
	}{
		// TODO: Add test cases.
		{
			name: "Test NewUseCaseAdmin",
			args: args{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			want: useCase, // Define the expected UseCaseAdmin instance here,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUseCaseAdmin(tt.args.adminRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUseCaseAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseAdmin_CreateAdmin(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		admin AdminParam
	}

	mockRepo := &mocks.ActorInterfaceRepo{}

	// Prepare the input data
	admin := AdminParam{
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

	// Set the expectation for the CreateActor method
	mockRepo.On("CreateActor", mock.AnythingOfType("*entity.Actor")).Return(&expectedAdmin, nil)
	isSame := bcrypt.CompareHashAndPassword([]byte(password), []byte(admin.Password+salt))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test CreateAdmin",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				admin: admin, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedAdmin, // Define the expected entity.Actor instance here,
			wantErr: false,         // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.CreateAdmin(tt.args.admin)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}
			//assert.Equal(t, tt.want, got, "CreateAdmin() got = %v, want %v", got, tt.want)

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Username, got.Username)
			assert.Equal(t, tt.want.RoleID, got.RoleID)
			assert.Equal(t, tt.want.IsVerified, got.IsVerified)
			assert.Equal(t, tt.want.IsActive, got.IsActive)
			assert.WithinDuration(t, tt.want.CreatedAt, got.CreatedAt, time.Second)
			assert.WithinDuration(t, tt.want.UpdatedAt, got.UpdatedAt, time.Second)
			// Add password hashing comparison
			assert.True(t, isSame == nil)
		})
	}

	// Assert that all expected method calls were made
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_DeleteAdminById(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		id uint
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Prepare the input data
	id := uint(1)

	// Prepare the expected output
	expectedAdmin := entity.Actor{
		ID: 1,
	}

	// Set up expectations
	mockRepo.On("DeleteActorById", id).Return(expectedAdmin, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test DeleteActorById",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				id: id, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedAdmin, // Define the expected entity.Actor instance here,
			wantErr: false,         // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.DeleteAdminById(tt.args.id)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("DeleteAdminById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_GetAdminById(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		id uint
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

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

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test GetActorById",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				id: adminID, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedAdmin, // Define the expected entity.Actor instance here,
			wantErr: false,         // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.GetAdminById(tt.args.id)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("GetActorById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.Username, got.Username)
			assert.Equal(t, tt.want.RoleID, got.RoleID)
			assert.WithinDuration(t, tt.want.CreatedAt, got.CreatedAt, time.Second)
			assert.WithinDuration(t, tt.want.UpdatedAt, got.UpdatedAt, time.Second)
		})
	}

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_GetAdmins(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		username string
		page     uint
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

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

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test GetActors",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				username: username, // Provide the necessary AdminParam values for the test case,
				page:     page,
			},
			want:    expectedAdmins, // Define the expected entity.Actor instance here,
			wantErr: false,          // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.GetAdmins(tt.args.username, tt.args.page)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("GetActors() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, len(tt.want), len(got))
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i].ID, got[i].ID)
				assert.Equal(t, tt.want[i].Username, got[i].Username)
				assert.Equal(t, tt.want[i].RoleID, got[i].RoleID)
				assert.WithinDuration(t, tt.want[i].CreatedAt, got[i].CreatedAt, time.Second)
				assert.WithinDuration(t, tt.want[i].UpdatedAt, got[i].UpdatedAt, time.Second)
			}
		})
	}

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_GetRegisterApproval(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	// Create a mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create the expected result
	expectedApproval := []entity.RegisterApproval{
		{ID: 1, AdminID: 2, SuperAdminID: 1, Status: "pending"},
		{ID: 2, AdminID: 3, SuperAdminID: 1, Status: "approved"},
		{ID: 3, AdminID: 4, SuperAdminID: 1, Status: "rejected"},
	}

	// Set up the mock repository expectation
	mockRepo.On("GetRegisterApproval").Return(expectedApproval, nil)

	tests := []struct {
		name    string
		fields  fields
		want    []entity.RegisterApproval
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test GetRegisterApproval",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			want:    expectedApproval, // Define the expected entity.Actor instance here,
			wantErr: false,            // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.GetRegisterApproval()
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("GetRegisterApproval() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_Login(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		admin LoginParam
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Prepare the input data
	admin := LoginParam{
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

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Login",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				admin: admin, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedActor, // Define the expected entity.Actor instance here,
			wantErr: false,         // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.Login(tt.args.admin)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_Register(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		admin RegisterParam
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create an instance of the use case with the mock repository

	// Prepare the input data
	admin := RegisterParam{
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
	isSame := bcrypt.CompareHashAndPassword([]byte(password), []byte(admin.Password+salt))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Register",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				admin: admin, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedActor, // Define the expected entity.Actor instance here,
			wantErr: false,         // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.Register(tt.args.admin)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Username, got.Username)
			assert.Equal(t, tt.want.RoleID, got.RoleID)
			assert.Equal(t, tt.want.IsVerified, got.IsVerified)
			assert.Equal(t, tt.want.IsActive, got.IsActive)
			assert.WithinDuration(t, tt.want.UpdatedAt, got.UpdatedAt, time.Second)
			assert.True(t, isSame == nil)
		})
	}

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_SetActivateAdminById(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		id uint
	}

	// Create a mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create the input parameter
	id := uint(1)

	// Create the expected result
	expectedAdmin := entity.Actor{
		ID:       id,
		IsActive: "true",
	}

	// Set up the mock repository expectation
	mockRepo.On("SetActivateAdminById", id).Return(expectedAdmin, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test SetActivateAdminById",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				id: id, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedAdmin, // Define the expected entity.Actor instance here,
			wantErr: false,         // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.SetActivateAdminById(tt.args.id)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("SetActivateAdminById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_SetDeactivateAdminById(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		id uint
	}
	// Create a mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create the input parameter
	id := uint(1)

	// Create the expected result
	expectedAdmin := entity.Actor{
		ID:       id,
		IsActive: "false",
	}

	// Set up the mock repository expectation
	mockRepo.On("SetDeactivateAdminById", id).Return(expectedAdmin, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test SetDeactivateAdminById",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				id: id, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedAdmin, // Define the expected entity.Actor instance here,
			wantErr: false,         // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.SetDeactivateAdminById(tt.args.id)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("SetDeactivateAdminById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_UpdateAdminById(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		admin UpdateAdmin
		id    uint
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Prepare the input data
	admin := UpdateAdmin{
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
	isSame := bcrypt.CompareHashAndPassword([]byte(password), []byte(admin.Password+salt))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test UpdateActorById",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				admin: admin, // Provide the necessary AdminParam values for the test case,
				id:    adminId,
			},
			want:    expectedAdmin, // Define the expected entity.Actor instance here,
			wantErr: false,         // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.UpdateAdminById(tt.args.admin, tt.args.id)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("UpdateActorById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Username, got.Username)
			assert.Equal(t, tt.want.RoleID, got.RoleID)
			assert.Equal(t, tt.want.IsVerified, got.IsVerified)
			assert.Equal(t, tt.want.IsActive, got.IsActive)
			assert.WithinDuration(t, tt.want.UpdatedAt, got.UpdatedAt, time.Second)
			// Add password hashing comparison
			assert.True(t, isSame == nil)
		})
	}
	// Assert the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseAdmin_UpdateRegisterApprovalStatusById(t *testing.T) {
	type fields struct {
		adminRepo repository.ActorInterfaceRepo
	}
	type args struct {
		reg       UpdateRegisterApproval
		id        uint
		adminInfo map[string]uint
	}
	// Create a mock repository
	mockRepo := &mocks.ActorInterfaceRepo{}

	// Create the input parameters
	reg := UpdateRegisterApproval{
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

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.RegisterApproval
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test UpdateRegisterApprovalStatusById",
			fields: fields{
				adminRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				reg:       reg, // Provide the necessary AdminParam values for the test case,
				id:        id,
				adminInfo: adminInfo,
			},
			want:    expectedApproval, // Define the expected entity.Actor instance here,
			wantErr: false,            // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.UpdateRegisterApprovalStatusById(tt.args.reg, tt.args.id, tt.args.adminInfo)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("UpdateActorById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
