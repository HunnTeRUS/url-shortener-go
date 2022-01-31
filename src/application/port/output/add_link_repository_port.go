package output

import (
	"github.com/HunnTeRUS/url-shortener-go/src/application/domain"
	"github.com/HunnTeRUS/url-shortener-go/src/configuration/rest_errors"
)

type AddLinkInterfacePort interface {
	AddLinkRepository(originalURL string, shortenedURL string) (*domain.URLDomain, *rest_errors.RestErr)
	GetLinkRepositoryByOriginalURL(originalURL string) (*domain.URLDomain, *rest_errors.RestErr)
	ListAllLinksRepositoryByOriginalURL() (*[]domain.URLDomain, *rest_errors.RestErr)
}
