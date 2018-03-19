package vcontainerclient

import (
	"time"

	"code.cloudfoundry.org/cfhttp"
	"code.cloudfoundry.org/lager"
	"github.com/virtualcloudfoundry/vcontainercommon/vcontainermodels"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewConnSkipCertVerify(logger lager.Logger, config vcontainermodels.VContainerClientConfig) (*grpc.ClientConn, error) {
	return newConnInternal(logger, config, true)
}

func NewConn(logger lager.Logger, config vcontainermodels.VContainerClientConfig) (*grpc.ClientConn, error) {
	return newConnInternal(logger, config, false)
}

func newConnInternal(logger lager.Logger, config vcontainermodels.VContainerClientConfig, skipCertVerify bool) (*grpc.ClientConn, error) {
	logger.Info("new-conn-internal", lager.Data{"config": config})
	if config.VContainerAddress == "" {
		logger.Fatal("invalid-vcontainer-config", nil)
	}

	vcontainerTLSConfig, err := cfhttp.NewTLSConfig(config.VContainerClientCertFile, config.VContainerClientKeyFile, config.VContainerCACertFile)
	if err != nil {
		logger.Error("failed-to-open-tls-config", err, lager.Data{"keypath": config.VContainerClientKeyFile, "certpath": config.VContainerClientCertFile, "capath": config.VContainerCACertFile})
		return nil, err
	}
	vcontainerTLSConfig.InsecureSkipVerify = skipCertVerify

	conn, err := grpc.Dial(
		config.VContainerAddress,
		grpc.WithTransportCredentials(credentials.NewTLS(vcontainerTLSConfig)),
		grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),
	)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
