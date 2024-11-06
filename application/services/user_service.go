package services

import (
	"fmt"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/requests"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/responses"
	"github.com/MasDev-12/mechta.testapi/application/helpers"
	"github.com/MasDev-12/mechta.testapi/domain/entities"
	"github.com/MasDev-12/mechta.testapi/infrastructure/repositories"
	"github.com/google/uuid"
	"strings"
	"time"
)

type UserService struct {
	UserRepository *repositories.UserRepository
	Argon2         *helpers.Argon2Helper
}

func NewUserService(userRepository *repositories.UserRepository,
	argon2 *helpers.Argon2Helper) *UserService {
	return &UserService{
		UserRepository: userRepository,
		Argon2:         argon2,
	}
}

func (service *UserService) Create(request requests.CreateUserRequest) responses.CreateUserResponse {
	responseChan := make(chan responses.CreateUserResponse)

	go func() {
		passwordHash, err := service.Argon2.GetPasswordHash(request.Password)
		user := entities.User{
			Id:           uuid.New(),
			Username:     request.Username,
			Email:        strings.TrimSpace(strings.ToLower(request.Email)),
			PasswordHash: passwordHash,
			IsActive:     true,
			CreatedAt:    time.Now(),
		}
		response, err := service.UserRepository.Add(user)

		if err != nil {
			responseChan <- responses.CreateUserResponse{
				Result: response,
				Error:  err,
			}
			return
		}
		responseChan <- responses.CreateUserResponse{
			Result: response,
			Error:  nil,
		}
		return
	}()
	select {
	case response := <-responseChan:
		return response
	case <-time.After(time.Second * 10):
		return responses.CreateUserResponse{
			Result: false,
			Error:  fmt.Errorf("timeout: could not add user in time"),
		}
	}
}

func (service *UserService) GetById(request requests.GetUserRequest) responses.GetUserResponse {
	responseChan := make(chan responses.GetUserResponse)

	go func() {
		response, err := service.UserRepository.GetById(request.Id)
		if err != nil {
			responseChan <- responses.GetUserResponse{
				Id:       nil,
				Username: nil,
				Email:    nil,
				Error:    fmt.Errorf("failed to get account: %w", err),
			}
			return
		}
		responseChan <- responses.GetUserResponse{
			Id:       &response.Id,
			Username: &response.Username,
			Email:    &response.Email,
			Error:    nil,
		}
	}()
	select {
	case response := <-responseChan:
		return response
	case <-time.After(time.Second * 10):
		return responses.GetUserResponse{
			Id:       nil,
			Username: nil,
			Email:    nil,
			Error:    fmt.Errorf("timeout: could not update user in time"),
		}
	}
}
