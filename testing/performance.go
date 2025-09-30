package testing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// BenchmarkResult represents a single benchmark measurement.
type BenchmarkResult struct {
	Name           string    `json:"name"`
	NsPerOp        float64   `json:"ns_per_op"`
	AllocsPerOp    uint64    `json:"allocs_per_op"`
	BytesPerOp     uint64    `json:"bytes_per_op"`
	Timestamp      time.Time `json:"timestamp"`
	GitCommit      string    `json:"git_commit,omitempty"`
	GitBranch      string    `json:"git_branch,omitempty"`
	GoVersion      string    `json:"go_version,omitempty"`
	OS             string    `json:"os,omitempty"`
	Arch           string    `json:"arch,omitempty"`
	CPUModel       string    `json:"cpu_model,omitempty"`
	MemoryTotal    uint64    `json:"memory_total,omitempty"`
	IterationCount int       `json:"iteration_count"`
}

// BenchmarkReport contains a collection of benchmark results.
type BenchmarkReport struct {
	Results   []BenchmarkResult `json:"results"`
	Timestamp time.Time         `json:"timestamp"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// PerformanceBaseline stores baseline performance expectations.
type PerformanceBaseline struct {
	Name        string  `json:"name"`
	MaxNsPerOp  float64 `json:"max_ns_per_op"`
	MaxAllocs   uint64  `json:"max_allocs"`
	MaxBytes    uint64  `json:"max_bytes"`
	Description string  `json:"description,omitempty"`
}

// RegressionThresholds defines acceptable performance degradation.
type RegressionThresholds struct {
	TimeIncrease       float64 `json:"time_increase"`                    // e.g., 1.1 = 10% slower is acceptable
	AllocIncrease      float64 `json:"alloc_increase"`                   // e.g., 1.2 = 20% more allocations
	BytesIncrease      float64 `json:"bytes_increase"`                   // e.g., 1.15 = 15% more bytes
	AbsoluteMaxNsPerOp float64 `json:"absolute_max_ns_per_op,omitempty"` // Hard limit
}

// DefaultRegressionThresholds returns sensible defaults for regression detection.
func DefaultRegressionThresholds() RegressionThresholds {
	return RegressionThresholds{
		TimeIncrease:  1.10, // 10% slower
		AllocIncrease: 1.20, // 20% more allocations
		BytesIncrease: 1.15, // 15% more memory
	}
}

// PerformanceRegression represents a detected performance regression.
type PerformanceRegression struct {
	BenchmarkName string  `json:"benchmark_name"`
	Metric        string  `json:"metric"` // "time", "allocs", "bytes"
	BaselineValue float64 `json:"baseline_value"`
	CurrentValue  float64 `json:"current_value"`
	Increase      float64 `json:"increase"` // Percentage increase
	Threshold     float64 `json:"threshold"`
	Severity      string  `json:"severity"` // "warning", "critical"
}

// RegressionReport contains all detected regressions.
type RegressionReport struct {
	Regressions []PerformanceRegression `json:"regressions"`
	Timestamp   time.Time               `json:"timestamp"`
	TotalChecks int                     `json:"total_checks"`
	Passed      int                     `json:"passed"`
	Failed      int                     `json:"failed"`
}

// PerformanceTracker tracks and analyzes benchmark results over time.
type PerformanceTracker struct {
	resultsDir string
	baselines  map[string]PerformanceBaseline
	thresholds RegressionThresholds
}

// NewPerformanceTracker creates a new performance tracker.
func NewPerformanceTracker(resultsDir string) *PerformanceTracker {
	return &PerformanceTracker{
		resultsDir: resultsDir,
		baselines:  make(map[string]PerformanceBaseline),
		thresholds: DefaultRegressionThresholds(),
	}
}

// SetThresholds updates the regression detection thresholds.
func (pt *PerformanceTracker) SetThresholds(thresholds RegressionThresholds) {
	pt.thresholds = thresholds
}

// AddBaseline adds or updates a performance baseline.
func (pt *PerformanceTracker) AddBaseline(baseline PerformanceBaseline) {
	pt.baselines[baseline.Name] = baseline
}

// LoadBaselines loads baselines from a JSON file.
func (pt *PerformanceTracker) LoadBaselines(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read baselines: %w", err)
	}

	var baselines []PerformanceBaseline
	if err := json.Unmarshal(data, &baselines); err != nil {
		return fmt.Errorf("failed to parse baselines: %w", err)
	}

	for _, baseline := range baselines {
		pt.baselines[baseline.Name] = baseline
	}

	return nil
}

// SaveBaselines saves current baselines to a JSON file.
func (pt *PerformanceTracker) SaveBaselines(path string) error {
	baselines := make([]PerformanceBaseline, 0, len(pt.baselines))
	for _, baseline := range pt.baselines {
		baselines = append(baselines, baseline)
	}

	data, err := json.MarshalIndent(baselines, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal baselines: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write baselines: %w", err)
	}

	return nil
}

// SaveReport saves a benchmark report to the results directory.
func (pt *PerformanceTracker) SaveReport(report BenchmarkReport) error {
	if err := os.MkdirAll(pt.resultsDir, 0o700); err != nil {
		return fmt.Errorf("failed to create results directory: %w", err)
	}

	timestamp := report.Timestamp.Format("20060102-150405")
	filename := filepath.Join(pt.resultsDir, fmt.Sprintf("benchmark-%s.json", timestamp))

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	if err := os.WriteFile(filename, data, 0o600); err != nil {
		return fmt.Errorf("failed to write report: %w", err)
	}

	return nil
}

// LoadLatestReport loads the most recent benchmark report.
func (pt *PerformanceTracker) LoadLatestReport() (*BenchmarkReport, error) {
	entries, err := os.ReadDir(pt.resultsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read results directory: %w", err)
	}

	var latestFile string
	var latestTime time.Time

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if info, err := entry.Info(); err == nil {
			if info.ModTime().After(latestTime) {
				latestTime = info.ModTime()
				latestFile = filepath.Join(pt.resultsDir, entry.Name())
			}
		}
	}

	if latestFile == "" {
		return nil, fmt.Errorf("no benchmark reports found")
	}

	data, err := os.ReadFile(latestFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read report: %w", err)
	}

	var report BenchmarkReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("failed to parse report: %w", err)
	}

	return &report, nil
}

// CheckRegressions compares current results against baselines.
func (pt *PerformanceTracker) CheckRegressions(results []BenchmarkResult) RegressionReport {
	report := RegressionReport{
		Timestamp:   time.Now(),
		TotalChecks: len(results),
	}

	for _, result := range results {
		baseline, exists := pt.baselines[result.Name]
		if !exists {
			continue
		}

		// Check execution time
		if baseline.MaxNsPerOp > 0 {
			if result.NsPerOp > baseline.MaxNsPerOp*pt.thresholds.TimeIncrease {
				increase := (result.NsPerOp / baseline.MaxNsPerOp) - 1.0
				severity := "warning"
				if pt.thresholds.AbsoluteMaxNsPerOp > 0 && result.NsPerOp > pt.thresholds.AbsoluteMaxNsPerOp {
					severity = "critical"
				}

				report.Regressions = append(report.Regressions, PerformanceRegression{
					BenchmarkName: result.Name,
					Metric:        "time",
					BaselineValue: baseline.MaxNsPerOp,
					CurrentValue:  result.NsPerOp,
					Increase:      increase * 100,
					Threshold:     (pt.thresholds.TimeIncrease - 1.0) * 100,
					Severity:      severity,
				})
				report.Failed++
				continue
			}
		}

		// Check allocations
		if baseline.MaxAllocs > 0 {
			if result.AllocsPerOp > uint64(float64(baseline.MaxAllocs)*pt.thresholds.AllocIncrease) {
				increase := (float64(result.AllocsPerOp) / float64(baseline.MaxAllocs)) - 1.0

				report.Regressions = append(report.Regressions, PerformanceRegression{
					BenchmarkName: result.Name,
					Metric:        "allocs",
					BaselineValue: float64(baseline.MaxAllocs),
					CurrentValue:  float64(result.AllocsPerOp),
					Increase:      increase * 100,
					Threshold:     (pt.thresholds.AllocIncrease - 1.0) * 100,
					Severity:      "warning",
				})
				report.Failed++
				continue
			}
		}

		// Check memory usage
		if baseline.MaxBytes > 0 {
			if result.BytesPerOp > uint64(float64(baseline.MaxBytes)*pt.thresholds.BytesIncrease) {
				increase := (float64(result.BytesPerOp) / float64(baseline.MaxBytes)) - 1.0

				report.Regressions = append(report.Regressions, PerformanceRegression{
					BenchmarkName: result.Name,
					Metric:        "bytes",
					BaselineValue: float64(baseline.MaxBytes),
					CurrentValue:  float64(result.BytesPerOp),
					Increase:      increase * 100,
					Threshold:     (pt.thresholds.BytesIncrease - 1.0) * 100,
					Severity:      "warning",
				})
				report.Failed++
				continue
			}
		}

		report.Passed++
	}

	return report
}

// GenerateBaselineFromResults creates baseline expectations from benchmark results.
// Use a safety factor (e.g., 1.1) to allow some variance.
func (pt *PerformanceTracker) GenerateBaselineFromResults(results []BenchmarkResult, safetyFactor float64) {
	if safetyFactor < 1.0 {
		safetyFactor = 1.0
	}

	for _, result := range results {
		baseline := PerformanceBaseline{
			Name:        result.Name,
			MaxNsPerOp:  result.NsPerOp * safetyFactor,
			MaxAllocs:   uint64(float64(result.AllocsPerOp) * safetyFactor),
			MaxBytes:    uint64(float64(result.BytesPerOp) * safetyFactor),
			Description: fmt.Sprintf("Generated from benchmark run at %s", result.Timestamp.Format(time.RFC3339)),
		}
		pt.AddBaseline(baseline)
	}
}

// CompareReports compares two benchmark reports and returns performance changes.
func CompareReports(baseline, current BenchmarkReport) map[string]map[string]float64 {
	baselineMap := make(map[string]BenchmarkResult)
	for _, result := range baseline.Results {
		baselineMap[result.Name] = result
	}

	changes := make(map[string]map[string]float64)

	for _, currResult := range current.Results {
		baseResult, exists := baselineMap[currResult.Name]
		if !exists {
			continue
		}

		benchChanges := make(map[string]float64)

		// Calculate percentage changes
		if baseResult.NsPerOp > 0 {
			benchChanges["time_change"] = ((currResult.NsPerOp / baseResult.NsPerOp) - 1.0) * 100
		}
		if baseResult.AllocsPerOp > 0 {
			benchChanges["allocs_change"] = ((float64(currResult.AllocsPerOp) / float64(baseResult.AllocsPerOp)) - 1.0) * 100
		}
		if baseResult.BytesPerOp > 0 {
			benchChanges["bytes_change"] = ((float64(currResult.BytesPerOp) / float64(baseResult.BytesPerOp)) - 1.0) * 100
		}

		changes[currResult.Name] = benchChanges
	}

	return changes
}
