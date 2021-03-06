package run

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

// Globally used constants
const (
	DefaultRegistry    = "docker.io"
	DefaultServerCount = 1
)

// defaultNodes describes the type of nodes on which a port should be exposed by default
const defaultNodes = "server"

// mapping a node role to groups that should be applied to it
var nodeRuleGroupsMap = map[string][]string{
	"worker": {"all", "workers", "agents"},
	"server": {"all", "server", "master"},
}

// Cluster describes an existing cluster
type Cluster struct {
	name        string
	image       string
	status      string
	serverPorts []string
	server      types.Container
	workers     []types.Container
}

// ClusterSpec defines the specs for a cluster that's up for creation
type ClusterSpec struct {
	AgentArgs         []string
	APIPort           apiPort
	AutoRestart       bool
	ClusterName       string
	Env               []string
	Image             string
	NodeToPortSpecMap map[string][]string
	PortAutoOffset    int
	ServerArgs        []string
	Volumes           *Volumes
}

// PublishedPorts is a struct used for exposing container ports on the host system
type PublishedPorts struct {
	ExposedPorts map[nat.Port]struct{}
	PortBindings map[nat.Port][]nat.PortBinding
}
