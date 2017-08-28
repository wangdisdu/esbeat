// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period   time.Duration `config:"period"`
	Timeout  time.Duration `config:"timeout"`
	Protocol string        `config:"protocol"`
	Hosts    []string      `config:"hosts"`
	Stats    []string      `config:"stats"`
	Username string        `config:"username"`
	Password string        `config:"password"`
}

var DefaultConfig = Config{
	Period:   5 * time.Second,
	Timeout:  2 * time.Second,
	Protocol: "http",
	Hosts:    []string{"localhost:9200"},
	Stats:    []string{"nodes", "nodes_stats"},
}
