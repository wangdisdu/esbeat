package beater

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/wangdisdu/esbeat/beater/helper"
	"net/url"
	"strings"
)

const (
	NODES_ALL_HTTP_PATH = "/_nodes/_all/http"
)

type NodeHttp struct {
	ClusterName string `json:"cluster_name"`
	Nodes       map[string]struct {
		Name string `json:"name"`
		Http struct {
			PublishAddress string `json:"publish_address"`
		} `json:"http"`
	} `json:"nodes"`
}

type NodeUrl struct {
	ClusterName string
	NodeName    string
	Url         *url.URL
}

func (bt *Esbeat) GetNodeUrls() ([]NodeUrl, error) {
	http := helper.NewHTTP(bt.config)

	var err error
	var body []byte
	urlIndex := 0

	for {
		if urlIndex >= len(bt.urls) {
			return nil, fmt.Errorf("All hosts fetch error: %v", bt.urls)
		}
		body, err = http.FetchContent(strings.TrimSuffix(bt.urls[urlIndex].String(), "/") + NODES_ALL_HTTP_PATH)
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

	address := NodeHttp{}
	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}
	var list []NodeUrl
	for _, node := range address.Nodes {
		u, err := url.Parse(fmt.Sprintf("%s://%s", bt.config.Protocol, node.Http.PublishAddress))
		if err != nil {
			logp.Err("Error gen url : %v", err)
			continue
		}

		nodeUrl := NodeUrl{}
		nodeUrl.ClusterName = address.ClusterName
		nodeUrl.NodeName = node.Name
		nodeUrl.Url = u
		list = append(list, nodeUrl)
	}

	return list, nil
}
