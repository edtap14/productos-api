package models

import "time"

/*ProductSearchResponse estructura de respuesta*/
type ProductSearchResponse struct {
	RequestInfo struct {
		Success          bool `json:"success"`
		CreditsUsed      int  `json:"credits_used"`
		CreditsRemaining int  `json:"credits_remaining"`
	} `json:"request_info"`
	RequestMetadata struct {
		CreatedAt      time.Time `json:"created_at"`
		ProcessedAt    time.Time `json:"processed_at"`
		TotalTimeTaken float64   `json:"total_time_taken"`
		AmazonURL      string    `json:"amazon_url"`
		Timing         []string  `json:"timing"`
	} `json:"request_metadata"`
	RequestParameters struct {
		Type         string `json:"type"`
		AmazonDomain string `json:"amazon_domain"`
		SearchTerm   string `json:"search_term"`
	} `json:"request_parameters"`
	SearchResults []struct {
		Position   int    `json:"position"`
		Title      string `json:"title"`
		Asin       string `json:"asin"`
		Link       string `json:"link"`
		Categories []struct {
			Name string `json:"name"`
		} `json:"categories"`
		Image        string  `json:"image"`
		IsPrime      bool    `json:"is_prime"`
		Rating       float64 `json:"rating,omitempty"`
		RatingsTotal int     `json:"ratings_total,omitempty"`
		Sponsored    bool    `json:"sponsored"`
		Prices       []struct {
			Symbol   string  `json:"symbol"`
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
			Raw      string  `json:"raw"`
			Name     string  `json:"name"`
		} `json:"prices,omitempty"`
	} `json:"search_results"`
	Pagination struct {
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
	} `json:"pagination"`
}
