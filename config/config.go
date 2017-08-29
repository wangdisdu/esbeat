// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period   time.Duration `config:"period"`
	Timeout  time.Duration `config:"timeout"`
	Protocol string        `config:"protocol"`
	Hosts    []string      `config:"hosts"`
	Username string        `config:"username"`
	Password string        `config:"password"`
	Stats    ConfigStats   `config:"stats"`
}

type ConfigStats struct {
	Node          bool `config:"node"`
	Nodestats     bool `config:"nodestats"`
	Clusterhealth bool `config:"clusterhealth"`
	Clusterstats  bool `config:"clusterstats"`
}

var DefaultConfig = Config{
	Period:   5 * time.Second,
	Timeout:  2 * time.Second,
	Protocol: "http",
	Hosts:    []string{"localhost:9200"},
	Stats: ConfigStats{
		Node:          true,
		Nodestats:     true,
		Clusterhealth: true,
		Clusterstats:  true,
	},
}
