package domain

type URLDomain struct {
	OriginalURL  string `json:"original_url"`
	ShortenedURL string `json:"shortened_url"`
}
