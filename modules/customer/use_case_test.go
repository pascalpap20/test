package customer

import (
	"crud/entity"
	mocks "crud/mock"
	"crud/repository"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

func TestNewUseCaseCustomer(t *testing.T) {
	type args struct {
		customerRepo repository.CustomerInterfaceRepo
	}

	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := NewUseCaseCustomer(mockRepo)

	tests := []struct {
		name string
		args args
		want UseCaseCustomer
	}{
		{
			name: "Test NewUseCaseAdmin",
			args: args{
				customerRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			want: useCase, // Define the expected UseCaseAdmin instance here,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUseCaseCustomer(tt.args.customerRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUseCaseCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseCustomer_CreateCustomer(t *testing.T) {
	type fields struct {
		customerRepo repository.CustomerInterfaceRepo
	}
	type args struct {
		customer CustomerParam
	}

	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Prepare the input data
	param := CustomerParam{
		FirstName: "Peter",
		LastName:  "Parker",
		Email:     "peter.parker@gmail.com",
		Avatar:    "https://upload.wikimedia.org/wikipedia/en/0/0f/Tom_Holland_as_Spider-Man.jpg",
	}

	expectedCustomer := entity.Customer{
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Email:     param.Email,
		Avatar:    param.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set up expectations
	mockRepo.On("CreateCustomer", mock.AnythingOfType("*entity.Customer")).Return(&expectedCustomer, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Customer
		wantErr bool
	}{
		{
			name: "Test CreateCustomer",
			fields: fields{
				customerRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				customer: param, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedCustomer, // Define the expected entity.Actor instance here,
			wantErr: false,            // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: tt.fields.customerRepo,
			}
			got, err := uc.CreateCustomer(tt.args.customer)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}
			//assert.Equal(t, tt.want, got, "CreateAdmin() got = %v, want %v", got, tt.want)

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want.FirstName, got.FirstName)
			assert.Equal(t, tt.want.LastName, got.LastName)
			assert.Equal(t, tt.want.Email, got.Email)
			assert.Equal(t, tt.want.Avatar, got.Avatar)
			assert.WithinDuration(t, tt.want.CreatedAt, got.CreatedAt, time.Second)
			assert.WithinDuration(t, tt.want.UpdatedAt, got.UpdatedAt, time.Second)
		})
	}

	// Assert that all expected method calls were made
	mockRepo.AssertExpectations(t)
}

func Test_useCaseCustomer_DeleteCustomerById(t *testing.T) {
	type fields struct {
		customerRepo repository.CustomerInterfaceRepo
	}
	type args struct {
		id uint
	}

	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Prepare the input data
	id := uint(1)

	// Prepare the expected output
	expectedCustomer := entity.Customer{
		ID: 1,
	}

	// Set up expectations
	mockRepo.On("DeleteCustomerById", id).Return(expectedCustomer, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Customer
		wantErr bool
	}{
		{
			name: "Test DeleteCustomerById",
			fields: fields{
				customerRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				id: id, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedCustomer, // Define the expected entity.Actor instance here,
			wantErr: false,            // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: tt.fields.customerRepo,
			}
			got, err := uc.DeleteCustomerById(tt.args.id)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("DeleteCustomerById() error = %v, wantErr %v", err, tt.wantErr)
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

func Test_useCaseCustomer_GetCustomerById(t *testing.T) {
	type fields struct {
		customerRepo repository.CustomerInterfaceRepo
	}
	type args struct {
		id uint
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Prepare the input data
	customerID := uint(1)

	// Prepare the expected output
	expectedCustomer := entity.Customer{
		ID:        customerID,
		FirstName: "Peter",
		LastName:  "Parker",
		Email:     "peter.parker@gmail.com",
		Avatar:    "https://upload.wikimedia.org/wikipedia/en/0/0f/Tom_Holland_as_Spider-Man.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set up expectations
	mockRepo.On("GetCustomerById", customerID).Return(expectedCustomer, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Customer
		wantErr bool
	}{
		{
			name: "Test GetCustomerById",
			fields: fields{
				customerRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				id: customerID, // Provide the necessary AdminParam values for the test case,
			},
			want:    expectedCustomer, // Define the expected entity.Actor instance here,
			wantErr: false,            // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: tt.fields.customerRepo,
			}
			got, err := uc.GetCustomerById(tt.args.id)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("GetCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.FirstName, got.FirstName)
			assert.Equal(t, tt.want.LastName, got.LastName)
			assert.Equal(t, tt.want.Email, got.Email)
			assert.Equal(t, tt.want.Avatar, got.Avatar)
			assert.WithinDuration(t, tt.want.CreatedAt, got.CreatedAt, time.Second)
			assert.WithinDuration(t, tt.want.UpdatedAt, got.UpdatedAt, time.Second)
		})
	}
	// Assert the expectations were met
	mockRepo.AssertExpectations(t)
}

func Test_useCaseCustomer_GetCustomers(t *testing.T) {
	type fields struct {
		customerRepo repository.CustomerInterfaceRepo
	}
	type args struct {
		name  string
		email string
		page  uint
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Prepare the input data
	name := ""
	email := ""
	page := uint(1)

	// Prepare the expected output
	expectedCustomers := []entity.Customer{
		{
			ID:        1,
			FirstName: "peter",
			LastName:  "parker",
			Email:     "peter.parker@gmail.com",
			Avatar:    "https://upload.wikimedia.org/wikipedia/en/0/0f/Tom_Holland_as_Spider-Man.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			FirstName: "Miles",
			LastName:  "Morales",
			Email:     "miles.morales@gmail.com",
			Avatar:    "https://upload.wikimedia.org/wikipedia/en/0/0f/Tom_Holland_as_Spider-Man.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Set up expectations
	mockRepo.On("GetCustomers", name, email, page).Return(expectedCustomers, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Customer
		wantErr bool
	}{
		{
			name: "Test GetCustomers",
			fields: fields{
				customerRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				name:  name, // Provide the necessary AdminParam values for the test case,
				email: email,
				page:  page,
			},
			want:    expectedCustomers, // Define the expected entity.Actor instance here,
			wantErr: false,             // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: tt.fields.customerRepo,
			}
			got, err := uc.GetCustomers(tt.args.name, tt.args.email, tt.args.page)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("GetCustomers() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, len(tt.want), len(got))
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i].ID, got[i].ID)
				assert.Equal(t, tt.want[i].FirstName, got[i].FirstName)
				assert.Equal(t, tt.want[i].LastName, got[i].LastName)
				assert.Equal(t, tt.want[i].Email, got[i].Email)
				assert.Equal(t, tt.want[i].Avatar, got[i].Avatar)
				assert.WithinDuration(t, tt.want[i].CreatedAt, got[i].CreatedAt, time.Second)
				assert.WithinDuration(t, tt.want[i].UpdatedAt, got[i].UpdatedAt, time.Second)
			}
		})
	}
	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

}

func Test_useCaseCustomer_UpdateCustomerById(t *testing.T) {
	type fields struct {
		customerRepo repository.CustomerInterfaceRepo
	}
	type args struct {
		customer UpdateCustomer
		id       uint
	}
	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Prepare the input data
	param := UpdateCustomer{
		FirstName: "Miles",
		LastName:  "Morales",
		Email:     "miles.morales@gmail.com",
		Avatar:    "https://upload.wikimedia.org/wikipedia/en/0/0f/Tom_Holland_as_Spider-Man.jpg",
	}

	customerId := uint(1)

	expectedCustomer := entity.Customer{
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Email:     param.Email,
		Avatar:    param.Avatar,
	}

	// Set up expectations
	mockRepo.On("UpdateCustomerById", mock.AnythingOfType("*entity.Customer"), customerId).Return(&expectedCustomer, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Customer
		wantErr bool
	}{
		{
			name: "Test GetCustomers",
			fields: fields{
				customerRepo: mockRepo, // Provide a mock or stub for the repository.ActorInterfaceRepo interface,
			},
			args: args{
				id:       customerId, // Provide the necessary AdminParam values for the test case,
				customer: param,
			},
			want:    expectedCustomer, // Define the expected entity.Actor instance here,
			wantErr: false,            // Define whether the test case expects an error or not,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: tt.fields.customerRepo,
			}
			got, err := uc.UpdateCustomerById(tt.args.customer, tt.args.id)
			if (err != nil) != tt.wantErr {
				errMsg := fmt.Errorf("UpdateCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				assert.Fail(t, errMsg.Error())
				return
			}

			// Assert the result using testify assertions
			assert.NoError(t, err)
			assert.Equal(t, tt.want.FirstName, got.FirstName)
			assert.Equal(t, tt.want.LastName, got.LastName)
			assert.Equal(t, tt.want.Email, got.Email)
			assert.Equal(t, tt.want.Avatar, got.Avatar)
			assert.WithinDuration(t, tt.want.UpdatedAt, got.UpdatedAt, time.Second)
		})
	}
	// Assert the expectations were met
	mockRepo.AssertExpectations(t)
}
