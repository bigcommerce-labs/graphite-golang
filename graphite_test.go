package graphite

import (
	"net"
	"testing"
)

// Change these to be your own graphite server if you so please
var graphiteHost = "carbon.hostedgraphite.com"
var graphitePort = 2003

func TestNewGraphite(t *testing.T) {
	gh, err := NewGraphite(graphiteHost, graphitePort)
	if err != nil {
		t.Error(err)
	}

	if _, ok := gh.conn.(*net.TCPConn); !ok {
		t.Error("GraphiteHost.conn is not a TCP connection")
	}
}

// Uncomment the following method to test sending an actual metric to graphite
func TestSendMetric(t *testing.T) {
	gh, err := NewGraphite(graphiteHost, graphitePort)
	defer gh.Close()

	if err != nil {
		t.Error(err)
	}
	err = gh.SimpleSend("e6c05da9-cf17-4dac-9696-1a3f8e0f1967.stats.test.metric11", "1")
	if err != nil {
		t.Error(err)
	}
}
