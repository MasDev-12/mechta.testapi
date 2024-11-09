package services

import (
	"fmt"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/requests"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/responses"
	"github.com/MasDev-12/mechta.testapi/application/helpers"
	"github.com/MasDev-12/mechta.testapi/domain/dto"
	"github.com/MasDev-12/mechta.testapi/domain/entities"
	"github.com/MasDev-12/mechta.testapi/infrastructure/repositories"
	"github.com/google/uuid"
	"sync"
	"time"
)

type URLService struct {
	URLRepository *repositories.URLRepository
}

func NewURLService(urlRepository *repositories.URLRepository) *URLService {
	return &URLService{URLRepository: urlRepository}
}

func (service *URLService) Create(request requests.CreateURLRequest) responses.CreateUrlResponse {
	responseChan := make(chan responses.CreateUrlResponse)

	go func() {
		url := entities.URL{
			Id:             uuid.New(),
			OriginalURL:    request.OriginalURL,
			ShortURL:       helpers.GenerateShortURL(request.OriginalURL),
			UserId:         request.UserId,
			CreatedAt:      time.Now(),
			IsActive:       true,
			ExpiresAt:      time.Now().AddDate(0, 0, 30),
			ClickCount:     0,
			LastAccessedAt: nil,
		}
		response, err := service.URLRepository.Add(url)
		if err != nil {
			responseChan <- responses.CreateUrlResponse{
				Id:       nil,
				ShortURL: nil,
				Error:    err,
			}
			return
		}
		responseChan <- responses.CreateUrlResponse{
			Id:       &response.Id,
			ShortURL: &response.ShortURL,
			Error:    nil,
		}
		return
	}()
	select {
	case response := <-responseChan:
		return response
	case <-time.After(time.Second * 10):
		return responses.CreateUrlResponse{
			Id:       nil,
			ShortURL: nil,
			Error:    fmt.Errorf("timeout: could not add user in time"),
		}
	}
}

func (service *URLService) GetUserUrls(request requests.GetUserUrlsRequest) responses.GetUserUrlsResponse {
	responseChan := make(chan responses.GetUserUrlsResponse)
	go func() {
		response, err := service.URLRepository.GetUserUrls(request.UserId)

		if err != nil {
			responseChan <- responses.GetUserUrlsResponse{
				Urls:  nil,
				Error: fmt.Errorf("failed to get urls: %w", err),
			}
			return
		}
		responseChan <- responses.GetUserUrlsResponse{
			Urls:  convertToDTO(response),
			Error: nil,
		}
		return
	}()
	select {
	case response := <-responseChan:
		return response
	case <-time.After(time.Second * 10):
		return responses.GetUserUrlsResponse{
			Urls:  nil,
			Error: fmt.Errorf("timeout: could not fetch accounts in time"),
		}
	}
}

func convertToDTO(urls []entities.URL) []dto.URLDto {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	resultChan := make(chan dto.URLDto)
	for _, url := range urls {
		go func(url entities.URL) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					// Обработка паники, например, логирование
					fmt.Println("Recovered in goroutine:", r)
				}
			}()
			resultChan <- dto.URLDto{
				Id:             url.Id,
				OriginalURL:    url.OriginalURL,
				ShortURL:       url.ShortURL,
				UserId:         url.UserId,
				IsActive:       url.IsActive,
				ExpiresAt:      url.ExpiresAt,
				ClickCount:     url.ClickCount,
				LastAccessedAt: url.LastAccessedAt,
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var response []dto.URLDto
	for result := range resultChan {
		response = append(response, result)
	}
	return response
}

