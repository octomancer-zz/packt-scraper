package books

import "gitlab.com/octomancer/packt-scraper/book"

type EntitlementBooks struct {
	Count int                     `json:"count"`
	Books []*book.EntitlementBook `json:"data"`
}
