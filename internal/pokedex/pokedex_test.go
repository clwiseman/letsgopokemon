package pokedex

import (
	"context"
	"testing"

	"github.com/go-pg/pg/v9"
	gm "github.com/onsi/gomega"

	"github.com/clwiseman/letsgopokemon/internal/db_models"
)

func TestPokedex_ListGenerations(t *testing.T) {
	type expectation struct {
		want []db_models.Generation
		err  bool
	}
	tcs := map[string]struct {
		name     string
		expected expectation
	}{
		"Success": {
			expected: expectation{
				want: []db_models.Generation{
					{
						BaseModel: db_models.BaseModel{
							Id: 1,
						},
						Name: "Gen I",
					},
					{
						BaseModel: db_models.BaseModel{
							Id: 2,
						},
						Name: "Gen II",
					},
					{
						BaseModel: db_models.BaseModel{
							Id: 3,
						},
						Name: "Gen III",
					},
					{
						BaseModel: db_models.BaseModel{
							Id: 4,
						},
						Name: "Gen IV",
					},
					{
						BaseModel: db_models.BaseModel{
							Id: 5,
						},
						Name: "Gen V",
					},
					{
						BaseModel: db_models.BaseModel{
							Id: 6,
						},
						Name: "Gen VI",
					},
					{
						BaseModel: db_models.BaseModel{
							Id: 7,
						},
						Name: "Gen VII",
					},
				},
			},
		},
	}
	for scenario := range tcs {
		t.Run(scenario, func(t *testing.T) {
			ctx := context.Background()

			tx, err := t.db.Begin()
			gm.Expect(got).To(gm.BeNil())
			defer tx.Rollback()

			pgDB := pg.Connect(&pg.Options{
				User:     "postgres",
				Password: "",
				Database: "postgres",
			})
			defer pgDB.Close()

			pd := Pokedex{
				db: pgDB,
			}

			got, err := pd.ListGenerations(ctx)

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
