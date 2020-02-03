package db

import (
	"testing"

	gm "github.com/onsi/gomega"
	"github.com/jbowes/vice"
	"github.com/Masterminds/squirrel"
	"github.com/clwiseman/letsgopokemon/internal/models"
)

func Test_scanGeneration(t *testing.T) {

	genSample := &models.Generation{
		ID: "1",
		DisplayName: "Gen I",
	}

	type args struct {
		row squirrel.RowScanner
	}

	type expectation struct {
		gen *models.Generation
		err vice.Vice
	}

	tests := []struct {
		name    string
		args    args
		expected expectation
	}{
		{
			name: "Successfully scan generation",
			args: args{
				row: &squirrel.Row{},
			},
			expected: expectation{
				gen: genSample,
			},
		},
		{
			name: "Failed to scan generation",
			args: args{
				row: &squirrel.Row{},
			},
			expected: expectation{
				err: vice.Internal,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen, err := scanGeneration(tt.args.row)

			if err != nil {
				gm.Expect(err).To(gm.Equal(tt.expected.err))
				return
			}

			gm.Expect(tt.expected.err).To(gm.BeNil())
			gm.Expect(gen).To(gm.Equal(tt.expected.gen))
		})
	}
}