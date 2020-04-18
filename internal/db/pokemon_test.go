package db

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/clwiseman/letsgopokemon/internal/models"
	"github.com/jbowes/vice"
	gm "github.com/onsi/gomega"
)

func Test_ListGenerations(t *testing.T) {
	genSample := &models.Generation{
		ID:          "1",
		DisplayName: "Gen I",
	}

	type expectation struct {
		generations []*models.Generation
		err vice.Vice
	}

	tests := []struct {
		name     string
		expected expectation
	}{
		{
			name: "Successfully return generations",
			expected: expectation{
				generations: genSample,
			},
		},
		{
			name: "Failed to return generations",
			expected: expectation{
				err: vice.Internal,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			h := New(Config{DB: tx, 

			ctx := context.Background()

			gen, err := ListGenerations(ctx)

			if err != nil {
				gm.Expect(err).To(gm.Equal(tt.expected.err))
				return
			}

			gm.Expect(tt.expected.err).To(gm.BeNil())
			gm.Expect(gen[0]).To(gm.Equal(tt.expected.generations))
		})
	}
}
