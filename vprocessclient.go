package vcontainerclient

import (
	"code.cloudfoundry.org/lager"
	"github.com/virtualcloudfoundry/vcontainercommon/vcontainermodels"
)

func NewVProcessClientSkipCertVerify(logger lager.Logger, config vcontainermodels.VContainerClientConfig) (vcontainermodels.VProcessClient, error) {
	return newVProcessClientInternal(logger, config, true)
}

func NewVProcessClient(logger lager.Logger, config vcontainermodels.VContainerClientConfig) (vcontainermodels.VProcessClient, error) {
	return newVProcessClientInternal(logger, config, false)
}

func newVProcessClientInternal(logger lager.Logger, config vcontainermodels.VContainerClientConfig, skipCertVerify bool) (vcontainermodels.VProcessClient, error) {
	conn, err := NewConn(logger, config)
	if err != nil {
		return nil, err
	}
	return vcontainermodels.NewVProcessClient(conn), nil
}
