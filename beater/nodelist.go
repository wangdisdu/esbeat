package beater

import (
	"encoding/json"
	"fmt"
	//"net/http"
	//"net/url"
	"strings"
)

const (
	NODES_ALL_HTTP_PATH = "/_nodes/_all/http"
)

type NodesHttpAddress struct {
	ClusterName string `json:"cluster_name"`
	Nodes       map[string]struct {
		Http struct {
			PublishAddress string `json:"publish_address"`
		}
	}
}

type NodesList struct {
	ClusterName string
	Nodes       []string
}

func (bt *Esbeat) GetNodesHttpAddress() (*NodesList, error) {
	address := NodesHttpAddress{}
	nodesList := NodesList{}
	http := NewHTTP(bt.config)

	var err error
	var body []byte
	urlIndex := 0

	for {
		if urlIndex >= len(bt.urls) {
			return nil, fmt.Errorf("All hosts fetch error: %v", bt.urls)
		}
		url := strings.TrimSuffix(bt.urls[urlIndex].String(), "/") + NODES_ALL_HTTP_PATH
		body, err = http.FetchContent(url)
		if err != nil {
			urlIndex = urlIndex + 1
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}

	//TODO sort bt.urls

	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}
	var list []string
	for name, node := range address.Nodes {
		fmt.Println("Node ID:", name)
		list = append(list, node.Http.PublishAddress)
	}

	nodesList.ClusterName = address.ClusterName
	nodesList.Nodes = list

	return &nodesList, nil
}