func (service *URLService) GetUrlByShortName(request requests.GetUrlByShortNameRequest) responses.GetUrlByShortNameResponse {
	responseChan := make(chan responses.GetUrlByShortNameResponse)

	go func() {
		response, err := service.URLRepository.GetUrlByShortName(request.ShortName)

		if err != nil {
			responseChan <- responses.GetUrlByShortNameResponse{
				Url:   nil,
				Error: fmt.Errorf("failed to get account: %w", err),
			}
			return
		}
		responseChan <- responses.GetUrlByShortNameResponse{
			Url: &dto.URLDto{
				Id:             response.Id,
				OriginalURL:    response.OriginalURL,
				ShortURL:       response.ShortURL,
				UserId:         response.UserId,
				IsActive:       response.IsActive,
				ExpiresAt:      response.ExpiresAt,
				ClickCount:     response.ClickCount,
				LastAccessedAt: response.LastAccessedAt,
			},
			Error: nil,
		}
	}()

	select {
	case response := <-responseChan:
		return response
	case <-time.After(time.Second * 10):
		return responses.GetUrlByShortNameResponse{
			Url:   nil,
			Error: fmt.Errorf("timeout: could not update user in time"),
		}
	}
}

func (service *URLService) DeleteUrlByShortName(request requests.DeleteByShortNameRequest) responses.DeleteUrlByShortNameResponse {
	responseChan := make(chan responses.DeleteUrlByShortNameResponse)

	go func() {
		response, err := service.URLRepository.DeleteUrlByShortName(request.ShortName)
		if err != nil {
			responseChan <- responses.DeleteUrlByShortNameResponse{
				Result: response,
				Error:  err,
			}
			return
		}
		responseChan <- responses.DeleteUrlByShortNameResponse{
			Result: response,
			Error:  nil,
		}
		return
	}()
	select {
	case response := <-responseChan:
		return response
	case <-time.After(time.Second * 10):
		return responses.DeleteUrlByShortNameResponse{
			Result: false,
			Error:  fmt.Errorf("timeout: could not update user in time"),
		}
	}
}

func (service *URLService) GetUrlStatByShortName(request requests.GetUrlStatByShortNameRequest) responses.GetUrlStatByShortNameResponse {
	responseChan := make(chan responses.GetUrlStatByShortNameResponse)

	go func() {
		response, err := service.URLRepository.GetUrlByShortName(request.ShortName)

		if err != nil {
			responseChan <- responses.GetUrlStatByShortNameResponse{
				OriginalURL:    nil,
				IsActive:       nil,
				ExpiresAt:      nil,
				ClickCount:     nil,
				LastAccessedAt: nil,
				Error:          fmt.Errorf("failed to get account: %w", err),
			}
			return
		}
		responseChan <- responses.GetUrlStatByShortNameResponse{
			OriginalURL:    &response.OriginalURL,
			IsActive:       &response.IsActive,
			ExpiresAt:      &response.ExpiresAt,
			ClickCount:     &response.ClickCount,
			LastAccessedAt: response.LastAccessedAt,
			Error:          nil,
		}
		return
	}()

	select {
	case response := <-responseChan:
		return response
	case <-time.After(time.Second * 10):
		return responses.GetUrlStatByShortNameResponse{
			OriginalURL:    nil,
			IsActive:       nil,
			ExpiresAt:      nil,
			ClickCount:     nil,
			LastAccessedAt: nil,
			Error:          fmt.Errorf("timeout: could not update user in time"),
		}
	}
}

func (service *URLService) GetUrlByOriginalName(request requests.GetUrlByOriginalNameRequest) responses.GetUrlByOriginalNameResponse {
	responseChan := make(chan responses.GetUrlByOriginalNameResponse)

	go func() {
		response, err := service.URLRepository.GetUrlByOriginalName(request.OriginalName)

		if err != nil {
			responseChan <- responses.GetUrlByOriginalNameResponse{
				OriginalURL: nil,
				Error:       fmt.Errorf("failed to get account: %w", err),
			}
			return
		}
		responseChan <- responses.GetUrlByOriginalNameResponse{
			OriginalURL: &response.OriginalURL,
			Error:       nil,
		}
		return
	}()

	select {
	case response := <-responseChan:
		return response
	case <-time.After(time.Second * 10):
		return responses.GetUrlByOriginalNameResponse{
			OriginalURL: nil,
			Error:       fmt.Errorf("timeout: could not update user in time"),
		}
	}
}
