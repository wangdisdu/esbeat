package stats

import (
	"encoding/json"
	"github.com/wangdisdu/esbeat/beater/helper"
	"net/url"
	"strings"
)

const (
	CLUSTER_STATS_PATH = "/_cluster/stats"
)

type ClusterStats struct {
	Timestamp   int64  `json:"timestamp"`
	ClusterName string `json:"cluster_name"`
	Status      string `json:"status"`
	Indices     struct {
		Count  int64 `json:"count"`
		Shards struct {
			Total       int64   `json:"total"`
			Primaries   int64   `json:"primaries"`
			Replication float64 `json:"replication"`
			Index       struct {
				Shards struct {
					Min float64 `json:"min"`
					Max float64 `json:"max"`
					Avg float64 `json:"avg"`
				} `json:"shards"`
				Primaries struct {
					Min float64 `json:"min"`
					Max float64 `json:"max"`
					Avg float64 `json:"avg"`
				} `json:"primaries"`
				Replication struct {
					Min float64 `json:"min"`
					Max float64 `json:"max"`
					Avg float64 `json:"avg"`
				} `json:"replication"`
			}
		} `json:"shards"`
		Docs struct {
			Count   int64 `json:"count"`
			Deleted int64 `json:"deleted"`
		} `json:"docs"`
		Store struct {
			SizeInBytes          int64 `json:"size_in_bytes"`
			ThrottleTimeInMillis int64 `json:"throttle_time_in_millis"`
		} `json:"store"`
		Fielddata struct {
			MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
			Evictions         int64 `json:"evictions"`
		} `json:"fielddata"`
		QueryCache struct {
			MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
			TotalCount        int64 `json:"total_count"`
			HitCount          int64 `json:"hit_count"`
			MissCount         int64 `json:"miss_count"`
			CacheSize         int64 `json:"cache_size"`
			CacheCount        int64 `json:"cache_count"`
			Evictions         int64 `json:"evictions"`
		} `json:"query_cache"`
		Completion struct {
			SizeInBytes int64 `json:"size_in_bytes"`
		} `json:"completion"`
		Segments struct {
			Count                     int64 `json:"count"`
			MemoryInBytes             int64 `json:"memory_in_bytes"`
			TermsMemoryInBytes        int64 `json:"terms_memory_in_bytes"`
			StoredFieldsMemoryInBytes int64 `json:"stored_fields_memory_in_bytes"`
			TermVectorsMemoryInBytes  int64 `json:"term_vectors_memory_in_bytes"`
			NormsMemoryInBytes        int64 `json:"norms_memory_in_bytes"`
			DocValuesMemoryInBytes    int64 `json:"doc_values_memory_in_bytes"`
			IndexWriterMemoryInBytes  int64 `json:"index_writer_memory_in_bytes"`
			VersionMapMemoryInBytes   int64 `json:"version_map_memory_in_bytes"`
			FixedBitSetMemoryInBytes  int64 `json:"fixed_bit_set_memory_in_bytes"`
		} `json:"segments"`
	} `json:"indices"`
	Nodes struct {
		Count struct {
			Total int64 `json:"total"`
		} `json:"count"`
		Os struct {
			AvailableProcessors int64 `json:"available_processors"`
			AllocatedProcessors int64 `json:"allocated_processors"`
			Mem                 struct {
				Total_in_bytes int64 `json:"total_in_bytes"`
			} `json:"mem"`
		} `json:"os"`
		Process struct {
			Cpu struct {
				Percent int64 `json:"percent"`
			} `json:"cpu"`
			OpenFileDescriptors struct {
				Min float64 `json:"min"`
				Max float64 `json:"max"`
				Avg float64 `json:"avg"`
			} `json:"open_file_descriptors"`
		} `json:"process"`
		Jvm struct {
			MaxUptimeInMillis int64 `json:"max_uptime_in_millis"`
			Mem               struct {
				HeapUsedInBytes int64 `json:"heap_used_in_bytes"`
				HeapMaxInBytes  int64 `json:"heap_max_in_bytes"`
			} `json:"mem"`
			Threads int64 `json:"threads"`
		} `json:"jvm"`
		Fs struct {
			Total_in_bytes     int64  `json:"total_in_bytes"`
			Free_in_bytes      int64  `json:"free_in_bytes"`
			Available_in_bytes int64  `json:"available_in_bytes"`
			Spins              string `json:"spins"`
		} `json:"fs"`
	} `json:"nodes"`
}

func FetchClusterStats(http *helper.HTTP, host *url.URL) (interface{}, error) {
	uri := strings.TrimSuffix(host.String(), "/") + CLUSTER_STATS_PATH
	var clusterStats ClusterStats

	body, err := http.FetchContent(uri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &clusterStats)
	if err != nil {
		return nil, err
	}

	return &clusterStats, nil
}
