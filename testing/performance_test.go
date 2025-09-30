package testing

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPerformanceTracker_AddBaseline(t *testing.T) {
	tracker := NewPerformanceTracker("test-results")

	baseline := PerformanceBaseline{
		Name:       "BenchmarkCircuitBreaker",
		MaxNsPerOp: 1000,
		MaxAllocs:  5,
		MaxBytes:   512,
	}

	tracker.AddBaseline(baseline)

	if len(tracker.baselines) != 1 {
		t.Errorf("expected 1 baseline, got %d", len(tracker.baselines))
	}

	stored, exists := tracker.baselines["BenchmarkCircuitBreaker"]
	if !exists {
		t.Fatal("baseline not found")
	}

	if stored.MaxNsPerOp != 1000 {
		t.Errorf("expected MaxNsPerOp=1000, got %f", stored.MaxNsPerOp)
	}
}

func TestPerformanceTracker_SaveAndLoadBaselines(t *testing.T) {
	tempDir := t.TempDir()
	baselinesFile := filepath.Join(tempDir, "baselines.json")

	tracker := NewPerformanceTracker(tempDir)
	tracker.AddBaseline(PerformanceBaseline{
		Name:        "BenchmarkRetry",
		MaxNsPerOp:  2000,
		MaxAllocs:   10,
		MaxBytes:    1024,
		Description: "Retry benchmark baseline",
	})

	// Save baselines
	if err := tracker.SaveBaselines(baselinesFile); err != nil {
		t.Fatalf("failed to save baselines: %v", err)
	}

	// Create new tracker and load baselines
	tracker2 := NewPerformanceTracker(tempDir)
	if err := tracker2.LoadBaselines(baselinesFile); err != nil {
		t.Fatalf("failed to load baselines: %v", err)
	}

	baseline, exists := tracker2.baselines["BenchmarkRetry"]
	if !exists {
		t.Fatal("baseline not found after loading")
	}

	if baseline.MaxNsPerOp != 2000 {
		t.Errorf("expected MaxNsPerOp=2000, got %f", baseline.MaxNsPerOp)
	}
	if baseline.Description != "Retry benchmark baseline" {
		t.Errorf("description mismatch: %s", baseline.Description)
	}
}

func TestPerformanceTracker_SaveReport(t *testing.T) {
	tempDir := t.TempDir()
	tracker := NewPerformanceTracker(tempDir)

	report := BenchmarkReport{
		Timestamp: time.Now(),
		Results: []BenchmarkResult{
			{
				Name:        "BenchmarkTimeout",
				NsPerOp:     1500,
				AllocsPerOp: 8,
				BytesPerOp:  768,
				Timestamp:   time.Now(),
			},
		},
	}

	if err := tracker.SaveReport(report); err != nil {
		t.Fatalf("failed to save report: %v", err)
	}

	// Verify file was created
	entries, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("failed to read temp dir: %v", err)
	}

	found := false
	for _, entry := range entries {
		if !entry.IsDir() {
			found = true
			break
		}
	}

	if !found {
		t.Error("report file not created")
	}
}

func TestPerformanceTracker_LoadLatestReport(t *testing.T) {
	tempDir := t.TempDir()
	tracker := NewPerformanceTracker(tempDir)

	// Save multiple reports
	for i := 0; i < 3; i++ {
		report := BenchmarkReport{
			Timestamp: time.Now().Add(time.Duration(i) * time.Second),
			Results: []BenchmarkResult{
				{
					Name:           "BenchmarkRateLimit",
					NsPerOp:        float64(1000 + i*100),
					AllocsPerOp:    5,
					BytesPerOp:     512,
					Timestamp:      time.Now(),
					IterationCount: 1000,
				},
			},
		}
		if err := tracker.SaveReport(report); err != nil {
			t.Fatalf("failed to save report %d: %v", i, err)
		}
		time.Sleep(10 * time.Millisecond) // Ensure different timestamps
	}

	// Load latest
	latest, err := tracker.LoadLatestReport()
	if err != nil {
		t.Fatalf("failed to load latest report: %v", err)
	}

	if len(latest.Results) == 0 {
		t.Fatal("no results in latest report")
	}

	// Should be the last one saved
	if latest.Results[0].NsPerOp < 1000 {
		t.Error("loaded report doesn't appear to be the latest")
	}
}

