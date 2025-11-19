package drivers_test

import (
	"testing"

	"github.com/james-lawrence/genieql"
	. "github.com/james-lawrence/genieql/internal/drivers"
	"github.com/stretchr/testify/require"
)

func TestPGVector(t *testing.T) {
	t.Run("driver registration", func(t *testing.T) {
		driver, err := genieql.LookupDriver(PGVector)
		require.NoError(t, err)
		require.NotNil(t, driver)
	})

	t.Run("vector type lookup", func(t *testing.T) {
		driver, err := genieql.LookupDriver(PGVector)
		require.NoError(t, err)

		testCases := []struct {
			name           string
			typeName       string
			expectedColumn string
			expectedNative string
			shouldFail     bool
		}{
			{
				name:           "vector type",
				typeName:       "pgvector.Vector",
				expectedColumn: "pgvector.Vector",
				expectedNative: "[]float32",
			},
			{
				name:           "vector db type name",
				typeName:       "vector",
				expectedColumn: "pgvector.Vector",
				expectedNative: "[]float32",
			},
			{
				name:           "halfvec type",
				typeName:       "pgvector.HalfVector",
				expectedColumn: "pgvector.HalfVector",
				expectedNative: "[]float32",
			},
			{
				name:           "halfvec db type name",
				typeName:       "halfvec",
				expectedColumn: "pgvector.HalfVector",
				expectedNative: "[]float32",
			},
			{
				name:           "sparsevec type",
				typeName:       "pgvector.SparseVector",
				expectedColumn: "pgvector.SparseVector",
				expectedNative: "[]float32",
			},
			{
				name:           "sparsevec db type name",
				typeName:       "sparsevec",
				expectedColumn: "pgvector.SparseVector",
				expectedNative: "[]float32",
			},
			{
				name:       "unknown type",
				typeName:   "pgvector.Unknown",
				shouldFail: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				def, err := driver.LookupType(tc.typeName)
				if tc.shouldFail {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
				require.Equal(t, tc.expectedColumn, def.ColumnType)
				require.Equal(t, tc.expectedNative, def.Native)
			})
		}
	})

	t.Run("pgx driver includes pgvector types", func(t *testing.T) {
		driver, err := genieql.LookupDriver(PGX)
		require.NoError(t, err)

		testCases := []struct {
			name           string
			typeName       string
			expectedColumn string
		}{
			{
				name:           "vector db type name via pgx",
				typeName:       "vector",
				expectedColumn: "pgvector.Vector",
			},
			{
				name:           "vector type via pgx",
				typeName:       "pgvector.Vector",
				expectedColumn: "pgvector.Vector",
			},
			{
				name:           "halfvec db type name via pgx",
				typeName:       "halfvec",
				expectedColumn: "pgvector.HalfVector",
			},
			{
				name:           "halfvec type via pgx",
				typeName:       "pgvector.HalfVector",
				expectedColumn: "pgvector.HalfVector",
			},
			{
				name:           "sparsevec db type name via pgx",
				typeName:       "sparsevec",
				expectedColumn: "pgvector.SparseVector",
			},
			{
				name:           "sparsevec type via pgx",
				typeName:       "pgvector.SparseVector",
				expectedColumn: "pgvector.SparseVector",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				def, err := driver.LookupType(tc.typeName)
				require.NoError(t, err)
				require.Equal(t, tc.expectedColumn, def.ColumnType)
			})
		}
	})
}
