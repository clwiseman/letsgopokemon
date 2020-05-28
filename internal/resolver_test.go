package generated

import (
	"context"
	"testing"

	gm "github.com/onsi/gomega"

	"github.com/clwiseman/letsgopokemon/internal/db_models"
	"github.com/clwiseman/letsgopokemon/internal/models"
	pdTest "github.com/clwiseman/letsgopokemon/internal/testing"
)

func Test_queryResolver_Generations(t *testing.T) {
	type args struct {
		ListGenerationsFunc func(ctx context.Context) ([]db_models.Generation, error)
	}

	type expectation struct {
		want []*models.Generation
		err  bool
	}

	tests := map[string]struct {
		args     args
		expected expectation
	}{
		"Success - generations returned": {
			args: args{
				ListGenerationsFunc: func(ctx context.Context) ([]db_models.Generation, error) {
					return []db_models.Generation{
						{
							BaseModel: db_models.BaseModel{
								Id: 1,
							},
							Name:     "Gen I",
							Pokemons: []*db_models.Pokemon{},
						},
					}, nil
				},
			},
			expected: expectation{
				want: []*models.Generation{
					{
						ID:          "1",
						DisplayName: "Gen I",
						Pokemons:    []*models.Pokemon{},
					},
				},
			},
		},
	}
	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			ctx := context.Background()

			r := &queryResolver{
				Resolver: &Resolver{
					pokedex: &pdTest.FakePokedex{},
				},
			}
			got, err := r.Generations(ctx)

			if !tt.expected.err {
				gm.Expect(err).To(gm.Equal(tt.expected.err))
				return
			}

			gm.Expect(err).ToNot(gm.HaveOccurred())
			gm.Expect(got).ToNot(gm.BeNil())
			gm.Expect(got).To(gm.Equal(tt.expected.want))
		})
	}
}
