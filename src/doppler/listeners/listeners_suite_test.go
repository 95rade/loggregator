package listeners_test

import (
	"github.com/cloudfoundry/dropsonde/emitter/fake"
	"github.com/cloudfoundry/dropsonde/metric_sender"
	"github.com/cloudfoundry/dropsonde/metricbatcher"
	"github.com/cloudfoundry/dropsonde/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"time"
)

func TestTcplistener(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Listeners Suite")
}

var (
	fakeEventEmitter = fake.NewFakeEventEmitter("doppler")
	metricBatcher    *metricbatcher.MetricBatcher
)

var _ = BeforeSuite(func() {
	sender := metric_sender.NewMetricSender(fakeEventEmitter)
	metricBatcher = metricbatcher.New(sender, 100*time.Millisecond)
	metrics.Initialize(sender, metricBatcher)
})
