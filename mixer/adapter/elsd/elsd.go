//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -f mixer/adapter/elsd/config/config.proto

package elsd

import (
	"context"
	"os"
	"path/filepath"

	"fmt"
	"istio.io/istio/mixer/adapter/elsd/config"
	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/istio/mixer/template/metric"

	"github.com/hpcwp/elsd/pkg/api"
	"google.golang.org/grpc"
	"time"
)

type (
	builder struct {
		adpCfg      *config.Params
		metricTypes map[string]*metric.Type
	}

	handler struct {
		f           *os.File
		elsdUrl 	string
		metricTypes map[string]*metric.Type
		env         adapter.Env
	}
)

// ensure types implement the requisite interfaces
var _ metric.HandlerBuilder = &builder{}
var _ metric.Handler = &handler{}

///////////////// Configuration-time Methods ///////////////

// adapter.HandlerBuilder#Build
func (b *builder) Build(ctx context.Context, env adapter.Env) (adapter.Handler, error) {
	var err error
	var file *os.File
	file, err = os.Create(b.adpCfg.FilePath)

	return &handler{f: file, metricTypes: b.metricTypes, env: env}, err
}

// adapter.HandlerBuilder#SetAdapterConfig
func (b *builder) SetAdapterConfig(cfg adapter.Config) {
	b.adpCfg = cfg.(*config.Params)

}

// adapter.HandlerBuilder#Validate
func (b *builder) Validate() (ce *adapter.ConfigErrors) {
	// Check if the path is valid
	if _, err := filepath.Abs(b.adpCfg.FilePath); err != nil {
		ce = ce.Append("file_path", err)
	}
	return
}

// metric.HandlerBuilder#SetMetricTypes
func (b *builder) SetMetricTypes(types map[string]*metric.Type) {
	b.metricTypes = types
}

////////////////// Request-time Methods //////////////////////////
// metric.Handler#HandleMetric
func (h *handler) HandleMetric(ctx context.Context, insts []*metric.Instance) error {

	// Just logging
	h.env.Logger().Infof("Handle metrics")

	// Call elsd
	h.elsdPing()

	h.env.Logger().Infof("Elsd healthy")

	for _, inst := range insts {
		if _, ok := h.metricTypes[inst.Name]; !ok {
			h.env.Logger().Errorf("Cannot find Type for instance %s", inst.Name)
			continue
		}
		h.f.WriteString(fmt.Sprintf(`HandleMetric invoke for :
		Instance Name  :'%s'
		Instance Value : %v,
		Type           : %v`, inst.Name, *inst, *h.metricTypes[inst.Name]))
	}
	return nil
}

// adapter.Handler#Close
func (h *handler) Close() error {
	return h.f.Close()
}

func (h *handler) elsdPing() {
	// Elsd address hardcoded for now - todo part of the adapter configuration
	grpcAddr := "localhost:8082"

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	healthClient := api.NewHealthClient(conn)
	_, err = healthClient.Check(context.Background(), &api.HealthCheckRequest{"els"})

	if err != nil {
		h.f.WriteString(fmt.Sprintf("elsd service is not responding \n"))
		os.Exit(1)
	}
	h.f.WriteString(fmt.Sprintf("elsd is ready \n"))

}

////////////////// Bootstrap //////////////////////////
// GetInfo returns the adapter.Info specific to this adapter.
func GetInfo() adapter.Info {
	return adapter.Info{
		Name:        "elsd",
		Description: "Logs the metric calls into a file",
		SupportedTemplates: []string{
			metric.TemplateName,
		},
		NewBuilder:    func() adapter.HandlerBuilder { return &builder{} },
		DefaultConfig: &config.Params{},
	}
}