func TestPerformanceTracker_CheckRegressions_NoRegression(t *testing.T) {
	tracker := NewPerformanceTracker("test-results")
	tracker.AddBaseline(PerformanceBaseline{
		Name:       "BenchmarkBulkhead",
		MaxNsPerOp: 2000,
		MaxAllocs:  10,
		MaxBytes:   1024,
	})

	results := []BenchmarkResult{
		{
			Name:        "BenchmarkBulkhead",
			NsPerOp:     1800, // Within threshold
			AllocsPerOp: 9,    // Within threshold
			BytesPerOp:  1000, // Within threshold
		},
	}

	report := tracker.CheckRegressions(results)

	if report.Failed != 0 {
		t.Errorf("expected 0 failures, got %d", report.Failed)
	}
	if report.Passed != 1 {
		t.Errorf("expected 1 pass, got %d", report.Passed)
	}
	if len(report.Regressions) != 0 {
		t.Errorf("expected 0 regressions, got %d", len(report.Regressions))
	}
}

func TestPerformanceTracker_CheckRegressions_TimeRegression(t *testing.T) {
	tracker := NewPerformanceTracker("test-results")
	tracker.AddBaseline(PerformanceBaseline{
		Name:       "BenchmarkFallback",
		MaxNsPerOp: 1000,
		MaxAllocs:  5,
		MaxBytes:   512,
	})

	results := []BenchmarkResult{
		{
			Name:        "BenchmarkFallback",
			NsPerOp:     1500, // 50% slower - exceeds 10% threshold
			AllocsPerOp: 5,
			BytesPerOp:  512,
		},
	}

	report := tracker.CheckRegressions(results)

	if report.Failed != 1 {
		t.Errorf("expected 1 failure, got %d", report.Failed)
	}
	if len(report.Regressions) != 1 {
		t.Fatalf("expected 1 regression, got %d", len(report.Regressions))
	}

	regression := report.Regressions[0]
	if regression.Metric != "time" {
		t.Errorf("expected time regression, got %s", regression.Metric)
	}
	if regression.Increase < 45 || regression.Increase > 55 {
		t.Errorf("expected ~50%% increase, got %.2f%%", regression.Increase)
	}
}

func TestPerformanceTracker_CheckRegressions_AllocRegression(t *testing.T) {
	tracker := NewPerformanceTracker("test-results")
	tracker.AddBaseline(PerformanceBaseline{
		Name:       "BenchmarkMetrics",
		MaxNsPerOp: 1000,
		MaxAllocs:  10,
		MaxBytes:   512,
	})

	results := []BenchmarkResult{
		{
			Name:        "BenchmarkMetrics",
			NsPerOp:     1000,
			AllocsPerOp: 15, // 50% more allocations - exceeds 20% threshold
			BytesPerOp:  512,
		},
	}

	report := tracker.CheckRegressions(results)

	if report.Failed != 1 {
		t.Errorf("expected 1 failure, got %d", report.Failed)
	}
	if len(report.Regressions) != 1 {
		t.Fatalf("expected 1 regression, got %d", len(report.Regressions))
	}

	regression := report.Regressions[0]
	if regression.Metric != "allocs" {
		t.Errorf("expected allocs regression, got %s", regression.Metric)
	}
}

func TestPerformanceTracker_CheckRegressions_BytesRegression(t *testing.T) {
	tracker := NewPerformanceTracker("test-results")
	tracker.AddBaseline(PerformanceBaseline{
		Name:       "BenchmarkChaos",
		MaxNsPerOp: 1000,
		MaxAllocs:  5,
		MaxBytes:   512,
	})

	results := []BenchmarkResult{
		{
			Name:        "BenchmarkChaos",
			NsPerOp:     1000,
			AllocsPerOp: 5,
			BytesPerOp:  700, // 37% more memory - exceeds 15% threshold
		},
	}

	report := tracker.CheckRegressions(results)

	if report.Failed != 1 {
		t.Errorf("expected 1 failure, got %d", report.Failed)
	}
	if len(report.Regressions) != 1 {
		t.Fatalf("expected 1 regression, got %d", len(report.Regressions))
	}

	regression := report.Regressions[0]
	if regression.Metric != "bytes" {
		t.Errorf("expected bytes regression, got %s", regression.Metric)
	}
}

