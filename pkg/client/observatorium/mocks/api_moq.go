// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	pV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"sync"
	"time"
)

// Ensure, that APIMock does implement API.
// If this is not the case, regenerate this file with moq.
var _ API = &APIMock{}

// APIMock is a mock implementation of API.
//
//	func TestSomethingThatUsesAPI(t *testing.T) {
//
//		// make and configure a mocked API
//		mockedAPI := &APIMock{
//			AlertManagersFunc: func(ctx context.Context) (pV1.AlertManagersResult, error) {
//				panic("mock out the AlertManagers method")
//			},
//			AlertsFunc: func(ctx context.Context) (pV1.AlertsResult, error) {
//				panic("mock out the Alerts method")
//			},
//			BuildinfoFunc: func(ctx context.Context) (pV1.BuildinfoResult, error) {
//				panic("mock out the Buildinfo method")
//			},
//			CleanTombstonesFunc: func(ctx context.Context) error {
//				panic("mock out the CleanTombstones method")
//			},
//			ConfigFunc: func(ctx context.Context) (pV1.ConfigResult, error) {
//				panic("mock out the Config method")
//			},
//			DeleteSeriesFunc: func(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) error {
//				panic("mock out the DeleteSeries method")
//			},
//			FlagsFunc: func(ctx context.Context) (pV1.FlagsResult, error) {
//				panic("mock out the Flags method")
//			},
//			LabelNamesFunc: func(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]string, pV1.Warnings, error) {
//				panic("mock out the LabelNames method")
//			},
//			LabelValuesFunc: func(ctx context.Context, label string, matches []string, startTime time.Time, endTime time.Time) (model.LabelValues, pV1.Warnings, error) {
//				panic("mock out the LabelValues method")
//			},
//			MetadataFunc: func(ctx context.Context, metric string, limit string) (map[string][]pV1.Metadata, error) {
//				panic("mock out the Metadata method")
//			},
//			QueryFunc: func(ctx context.Context, query string, ts time.Time, opts ...pV1.Option) (model.Value, pV1.Warnings, error) {
//				panic("mock out the Query method")
//			},
//			QueryExemplarsFunc: func(ctx context.Context, query string, startTime time.Time, endTime time.Time) ([]pV1.ExemplarQueryResult, error) {
//				panic("mock out the QueryExemplars method")
//			},
//			QueryRangeFunc: func(ctx context.Context, query string, r pV1.Range, opts ...pV1.Option) (model.Value, pV1.Warnings, error) {
//				panic("mock out the QueryRange method")
//			},
//			RulesFunc: func(ctx context.Context) (pV1.RulesResult, error) {
//				panic("mock out the Rules method")
//			},
//			RuntimeinfoFunc: func(ctx context.Context) (pV1.RuntimeinfoResult, error) {
//				panic("mock out the Runtimeinfo method")
//			},
//			SeriesFunc: func(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]model.LabelSet, pV1.Warnings, error) {
//				panic("mock out the Series method")
//			},
//			SnapshotFunc: func(ctx context.Context, skipHead bool) (pV1.SnapshotResult, error) {
//				panic("mock out the Snapshot method")
//			},
//			TSDBFunc: func(ctx context.Context) (pV1.TSDBResult, error) {
//				panic("mock out the TSDB method")
//			},
//			TargetsFunc: func(ctx context.Context) (pV1.TargetsResult, error) {
//				panic("mock out the Targets method")
//			},
//			TargetsMetadataFunc: func(ctx context.Context, matchTarget string, metric string, limit string) ([]pV1.MetricMetadata, error) {
//				panic("mock out the TargetsMetadata method")
//			},
//			WalReplayFunc: func(ctx context.Context) (pV1.WalReplayStatus, error) {
//				panic("mock out the WalReplay method")
//			},
//		}
//
//		// use mockedAPI in code that requires API
//		// and then make assertions.
//
//	}
type APIMock struct {
	// AlertManagersFunc mocks the AlertManagers method.
	AlertManagersFunc func(ctx context.Context) (pV1.AlertManagersResult, error)

	// AlertsFunc mocks the Alerts method.
	AlertsFunc func(ctx context.Context) (pV1.AlertsResult, error)

	// BuildinfoFunc mocks the Buildinfo method.
	BuildinfoFunc func(ctx context.Context) (pV1.BuildinfoResult, error)

	// CleanTombstonesFunc mocks the CleanTombstones method.
	CleanTombstonesFunc func(ctx context.Context) error

	// ConfigFunc mocks the Config method.
	ConfigFunc func(ctx context.Context) (pV1.ConfigResult, error)

	// DeleteSeriesFunc mocks the DeleteSeries method.
	DeleteSeriesFunc func(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) error

	// FlagsFunc mocks the Flags method.
	FlagsFunc func(ctx context.Context) (pV1.FlagsResult, error)

	// LabelNamesFunc mocks the LabelNames method.
	LabelNamesFunc func(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]string, pV1.Warnings, error)

	// LabelValuesFunc mocks the LabelValues method.
	LabelValuesFunc func(ctx context.Context, label string, matches []string, startTime time.Time, endTime time.Time) (model.LabelValues, pV1.Warnings, error)

	// MetadataFunc mocks the Metadata method.
	MetadataFunc func(ctx context.Context, metric string, limit string) (map[string][]pV1.Metadata, error)

	// QueryFunc mocks the Query method.
	QueryFunc func(ctx context.Context, query string, ts time.Time, opts ...pV1.Option) (model.Value, pV1.Warnings, error)

	// QueryExemplarsFunc mocks the QueryExemplars method.
	QueryExemplarsFunc func(ctx context.Context, query string, startTime time.Time, endTime time.Time) ([]pV1.ExemplarQueryResult, error)

	// QueryRangeFunc mocks the QueryRange method.
	QueryRangeFunc func(ctx context.Context, query string, r pV1.Range, opts ...pV1.Option) (model.Value, pV1.Warnings, error)

	// RulesFunc mocks the Rules method.
	RulesFunc func(ctx context.Context) (pV1.RulesResult, error)

	// RuntimeinfoFunc mocks the Runtimeinfo method.
	RuntimeinfoFunc func(ctx context.Context) (pV1.RuntimeinfoResult, error)

	// SeriesFunc mocks the Series method.
	SeriesFunc func(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]model.LabelSet, pV1.Warnings, error)

	// SnapshotFunc mocks the Snapshot method.
	SnapshotFunc func(ctx context.Context, skipHead bool) (pV1.SnapshotResult, error)

	// TSDBFunc mocks the TSDB method.
	TSDBFunc func(ctx context.Context) (pV1.TSDBResult, error)

	// TargetsFunc mocks the Targets method.
	TargetsFunc func(ctx context.Context) (pV1.TargetsResult, error)

	// TargetsMetadataFunc mocks the TargetsMetadata method.
	TargetsMetadataFunc func(ctx context.Context, matchTarget string, metric string, limit string) ([]pV1.MetricMetadata, error)

	// WalReplayFunc mocks the WalReplay method.
	WalReplayFunc func(ctx context.Context) (pV1.WalReplayStatus, error)

	// calls tracks calls to the methods.
	calls struct {
		// AlertManagers holds details about calls to the AlertManagers method.
		AlertManagers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Alerts holds details about calls to the Alerts method.
		Alerts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Buildinfo holds details about calls to the Buildinfo method.
		Buildinfo []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// CleanTombstones holds details about calls to the CleanTombstones method.
		CleanTombstones []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Config holds details about calls to the Config method.
		Config []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// DeleteSeries holds details about calls to the DeleteSeries method.
		DeleteSeries []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Matches is the matches argument value.
			Matches []string
			// StartTime is the startTime argument value.
			StartTime time.Time
			// EndTime is the endTime argument value.
			EndTime time.Time
		}
		// Flags holds details about calls to the Flags method.
		Flags []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// LabelNames holds details about calls to the LabelNames method.
		LabelNames []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Matches is the matches argument value.
			Matches []string
			// StartTime is the startTime argument value.
			StartTime time.Time
			// EndTime is the endTime argument value.
			EndTime time.Time
		}
		// LabelValues holds details about calls to the LabelValues method.
		LabelValues []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Label is the label argument value.
			Label string
			// Matches is the matches argument value.
			Matches []string
			// StartTime is the startTime argument value.
			StartTime time.Time
			// EndTime is the endTime argument value.
			EndTime time.Time
		}
		// Metadata holds details about calls to the Metadata method.
		Metadata []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Metric is the metric argument value.
			Metric string
			// Limit is the limit argument value.
			Limit string
		}
		// Query holds details about calls to the Query method.
		Query []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Ts is the ts argument value.
			Ts time.Time
			// Opts is the opts argument value.
			Opts []pV1.Option
		}
		// QueryExemplars holds details about calls to the QueryExemplars method.
		QueryExemplars []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// StartTime is the startTime argument value.
			StartTime time.Time
			// EndTime is the endTime argument value.
			EndTime time.Time
		}
		// QueryRange holds details about calls to the QueryRange method.
		QueryRange []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// R is the r argument value.
			R pV1.Range
			// Opts is the opts argument value.
			Opts []pV1.Option
		}
		// Rules holds details about calls to the Rules method.
		Rules []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Runtimeinfo holds details about calls to the Runtimeinfo method.
		Runtimeinfo []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Series holds details about calls to the Series method.
		Series []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Matches is the matches argument value.
			Matches []string
			// StartTime is the startTime argument value.
			StartTime time.Time
			// EndTime is the endTime argument value.
			EndTime time.Time
		}
		// Snapshot holds details about calls to the Snapshot method.
		Snapshot []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SkipHead is the skipHead argument value.
			SkipHead bool
		}
		// TSDB holds details about calls to the TSDB method.
		TSDB []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Targets holds details about calls to the Targets method.
		Targets []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// TargetsMetadata holds details about calls to the TargetsMetadata method.
		TargetsMetadata []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// MatchTarget is the matchTarget argument value.
			MatchTarget string
			// Metric is the metric argument value.
			Metric string
			// Limit is the limit argument value.
			Limit string
		}
		// WalReplay holds details about calls to the WalReplay method.
		WalReplay []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockAlertManagers   sync.RWMutex
	lockAlerts          sync.RWMutex
	lockBuildinfo       sync.RWMutex
	lockCleanTombstones sync.RWMutex
	lockConfig          sync.RWMutex
	lockDeleteSeries    sync.RWMutex
	lockFlags           sync.RWMutex
	lockLabelNames      sync.RWMutex
	lockLabelValues     sync.RWMutex
	lockMetadata        sync.RWMutex
	lockQuery           sync.RWMutex
	lockQueryExemplars  sync.RWMutex
	lockQueryRange      sync.RWMutex
	lockRules           sync.RWMutex
	lockRuntimeinfo     sync.RWMutex
	lockSeries          sync.RWMutex
	lockSnapshot        sync.RWMutex
	lockTSDB            sync.RWMutex
	lockTargets         sync.RWMutex
	lockTargetsMetadata sync.RWMutex
	lockWalReplay       sync.RWMutex
}

