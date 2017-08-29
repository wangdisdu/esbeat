package beater

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/wangdisdu/esbeat/beater/helper"
	"github.com/wangdisdu/esbeat/beater/stats"
	"github.com/wangdisdu/esbeat/config"
)

type Esbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
	urls   []*url.URL
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Esbeat{
		done:   make(chan struct{}),
		config: config,
	}

	bt.urls = make([]*url.URL, len(config.Hosts))
	for i := 0; i < len(config.Hosts); i++ {
		host := fmt.Sprintf("%s://%s", config.Protocol, config.Hosts[i])
		u, err := url.Parse(host)
		if err != nil {
			logp.Err("Invalid ElasticSearch Host: %v", err)
			return nil, err
		}
		bt.urls[i] = u
	}

	return bt, nil
}

func (bt *Esbeat) Run(b *beat.Beat) error {
	logp.Info("esbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()

	nodeUrls, err := bt.GetNodeUrls()
	if err != nil {
		logp.Info("esbeat can not fetch cluster nodes.")
		return err
	}
	if len(nodeUrls) == 0 {
		logp.Info("esbeat can not gen urls for cluster nodes.")
		return err
	}

	var wg sync.WaitGroup
	for id, u := range nodeUrls {
		host := u.Url
		if bt.config.Stats.Node {
			wg.Add(1)
			go func(host *url.URL) {
				defer wg.Add(-1)
				bt.Polling("node", host, stats.FetchNode)
			}(host)
		}
		if bt.config.Stats.Nodestats {
			wg.Add(1)
			go func(host *url.URL) {
				defer wg.Add(-1)
				bt.Polling("nodestats", host, stats.FetchNodeStats)
			}(host)
		}
		if id == 0 && bt.config.Stats.Clusterhealth {
			wg.Add(1)
			go func(host *url.URL) {
				defer wg.Add(-1)
				bt.Polling("clusterhealth", host, stats.FetchClusterHealth)
			}(host)
		}
		if id == 0 && bt.config.Stats.Clusterstats {
			wg.Add(1)
			go func(host *url.URL) {
				defer wg.Add(-1)
				bt.Polling("clusterstats", host, stats.FetchClusterStats)
			}(host)
		}
	}

	wg.Wait()
	logp.Info("esbeat is stopping")
	return nil
}

func (bt *Esbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

type FuncFetchData func(http *helper.HTTP, host *url.URL) (interface{}, error)

//you should run it in goroutine
func (bt *Esbeat) Polling(name string, host *url.URL, fetchData FuncFetchData) error {
	logp.Info("esbeat-%s-%s is running", name, host.String())

	http := helper.NewHTTP(bt.config)
	ticker := time.NewTicker(bt.config.Period)

	for {
		select {
		case <-bt.done:
			logp.Info("esbeat-%s-%s is stopping", name, host.String())
			return nil
		case <-ticker.C:
		}

		body, err := fetchData(http, host)
		if err != nil {
			logp.Err("Error reading cluster node: %v", err)
		} else {
			event := common.MapStr{
				"@timestamp": common.Time(time.Now()),
				"type":       name,
				"url":        host.String(),
				name:         body,
			}
			bt.client.PublishEvent(event)
		}
	}
}
