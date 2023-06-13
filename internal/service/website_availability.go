package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/begenov/test-task/internal/domain"
)

type AvailabilityService struct {
	webSite      []string
	availability map[string]*domain.SiteAvailability
	sync.RWMutex
}

func NewAvailabilityService(website []string) *AvailabilityService {
	return &AvailabilityService{
		webSite:      website,
		availability: make(map[string]*domain.SiteAvailability),
	}
}

func (s *AvailabilityService) CheckAvailability() {

	for {
		for _, site := range s.webSite {

			go func(site string) {
				_, err := http.Get("http://" + site)
				s.Lock()
				if err == nil {
					s.availability[site] = &domain.SiteAvailability{URL: site, LastTime: time.Now()}
				} else {
					delete(s.availability, site)
				}
				s.Unlock()
			}(site)

		}
		time.Sleep(1 * time.Minute)
	}
}

func (s *AvailabilityService) GetSite(url string) (*domain.SiteAvailability, error) {
	s.RLock()
	defer s.RUnlock()

	site, ok := s.availability[url]
	if !ok {
		return nil, fmt.Errorf("caйт не найден: %s", url)
	}

	return site, nil
}

func (s *AvailabilityService) GetSiteWithMinAvailability() (string, error) {
	s.RLock()
	defer s.RUnlock()

	var minSite string
	minTime := time.Now().AddDate(0, 0, 1)

	for _, site := range s.availability {
		if site.LastTime.Before(minTime) {
			minSite = site.URL
			minTime = site.LastTime
		}
	}

	if minSite == "" {
		return "", fmt.Errorf("сайт с доступностью не найден")
	}

	return minSite, nil
}

func (s *AvailabilityService) GetSiteWithMaxAvailability() (string, error) {
	s.RLock()
	defer s.RUnlock()

	var maxSite string
	maxTime := time.Time{}

	for _, site := range s.availability {
		if site.LastTime.After(maxTime) {
			maxSite = site.URL
			maxTime = site.LastTime
		}
	}

	if maxSite == "" {
		return "", fmt.Errorf("сайт с доступностью не найден")
	}

	return maxSite, nil
}
