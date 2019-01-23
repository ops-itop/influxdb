package querytest

import (
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/querytest"
	"github.com/influxdata/flux/stdlib/influxdata/influxdb/v1"
	"github.com/influxdata/influxdb/query/stdlib/influxdata/influxdb"
)

// FromInfluxJSONCompiler wraps a compiler and replaces all From operations with FromJSON
func FromInfluxJSONCompiler(c flux.Compiler, jsonFile string) *querytest.ReplaceSpecCompiler {
	return querytest.NewReplaceSpecCompiler(c, func(op *flux.Operation) {
		if op.Spec.Kind() == influxdb.FromKind {
			op.Spec = &v1.FromInfluxJSONOpSpec{
				File: jsonFile,
			}
		}
	})
}
