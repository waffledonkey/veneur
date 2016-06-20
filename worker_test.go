package veneur

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	ReadConfig("example.yaml")
	w := NewWorker(1, nil, logrus.New(), Config.Percentiles, Config.HistCounters, Config.SetSize, Config.SetAccuracy)

	m := Metric{Name: "a.b.c", Value: 1.0, Digest: 12345, Type: "counter", SampleRate: 1.0}
	w.ProcessMetric(&m)

	ddmetrics := w.Flush(Config.Interval)
	assert.Len(t, ddmetrics, 1, "Number of flushed metrics")

	nometrics := w.Flush(Config.Interval)
	assert.Len(t, nometrics, 0, "Should flush no metrics")
}
