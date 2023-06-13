package service

import "github.com/begenov/test-task/internal/domain"

type Availability interface {
	CheckAvailability()
	GetSiteWithMaxAvailability() (string, error)
	GetSite(url string) (*domain.SiteAvailability, error)
	GetSiteWithMinAvailability() (string, error)
}

type Service struct {
	Availability Availability
}

func NewService(website []string) *Service {
	return &Service{
		Availability: NewAvailabilityService(website),
	}
}
