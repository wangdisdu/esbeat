package stats

import (
	"encoding/json"
	"github.com/wangdisdu/esbeat/beater/helper"
	"net/url"
	"strings"
)

const (
	NODES_LOCAL_STATS = "/_nodes/_local/stats"
)

type NodeStatsBody struct {
	ClusterName string                   `json:"cluster_name"`
	Nodes       map[string]NodeStatsInfo `json:"nodes"`
}

type NodeStatsInfo struct {
	Timestamp   int64  `json:"timestamp"`
	ClusterName string `json:"cluster_name"`
	Name        string `json:"name"`
	//TransportAddress string `json:"transport_address"`
	Host string `json:"host"`
	//Ip               string `json:"ip"`
	Indices struct {
		Docs struct {
			Count   int64 `json:"count"`
			Deleted int64 `json:"deleted"`
		} `json:"docs"`
		Store struct {
			SizeInBytes          int64 `json:"size_in_bytes"`
			ThrottleTimeInMillis int64 `json:"throttle_time_in_millis"`
		} `json:"store"`
		Indexing struct {
			IndexTotal         int64 `json:"index_total"`
			IndexTimeInMillis  int64 `json:"index_time_in_millis"`
			IndexCurrent       int64 `json:"index_current"`
			DeleteTotal        int64 `json:"delete_total"`
			DeleteTimeInMillis int64 `json:"delete_time_in_millis"`
			DeleteCurrent      int64 `json:"delete_current"`
		} `json:"indexing"`
		Get struct {
			Total               int64 `json:"total"`
			TimeInMillis        int64 `json:"time_in_millis"`
			ExistsTotal         int64 `json:"exists_total"`
			ExistsTimeInMillis  int64 `json:"exists_time_in_millis"`
			MissingTotal        int64 `json:"missing_total"`
			MissingTimeInMillis int64 `json:"missing_time_in_millis"`
			Current             int64 `json:"current"`
		} `json:"get"`
		Search struct {
			OpenContexts      int64 `json:"open_contexts"`
			QueryTotal        int64 `json:"query_total"`
			QueryTimeInMillis int64 `json:"query_time_in_millis"`
			QueryCurrent      int64 `json:"query_current"`
			FetchTotal        int64 `json:"fetch_total"`
			FetchTimeInMillis int64 `json:"fetch_time_in_millis"`
			FetchCurrent      int64 `json:"fetch_current"`
		} `json:"search"`
		Merges struct {
			Current            int64 `json:"current"`
			CurrentDocs        int64 `json:"current_docs"`
			CurrentSizeInBytes int64 `json:"current_size_in_bytes"`
			Total              int64 `json:"total"`
			TotalTimeInMillis  int64 `json:"total_time_in_millis"`
			TotalDocs          int64 `json:"total_docs"`
			TotalSizeInBytes   int64 `json:"total_size_in_bytes"`
		} `json:"merges"`
		FilterCache struct {
			MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
			Evictions         int64 `json:"evictions"`
		} `json:"filter_cache"`
		RequestCache struct {
			MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
			Evictions         int64 `json:"evictions"`
		} `json:"request_cache"`
		Segments struct {
			Count         int64 `json:"count"`
			MemoryInBytes int64 `json:"memory_in_bytes"`
		} `json:"segments"`
	} `json:"indices"`
	Os struct {
		Timestamp   int64    `json:"timestamp"`
		LoadAverage float64  `json:"load_average"` //for es 2x version
		Cpu         struct { //for es 5x version
			Percent     int64 `json:"percent"`
			LoadAverage struct {
				OneMinute     float64 `json:"1m"`
				FiveMinute    float64 `json:"5m"`
				FifteenMinute float64 `json:"15m"`
			} `json:"load_average"`
		} `json:"cpu"`
		Mem struct {
			TotalInBytes int64 `json:"total_in_bytes"`
			FreeInBytes  int64 `json:"free_in_bytes"`
			UsedInBytes  int64 `json:"used_in_bytes"`
			FreePercent  int64 `json:"free_percent"`
			UsedPercent  int64 `json:"used_percent"`
		} `json:"mem"`
		Swap struct {
			TotalInBytes int64 `json:"total_in_bytes"`
			FreeInBytes  int64 `json:"free_in_bytes"`
			UsedInBytes  int64 `json:"used_in_bytes"`
		} `json:"swap"`
	} `json:"os"`
	Process struct {
		Timestamp           int64 `json:"timestamp"`
		OpenFileDescriptors int64 `json:"open_file_descriptors"`
		MaxFileDescriptors  int64 `json:"max_file_descriptors"`
		Cpu                 struct {
			Percent       int64 `json:"percent"`
			TotalInMillis int64 `json:"total_in_millis"`
		} `json:"cpu"`
		Mem struct {
			TotalVirtualInBytes int64 `json:"total_virtual_in_bytes"`
		} `json:"mem"`
	} `json:"process"`
	Jvm struct {
		Timestamp      int64 `json:"timestamp"`
		UptimeInMillis int64 `json:"uptime_in_millis"`
		Mem            struct {
			HeapUsedInBytes         int64 `json:"heap_used_in_bytes"`
			HeapUsedPercent         int64 `json:"heap_used_percent"`
			HeapCommittedInBytes    int64 `json:"heap_committed_in_bytes"`
			HeapMaxInBytes          int64 `json:"heap_max_in_bytes"`
			NonHeapUsedInBytes      int64 `json:"non_heap_used_in_bytes"`
			NonHeapCommittedInBytes int64 `json:"non_heap_committed_in_bytes"`
			Pools                   struct {
				Young struct {
					UsedInBytes     int64 `json:"used_in_bytes"`
					MaxInBytes      int64 `json:"max_in_bytes"`
					PeakUsedInBytes int64 `json:"peak_used_in_bytes"`
					PeakMaxInBytes  int64 `json:"peak_max_in_bytes"`
				} `json:"young"`
				Survivor struct {
					UsedInBytes     int64 `json:"used_in_bytes"`
					MaxInBytes      int64 `json:"max_in_bytes"`
					PeakUsedInBytes int64 `json:"peak_used_in_bytes"`
					PeakMaxInBytes  int64 `json:"peak_max_in_bytes"`
				} `json:"survivor"`
				Old struct {
					UsedInBytes     int64 `json:"used_in_bytes"`
					MaxInBytes      int64 `json:"max_in_bytes"`
					PeakUsedInBytes int64 `json:"peak_used_in_bytes"`
					PeakMaxInBytes  int64 `json:"peak_max_in_bytes"`
				} `json:"old"`
			} `json:"pools"`
		} `json:"mem"`
		Threads struct {
			Count     int64 `json:"count"`
			PeakCount int64 `json:"peak_count"`
		} `json:"threads"`
		Gc struct {
			Collectors struct {
				Young struct {
					CollectionCount        int64 `json:"collection_count"`
					CollectionTimeInMillis int64 `json:"collection_time_in_millis"`
				} `json:"young"`
				Old struct {
					Collection_count       int64 `json:"collection_count"`
					CollectionTimeInMillis int64 `json:"collection_time_in_millis"`
				} `json:"old"`
			} `json:"collectors"`
		} `json:"gc"`
		BufferPools struct {
			Direct struct {
				Count                int64 `json:"count"`
				UsedInBytes          int64 `json:"used_in_bytes"`
				TotalCapacityInBytes int64 `json:"total_capacity_in_bytes"`
			} `json:"direct"`
			Mapped struct {
				Count                int64 `json:"count"`
				UsedInBytes          int64 `json:"used_in_bytes"`
				TotalCapacityInBytes int64 `json:"total_capacity_in_bytes"`
			} `json:"mapped"`
		} `json:"buffer_pools"`
	} `json:"jvm"`
	ThreadPool struct { //focus on are index, bulk, search, and merge, if someone want to collect more detail, tell me!
		Bulk struct {
			Threads   int64 `json:"threads"`
			Queue     int64 `json:"queue"`
			Active    int64 `json:"active"`
			Rejected  int64 `json:"rejected"`
			Largest   int64 `json:"largest"`
			Completed int64 `json:"completed"`
		} `json:"bulk"`
		Flush struct {
			Threads   int64 `json:"threads"`
			Queue     int64 `json:"queue"`
			Active    int64 `json:"active"`
			Rejected  int64 `json:"rejected"`
			Largest   int64 `json:"largest"`
			Completed int64 `json:"completed"`
		} `json:"flush"`
		ForceMerge struct {
			Threads   int64 `json:"threads"`
			Queue     int64 `json:"queue"`
			Active    int64 `json:"active"`
			Rejected  int64 `json:"rejected"`
			Largest   int64 `json:"largest"`
			Completed int64 `json:"completed"`
		} `json:"force_merge"`
		Get struct {
			Threads   int64 `json:"threads"`
			Queue     int64 `json:"queue"`
			Active    int64 `json:"active"`
			Rejected  int64 `json:"rejected"`
			Largest   int64 `json:"largest"`
			Completed int64 `json:"completed"`
		} `json:"get"`
		Index struct {
			Threads   int64 `json:"threads"`
			Queue     int64 `json:"queue"`
			Active    int64 `json:"active"`
			Rejected  int64 `json:"rejected"`
			Largest   int64 `json:"largest"`
			Completed int64 `json:"completed"`
		} `json:"index"`
		Search struct {
			Threads   int64 `json:"threads"`
			Queue     int64 `json:"queue"`
			Active    int64 `json:"active"`
			Rejected  int64 `json:"rejected"`
			Largest   int64 `json:"largest"`
			Completed int64 `json:"completed"`
		} `json:"search"`
	} `json:"thread_pool"`
	Fs struct {
		Timestamp int64 `json:"timestamp"`
		Total     struct {
			TotalInBytes     int64  `json:"total_in_bytes"`
			FreeInBytes      int64  `json:"free_in_bytes"`
			AvailableInBytes int64  `json:"available_in_bytes"`
			Spins            string `json:"spins"`
		} `json:"total"`
	} `json:"fs"`
	Transport struct {
		ServerOpen    int64 `json:"server_open"`
		RxCount       int64 `json:"rx_count"`
		RxSizeInBytes int64 `json:"rx_size_in_bytes"`
		TxCount       int64 `json:"tx_count"`
		TxSizeInBytes int64 `json:"tx_size_in_bytes"`
	} `json:"transport"`
	Http struct {
		CurrentOpen int64 `json:"current_open"`
		TotalOpened int64 `json:"total_opened"`
	} `json:"http"`
}

func FetchNodeStats(http *helper.HTTP, url *url.URL) (interface{}, error) {
	uri := strings.TrimSuffix(url.String(), "/") + NODES_LOCAL_STATS
	nodeBody := NodeStatsBody{}

	body, err := http.FetchContent(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &nodeBody)
	if err != nil {
		return nil, err
	}
	var nodeInfo NodeStatsInfo
	for _, node := range nodeBody.Nodes {
		nodeInfo = node
		nodeInfo.ClusterName = nodeBody.ClusterName
		break
	}

	return &nodeInfo, nil
}
