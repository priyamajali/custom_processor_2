// custom_processor.go
package resourceProcessor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.opentelemetry.io/collector/processor"
	// "go.opentelemetry.io/collector/config/configmodels"
	

	

	"go.uber.org/zap"
	
)
var processorCapabilities = consumer.Capabilities{MutatesData: true}

// type Config struct {

// 	// AttributesActions specifies the list of actions to be applied on resource attributes.
// 	// The set of actions are {INSERT, UPDATE, UPSERT, DELETE, HASH, EXTRACT}.
// 	// AttributesActions = "trial"
	
// }

type Config struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {
	
	return nil
}
type resourceProcessor struct {
	logger   *zap.Logger
	
}

func (rp *resourceProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	rms := md.ResourceMetrics()
	// rp.logger.Info("Name:", rp.Config.Name, "Value:", rp.Config.Value)
	// fmt.Println("Name:", rp.Config.Name, "Value:", rp.Config.Value)
	fmt.Println("Hello, World!")
	for i := 0; i < rms.Len(); i++ {
		rp.logger.Info("Hello, World!")

	}
	return md, nil
}



// func NewFactory() component.ProcessorFactory {
// 	return processor.NewFactory(
// 		"resource",
// 		// createDefaultConfig,
// 		processorhelper.WithMetrics(createMetricsProcessor))
// }

func NewFactory() processor.Factory {
	return processor.NewFactory(
		"resource",
		createDefaultConfig,
		processor.WithMetrics(createMetricsProcessor, component.StabilityLevelStable),
	)
}

// func createDefaultConfig()  {
// 	return 
// }

func createDefaultConfig() component.Config {
	return &Config{
		Name:  "defaultName",
		Value: "defaultValue",
	}
}


func createMetricsProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Metrics) (processor.Metrics, error) {
	
	proc := &resourceProcessor{logger: set.Logger}
	return processorhelper.NewMetricsProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		proc.processMetrics,
		processorhelper.WithCapabilities(processorCapabilities))
}

