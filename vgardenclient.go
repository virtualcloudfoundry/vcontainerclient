package vcontainerclient

import (
	"code.cloudfoundry.org/lager"
	"github.com/virtualcloudfoundry/vcontainercommon/vcontainermodels"
)

func NewVGardenClientSkipCertVerify(logger lager.Logger, config vcontainermodels.VContainerClientConfig) (vcontainermodels.VGardenClient, error) {
	return newVGardenClientInternal(logger, config, true)
}

func NewVGardenClient(logger lager.Logger, config vcontainermodels.VContainerClientConfig) (vcontainermodels.VGardenClient, error) {
	return newVGardenClientInternal(logger, config, false)
}

func newVGardenClientInternal(logger lager.Logger, config vcontainermodels.VContainerClientConfig, skipCertVerify bool) (vcontainermodels.VGardenClient, error) {
	conn, err := NewConn(logger, config)
	if err != nil {
		return nil, err
	}
	return vcontainermodels.NewVGardenClient(conn), nil
}
