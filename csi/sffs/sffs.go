package sffs

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	log "github.com/sirupsen/logrus"

	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
)

type SFFS struct {
	driver *csicommon.CSIDriver

	idServer         *csicommon.DefaultIdentityServer
	controllerServer csi.ControllerServer

	endpoint string
}

func NewDriver(endpoint string) *SFFS {
	log.Info("Driver: ")

	d := &SFFS{}
	d.endpoint = endpoint
	return d
}

func newNodeServer(d *SFFS) *nodeServer {
	return &nodeServer{}
}

func (d *SFFS) Run() {
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(d.endpoint,
		csicommon.NewDefaultIdentityServer(d.driver),
		d.controllerServer,
		newNodeServer(d))
	s.Wait()
}
