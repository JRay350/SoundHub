package search

import (
	"soundhub/internal/models"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Search(query string) ([]models.SearchResult, error) {
	// placeholder for now
	return []models.SearchResult{
		{
			Source: "mock",
			Title:  query,
			Artist: "test artist",
		},
	}, nil
}