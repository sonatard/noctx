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
		{desc: "sql"},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			analysistest.Run(t, analysistest.TestData(), noctx.Analyzer, test.desc)
		})
	}
}
