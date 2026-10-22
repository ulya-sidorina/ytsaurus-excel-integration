package app

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestEnableStreamingWriteConfig(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		t.Run(fmt.Sprintf("enabled=%t", enabled), func(t *testing.T) {
			flag := ""
			if enabled {
				flag = "enable_streaming_write: true\n"
			}
			data := []byte(flag + `http_addr: ":6029"
debug_http_addr: ":6060"
clusters:
  - proxy: planck
`)

			var config Config
			require.NoError(t, yaml.Unmarshal(data, &config))
			require.Equal(t, enabled, config.EnableStreamingWrite)
			require.Equal(t, enabled, config.Clusters[0].enableStreamingWrite)
		})
	}
}