func TestPerformanceTracker_CheckRegressions_MultipleRegressions(t *testing.T) {
	tracker := NewPerformanceTracker("test-results")
	tracker.AddBaseline(PerformanceBaseline{
		Name:       "BenchmarkAll",
		MaxNsPerOp: 1000,
		MaxAllocs:  10,
		MaxBytes:   512,
	})

	results := []BenchmarkResult{
		{
			Name:        "BenchmarkAll",
			NsPerOp:     1500, // Time regression
			AllocsPerOp: 15,   // Would regress but caught by time first
			BytesPerOp:  700,  // Would regress but caught by time first
		},
	}

	report := tracker.CheckRegressions(results)

	// Only first regression is caught (function continues after finding one)
	if report.Failed != 1 {
		t.Errorf("expected 1 failure, got %d", report.Failed)
	}
	if len(report.Regressions) != 1 {
		t.Errorf("expected 1 regression, got %d", len(report.Regressions))
	}
}

func TestPerformanceTracker_GenerateBaselineFromResults(t *testing.T) {
	tracker := NewPerformanceTracker("test-results")

	results := []BenchmarkResult{
		{
			Name:        "BenchmarkNew",
			NsPerOp:     1000,
			AllocsPerOp: 5,
			BytesPerOp:  512,
			Timestamp:   time.Now(),
		},
	}

	tracker.GenerateBaselineFromResults(results, 1.1)

	baseline, exists := tracker.baselines["BenchmarkNew"]
	if !exists {
		t.Fatal("baseline not generated")
	}

	expectedNs := 1000.0 * 1.1
	if baseline.MaxNsPerOp != expectedNs {
		t.Errorf("expected MaxNsPerOp=%.0f, got %.0f", expectedNs, baseline.MaxNsPerOp)
	}

	// uint64 conversion truncates, so 5 * 1.1 = 5.5 becomes 5
	if baseline.MaxAllocs != 5 {
		t.Errorf("expected MaxAllocs=5, got %d", baseline.MaxAllocs)
	}
}

func TestPerformanceTracker_SetThresholds(t *testing.T) {
	tracker := NewPerformanceTracker("test-results")

	customThresholds := RegressionThresholds{
		TimeIncrease:  1.05, // 5%
		AllocIncrease: 1.10, // 10%
		BytesIncrease: 1.08, // 8%
	}

	tracker.SetThresholds(customThresholds)

	if tracker.thresholds.TimeIncrease != 1.05 {
		t.Errorf("threshold not updated: %f", tracker.thresholds.TimeIncrease)
	}
}

func TestCompareReports(t *testing.T) {
	baseline := BenchmarkReport{
		Results: []BenchmarkResult{
			{
				Name:        "BenchmarkCompare",
				NsPerOp:     1000,
				AllocsPerOp: 10,
				BytesPerOp:  512,
			},
		},
	}

	current := BenchmarkReport{
		Results: []BenchmarkResult{
			{
				Name:        "BenchmarkCompare",
				NsPerOp:     1200, // 20% slower
				AllocsPerOp: 12,   // 20% more allocs
				BytesPerOp:  600,  // 17.2% more bytes
			},
		},
	}

	changes := CompareReports(baseline, current)

	benchChanges, exists := changes["BenchmarkCompare"]
	if !exists {
		t.Fatal("benchmark comparison not found")
	}

	timeChange := benchChanges["time_change"]
	if timeChange < 19 || timeChange > 21 {
		t.Errorf("expected ~20%% time change, got %.2f%%", timeChange)
	}

	allocsChange := benchChanges["allocs_change"]
	if allocsChange < 19 || allocsChange > 21 {
		t.Errorf("expected ~20%% allocs change, got %.2f%%", allocsChange)
	}
}

func TestDefaultRegressionThresholds(t *testing.T) {
	thresholds := DefaultRegressionThresholds()

	if thresholds.TimeIncrease != 1.10 {
		t.Errorf("unexpected default TimeIncrease: %f", thresholds.TimeIncrease)
	}
	if thresholds.AllocIncrease != 1.20 {
		t.Errorf("unexpected default AllocIncrease: %f", thresholds.AllocIncrease)
	}
	if thresholds.BytesIncrease != 1.15 {
		t.Errorf("unexpected default BytesIncrease: %f", thresholds.BytesIncrease)
	}
}
