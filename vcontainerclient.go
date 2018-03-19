package vcontainerclient

import (
	"code.cloudfoundry.org/lager"
	"github.com/virtualcloudfoundry/vcontainercommon/vcontainermodels"
)

func NewVContainerClientSkipCertVerify(logger lager.Logger, config vcontainermodels.VContainerClientConfig) (vcontainermodels.VContainerClient, error) {
	return newVContainerClientInternal(logger, config, true)
}

func NewVContainerClient(logger lager.Logger, config vcontainermodels.VContainerClientConfig) (vcontainermodels.VContainerClient, error) {
	return newVContainerClientInternal(logger, config, false)
}

func newVContainerClientInternal(logger lager.Logger, config vcontainermodels.VContainerClientConfig, skipCertVerify bool) (vcontainermodels.VContainerClient, error) {
	conn, err := NewConn(logger, config)
	if err != nil {
		return nil, err
	}
	return vcontainermodels.NewVContainerClient(conn), nil
}
