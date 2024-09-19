// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testDataSet int

const (
	testDataSetDefault testDataSet = iota
	testDataSetAll
	testDataSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name        string
		metricsSet  testDataSet
		resAttrsSet testDataSet
		expectEmpty bool
	}{
		{
			name: "default",
		},
		{
			name:        "all_set",
			metricsSet:  testDataSetAll,
			resAttrsSet: testDataSetAll,
		},
		{
			name:        "none_set",
			metricsSet:  testDataSetNone,
			resAttrsSet: testDataSetNone,
			expectEmpty: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			allMetricsCount++
			mb.RecordNtpFrequencyOffsetDataPoint(ts, 1, AttributeLeapStatusNormal)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNtpSkewDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordNtpStratumDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNtpTimeCorrectionDataPoint(ts, 1, AttributeLeapStatusNormal)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordNtpTimeLastOffsetDataPoint(ts, 1, AttributeLeapStatusNormal)

			allMetricsCount++
			mb.RecordNtpTimeRmsOffsetDataPoint(ts, 1, AttributeLeapStatusNormal)

			allMetricsCount++
			mb.RecordNtpTimeRootDelayDataPoint(ts, 1, AttributeLeapStatusNormal)

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if test.expectEmpty {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.metricsSet == testDataSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.metricsSet == testDataSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "ntp.frequency.offset":
					assert.False(t, validatedMetrics["ntp.frequency.offset"], "Found a duplicate in the metrics slice: ntp.frequency.offset")
					validatedMetrics["ntp.frequency.offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The frequency is the rate by which the system s clock would be wrong if chronyd was not correcting it.", ms.At(i).Description())
					assert.Equal(t, "ppm", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("leap.status")
					assert.True(t, ok)
					assert.EqualValues(t, "normal", attrVal.Str())
				case "ntp.skew":
					assert.False(t, validatedMetrics["ntp.skew"], "Found a duplicate in the metrics slice: ntp.skew")
					validatedMetrics["ntp.skew"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "This is the estimated error bound on the frequency.", ms.At(i).Description())
					assert.Equal(t, "ppm", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
				case "ntp.stratum":
					assert.False(t, validatedMetrics["ntp.stratum"], "Found a duplicate in the metrics slice: ntp.stratum")
					validatedMetrics["ntp.stratum"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of hops away from the reference system keeping the reference time", ms.At(i).Description())
					assert.Equal(t, "{count}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "ntp.time.correction":
					assert.False(t, validatedMetrics["ntp.time.correction"], "Found a duplicate in the metrics slice: ntp.time.correction")
					validatedMetrics["ntp.time.correction"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The number of seconds difference between the system's clock and the reference clock", ms.At(i).Description())
					assert.Equal(t, "seconds", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("leap.status")
					assert.True(t, ok)
					assert.EqualValues(t, "normal", attrVal.Str())
				case "ntp.time.last_offset":
					assert.False(t, validatedMetrics["ntp.time.last_offset"], "Found a duplicate in the metrics slice: ntp.time.last_offset")
					validatedMetrics["ntp.time.last_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The estimated local offset on the last clock update", ms.At(i).Description())
					assert.Equal(t, "seconds", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("leap.status")
					assert.True(t, ok)
					assert.EqualValues(t, "normal", attrVal.Str())
				case "ntp.time.rms_offset":
					assert.False(t, validatedMetrics["ntp.time.rms_offset"], "Found a duplicate in the metrics slice: ntp.time.rms_offset")
					validatedMetrics["ntp.time.rms_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "the long term average of the offset value", ms.At(i).Description())
					assert.Equal(t, "seconds", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("leap.status")
					assert.True(t, ok)
					assert.EqualValues(t, "normal", attrVal.Str())
				case "ntp.time.root_delay":
					assert.False(t, validatedMetrics["ntp.time.root_delay"], "Found a duplicate in the metrics slice: ntp.time.root_delay")
					validatedMetrics["ntp.time.root_delay"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "This is the total of the network path delays to the stratum-1 system from which the system is ultimately synchronised.", ms.At(i).Description())
					assert.Equal(t, "seconds", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("leap.status")
					assert.True(t, ok)
					assert.EqualValues(t, "normal", attrVal.Str())
				}
			}
		})
	}
}