// AlertManagers calls AlertManagersFunc.
func (mock *APIMock) AlertManagers(ctx context.Context) (pV1.AlertManagersResult, error) {
	if mock.AlertManagersFunc == nil {
		panic("APIMock.AlertManagersFunc: method is nil but API.AlertManagers was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockAlertManagers.Lock()
	mock.calls.AlertManagers = append(mock.calls.AlertManagers, callInfo)
	mock.lockAlertManagers.Unlock()
	return mock.AlertManagersFunc(ctx)
}

// AlertManagersCalls gets all the calls that were made to AlertManagers.
// Check the length with:
//
//	len(mockedAPI.AlertManagersCalls())
func (mock *APIMock) AlertManagersCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockAlertManagers.RLock()
	calls = mock.calls.AlertManagers
	mock.lockAlertManagers.RUnlock()
	return calls
}

// Alerts calls AlertsFunc.
func (mock *APIMock) Alerts(ctx context.Context) (pV1.AlertsResult, error) {
	if mock.AlertsFunc == nil {
		panic("APIMock.AlertsFunc: method is nil but API.Alerts was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockAlerts.Lock()
	mock.calls.Alerts = append(mock.calls.Alerts, callInfo)
	mock.lockAlerts.Unlock()
	return mock.AlertsFunc(ctx)
}

// AlertsCalls gets all the calls that were made to Alerts.
// Check the length with:
//
//	len(mockedAPI.AlertsCalls())
func (mock *APIMock) AlertsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockAlerts.RLock()
	calls = mock.calls.Alerts
	mock.lockAlerts.RUnlock()
	return calls
}

// Buildinfo calls BuildinfoFunc.
func (mock *APIMock) Buildinfo(ctx context.Context) (pV1.BuildinfoResult, error) {
	if mock.BuildinfoFunc == nil {
		panic("APIMock.BuildinfoFunc: method is nil but API.Buildinfo was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockBuildinfo.Lock()
	mock.calls.Buildinfo = append(mock.calls.Buildinfo, callInfo)
	mock.lockBuildinfo.Unlock()
	return mock.BuildinfoFunc(ctx)
}

// BuildinfoCalls gets all the calls that were made to Buildinfo.
// Check the length with:
//
//	len(mockedAPI.BuildinfoCalls())
func (mock *APIMock) BuildinfoCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockBuildinfo.RLock()
	calls = mock.calls.Buildinfo
	mock.lockBuildinfo.RUnlock()
	return calls
}

// CleanTombstones calls CleanTombstonesFunc.
func (mock *APIMock) CleanTombstones(ctx context.Context) error {
	if mock.CleanTombstonesFunc == nil {
		panic("APIMock.CleanTombstonesFunc: method is nil but API.CleanTombstones was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCleanTombstones.Lock()
	mock.calls.CleanTombstones = append(mock.calls.CleanTombstones, callInfo)
	mock.lockCleanTombstones.Unlock()
	return mock.CleanTombstonesFunc(ctx)
}

// CleanTombstonesCalls gets all the calls that were made to CleanTombstones.
// Check the length with:
//
//	len(mockedAPI.CleanTombstonesCalls())
func (mock *APIMock) CleanTombstonesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockCleanTombstones.RLock()
	calls = mock.calls.CleanTombstones
	mock.lockCleanTombstones.RUnlock()
	return calls
}

// Config calls ConfigFunc.
func (mock *APIMock) Config(ctx context.Context) (pV1.ConfigResult, error) {
	if mock.ConfigFunc == nil {
		panic("APIMock.ConfigFunc: method is nil but API.Config was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockConfig.Lock()
	mock.calls.Config = append(mock.calls.Config, callInfo)
	mock.lockConfig.Unlock()
	return mock.ConfigFunc(ctx)
}

// ConfigCalls gets all the calls that were made to Config.
// Check the length with:
//
//	len(mockedAPI.ConfigCalls())
func (mock *APIMock) ConfigCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockConfig.RLock()
	calls = mock.calls.Config
	mock.lockConfig.RUnlock()
	return calls
}

// DeleteSeries calls DeleteSeriesFunc.
func (mock *APIMock) DeleteSeries(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) error {
	if mock.DeleteSeriesFunc == nil {
		panic("APIMock.DeleteSeriesFunc: method is nil but API.DeleteSeries was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Matches   []string
		StartTime time.Time
		EndTime   time.Time
	}{
		Ctx:       ctx,
		Matches:   matches,
		StartTime: startTime,
		EndTime:   endTime,
	}
	mock.lockDeleteSeries.Lock()
	mock.calls.DeleteSeries = append(mock.calls.DeleteSeries, callInfo)
	mock.lockDeleteSeries.Unlock()
	return mock.DeleteSeriesFunc(ctx, matches, startTime, endTime)
}

// DeleteSeriesCalls gets all the calls that were made to DeleteSeries.
// Check the length with:
//
//	len(mockedAPI.DeleteSeriesCalls())
func (mock *APIMock) DeleteSeriesCalls() []struct {
	Ctx       context.Context
	Matches   []string
	StartTime time.Time
	EndTime   time.Time
} {
	var calls []struct {
		Ctx       context.Context
		Matches   []string
		StartTime time.Time
		EndTime   time.Time
	}
	mock.lockDeleteSeries.RLock()
	calls = mock.calls.DeleteSeries
	mock.lockDeleteSeries.RUnlock()
	return calls
}

// Flags calls FlagsFunc.
func (mock *APIMock) Flags(ctx context.Context) (pV1.FlagsResult, error) {
	if mock.FlagsFunc == nil {
		panic("APIMock.FlagsFunc: method is nil but API.Flags was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFlags.Lock()
	mock.calls.Flags = append(mock.calls.Flags, callInfo)
	mock.lockFlags.Unlock()
	return mock.FlagsFunc(ctx)
}

// FlagsCalls gets all the calls that were made to Flags.
// Check the length with:
//
//	len(mockedAPI.FlagsCalls())
func (mock *APIMock) FlagsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFlags.RLock()
	calls = mock.calls.Flags
	mock.lockFlags.RUnlock()
	return calls
}

// LabelNames calls LabelNamesFunc.
func (mock *APIMock) LabelNames(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]string, pV1.Warnings, error) {
	if mock.LabelNamesFunc == nil {
		panic("APIMock.LabelNamesFunc: method is nil but API.LabelNames was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Matches   []string
		StartTime time.Time
		EndTime   time.Time
	}{
		Ctx:       ctx,
		Matches:   matches,
		StartTime: startTime,
		EndTime:   endTime,
	}
	mock.lockLabelNames.Lock()
	mock.calls.LabelNames = append(mock.calls.LabelNames, callInfo)
	mock.lockLabelNames.Unlock()
	return mock.LabelNamesFunc(ctx, matches, startTime, endTime)
}

// LabelNamesCalls gets all the calls that were made to LabelNames.
// Check the length with:
//
//	len(mockedAPI.LabelNamesCalls())
func (mock *APIMock) LabelNamesCalls() []struct {
	Ctx       context.Context
	Matches   []string
	StartTime time.Time
	EndTime   time.Time
} {
	var calls []struct {
		Ctx       context.Context
		Matches   []string
		StartTime time.Time
		EndTime   time.Time
	}
	mock.lockLabelNames.RLock()
	calls = mock.calls.LabelNames
	mock.lockLabelNames.RUnlock()
	return calls
}

// LabelValues calls LabelValuesFunc.
func (mock *APIMock) LabelValues(ctx context.Context, label string, matches []string, startTime time.Time, endTime time.Time) (model.LabelValues, pV1.Warnings, error) {
	if mock.LabelValuesFunc == nil {
		panic("APIMock.LabelValuesFunc: method is nil but API.LabelValues was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Label     string
		Matches   []string
		StartTime time.Time
		EndTime   time.Time
	}{
		Ctx:       ctx,
		Label:     label,
		Matches:   matches,
		StartTime: startTime,
		EndTime:   endTime,
	}
	mock.lockLabelValues.Lock()
	mock.calls.LabelValues = append(mock.calls.LabelValues, callInfo)
	mock.lockLabelValues.Unlock()
	return mock.LabelValuesFunc(ctx, label, matches, startTime, endTime)
}

// LabelValuesCalls gets all the calls that were made to LabelValues.
// Check the length with:
//
//	len(mockedAPI.LabelValuesCalls())
func (mock *APIMock) LabelValuesCalls() []struct {
	Ctx       context.Context
	Label     string
	Matches   []string
	StartTime time.Time
	EndTime   time.Time
} {
	var calls []struct {
		Ctx       context.Context
		Label     string
		Matches   []string
		StartTime time.Time
		EndTime   time.Time
	}
	mock.lockLabelValues.RLock()
	calls = mock.calls.LabelValues
	mock.lockLabelValues.RUnlock()
	return calls
}

// Metadata calls MetadataFunc.
func (mock *APIMock) Metadata(ctx context.Context, metric string, limit string) (map[string][]pV1.Metadata, error) {
	if mock.MetadataFunc == nil {
		panic("APIMock.MetadataFunc: method is nil but API.Metadata was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Metric string
		Limit  string
	}{
		Ctx:    ctx,
		Metric: metric,
		Limit:  limit,
	}
	mock.lockMetadata.Lock()
	mock.calls.Metadata = append(mock.calls.Metadata, callInfo)
	mock.lockMetadata.Unlock()
	return mock.MetadataFunc(ctx, metric, limit)
}

// MetadataCalls gets all the calls that were made to Metadata.
// Check the length with:
//
//	len(mockedAPI.MetadataCalls())
func (mock *APIMock) MetadataCalls() []struct {
	Ctx    context.Context
	Metric string
	Limit  string
} {
	var calls []struct {
		Ctx    context.Context
		Metric string
		Limit  string
	}
	mock.lockMetadata.RLock()
	calls = mock.calls.Metadata
	mock.lockMetadata.RUnlock()
	return calls
}

// Query calls QueryFunc.
func (mock *APIMock) Query(ctx context.Context, query string, ts time.Time, opts ...pV1.Option) (model.Value, pV1.Warnings, error) {
	if mock.QueryFunc == nil {
		panic("APIMock.QueryFunc: method is nil but API.Query was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Ts    time.Time
		Opts  []pV1.Option
	}{
		Ctx:   ctx,
		Query: query,
		Ts:    ts,
		Opts:  opts,
	}
	mock.lockQuery.Lock()
	mock.calls.Query = append(mock.calls.Query, callInfo)
	mock.lockQuery.Unlock()
	return mock.QueryFunc(ctx, query, ts, opts...)
}

// QueryCalls gets all the calls that were made to Query.
// Check the length with:
//
//	len(mockedAPI.QueryCalls())
func (mock *APIMock) QueryCalls() []struct {
	Ctx   context.Context
	Query string
	Ts    time.Time
	Opts  []pV1.Option
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Ts    time.Time
		Opts  []pV1.Option
	}
	mock.lockQuery.RLock()
	calls = mock.calls.Query
	mock.lockQuery.RUnlock()
	return calls
}

// QueryExemplars calls QueryExemplarsFunc.
func (mock *APIMock) QueryExemplars(ctx context.Context, query string, startTime time.Time, endTime time.Time) ([]pV1.ExemplarQueryResult, error) {
	if mock.QueryExemplarsFunc == nil {
		panic("APIMock.QueryExemplarsFunc: method is nil but API.QueryExemplars was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Query     string
		StartTime time.Time
		EndTime   time.Time
	}{
		Ctx:       ctx,
		Query:     query,
		StartTime: startTime,
		EndTime:   endTime,
	}
	mock.lockQueryExemplars.Lock()
	mock.calls.QueryExemplars = append(mock.calls.QueryExemplars, callInfo)
	mock.lockQueryExemplars.Unlock()
	return mock.QueryExemplarsFunc(ctx, query, startTime, endTime)
}

// QueryExemplarsCalls gets all the calls that were made to QueryExemplars.
// Check the length with:
//
//	len(mockedAPI.QueryExemplarsCalls())
func (mock *APIMock) QueryExemplarsCalls() []struct {
	Ctx       context.Context
	Query     string
	StartTime time.Time
	EndTime   time.Time
} {
	var calls []struct {
		Ctx       context.Context
		Query     string
		StartTime time.Time
		EndTime   time.Time
	}
	mock.lockQueryExemplars.RLock()
	calls = mock.calls.QueryExemplars
	mock.lockQueryExemplars.RUnlock()
	return calls
}

// QueryRange calls QueryRangeFunc.
func (mock *APIMock) QueryRange(ctx context.Context, query string, r pV1.Range, opts ...pV1.Option) (model.Value, pV1.Warnings, error) {
	if mock.QueryRangeFunc == nil {
		panic("APIMock.QueryRangeFunc: method is nil but API.QueryRange was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		R     pV1.Range
		Opts  []pV1.Option
	}{
		Ctx:   ctx,
		Query: query,
		R:     r,
		Opts:  opts,
	}
	mock.lockQueryRange.Lock()
	mock.calls.QueryRange = append(mock.calls.QueryRange, callInfo)
	mock.lockQueryRange.Unlock()
	return mock.QueryRangeFunc(ctx, query, r, opts...)
}

// QueryRangeCalls gets all the calls that were made to QueryRange.
// Check the length with:
//
//	len(mockedAPI.QueryRangeCalls())
func (mock *APIMock) QueryRangeCalls() []struct {
	Ctx   context.Context
	Query string
	R     pV1.Range
	Opts  []pV1.Option
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		R     pV1.Range
		Opts  []pV1.Option
	}
	mock.lockQueryRange.RLock()
	calls = mock.calls.QueryRange
	mock.lockQueryRange.RUnlock()
	return calls
}

// Rules calls RulesFunc.
func (mock *APIMock) Rules(ctx context.Context) (pV1.RulesResult, error) {
	if mock.RulesFunc == nil {
		panic("APIMock.RulesFunc: method is nil but API.Rules was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockRules.Lock()
	mock.calls.Rules = append(mock.calls.Rules, callInfo)
	mock.lockRules.Unlock()
	return mock.RulesFunc(ctx)
}

// RulesCalls gets all the calls that were made to Rules.
// Check the length with:
//
//	len(mockedAPI.RulesCalls())
func (mock *APIMock) RulesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockRules.RLock()
	calls = mock.calls.Rules
	mock.lockRules.RUnlock()
	return calls
}

// Runtimeinfo calls RuntimeinfoFunc.
func (mock *APIMock) Runtimeinfo(ctx context.Context) (pV1.RuntimeinfoResult, error) {
	if mock.RuntimeinfoFunc == nil {
		panic("APIMock.RuntimeinfoFunc: method is nil but API.Runtimeinfo was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockRuntimeinfo.Lock()
	mock.calls.Runtimeinfo = append(mock.calls.Runtimeinfo, callInfo)
	mock.lockRuntimeinfo.Unlock()
	return mock.RuntimeinfoFunc(ctx)
}

// RuntimeinfoCalls gets all the calls that were made to Runtimeinfo.
// Check the length with:
//
//	len(mockedAPI.RuntimeinfoCalls())
func (mock *APIMock) RuntimeinfoCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockRuntimeinfo.RLock()
	calls = mock.calls.Runtimeinfo
	mock.lockRuntimeinfo.RUnlock()
	return calls
}

// Series calls SeriesFunc.
func (mock *APIMock) Series(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]model.LabelSet, pV1.Warnings, error) {
	if mock.SeriesFunc == nil {
		panic("APIMock.SeriesFunc: method is nil but API.Series was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Matches   []string
		StartTime time.Time
		EndTime   time.Time
	}{
		Ctx:       ctx,
		Matches:   matches,
		StartTime: startTime,
		EndTime:   endTime,
	}
	mock.lockSeries.Lock()
	mock.calls.Series = append(mock.calls.Series, callInfo)
	mock.lockSeries.Unlock()
	return mock.SeriesFunc(ctx, matches, startTime, endTime)
}

// SeriesCalls gets all the calls that were made to Series.
// Check the length with:
//
//	len(mockedAPI.SeriesCalls())
func (mock *APIMock) SeriesCalls() []struct {
	Ctx       context.Context
	Matches   []string
	StartTime time.Time
	EndTime   time.Time
} {
	var calls []struct {
		Ctx       context.Context
		Matches   []string
		StartTime time.Time
		EndTime   time.Time
	}
	mock.lockSeries.RLock()
	calls = mock.calls.Series
	mock.lockSeries.RUnlock()
	return calls
}

// Snapshot calls SnapshotFunc.
func (mock *APIMock) Snapshot(ctx context.Context, skipHead bool) (pV1.SnapshotResult, error) {
	if mock.SnapshotFunc == nil {
		panic("APIMock.SnapshotFunc: method is nil but API.Snapshot was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		SkipHead bool
	}{
		Ctx:      ctx,
		SkipHead: skipHead,
	}
	mock.lockSnapshot.Lock()
	mock.calls.Snapshot = append(mock.calls.Snapshot, callInfo)
	mock.lockSnapshot.Unlock()
	return mock.SnapshotFunc(ctx, skipHead)
}

// SnapshotCalls gets all the calls that were made to Snapshot.
// Check the length with:
//
//	len(mockedAPI.SnapshotCalls())
func (mock *APIMock) SnapshotCalls() []struct {
	Ctx      context.Context
	SkipHead bool
} {
	var calls []struct {
		Ctx      context.Context
		SkipHead bool
	}
	mock.lockSnapshot.RLock()
	calls = mock.calls.Snapshot
	mock.lockSnapshot.RUnlock()
	return calls
}

// TSDB calls TSDBFunc.
func (mock *APIMock) TSDB(ctx context.Context) (pV1.TSDBResult, error) {
	if mock.TSDBFunc == nil {
		panic("APIMock.TSDBFunc: method is nil but API.TSDB was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockTSDB.Lock()
	mock.calls.TSDB = append(mock.calls.TSDB, callInfo)
	mock.lockTSDB.Unlock()
	return mock.TSDBFunc(ctx)
}

// TSDBCalls gets all the calls that were made to TSDB.
// Check the length with:
//
//	len(mockedAPI.TSDBCalls())
func (mock *APIMock) TSDBCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockTSDB.RLock()
	calls = mock.calls.TSDB
	mock.lockTSDB.RUnlock()
	return calls
}

// Targets calls TargetsFunc.
func (mock *APIMock) Targets(ctx context.Context) (pV1.TargetsResult, error) {
	if mock.TargetsFunc == nil {
		panic("APIMock.TargetsFunc: method is nil but API.Targets was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockTargets.Lock()
	mock.calls.Targets = append(mock.calls.Targets, callInfo)
	mock.lockTargets.Unlock()
	return mock.TargetsFunc(ctx)
}

// TargetsCalls gets all the calls that were made to Targets.
// Check the length with:
//
//	len(mockedAPI.TargetsCalls())
func (mock *APIMock) TargetsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockTargets.RLock()
	calls = mock.calls.Targets
	mock.lockTargets.RUnlock()
	return calls
}

// TargetsMetadata calls TargetsMetadataFunc.
func (mock *APIMock) TargetsMetadata(ctx context.Context, matchTarget string, metric string, limit string) ([]pV1.MetricMetadata, error) {
	if mock.TargetsMetadataFunc == nil {
		panic("APIMock.TargetsMetadataFunc: method is nil but API.TargetsMetadata was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		MatchTarget string
		Metric      string
		Limit       string
	}{
		Ctx:         ctx,
		MatchTarget: matchTarget,
		Metric:      metric,
		Limit:       limit,
	}
	mock.lockTargetsMetadata.Lock()
	mock.calls.TargetsMetadata = append(mock.calls.TargetsMetadata, callInfo)
	mock.lockTargetsMetadata.Unlock()
	return mock.TargetsMetadataFunc(ctx, matchTarget, metric, limit)
}

// TargetsMetadataCalls gets all the calls that were made to TargetsMetadata.
// Check the length with:
//
//	len(mockedAPI.TargetsMetadataCalls())
func (mock *APIMock) TargetsMetadataCalls() []struct {
	Ctx         context.Context
	MatchTarget string
	Metric      string
	Limit       string
} {
	var calls []struct {
		Ctx         context.Context
		MatchTarget string
		Metric      string
		Limit       string
	}
	mock.lockTargetsMetadata.RLock()
	calls = mock.calls.TargetsMetadata
	mock.lockTargetsMetadata.RUnlock()
	return calls
}

// WalReplay calls WalReplayFunc.
func (mock *APIMock) WalReplay(ctx context.Context) (pV1.WalReplayStatus, error) {
	if mock.WalReplayFunc == nil {
		panic("APIMock.WalReplayFunc: method is nil but API.WalReplay was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockWalReplay.Lock()
	mock.calls.WalReplay = append(mock.calls.WalReplay, callInfo)
	mock.lockWalReplay.Unlock()
	return mock.WalReplayFunc(ctx)
}

// WalReplayCalls gets all the calls that were made to WalReplay.
// Check the length with:
//
//	len(mockedAPI.WalReplayCalls())
func (mock *APIMock) WalReplayCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockWalReplay.RLock()
	calls = mock.calls.WalReplay
	mock.lockWalReplay.RUnlock()
	return calls
}