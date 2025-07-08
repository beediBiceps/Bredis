package config

type ClusterInfo struct {
	role                      string
	connectedSlaves           int
	masterReplID              string
	masterReplOffset          int64
	secondReplOffset          int64
	replBacklogActive         int
	replBacklogSize           int64
	replBacklogFirstByteOffset int64
	replBacklogHistlen        int64
}

type PortConfig map[string]*ClusterInfo

func NewClusterInfo() *ClusterInfo {
	return &ClusterInfo{}
}
