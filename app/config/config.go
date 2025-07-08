package config

import "sync"

type ClusterInfo struct{
    role string
}


type ServerConfig struct{
    port int
    ClusterInfo *ClusterInfo
    mu sync.RWMutex
}

var GlobalConfig *ServerConfig
var once sync.Once


func Initialize(port int) *ServerConfig{
    once.Do(func(){
        GlobalConfig = &ServerConfig{
            port: port,
            ClusterInfo: NewClusterInfo(),
        }
    })
    return GlobalConfig
}

func GetConfig() *ServerConfig {
	return GlobalConfig
}

func NewClusterInfo() *ClusterInfo {
	return &ClusterInfo{
		role:                      "master",
	}
}

func (sc *ServerConfig) GetRole() string {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.ClusterInfo.role
}

func (sc *ServerConfig) SetRole(role string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.ClusterInfo.role = role
}

func (sc *ServerConfig) GetClusterInfo() *ClusterInfo {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.ClusterInfo
}