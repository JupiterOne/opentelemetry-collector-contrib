// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlserverreceiver

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/confmap/confmaptest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver/internal/metadata"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		desc string
		cfg  *Config
	}{
		{
			desc: "valid config",
			cfg: &Config{
				Metrics: metadata.DefaultMetricsSettings(),
			},
		}, {
			desc: "valid config with no metric settings",
			cfg:  &Config{},
		},
		{
			desc: "default config is valid",
			cfg:  createDefaultConfig().(*Config),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actualErr := tc.cfg.Validate()
			require.NoError(t, actualErr)
		})
	}
}

func TestLoadConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	sub, err := cm.Sub("sqlserver")
	require.NoError(t, err)
	require.NoError(t, config.UnmarshalReceiver(sub, cfg))

	assert.NoError(t, cfg.Validate())
	assert.Equal(t, factory.CreateDefaultConfig(), cfg)
}
