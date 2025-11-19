package drivers

import (
	"github.com/james-lawrence/genieql"
	"github.com/james-lawrence/genieql/internal/errorsx"
)

const (
	float32SliceExpr = "[]float32"
)

const pgvectorDecode = `func() {
	var tmp {{ .Column.Definition.ColumnType }}
	if err := tmp.Scan({{ .From | expr }}); err != nil {
		return err
	}
	{{ .To | autodereference | expr }} = {{ if .Column.Definition.Nullable }}func() *[]float32 { s := tmp.Slice(); return &s }(){{ else }}tmp.Slice(){{ end }}
}`

const pgvectorEncode = `func() {
	{{ if .Column.Definition.Nullable }}
	if {{ .From | localident | expr }} == nil {
		if err := {{ .To | expr }}.Set(nil); err != nil {
			{{ error "err" | ast }}
		}
	} else {
		v := pgvector.NewVector(*{{ .From | localident | expr }})
		val, err := v.Value()
		if err != nil {
			{{ error "err" | ast }}
		}
		if err := {{ .To | expr }}.Set(val); err != nil {
			{{ error "err" | ast }}
		}
	}
	{{ else }}
	v := pgvector.NewVector({{ .From | localident | expr }})
	val, err := v.Value()
	if err != nil {
		{{ error "err" | ast }}
	}
	if err := {{ .To | expr }}.Set(val); err != nil {
		{{ error "err" | ast }}
	}
	{{ end }}
}`

const pghalfvectorDecode = `func() {
	var tmp {{ .Column.Definition.ColumnType }}
	if err := tmp.Scan({{ .From | expr }}); err != nil {
		return err
	}
	{{ .To | autodereference | expr }} = {{ if .Column.Definition.Nullable }}func() *[]float32 { s := tmp.Slice(); return &s }(){{ else }}tmp.Slice(){{ end }}
}`

const pghalfvectorEncode = `func() {
	{{ if .Column.Definition.Nullable }}
	if {{ .From | localident | expr }} == nil {
		if err := {{ .To | expr }}.Set(nil); err != nil {
			{{ error "err" | ast }}
		}
	} else {
		v := pgvector.NewHalfVector(*{{ .From | localident | expr }})
		val, err := v.Value()
		if err != nil {
			{{ error "err" | ast }}
		}
		if err := {{ .To | expr }}.Set(val); err != nil {
			{{ error "err" | ast }}
		}
	}
	{{ else }}
	v := pgvector.NewHalfVector({{ .From | localident | expr }})
	val, err := v.Value()
	if err != nil {
		{{ error "err" | ast }}
	}
	if err := {{ .To | expr }}.Set(val); err != nil {
		{{ error "err" | ast }}
	}
	{{ end }}
}`

const pgsparsevectorDecode = `func() {
	var tmp {{ .Column.Definition.ColumnType }}
	if err := tmp.Scan({{ .From | expr }}); err != nil {
		return err
	}
	{{ .To | autodereference | expr }} = {{ if .Column.Definition.Nullable }}func() *[]float32 { s := tmp.ToSlice(); return &s }(){{ else }}tmp.ToSlice(){{ end }}
}`

const pgsparsevectorEncode = `func() {
	{{ if .Column.Definition.Nullable }}
	if {{ .From | localident | expr }} == nil {
		if err := {{ .To | expr }}.Set(nil); err != nil {
			{{ error "err" | ast }}
		}
	} else {
		v := pgvector.NewSparseVector(*{{ .From | localident | expr }})
		val, err := v.Value()
		if err != nil {
			{{ error "err" | ast }}
		}
		if err := {{ .To | expr }}.Set(val); err != nil {
			{{ error "err" | ast }}
		}
	}
	{{ else }}
	v := pgvector.NewSparseVector({{ .From | localident | expr }})
	val, err := v.Value()
	if err != nil {
		{{ error "err" | ast }}
	}
	if err := {{ .To | expr }}.Set(val); err != nil {
		{{ error "err" | ast }}
	}
	{{ end }}
}`

var pgvectorTypes = []genieql.ColumnDefinition{
	{
		Type:       "pgvector.Vector",
		Native:     float32SliceExpr,
		ColumnType: "pgvector.Vector",
		DBTypeName: "vector",
		Decode:     pgvectorDecode,
		Encode:     pgvectorEncode,
	},
	{
		Type:       "pgvector.HalfVector",
		Native:     float32SliceExpr,
		ColumnType: "pgvector.HalfVector",
		DBTypeName: "halfvec",
		Decode:     pghalfvectorDecode,
		Encode:     pghalfvectorEncode,
	},
	{
		Type:       "pgvector.SparseVector",
		Native:     float32SliceExpr,
		ColumnType: "pgvector.SparseVector",
		DBTypeName: "sparsevec",
		Decode:     pgsparsevectorDecode,
		Encode:     pgsparsevectorEncode,
	},
}

func init() {
	errorsx.MaybePanic(genieql.RegisterDriver(PGVector, NewDriver("github.com/pgvector/pgvector-go", pgvectorTypes...)))
}

const PGVector = "github.com/pgvector/pgvector-go"
