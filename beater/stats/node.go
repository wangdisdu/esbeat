package stats

import (
	"encoding/json"
	"github.com/wangdisdu/esbeat/beater/helper"
	"net/url"
	"strings"
)

const (
	NODES_LOCAL_HTTP = "/_nodes/_local"
)

type NodeBody struct {
	ClusterName string              `json:"cluster_name"`
	Nodes       map[string]NodeInfo `json:"nodes"`
}

type NodeInfo struct {
	ClusterName string `json:"cluster_name"`
	Name        string `json:"name"`
	//TransportAddress string `json:"transport_address"`
	Host string `json:"host"`
	//Ip               string `json:"ip"`
	Version string `json:"version"`
	Os      struct {
		Name                string `json:"name"`
		AvailableProcessors int64  `json:"available_processors"`
		AllocatedProcessors int64  `json:"allocated_processors"`
	} `json:"os"`
	Process struct {
		Mlockall bool `json:"mlockall"`
	} `json:"process"`
	Jvm struct {
		Version string `json:"version"`
		Mem     struct {
			HeapInitInBytes    int64 `json:"heap_init_in_bytes"`
			HeapMaxInBytes     int64 `json:"heap_max_in_bytes"`
			NonHeapInitInBytes int64 `json:"non_heap_init_in_bytes"`
			NonHeapMaxInBytes  int64 `json:"non_heap_max_in_bytes"`
			DirectMaxInBytes   int64 `json:"direct_max_in_bytes"`
		} `json:"mem"`
	} `json:"jvm"`
}

func FetchNode(http *helper.HTTP, url *url.URL) (interface{}, error) {
	uri := strings.TrimSuffix(url.String(), "/") + NODES_LOCAL_HTTP
	nodeBody := NodeBody{}

	body, err := http.FetchContent(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &nodeBody)
	if err != nil {
		return nil, err
	}
	var nodeInfo NodeInfo
	for _, node := range nodeBody.Nodes {
		nodeInfo = node
		nodeInfo.ClusterName = nodeBody.ClusterName
		break
	}

	return &nodeInfo, nil
}
