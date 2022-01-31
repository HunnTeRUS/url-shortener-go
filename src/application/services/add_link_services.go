package services

import (
	"github.com/HunnTeRUS/url-shortener-go/src/application/domain"
	"github.com/HunnTeRUS/url-shortener-go/src/application/port/input"
	"github.com/HunnTeRUS/url-shortener-go/src/application/port/output"
	"github.com/HunnTeRUS/url-shortener-go/src/configuration/rest_errors"
	"github.com/HunnTeRUS/url-shortener-go/src/utils"
)

type addLink struct {
	repository output.AddLinkInterfacePort
}

// NewAddLinkInterface will receive a connection to the repository
// and returns a AddLinkInterface instance that can be used to call
// the url add methods
func NewAddLinkInterface(
	repository output.AddLinkInterfacePort,
) input.AddLinkInterface {
	return &addLink{
		repository: repository,
	}
}

// AddLinkService will save e/or return the shortened URL over the received
// url, always that already exists on database will simply return that and if
// it does not exists, will generate a new one, save on database and then, return
func (al *addLink) AddLinkService(
	url_domain *domain.URLDomain,
	chanErr chan<- *rest_errors.RestErr,
) {
	// channelRepoError is used do handle errors inside repository when trying to
	// verify if already exists some URL
	channelRepoError := make(chan *rest_errors.RestErr, 1)

	// channelGenerateURLError is used to handle errors inside generateURL method
	channelGenerateURLError := make(chan *rest_errors.RestErr, 1)

	// channelURL is used to wait and holds the value of generated/found URL
	channelURL := make(chan *domain.URLDomain, 1)

	// channelURLAux is used to communicate both goroutines running (because if
	// URL does not exists on database is not consider an error)
	channelURLAux := make(chan *domain.URLDomain, 1)

	// closeChannels will be called when all is done
	closeChannels := func() {
		close(channelGenerateURLError)
		close(channelGenerateURLError)
		close(chanErr)
	}

	// The first goroutine to be executed will verify if already exists in the
	// database the URL that was received
	go func() {
		domain, err := al.repository.GetLinkRepositoryByOriginalURL(
			url_domain.OriginalURL,
		)

		if err != nil {
			if err.Code == 404 {
				close(channelURLAux)
				return
			}

			channelRepoError <- err
			return
		}

		channelURL <- domain
	}()

	// The seconde goroutine is going to generate the URL while the last one is
	// verifing in the database, if exists, there is an atribuition to channelURL
	// already made in the first goroutine, if not, will save the generated URL in
	// the database and return
	go func() {
		generated_URL := utils.GenerateURL()

		<-channelURLAux

		inserted_domain, err := al.repository.AddLinkRepository(
			url_domain.OriginalURL,
			generated_URL,
		)

		if err != nil {
			channelGenerateURLError <- err
			return
		}

		channelURL <- inserted_domain
	}()

	// select will wait for some errors ou for some value into the domain value channel
	select {
	case err := <-channelGenerateURLError:
		chanErr <- err
		closeChannels()
	case domain := <-channelURL:
		url_domain = domain
		closeChannels()
	}
}
