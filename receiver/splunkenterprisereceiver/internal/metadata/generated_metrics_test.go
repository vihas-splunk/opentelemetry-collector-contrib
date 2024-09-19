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

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkAggregationQueueRatioDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkBucketsSearchableStatusDataPoint(ts, 1, "splunk.host-val", "splunk.indexer.searchable-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketCountDataPoint(ts, 1, "splunk.index.name-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketEventCountDataPoint(ts, 1, "splunk.index.name-val", "splunk.bucket.dir-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketHotCountDataPoint(ts, 1, "splunk.index.name-val", "splunk.bucket.dir-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedBucketWarmCountDataPoint(ts, 1, "splunk.index.name-val", "splunk.bucket.dir-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedEventCountDataPoint(ts, 1, "splunk.index.name-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedRawSizeDataPoint(ts, 1, "splunk.index.name-val")

			allMetricsCount++
			mb.RecordSplunkDataIndexesExtendedTotalSizeDataPoint(ts, 1, "splunk.index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexerAvgRateDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexerCPUTimeDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexerQueueRatioDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexerRawWriteTimeDataPoint(ts, 1, "splunk.host-val")

			allMetricsCount++
			mb.RecordSplunkIndexerThroughputDataPoint(ts, 1, "splunk.indexer.status-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexesAvgSizeDataPoint(ts, 1, "splunk.index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexesAvgUsageDataPoint(ts, 1, "splunk.index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexesBucketCountDataPoint(ts, 1, "splunk.index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexesMedianDataAgeDataPoint(ts, 1, "splunk.index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIndexesSizeDataPoint(ts, 1, "splunk.index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkIoAvgIopsDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkLicenseIndexUsageDataPoint(ts, 1, "splunk.index.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkParseQueueRatioDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkPipelineSetCountDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkSchedulerAvgExecutionLatencyDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkSchedulerAvgRunTimeDataPoint(ts, 1, "splunk.host-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkSchedulerCompletionRatioDataPoint(ts, 1, "splunk.host-val")

			allMetricsCount++
			mb.RecordSplunkServerIntrospectionQueuesCurrentDataPoint(ts, 1, "splunk.queue.name-val")

			allMetricsCount++
			mb.RecordSplunkServerIntrospectionQueuesCurrentBytesDataPoint(ts, 1, "splunk.queue.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordSplunkTypingQueueRatioDataPoint(ts, 1, "splunk.host-val")

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
				case "splunk.aggregation.queue.ratio":
					assert.False(t, validatedMetrics["splunk.aggregation.queue.ratio"], "Found a duplicate in the metrics slice: splunk.aggregation.queue.ratio")
					validatedMetrics["splunk.aggregation.queue.ratio"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the average indexer aggregation queue ration (%). *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "{%}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.buckets.searchable.status":
					assert.False(t, validatedMetrics["splunk.buckets.searchable.status"], "Found a duplicate in the metrics slice: splunk.buckets.searchable.status")
					validatedMetrics["splunk.buckets.searchable.status"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the number of buckets and their searchable status. *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "{count}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("splunk.indexer.searchable")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.indexer.searchable-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Count of buckets per index", ms.At(i).Description())
					assert.Equal(t, "{buckets}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.event.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.event.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.event.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.event.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Count of events in this bucket super-directory. *Note:** Must be pointed at specific indexer `endpoint`.", ms.At(i).Description())
					assert.Equal(t, "{events}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("splunk.bucket.dir")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.bucket.dir-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.hot.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.hot.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.hot.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.hot.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "(If size > 0) Number of hot buckets. *Note:** Must be pointed at specific indexer `endpoint`.", ms.At(i).Description())
					assert.Equal(t, "{buckets}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("splunk.bucket.dir")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.bucket.dir-val", attrVal.Str())
				case "splunk.data.indexes.extended.bucket.warm.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.bucket.warm.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.bucket.warm.count")
					validatedMetrics["splunk.data.indexes.extended.bucket.warm.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "(If size > 0) Number of warm buckets. *Note:** Must be pointed at specific indexer `endpoint` and gathers metrics from only that indexer.", ms.At(i).Description())
					assert.Equal(t, "{buckets}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("splunk.bucket.dir")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.bucket.dir-val", attrVal.Str())
				case "splunk.data.indexes.extended.event.count":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.event.count"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.event.count")
					validatedMetrics["splunk.data.indexes.extended.event.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Count of events for index, excluding frozen events. Approximately equal to the event_count sum of all buckets. *Note:** Must be pointed at specific indexer `endpoint` and gathers metrics from only that indexer.", ms.At(i).Description())
					assert.Equal(t, "{events}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.data.indexes.extended.raw.size":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.raw.size"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.raw.size")
					validatedMetrics["splunk.data.indexes.extended.raw.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Size in bytes on disk of the <bucket>/rawdata/ directories of all buckets in this index, excluding frozen *Note:** Must be pointed at specific indexer `endpoint` and gathers metrics from only that indexer.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.data.indexes.extended.total.size":
					assert.False(t, validatedMetrics["splunk.data.indexes.extended.total.size"], "Found a duplicate in the metrics slice: splunk.data.indexes.extended.total.size")
					validatedMetrics["splunk.data.indexes.extended.total.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Size in bytes on disk of this index *Note:** Must be pointed at specific indexer `endpoint` and gathers metrics from only that indexer.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.indexer.avg.rate":
					assert.False(t, validatedMetrics["splunk.indexer.avg.rate"], "Found a duplicate in the metrics slice: splunk.indexer.avg.rate")
					validatedMetrics["splunk.indexer.avg.rate"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the average rate of indexed data. **Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "KBy", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.indexer.cpu.time":
					assert.False(t, validatedMetrics["splunk.indexer.cpu.time"], "Found a duplicate in the metrics slice: splunk.indexer.cpu.time")
					validatedMetrics["splunk.indexer.cpu.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the number of indexing process cpu seconds per instance", ms.At(i).Description())
					assert.Equal(t, "{s}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.indexer.queue.ratio":
					assert.False(t, validatedMetrics["splunk.indexer.queue.ratio"], "Found a duplicate in the metrics slice: splunk.indexer.queue.ratio")
					validatedMetrics["splunk.indexer.queue.ratio"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the average indexer index queue ration (%). *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "{%}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.indexer.raw.write.time":
					assert.False(t, validatedMetrics["splunk.indexer.raw.write.time"], "Found a duplicate in the metrics slice: splunk.indexer.raw.write.time")
					validatedMetrics["splunk.indexer.raw.write.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the number of raw write seconds per instance", ms.At(i).Description())
					assert.Equal(t, "{s}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.indexer.throughput":
					assert.False(t, validatedMetrics["splunk.indexer.throughput"], "Found a duplicate in the metrics slice: splunk.indexer.throughput")
					validatedMetrics["splunk.indexer.throughput"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking average bytes per second throughput of indexer. *Note:** Must be pointed at specific indexer `endpoint` and gathers metrics from only that indexer.", ms.At(i).Description())
					assert.Equal(t, "By/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.indexer.status")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.indexer.status-val", attrVal.Str())
				case "splunk.indexes.avg.size":
					assert.False(t, validatedMetrics["splunk.indexes.avg.size"], "Found a duplicate in the metrics slice: splunk.indexes.avg.size")
					validatedMetrics["splunk.indexes.avg.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the indexes and their average size (gb). *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "Gb", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.indexes.avg.usage":
					assert.False(t, validatedMetrics["splunk.indexes.avg.usage"], "Found a duplicate in the metrics slice: splunk.indexes.avg.usage")
					validatedMetrics["splunk.indexes.avg.usage"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the indexes and their average usage (%). *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "{%}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.indexes.bucket.count":
					assert.False(t, validatedMetrics["splunk.indexes.bucket.count"], "Found a duplicate in the metrics slice: splunk.indexes.bucket.count")
					validatedMetrics["splunk.indexes.bucket.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the indexes and their bucket counts. *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "{count}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.indexes.median.data.age":
					assert.False(t, validatedMetrics["splunk.indexes.median.data.age"], "Found a duplicate in the metrics slice: splunk.indexes.median.data.age")
					validatedMetrics["splunk.indexes.median.data.age"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the indexes and their median data age (days). *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "{days}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.indexes.size":
					assert.False(t, validatedMetrics["splunk.indexes.size"], "Found a duplicate in the metrics slice: splunk.indexes.size")
					validatedMetrics["splunk.indexes.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the indexes and their total size (gb). *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "Gb", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.io.avg.iops":
					assert.False(t, validatedMetrics["splunk.io.avg.iops"], "Found a duplicate in the metrics slice: splunk.io.avg.iops")
					validatedMetrics["splunk.io.avg.iops"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the average IOPs used per instance", ms.At(i).Description())
					assert.Equal(t, "{iops}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.license.index.usage":
					assert.False(t, validatedMetrics["splunk.license.index.usage"], "Found a duplicate in the metrics slice: splunk.license.index.usage")
					validatedMetrics["splunk.license.index.usage"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the indexed license usage per index", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.index.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.index.name-val", attrVal.Str())
				case "splunk.parse.queue.ratio":
					assert.False(t, validatedMetrics["splunk.parse.queue.ratio"], "Found a duplicate in the metrics slice: splunk.parse.queue.ratio")
					validatedMetrics["splunk.parse.queue.ratio"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the average indexer parser queue ration (%). *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "{%}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.pipeline.set.count":
					assert.False(t, validatedMetrics["splunk.pipeline.set.count"], "Found a duplicate in the metrics slice: splunk.pipeline.set.count")
					validatedMetrics["splunk.pipeline.set.count"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the number of pipeline sets per indexer. **Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "KBy", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.scheduler.avg.execution.latency":
					assert.False(t, validatedMetrics["splunk.scheduler.avg.execution.latency"], "Found a duplicate in the metrics slice: splunk.scheduler.avg.execution.latency")
					validatedMetrics["splunk.scheduler.avg.execution.latency"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the average execution latency of scheduled searches", ms.At(i).Description())
					assert.Equal(t, "{ms}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.scheduler.avg.run.time":
					assert.False(t, validatedMetrics["splunk.scheduler.avg.run.time"], "Found a duplicate in the metrics slice: splunk.scheduler.avg.run.time")
					validatedMetrics["splunk.scheduler.avg.run.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the average runtime of scheduled searches", ms.At(i).Description())
					assert.Equal(t, "{ms}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.scheduler.completion.ratio":
					assert.False(t, validatedMetrics["splunk.scheduler.completion.ratio"], "Found a duplicate in the metrics slice: splunk.scheduler.completion.ratio")
					validatedMetrics["splunk.scheduler.completion.ratio"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the ratio of completed to skipped scheduled searches", ms.At(i).Description())
					assert.Equal(t, "{%}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				case "splunk.server.introspection.queues.current":
					assert.False(t, validatedMetrics["splunk.server.introspection.queues.current"], "Found a duplicate in the metrics slice: splunk.server.introspection.queues.current")
					validatedMetrics["splunk.server.introspection.queues.current"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking current length of queue. *Note:** Must be pointed at specific indexer `endpoint` and gathers metrics from only that indexer.", ms.At(i).Description())
					assert.Equal(t, "{queues}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.queue.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.queue.name-val", attrVal.Str())
				case "splunk.server.introspection.queues.current.bytes":
					assert.False(t, validatedMetrics["splunk.server.introspection.queues.current.bytes"], "Found a duplicate in the metrics slice: splunk.server.introspection.queues.current.bytes")
					validatedMetrics["splunk.server.introspection.queues.current.bytes"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking current bytes waiting in queue. *Note:** Must be pointed at specific indexer `endpoint` and gathers metrics from only that indexer.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("splunk.queue.name")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.queue.name-val", attrVal.Str())
				case "splunk.typing.queue.ratio":
					assert.False(t, validatedMetrics["splunk.typing.queue.ratio"], "Found a duplicate in the metrics slice: splunk.typing.queue.ratio")
					validatedMetrics["splunk.typing.queue.ratio"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Gauge tracking the average indexer typing queue ration (%). *Note:** Search is best run against a Cluster Manager.", ms.At(i).Description())
					assert.Equal(t, "{%}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.InDelta(t, float64(1), dp.DoubleValue(), 0.01)
					attrVal, ok := dp.Attributes().Get("splunk.host")
					assert.True(t, ok)
					assert.EqualValues(t, "splunk.host-val", attrVal.Str())
				}
			}
		})
	}
}
