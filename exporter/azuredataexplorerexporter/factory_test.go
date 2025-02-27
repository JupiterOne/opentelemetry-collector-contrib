// Copyright The OpenTelemetry Authors
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

package azuredataexplorerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtest"
	"go.opentelemetry.io/collector/service/servicetest"
)

// Given a new factory and no-op exporter , the NewMetric exporter should work.
// We could add additional failing tests if the config is wrong (using Validate) , but that is already done on config
func TestCreateMetricsExporter(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)
	factory := NewFactory()
	factories.Exporters[typeStr] = factory
	cfg, err := servicetest.LoadConfig(filepath.Join("testdata", "config.yaml"), factories)
	require.NoError(t, err)
	params := componenttest.NewNopExporterCreateSettings()
	exporter, err := factory.CreateMetricsExporter(context.Background(), params, cfg.Exporters[config.NewComponentID(typeStr)])
	assert.NotNil(t, err)
	assert.Nil(t, exporter)
}

// Given a new factory and no-op exporter , the NewMetric exporter should work.
// We could add additional failing tests if the config is wrong (using Validate) , but that is already done on config
func TestCreateMetricsExporterWhenIngestEmpty(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)
	factory := NewFactory()
	factories.Exporters[typeStr] = factory
	cfg, err := servicetest.LoadConfig(filepath.Join("testdata", "config.yaml"), factories)
	require.NoError(t, err)
	params := componenttest.NewNopExporterCreateSettings()
	// Load the #3 which has empty
	adxCfg := cfg.Exporters[config.NewComponentIDWithName(typeStr, "2")].(*Config)
	exporter, err := factory.CreateMetricsExporter(context.Background(), params, adxCfg)
	assert.NotNil(t, err)
	assert.Nil(t, exporter)
	// the fallback should be queued
	assert.Equal(t, queuedIngestTest, adxCfg.IngestionType)
}

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, configtest.CheckConfigStruct(cfg))
	assert.Equal(t, otelDb, cfg.Database)
	assert.Equal(t, queuedIngestTest, cfg.IngestionType)
}

// Given a new factory and no-op exporter , the LogExporter exporter should work.
// We could add additional failing tests if the config is wrong (using Validate) , but that is already done on config
func TestCreateLogsExporter(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)
	factory := NewFactory()
	factories.Exporters[typeStr] = factory
	cfg, err := servicetest.LoadConfig(filepath.Join("testdata", "config.yaml"), factories)
	require.NoError(t, err)
	params := componenttest.NewNopExporterCreateSettings()
	exporter, err := factory.CreateLogsExporter(context.Background(), params, cfg.Exporters[config.NewComponentID(typeStr)])
	assert.NotNil(t, err)
	assert.Nil(t, exporter)
}

// Given a new factory and no-op exporter , the NewLogs exporter should work.
// We could add additional failing tests if the config is wrong (using Validate) , but that is already done on config
func TestCreateLogsExporterWhenIngestEmpty(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)
	factory := NewFactory()
	factories.Exporters[typeStr] = factory
	cfg, err := servicetest.LoadConfig(filepath.Join("testdata", "config.yaml"), factories)
	require.NoError(t, err)
	params := componenttest.NewNopExporterCreateSettings()
	// Load the #3 which has empty
	adxCfg := cfg.Exporters[config.NewComponentIDWithName(typeStr, "2")].(*Config)
	exporter, err := factory.CreateLogsExporter(context.Background(), params, adxCfg)
	assert.NotNil(t, err)
	assert.Nil(t, exporter)
	// the fallback should be queued
	assert.Equal(t, queuedIngestTest, adxCfg.IngestionType)
}

// Given a new factory and no-op exporter , the LogExporter exporter should work.
// We could add additional failing tests if the config is wrong (using Validate) , but that is already done on config
func TestCreateTracesExporter(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)
	factory := NewFactory()
	factories.Exporters[typeStr] = factory
	cfg, err := servicetest.LoadConfig(filepath.Join("testdata", "config.yaml"), factories)
	require.NoError(t, err)
	params := componenttest.NewNopExporterCreateSettings()
	exporter, err := factory.CreateTracesExporter(context.Background(), params, cfg.Exporters[config.NewComponentID(typeStr)])
	assert.NotNil(t, err)
	assert.Nil(t, exporter)
}

// Given a new factory and no-op exporter , the NewLogs exporter should work.
// We could add additional failing tests if the config is wrong (using Validate) , but that is already done on config
func TestCreateTracesExporterWhenIngestEmpty(t *testing.T) {
	factories, err := componenttest.NopFactories()
	require.NoError(t, err)
	factory := NewFactory()
	factories.Exporters[typeStr] = factory
	cfg, err := servicetest.LoadConfig(filepath.Join("testdata", "config.yaml"), factories)
	require.NoError(t, err)
	params := componenttest.NewNopExporterCreateSettings()
	// Load the #3 which has empty
	adxCfg := cfg.Exporters[config.NewComponentIDWithName(typeStr, "2")].(*Config)
	exporter, err := factory.CreateTracesExporter(context.Background(), params, adxCfg)
	assert.NotNil(t, err)
	assert.Nil(t, exporter)
	// the fallback should be queued
	assert.Equal(t, queuedIngestTest, adxCfg.IngestionType)
}
