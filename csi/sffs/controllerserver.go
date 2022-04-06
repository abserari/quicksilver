package sffs

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type controllerServer struct {
	client kubernetes.Interface
	*csicommon.DefaultControllerServer
}

// used by check pvc is processed
var pvcProcessSuccess = map[string]*csi.Volume{}
var storageClassServerPos = map[string]int{}

// NewControllerServer is to create controller server
func NewControllerServer(d *csicommon.CSIDriver) csi.ControllerServer {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create config: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("NewControllerServer: Failed to create client: %v", err)
	}

	c := &controllerServer{
		client:                  clientset,
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
	}
	return c
}

func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	pvcUID := string(req.Name)
	if value, ok := pvcProcessSuccess[pvcUID]; ok {
		log.Warnf("CreateVolume: SFFS Volume %s has Created Already", req.Name)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

	pvName := req.Name
	for _, volCap := range req.VolumeCapabilities {
		var volCapMount *csi.VolumeCapability_Mount
		volCapMount, ok := ((*volCap).AccessType).(*csi.VolumeCapability_Mount)
		if !ok {
			return nil, fmt.Errorf("CreateVolume: Input sffs type error: volume: %v", req)
		}
		for _, mountFlag := range volCapMount.Mount.MountFlags {
			fmt.Println(mountFlag)
		}
	}

	// check provision mode
	fs := "ext4"
	for k, v := range req.Parameters {
		switch strings.ToLower(k) {
		case "FSTYPE":
			fs = strings.TrimSpace(v)
		default:
		}
	}
	volumeContext := map[string]string{}
	volumeContext["fileSystem"] = fs
	volumeContext["subPath"] = filepath.Join("sffs", pvName)
	if value, ok := req.Parameters["options"]; ok && value != "" {
		volumeContext["options"] = value
	}

	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	tmpVol := &csi.Volume{
		VolumeId:      req.Name,
		CapacityBytes: int64(volSizeBytes),
		VolumeContext: volumeContext,
	}

	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil

}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume is called, do nothing by now")
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume is called, do nothing by now")
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	log.Infof("CreateSnapshot is called, do nothing now")
	return &csi.CreateSnapshotResponse{}, nil
}

func (cs *controllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	log.Infof("DeleteSnapshot is called, do nothing now")
	return &csi.DeleteSnapshotResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest,
) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof("ControllerExpandVolume is called, do nothing now")
	return &csi.ControllerExpandVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	return &csi.ControllerGetVolumeResponse{}, nil
}
