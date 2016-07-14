package etcdstore

import (
	"fmt"

	"github.com/coreos/etcd/client"
	"gitlab.ricebook.net/platform/core/types"
)

var (
	allPodsKey       = "/eru-core/pod"
	podInfoKey       = "/eru-core/pod/%s/info"
	podNodesKey      = "/eru-core/pod/%s/node"
	nodeInfoKey      = "/eru-core/pod/%s/node/%s/info"
	nodeContainerKey = "/eru-core/pod/%s/node/%s/containers"
	containerInfoKey = "/eru-core/container/%s"
)

type krypton struct {
	etcd   client.KeysAPI
	config types.Config
}

func New(config types.Config) (*krypton, error) {
	if len(config.EtcdMachines) == 0 {
		return nil, fmt.Errorf("ETCD must be set")
	}

	cli, err := client.New(client.Config{Endpoints: config.EtcdMachines})
	if err != nil {
		return nil, err
	}

	etcd := client.NewKeysAPI(cli)
	return &krypton{etcd: etcd, config: config}, nil
}
