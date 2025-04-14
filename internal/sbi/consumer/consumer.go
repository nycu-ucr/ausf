package consumer

import (
	"github.com/nycu-ucr/ausf/pkg/app"
	Nnrf_NFDiscovery "github.com/nycu-ucr/openapi/nrf/NFDiscovery"
	Nnrf_NFManagement "github.com/nycu-ucr/openapi/nrf/NFManagement"
	Nudm_UEAuthentication "github.com/nycu-ucr/openapi/udm/UEAuthentication"
)

type ConsumerAusf interface {
	app.App
}

type Consumer struct {
	ConsumerAusf

	*nnrfService
	*nudmService
}

func NewConsumer(ausf ConsumerAusf) (*Consumer, error) {
	c := &Consumer{
		ConsumerAusf: ausf,
	}

	c.nnrfService = &nnrfService{
		consumer:        c,
		nfMngmntClients: make(map[string]*Nnrf_NFManagement.APIClient),
		nfDiscClients:   make(map[string]*Nnrf_NFDiscovery.APIClient),
	}

	c.nudmService = &nudmService{
		consumer:    c,
		ueauClients: make(map[string]*Nudm_UEAuthentication.APIClient),
	}

	return c, nil
}
