package reports

import "github.com/fvsystem/gomark/internal/adapter"

type Data struct {
	Results            []adapter.TestResult
	ResulstTransformed any
}
