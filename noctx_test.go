package noctx_test

import (
	"testing"

	"github.com/sonatard/noctx"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{desc: "crypto_tls"},
		{desc: "exec_cmd"},
		{desc: "http_client"},
		{desc: "http_request"},
		{desc: "network"},
		{desc: "slog"},
		{desc: "sql"},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			analysistest.Run(t, analysistest.TestData(), noctx.Analyzer, test.desc)
		})
	}
}

func TestAnalyzerWithExclusions(t *testing.T) {
	// Test with slog functions excluded
	t.Run("exclude_slog", func(t *testing.T) {
		// Save current flag value and restore after test
		oldExclude := noctx.Analyzer.Flags.Lookup("exclude")
		if oldExclude != nil {
			defer func() {
				_ = noctx.Analyzer.Flags.Set("exclude", oldExclude.Value.String())
			}()
		}

		// Set exclusion patterns
		if err := noctx.Analyzer.Flags.Set("exclude", "log/slog.*,(*log/slog.Logger).*"); err != nil {
			t.Fatalf("Failed to set exclude flag: %v", err)
		}
		analysistest.Run(t, analysistest.TestData(), noctx.Analyzer, "slog_excluded")
	})

	// Test with partial exclusion
	t.Run("exclude_partial", func(t *testing.T) {
		// Save current flag value and restore after test
		oldExclude := noctx.Analyzer.Flags.Lookup("exclude")
		if oldExclude != nil {
			defer func() {
				_ = noctx.Analyzer.Flags.Set("exclude", oldExclude.Value.String())
			}()
		}

		// Only exclude Debug methods
		if err := noctx.Analyzer.Flags.Set("exclude", "log/slog.Debug,(*log/slog.Logger).Debug"); err != nil {
			t.Fatalf("Failed to set exclude flag: %v", err)
		}

		analysistest.Run(t, analysistest.TestData(), noctx.Analyzer, "slog_partial")
	})
}
