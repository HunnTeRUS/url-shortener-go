package input

import (
	"github.com/HunnTeRUS/url-shortener-go/src/application/domain"
	"github.com/HunnTeRUS/url-shortener-go/src/configuration/rest_errors"
)

type AddLinkInterface interface {
	AddLinkService(
		*domain.URLDomain,
		chan<- *rest_errors.RestErr,
	)
}
