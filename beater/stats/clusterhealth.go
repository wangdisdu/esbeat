package stats

import (
	"encoding/json"
	"github.com/wangdisdu/esbeat/beater/helper"
	"net/url"
	"strings"
)

const (
	CLUSTER_HEALTH_PATH = "/_cluster/health"
)

type ClusterHealth struct {
	ClusterName                 string  `json:"cluster_name"`
	Status                      string  `json:"status"`
	TimedOut                    bool    `json:"timed_out"`
	NumberOfNodes               int64   `json:"number_of_nodes"`
	NumberOfDataNodes           int64   `json:"number_of_data_nodes"`
	ActivePrimaryShards         int64   `json:"active_primary_shards"`
	ActiveShards                int64   `json:"active_shards"`
	RelocatingShards            int64   `json:"relocating_shards"`
	InitializingShards          int64   `json:"intializing_shards"`
	UnassignedShards            int64   `json:"unassigned_shards"`
	DelayedUnassignedShards     int64   `json:"delayed_unassigned_shards"`
	NumberOfPendingTasks        int64   `json:"number_of_pending_tasks"`
	NumberOfInFlightFetch       int64   `json:"number_of_in_flight_fetch"`
	TaskMaxWaitingInQueueMillis int64   `json:"task_max_waiting_in_queue_millis"`
	ActiveShardsPercentAsNumber float64 `json:"active_shards_percent_as_number"`
}

func FetchClusterHealth(http *helper.HTTP, host *url.URL) (interface{}, error) {
	uri := strings.TrimSuffix(host.String(), "/") + CLUSTER_HEALTH_PATH
	var clusterHealth ClusterHealth

	body, err := http.FetchContent(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &clusterHealth)
	if err != nil {
		return nil, err
	}

	return &clusterHealth, nil
}
