package unit_test

import (
	mocks "crud/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"crud/entity" // Import the entity package
	"crud/modules/customer"
)

func TestCreateCustomer(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := customer.NewUseCaseCustomer(mockRepo)

	// Prepare the input data
	param := customer.CustomerParam{
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

	// Call the method under test
	result, err := useCase.CreateCustomer(param)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer.FirstName, result.FirstName)
	assert.Equal(t, expectedCustomer.LastName, result.LastName)
	assert.Equal(t, expectedCustomer.Email, result.Email)
	assert.Equal(t, expectedCustomer.Avatar, result.Avatar)
	assert.WithinDuration(t, expectedCustomer.CreatedAt, result.CreatedAt, time.Second)
	assert.WithinDuration(t, expectedCustomer.UpdatedAt, result.UpdatedAt, time.Second)
}

func TestGetCustomerById(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := customer.NewUseCaseCustomer(mockRepo)

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

	// Call the method under test
	admin, err := useCase.GetCustomerById(customerID)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer.ID, admin.ID)
	assert.Equal(t, expectedCustomer.FirstName, admin.FirstName)
	assert.Equal(t, expectedCustomer.LastName, admin.LastName)
	assert.Equal(t, expectedCustomer.Email, admin.Email)
	assert.Equal(t, expectedCustomer.Avatar, admin.Avatar)
	assert.WithinDuration(t, expectedCustomer.CreatedAt, admin.CreatedAt, time.Second)
	assert.WithinDuration(t, expectedCustomer.UpdatedAt, admin.UpdatedAt, time.Second)
}

func TestGetCustomers(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := customer.NewUseCaseCustomer(mockRepo)

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

	// Call the method under test
	admins, err := useCase.GetCustomers(name, email, page)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, len(expectedCustomers), len(admins))
	for i := 0; i < len(expectedCustomers); i++ {
		assert.Equal(t, expectedCustomers[i].ID, admins[i].ID)
		assert.Equal(t, expectedCustomers[i].FirstName, admins[i].FirstName)
		assert.Equal(t, expectedCustomers[i].LastName, admins[i].LastName)
		assert.Equal(t, expectedCustomers[i].Email, admins[i].Email)
		assert.Equal(t, expectedCustomers[i].Avatar, admins[i].Avatar)
		assert.WithinDuration(t, expectedCustomers[i].CreatedAt, admins[i].CreatedAt, time.Second)
		assert.WithinDuration(t, expectedCustomers[i].UpdatedAt, admins[i].UpdatedAt, time.Second)
	}
}

func TestCustomerAdminById(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := customer.NewUseCaseCustomer(mockRepo)

	// Prepare the input data
	param := customer.UpdateCustomer{
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

	// Call the method under test
	result, err := useCase.UpdateCustomerById(param, customerId)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer.FirstName, result.FirstName)
	assert.Equal(t, expectedCustomer.LastName, result.LastName)
	assert.Equal(t, expectedCustomer.Email, result.Email)
	assert.Equal(t, expectedCustomer.Avatar, result.Avatar)
	assert.WithinDuration(t, expectedCustomer.UpdatedAt, result.UpdatedAt, time.Second)
	// Add password hashing comparison
}

func TestDeleteCustomerById(t *testing.T) {
	// Create an instance of the mock repository
	mockRepo := &mocks.CustomerInterfaceRepo{}

	// Create an instance of the use case with the mock repository
	useCase := customer.NewUseCaseCustomer(mockRepo)

	// Prepare the input data
	id := uint(1)

	// Prepare the expected output
	expectedCustomer := entity.Customer{
		ID: 1,
	}

	// Set up expectations
	mockRepo.On("DeleteCustomerById", id).Return(expectedCustomer, nil)

	// Call the method under test
	deletedAdmin, err := useCase.DeleteCustomerById(id)

	// Assert the expectations were met
	mockRepo.AssertExpectations(t)

	// Assert the result using testify assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer.ID, deletedAdmin.ID)
}
