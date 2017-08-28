package node

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strings"
	"time"
)

const (
	NODES_LOCAL_HTTP = "_nodes/_local"
)
