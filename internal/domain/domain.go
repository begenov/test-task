package domain

import "time"

type SiteAvailability struct {
	URL      string
	LastTime time.Time
}
