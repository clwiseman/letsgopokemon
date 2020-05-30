package db_models

import (
	"github.com/go-pg/pg/v9"
)

// InsertDefaults inserts the base data for generations and pokemon.
func InsertDefaults(db *pg.DB) error {

	err := insertDefaultGenerations(db)
	if err != nil {
		return err
	}

	err = insertPokemon(db)
	if err != nil {
		return err
	}

	// linkPokemonToGenerationsSQL := `
	// UPDATE pokemons SET generation =
	// 	(CASE
	// 		WHEN id <= 151 THEN 1
	// 		WHEN id >= 152 AND id <= 251 THEN 2
	// 		WHEN id >= 252 AND id <= 386 THEN 3
	// 		WHEN id >= 387 AND id <= 493 THEN 4
	// 		WHEN id >= 494 AND id <= 649 THEN 5
	// 		WHEN id >= 650 AND id <= 721 THEN 6
	// 		WHEN id >= 722 AND id <= 809 THEN 7
	// 	END)
	// ;`

	// _, err = db.Exec(linkPokemonToGenerationsSQL)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func insertDefaultGenerations(db *pg.DB) error {
	generationDefaults := []Generation{
		{
			BaseModel: BaseModel{
				Id: 1,
			},
			Name: "Gen I",
		},
		{
			BaseModel: BaseModel{
				Id: 2,
			},
			Name: "Gen II",
		},
		{
			BaseModel: BaseModel{
				Id: 3,
			},
			Name: "Gen III",
		},
		{
			BaseModel: BaseModel{
				Id: 4,
			},
			Name: "Gen IV",
		},
		{
			BaseModel: BaseModel{
				Id: 5,
			},
			Name: "Gen V",
		},
		{
			BaseModel: BaseModel{
				Id: 6,
			},
			Name: "Gen VI",
		},
		{
			BaseModel: BaseModel{
				Id: 7,
			},
			Name: "Gen VII",
		},
	}

	for _, gen := range generationDefaults {
		err := db.Insert(&gen)
		if err != nil {
			continue
		}
	}

	var generation Generation
	err := db.Model(&generation).First()
	if err != nil {
		return err
	}

	return nil
}

func insertPokemon(db *pg.DB) error {
	pokemonDefaults := []Pokemon{
		{
			BaseModel: BaseModel{
				Id: 1,
			},
			Name: "Bulbasaur",
		},
		{
			BaseModel: BaseModel{
				Id: 2,
			},
			Name: "Ivysaur",
		},
		{
			BaseModel: BaseModel{
				Id: 3,
			},
			Name: "Venusaur",
		},
		{
			BaseModel: BaseModel{
				Id: 4,
			},
			Name: "Charmander",
		},
		{
			BaseModel: BaseModel{
				Id: 5,
			},
			Name: "Charmeleon",
		},
		{
			BaseModel: BaseModel{
				Id: 6,
			},
			Name: "Charizard",
		},
		{
			BaseModel: BaseModel{
				Id: 7,
			},
			Name: "Squirtle",
		},
		{
			BaseModel: BaseModel{
				Id: 8,
			},
			Name: "Wartortle",
		},
		{
			BaseModel: BaseModel{
				Id: 9,
			},
			Name: "Blastoise",
		},
		{
			BaseModel: BaseModel{
				Id: 10,
			},
			Name: "Caterpie",
		},
		{
			BaseModel: BaseModel{
				Id: 11,
			},
			Name: "Metapod",
		},
		{
			BaseModel: BaseModel{
				Id: 12,
			},
			Name: "Butterfree",
		},
		{
			BaseModel: BaseModel{
				Id: 13,
			},
			Name: "Weedle",
		},
		{
			BaseModel: BaseModel{
				Id: 14,
			},
			Name: "Kakuna",
		},
		{
			BaseModel: BaseModel{
				Id: 15,
			},
			Name: "Beedrill",
		},
		{
			BaseModel: BaseModel{
				Id: 16,
			},
			Name: "Pidgey",
		},
		{
			BaseModel: BaseModel{
				Id: 17,
			},
			Name: "Pidgeotto",
		},
		{
			BaseModel: BaseModel{
				Id: 18,
			},
			Name: "Pidgeot",
		},
		{
			BaseModel: BaseModel{
				Id: 19,
			},
			Name: "Rattata",
		},
		{
			BaseModel: BaseModel{
				Id: 20,
			},
			Name: "Raticate",
		},
		{
			BaseModel: BaseModel{
				Id: 21,
			},
			Name: "Spearow",
		},
		{
			BaseModel: BaseModel{
				Id: 22,
			},
			Name: "Fearow",
		},
		{
			BaseModel: BaseModel{
				Id: 23,
			},
			Name: "Ekans",
		},
		{
			BaseModel: BaseModel{
				Id: 24,
			},
			Name: "Arbok",
		},
		{
			BaseModel: BaseModel{
				Id: 25,
			},
			Name: "Pikachu",
		},
		{
			BaseModel: BaseModel{
				Id: 26,
			},
			Name: "Raichu",
		},
		{
			BaseModel: BaseModel{
				Id: 27,
			},
			Name: "Sandshrew",
		},
		{
			BaseModel: BaseModel{
				Id: 28,
			},
			Name: "Sandslash",
		},
		{
			BaseModel: BaseModel{
				Id: 29,
			},
			Name: "Nidoran♀",
		},
		{
			BaseModel: BaseModel{
				Id: 30,
			},
			Name: "Nidorina",
		},
		{
			BaseModel: BaseModel{
				Id: 31,
			},
			Name: "Nidoqueen",
		},
		{
			BaseModel: BaseModel{
				Id: 32,
			},
			Name: "Nidoran♂",
		},
		{
			BaseModel: BaseModel{
				Id: 33,
			},
			Name: "Nidorino",
		},
		{
			BaseModel: BaseModel{
				Id: 34,
			},
			Name: "Nidoking",
		},
		{
			BaseModel: BaseModel{
				Id: 35,
			},
			Name: "Clefairy",
		},
		{
			BaseModel: BaseModel{
				Id: 36,
			},
			Name: "Clefable",
		},
		{
			BaseModel: BaseModel{
				Id: 37,
			},
			Name: "Vulpix",
		},
		{
			BaseModel: BaseModel{
				Id: 38,
			},
			Name: "Ninetales",
		},
		{
			BaseModel: BaseModel{
				Id: 39,
			},
			Name: "Jigglypuff",
		},
		{
			BaseModel: BaseModel{
				Id: 40,
			},
			Name: "Wigglytuff",
		},
		{
			BaseModel: BaseModel{
				Id: 41,
			},
			Name: "Zubat",
		},
		{
			BaseModel: BaseModel{
				Id: 42,
			},
			Name: "Golbat",
		},
		{
			BaseModel: BaseModel{
				Id: 43,
			},
			Name: "Oddish",
		},
		{
			BaseModel: BaseModel{
				Id: 44,
			},
			Name: "Gloom",
		},
		{
			BaseModel: BaseModel{
				Id: 45,
			},
			Name: "Vileplume",
		},
		{
			BaseModel: BaseModel{
				Id: 46,
			},
			Name: "Paras",
		},
		{
			BaseModel: BaseModel{
				Id: 47,
			},
			Name: "Parasect",
		},
		{
			BaseModel: BaseModel{
				Id: 48,
			},
			Name: "Venonat",
		},
		{
			BaseModel: BaseModel{
				Id: 49,
			},
			Name: "Venomoth",
		},
		{
			BaseModel: BaseModel{
				Id: 50,
			},
			Name: "Diglett",
		},
		{
			BaseModel: BaseModel{
				Id: 51,
			},
			Name: "Dugtrio",
		},
		{
			BaseModel: BaseModel{
				Id: 52,
			},
			Name: "Meowth",
		},
		{
			BaseModel: BaseModel{
				Id: 53,
			},
			Name: "Persian",
		},
		{
			BaseModel: BaseModel{
				Id: 54,
			},
			Name: "Psyduck",
		},
		{
			BaseModel: BaseModel{
				Id: 55,
			},
			Name: "Golduck",
		},
		{
			BaseModel: BaseModel{
				Id: 56,
			},
			Name: "Mankey",
		},
		{
			BaseModel: BaseModel{
				Id: 57,
			},
			Name: "Primeape",
		},
		{
			BaseModel: BaseModel{
				Id: 58,
			},
			Name: "Growlithe",
		},
		{
			BaseModel: BaseModel{
				Id: 59,
			},
			Name: "Arcanine",
		},
		{
			BaseModel: BaseModel{
				Id: 60,
			},
			Name: "Poliwag",
		},
		{
			BaseModel: BaseModel{
				Id: 61,
			},
			Name: "Poliwhirl",
		},
		{
			BaseModel: BaseModel{
				Id: 62,
			},
			Name: "Poliwrath",
		},
		{
			BaseModel: BaseModel{
				Id: 63,
			},
			Name: "Abra",
		},
		{
			BaseModel: BaseModel{
				Id: 64,
			},
			Name: "Kadabra",
		},
		{
			BaseModel: BaseModel{
				Id: 65,
			},
			Name: "Alakazam",
		},
		{
			BaseModel: BaseModel{
				Id: 66,
			},
			Name: "Machop",
		},
		{
			BaseModel: BaseModel{
				Id: 67,
			},
			Name: "Machoke",
		},
		{
			BaseModel: BaseModel{
				Id: 68,
			},
			Name: "Machamp",
		},
		{
			BaseModel: BaseModel{
				Id: 69,
			},
			Name: "Bellsprout",
		},
		{
			BaseModel: BaseModel{
				Id: 70,
			},
			Name: "Weepinbell",
		},
		{
			BaseModel: BaseModel{
				Id: 71,
			},
			Name: "Victreebel",
		},
		{
			BaseModel: BaseModel{
				Id: 72,
			},
			Name: "Tentacool",
		},
		{
			BaseModel: BaseModel{
				Id: 73,
			},
			Name: "Tentacruel",
		},
		{
			BaseModel: BaseModel{
				Id: 74,
			},
			Name: "Geodude",
		},
		{
			BaseModel: BaseModel{
				Id: 75,
			},
			Name: "Graveler",
		},
		{
			BaseModel: BaseModel{
				Id: 76,
			},
			Name: "Golem",
		},
		{
			BaseModel: BaseModel{
				Id: 77,
			},
			Name: "Ponyta",
		},
		{
			BaseModel: BaseModel{
				Id: 78,
			},
			Name: "Rapidash",
		},
		{
			BaseModel: BaseModel{
				Id: 79,
			},
			Name: "Slowpoke",
		},
		{
			BaseModel: BaseModel{
				Id: 80,
			},
			Name: "Slowbro",
		},
		{
			BaseModel: BaseModel{
				Id: 81,
			},
			Name: "Magnemite",
		},
		{
			BaseModel: BaseModel{
				Id: 82,
			},
			Name: "Magneton",
		},
		{
			BaseModel: BaseModel{
				Id: 83,
			},
			Name: "Farfetch''d",
		},
		{
			BaseModel: BaseModel{
				Id: 84,
			},
			Name: "Doduo",
		},
		{
			BaseModel: BaseModel{
				Id: 85,
			},
			Name: "Dodrio",
		},
		{
			BaseModel: BaseModel{
				Id: 86,
			},
			Name: "Seel",
		},
		{
			BaseModel: BaseModel{
				Id: 87,
			},
			Name: "Dewgong",
		},
		{
			BaseModel: BaseModel{
				Id: 88,
			},
			Name: "Grimer",
		},
		{
			BaseModel: BaseModel{
				Id: 89,
			},
			Name: "Muk",
		},
		{
			BaseModel: BaseModel{
				Id: 90,
			},
			Name: "Shellder",
		},
		{
			BaseModel: BaseModel{
				Id: 91,
			},
			Name: "Cloyster",
		},
		{
			BaseModel: BaseModel{
				Id: 92,
			},
			Name: "Gastly",
		},
		{
			BaseModel: BaseModel{
				Id: 93,
			},
			Name: "Haunter",
		},
		{
			BaseModel: BaseModel{
				Id: 94,
			},
			Name: "Gengar",
		},
		{
			BaseModel: BaseModel{
				Id: 95,
			},
			Name: "Onix",
		},
		{
			BaseModel: BaseModel{
				Id: 96,
			},
			Name: "Drowzee",
		},
		{
			BaseModel: BaseModel{
				Id: 97,
			},
			Name: "Hypno",
		},
		{
			BaseModel: BaseModel{
				Id: 98,
			},
			Name: "Krabby",
		},
		{
			BaseModel: BaseModel{
				Id: 99,
			},
			Name: "Kingler",
		},
		{
			BaseModel: BaseModel{
				Id: 100,
			},
			Name: "Voltorb",
		},
		{
			BaseModel: BaseModel{
				Id: 101,
			},
			Name: "Electrode",
		},
		{
			BaseModel: BaseModel{
				Id: 102,
			},
			Name: "Exeggcute",
		},
		{
			BaseModel: BaseModel{
				Id: 103,
			},
			Name: "Exeggutor",
		},
		{
			BaseModel: BaseModel{
				Id: 104,
			},
			Name: "Cubone",
		},
		{
			BaseModel: BaseModel{
				Id: 105,
			},
			Name: "Marowak",
		},
		{
			BaseModel: BaseModel{
				Id: 106,
			},
			Name: "Hitmonlee",
		},
		{
			BaseModel: BaseModel{
				Id: 107,
			},
			Name: "Hitmonchan",
		},
		{
			BaseModel: BaseModel{
				Id: 108,
			},
			Name: "Lickitung",
		},
		{
			BaseModel: BaseModel{
				Id: 109,
			},
			Name: "Koffing",
		},
		{
			BaseModel: BaseModel{
				Id: 110,
			},
			Name: "Weezing",
		},
		{
			BaseModel: BaseModel{
				Id: 111,
			},
			Name: "Rhyhorn",
		},
		{
			BaseModel: BaseModel{
				Id: 112,
			},
			Name: "Rhydon",
		},
		{
			BaseModel: BaseModel{
				Id: 113,
			},
			Name: "Chansey",
		},
		{
			BaseModel: BaseModel{
				Id: 114,
			},
			Name: "Tangela",
		},
		{
			BaseModel: BaseModel{
				Id: 115,
			},
			Name: "Kangaskhan",
		},
		{
			BaseModel: BaseModel{
				Id: 116,
			},
			Name: "Horsea",
		},
		{
			BaseModel: BaseModel{
				Id: 117,
			},
			Name: "Seadra",
		},
		{
			BaseModel: BaseModel{
				Id: 118,
			},
			Name: "Goldeen",
		},
		{
			BaseModel: BaseModel{
				Id: 119,
			},
			Name: "Seaking",
		},
		{
			BaseModel: BaseModel{
				Id: 120,
			},
			Name: "Staryu",
		},
		{
			BaseModel: BaseModel{
				Id: 121,
			},
			Name: "Starmie",
		},
		{
			BaseModel: BaseModel{
				Id: 122,
			},
			Name: "Mr. Mime",
		},
		{
			BaseModel: BaseModel{
				Id: 123,
			},
			Name: "Scyther",
		},
		{
			BaseModel: BaseModel{
				Id: 124,
			},
			Name: "Jynx",
		},
		{
			BaseModel: BaseModel{
				Id: 125,
			},
			Name: "Electabuzz",
		},
		{
			BaseModel: BaseModel{
				Id: 126,
			},
			Name: "Magmar",
		},
		{
			BaseModel: BaseModel{
				Id: 127,
			},
			Name: "Pinsir",
		},
		{
			BaseModel: BaseModel{
				Id: 128,
			},
			Name: "Tauros",
		},
		{
			BaseModel: BaseModel{
				Id: 129,
			},
			Name: "Magikarp",
		},
		{
			BaseModel: BaseModel{
				Id: 130,
			},
			Name: "Gyarados",
		},
		{
			BaseModel: BaseModel{
				Id: 131,
			},
			Name: "Lapras",
		},
		{
			BaseModel: BaseModel{
				Id: 132,
			},
			Name: "Ditto",
		},
		{
			BaseModel: BaseModel{
				Id: 133,
			},
			Name: "Eevee",
		},
		{
			BaseModel: BaseModel{
				Id: 134,
			},
			Name: "Vaporeon",
		},
		{
			BaseModel: BaseModel{
				Id: 135,
			},
			Name: "Jolteon",
		},
		{
			BaseModel: BaseModel{
				Id: 136,
			},
			Name: "Flareon",
		},
		{
			BaseModel: BaseModel{
				Id: 137,
			},
			Name: "Porygon",
		},
		{
			BaseModel: BaseModel{
				Id: 138,
			},
			Name: "Omanyte",
		},
		{
			BaseModel: BaseModel{
				Id: 139,
			},
			Name: "Omastar",
		},
		{
			BaseModel: BaseModel{
				Id: 140,
			},
			Name: "Kabuto",
		},
		{
			BaseModel: BaseModel{
				Id: 141,
			},
			Name: "Kabutops",
		},
		{
			BaseModel: BaseModel{
				Id: 142,
			},
			Name: "Aerodactyl",
		},
		{
			BaseModel: BaseModel{
				Id: 143,
			},
			Name: "Snorlax",
		},
		{
			BaseModel: BaseModel{
				Id: 144,
			},
			Name: "Articuno",
		},
		{
			BaseModel: BaseModel{
				Id: 145,
			},
			Name: "Zapdos",
		},
		{
			BaseModel: BaseModel{
				Id: 146,
			},
			Name: "Moltres",
		},
		{
			BaseModel: BaseModel{
				Id: 147,
			},
			Name: "Dratini",
		},
		{
			BaseModel: BaseModel{
				Id: 148,
			},
			Name: "Dragonair",
		},
		{
			BaseModel: BaseModel{
				Id: 149,
			},
			Name: "Dragonite",
		},
		{
			BaseModel: BaseModel{
				Id: 150,
			},
			Name: "Mewtwo",
		},
		{
			BaseModel: BaseModel{
				Id: 151,
			},
			Name: "Mew",
		},
		{
			BaseModel: BaseModel{
				Id: 152,
			},
			Name: "Chikorita",
		},
		{
			BaseModel: BaseModel{
				Id: 153,
			},
			Name: "Bayleef",
		},
		{
			BaseModel: BaseModel{
				Id: 154,
			},
			Name: "Meganium",
		},
		{
			BaseModel: BaseModel{
				Id: 155,
			},
			Name: "Cyndaquil",
		},
		{
			BaseModel: BaseModel{
				Id: 156,
			},
			Name: "Quilava",
		},
		{
			BaseModel: BaseModel{
				Id: 157,
			},
			Name: "Typhlosion",
		},
		{
			BaseModel: BaseModel{
				Id: 158,
			},
			Name: "Totodile",
		},
		{
			BaseModel: BaseModel{
				Id: 159,
			},
			Name: "Croconaw",
		},
		{
			BaseModel: BaseModel{
				Id: 160,
			},
			Name: "Feraligatr",
		},
		{
			BaseModel: BaseModel{
				Id: 161,
			},
			Name: "Sentret",
		},
		{
			BaseModel: BaseModel{
				Id: 162,
			},
			Name: "Furret",
		},
		{
			BaseModel: BaseModel{
				Id: 163,
			},
			Name: "Hoothoot",
		},
		{
			BaseModel: BaseModel{
				Id: 164,
			},
			Name: "Noctowl",
		},
		{
			BaseModel: BaseModel{
				Id: 165,
			},
			Name: "Ledyba",
		},
		{
			BaseModel: BaseModel{
				Id: 166,
			},
			Name: "Ledian",
		},
		{
			BaseModel: BaseModel{
				Id: 167,
			},
			Name: "Spinarak",
		},
		{
			BaseModel: BaseModel{
				Id: 168,
			},
			Name: "Ariados",
		},
		{
			BaseModel: BaseModel{
				Id: 169,
			},
			Name: "Crobat",
		},
		{
			BaseModel: BaseModel{
				Id: 170,
			},
			Name: "Chinchou",
		},
		{
			BaseModel: BaseModel{
				Id: 171,
			},
			Name: "Lanturn",
		},
		{
			BaseModel: BaseModel{
				Id: 172,
			},
			Name: "Pichu",
		},
		{
			BaseModel: BaseModel{
				Id: 173,
			},
			Name: "Cleffa",
		},
		{
			BaseModel: BaseModel{
				Id: 174,
			},
			Name: "Igglybuff",
		},
		{
			BaseModel: BaseModel{
				Id: 175,
			},
			Name: "Togepi",
		},
		{
			BaseModel: BaseModel{
				Id: 176,
			},
			Name: "Togetic",
		},
		{
			BaseModel: BaseModel{
				Id: 177,
			},
			Name: "Natu",
		},
		{
			BaseModel: BaseModel{
				Id: 178,
			},
			Name: "Xatu",
		},
		{
			BaseModel: BaseModel{
				Id: 179,
			},
			Name: "Mareep",
		},
		{
			BaseModel: BaseModel{
				Id: 180,
			},
			Name: "Flaaffy",
		},
		{
			BaseModel: BaseModel{
				Id: 181,
			},
			Name: "Ampharos",
		},
		{
			BaseModel: BaseModel{
				Id: 182,
			},
			Name: "Bellossom",
		},
		{
			BaseModel: BaseModel{
				Id: 183,
			},
			Name: "Marill",
		},
		{
			BaseModel: BaseModel{
				Id: 184,
			},
			Name: "Azumarill",
		},
		{
			BaseModel: BaseModel{
				Id: 185,
			},
			Name: "Sudowoodo",
		},
		{
			BaseModel: BaseModel{
				Id: 186,
			},
			Name: "Politoed",
		},
		{
			BaseModel: BaseModel{
				Id: 187,
			},
			Name: "Hoppip",
		},
		{
			BaseModel: BaseModel{
				Id: 188,
			},
			Name: "Skiploom",
		},
		{
			BaseModel: BaseModel{
				Id: 189,
			},
			Name: "Jumpluff",
		},
		{
			BaseModel: BaseModel{
				Id: 190,
			},
			Name: "Aipom",
		},
		{
			BaseModel: BaseModel{
				Id: 191,
			},
			Name: "Sunkern",
		},
		{
			BaseModel: BaseModel{
				Id: 192,
			},
			Name: "Sunflora",
		},
		{
			BaseModel: BaseModel{
				Id: 193,
			},
			Name: "Yanma",
		},
		{
			BaseModel: BaseModel{
				Id: 194,
			},
			Name: "Wooper",
		},
		{
			BaseModel: BaseModel{
				Id: 195,
			},
			Name: "Quagsire",
		},
		{
			BaseModel: BaseModel{
				Id: 196,
			},
			Name: "Espeon",
		},
		{
			BaseModel: BaseModel{
				Id: 197,
			},
			Name: "Umbreon",
		},
		{
			BaseModel: BaseModel{
				Id: 198,
			},
			Name: "Murkrow",
		},
		{
			BaseModel: BaseModel{
				Id: 199,
			},
			Name: "Slowking",
		},
		{
			BaseModel: BaseModel{
				Id: 200,
			},
			Name: "Misdreavus",
		},
		{
			BaseModel: BaseModel{
				Id: 201,
			},
			Name: "Unown",
		},
		{
			BaseModel: BaseModel{
				Id: 202,
			},
			Name: "Wobbuffet",
		},
		{
			BaseModel: BaseModel{
				Id: 203,
			},
			Name: "Girafarig",
		},
		{
			BaseModel: BaseModel{
				Id: 204,
			},
			Name: "Pineco",
		},
		{
			BaseModel: BaseModel{
				Id: 205,
			},
			Name: "Forretress",
		},
		{
			BaseModel: BaseModel{
				Id: 206,
			},
			Name: "Dunsparce",
		},
		{
			BaseModel: BaseModel{
				Id: 207,
			},
			Name: "Gligar",
		},
		{
			BaseModel: BaseModel{
				Id: 208,
			},
			Name: "Steelix",
		},
		{
			BaseModel: BaseModel{
				Id: 209,
			},
			Name: "Snubbull",
		},
		{
			BaseModel: BaseModel{
				Id: 210,
			},
			Name: "Granbull",
		},
		{
			BaseModel: BaseModel{
				Id: 211,
			},
			Name: "Qwilfish",
		},
		{
			BaseModel: BaseModel{
				Id: 212,
			},
			Name: "Scizor",
		},
		{
			BaseModel: BaseModel{
				Id: 213,
			},
			Name: "Shuckle",
		},
		{
			BaseModel: BaseModel{
				Id: 214,
			},
			Name: "Heracross",
		},
		{
			BaseModel: BaseModel{
				Id: 215,
			},
			Name: "Sneasel",
		},
		{
			BaseModel: BaseModel{
				Id: 216,
			},
			Name: "Teddiursa",
		},
		{
			BaseModel: BaseModel{
				Id: 217,
			},
			Name: "Ursaring",
		},
		{
			BaseModel: BaseModel{
				Id: 218,
			},
			Name: "Slugma",
		},
		{
			BaseModel: BaseModel{
				Id: 219,
			},
			Name: "Magcargo",
		},
		{
			BaseModel: BaseModel{
				Id: 220,
			},
			Name: "Swinub",
		},
		{
			BaseModel: BaseModel{
				Id: 221,
			},
			Name: "Piloswine",
		},
		{
			BaseModel: BaseModel{
				Id: 222,
			},
			Name: "Corsola",
		},
		{
			BaseModel: BaseModel{
				Id: 223,
			},
			Name: "Remoraid",
		},
		{
			BaseModel: BaseModel{
				Id: 224,
			},
			Name: "Octillery",
		},
		{
			BaseModel: BaseModel{
				Id: 225,
			},
			Name: "Delibird",
		},
		{
			BaseModel: BaseModel{
				Id: 226,
			},
			Name: "Mantine",
		},
		{
			BaseModel: BaseModel{
				Id: 227,
			},
			Name: "Skarmory",
		},
		{
			BaseModel: BaseModel{
				Id: 228,
			},
			Name: "Houndour",
		},
		{
			BaseModel: BaseModel{
				Id: 229,
			},
			Name: "Houndoom",
		},
		{
			BaseModel: BaseModel{
				Id: 230,
			},
			Name: "Kingdra",
		},
		{
			BaseModel: BaseModel{
				Id: 231,
			},
			Name: "Phanpy",
		},
		{
			BaseModel: BaseModel{
				Id: 232,
			},
			Name: "Donphan",
		},
		{
			BaseModel: BaseModel{
				Id: 233,
			},
			Name: "Porygon2",
		},
		{
			BaseModel: BaseModel{
				Id: 234,
			},
			Name: "Stantler",
		},
		{
			BaseModel: BaseModel{
				Id: 235,
			},
			Name: "Smeargle",
		},
		{
			BaseModel: BaseModel{
				Id: 236,
			},
			Name: "Tyrogue",
		},
		{
			BaseModel: BaseModel{
				Id: 237,
			},
			Name: "Hitmontop",
		},
		{
			BaseModel: BaseModel{
				Id: 238,
			},
			Name: "Smoochum",
		},
		{
			BaseModel: BaseModel{
				Id: 239,
			},
			Name: "Elekid",
		},
		{
			BaseModel: BaseModel{
				Id: 240,
			},
			Name: "Magby",
		},
		{
			BaseModel: BaseModel{
				Id: 241,
			},
			Name: "Miltank",
		},
		{
			BaseModel: BaseModel{
				Id: 242,
			},
			Name: "Blissey",
		},
		{
			BaseModel: BaseModel{
				Id: 243,
			},
			Name: "Raikou",
		},
		{
			BaseModel: BaseModel{
				Id: 244,
			},
			Name: "Entei",
		},
		{
			BaseModel: BaseModel{
				Id: 245,
			},
			Name: "Suicune",
		},
		{
			BaseModel: BaseModel{
				Id: 246,
			},
			Name: "Larvitar",
		},
		{
			BaseModel: BaseModel{
				Id: 247,
			},
			Name: "Pupitar",
		},
		{
			BaseModel: BaseModel{
				Id: 248,
			},
			Name: "Tyranitar",
		},
		{
			BaseModel: BaseModel{
				Id: 249,
			},
			Name: "Lugia",
		},
		{
			BaseModel: BaseModel{
				Id: 250,
			},
			Name: "Ho-Oh",
		},
		{
			BaseModel: BaseModel{
				Id: 251,
			},
			Name: "Celebi",
		},
		{
			BaseModel: BaseModel{
				Id: 252,
			},
			Name: "Treecko",
		},
		{
			BaseModel: BaseModel{
				Id: 253,
			},
			Name: "Grovyle",
		},
		{
			BaseModel: BaseModel{
				Id: 254,
			},
			Name: "Sceptile",
		},
		{
			BaseModel: BaseModel{
				Id: 255,
			},
			Name: "Torchic",
		},
		{
			BaseModel: BaseModel{
				Id: 256,
			},
			Name: "Combusken",
		},
		{
			BaseModel: BaseModel{
				Id: 257,
			},
			Name: "Blaziken",
		},
		{
			BaseModel: BaseModel{
				Id: 258,
			},
			Name: "Mudkip",
		},
		{
			BaseModel: BaseModel{
				Id: 259,
			},
			Name: "Marshtomp",
		},
		{
			BaseModel: BaseModel{
				Id: 260,
			},
			Name: "Swampert",
		},
		{
			BaseModel: BaseModel{
				Id: 261,
			},
			Name: "Poochyena",
		},
		{
			BaseModel: BaseModel{
				Id: 262,
			},
			Name: "Mightyena",
		},
		{
			BaseModel: BaseModel{
				Id: 263,
			},
			Name: "Zigzagoon",
		},
		{
			BaseModel: BaseModel{
				Id: 264,
			},
			Name: "Linoone",
		},
		{
			BaseModel: BaseModel{
				Id: 265,
			},
			Name: "Wurmple",
		},
		{
			BaseModel: BaseModel{
				Id: 266,
			},
			Name: "Silcoon",
		},
		{
			BaseModel: BaseModel{
				Id: 267,
			},
			Name: "Beautifly",
		},
		{
			BaseModel: BaseModel{
				Id: 268,
			},
			Name: "Cascoon",
		},
		{
			BaseModel: BaseModel{
				Id: 269,
			},
			Name: "Dustox",
		},
		{
			BaseModel: BaseModel{
				Id: 270,
			},
			Name: "Lotad",
		},
		{
			BaseModel: BaseModel{
				Id: 271,
			},
			Name: "Lombre",
		},
		{
			BaseModel: BaseModel{
				Id: 272,
			},
			Name: "Ludicolo",
		},
		{
			BaseModel: BaseModel{
				Id: 273,
			},
			Name: "Seedot",
		},
		{
			BaseModel: BaseModel{
				Id: 274,
			},
			Name: "Nuzleaf",
		},
		{
			BaseModel: BaseModel{
				Id: 275,
			},
			Name: "Shiftry",
		},
		{
			BaseModel: BaseModel{
				Id: 276,
			},
			Name: "Taillow",
		},
		{
			BaseModel: BaseModel{
				Id: 277,
			},
			Name: "Swellow",
		},
		{
			BaseModel: BaseModel{
				Id: 278,
			},
			Name: "Wingull",
		},
		{
			BaseModel: BaseModel{
				Id: 279,
			},
			Name: "Pelipper",
		},
		{
			BaseModel: BaseModel{
				Id: 280,
			},
			Name: "Ralts",
		},
		{
			BaseModel: BaseModel{
				Id: 281,
			},
			Name: "Kirlia",
		},
		{
			BaseModel: BaseModel{
				Id: 282,
			},
			Name: "Gardevoir",
		},
		{
			BaseModel: BaseModel{
				Id: 283,
			},
			Name: "Surskit",
		},
		{
			BaseModel: BaseModel{
				Id: 284,
			},
			Name: "Masquerain",
		},
		{
			BaseModel: BaseModel{
				Id: 285,
			},
			Name: "Shroomish",
		},
		{
			BaseModel: BaseModel{
				Id: 286,
			},
			Name: "Breloom",
		},
		{
			BaseModel: BaseModel{
				Id: 287,
			},
			Name: "Slakoth",
		},
		{
			BaseModel: BaseModel{
				Id: 288,
			},
			Name: "Vigoroth",
		},
		{
			BaseModel: BaseModel{
				Id: 289,
			},
			Name: "Slaking",
		},
		{
			BaseModel: BaseModel{
				Id: 290,
			},
			Name: "Nincada",
		},
		{
			BaseModel: BaseModel{
				Id: 291,
			},
			Name: "Ninjask",
		},
		{
			BaseModel: BaseModel{
				Id: 292,
			},
			Name: "Shedinja",
		},
		{
			BaseModel: BaseModel{
				Id: 293,
			},
			Name: "Whismur",
		},
		{
			BaseModel: BaseModel{
				Id: 294,
			},
			Name: "Loudred",
		},
		{
			BaseModel: BaseModel{
				Id: 295,
			},
			Name: "Exploud",
		},
		{
			BaseModel: BaseModel{
				Id: 296,
			},
			Name: "Makuhita",
		},
		{
			BaseModel: BaseModel{
				Id: 297,
			},
			Name: "Hariyama",
		},
		{
			BaseModel: BaseModel{
				Id: 298,
			},
			Name: "Azurill",
		},
		{
			BaseModel: BaseModel{
				Id: 299,
			},
			Name: "Nosepass",
		},
		{
			BaseModel: BaseModel{
				Id: 300,
			},
			Name: "Skitty",
		},
		{
			BaseModel: BaseModel{
				Id: 301,
			},
			Name: "Delcatty",
		},
		{
			BaseModel: BaseModel{
				Id: 302,
			},
			Name: "Sableye",
		},
		{
			BaseModel: BaseModel{
				Id: 303,
			},
			Name: "Mawile",
		},
		{
			BaseModel: BaseModel{
				Id: 304,
			},
			Name: "Aron",
		},
		{
			BaseModel: BaseModel{
				Id: 305,
			},
			Name: "Lairon",
		},
		{
			BaseModel: BaseModel{
				Id: 306,
			},
			Name: "Aggron",
		},
		{
			BaseModel: BaseModel{
				Id: 307,
			},
			Name: "Meditite",
		},
		{
			BaseModel: BaseModel{
				Id: 308,
			},
			Name: "Medicham",
		},
		{
			BaseModel: BaseModel{
				Id: 309,
			},
			Name: "Electrike",
		},
		{
			BaseModel: BaseModel{
				Id: 310,
			},
			Name: "Manectric",
		},
		{
			BaseModel: BaseModel{
				Id: 311,
			},
			Name: "Plusle",
		},
		{
			BaseModel: BaseModel{
				Id: 312,
			},
			Name: "Minun",
		},
		{
			BaseModel: BaseModel{
				Id: 313,
			},
			Name: "Volbeat",
		},
		{
			BaseModel: BaseModel{
				Id: 314,
			},
			Name: "Illumise",
		},
		{
			BaseModel: BaseModel{
				Id: 315,
			},
			Name: "Roselia",
		},
		{
			BaseModel: BaseModel{
				Id: 316,
			},
			Name: "Gulpin",
		},
		{
			BaseModel: BaseModel{
				Id: 317,
			},
			Name: "Swalot",
		},
		{
			BaseModel: BaseModel{
				Id: 318,
			},
			Name: "Carvanha",
		},
		{
			BaseModel: BaseModel{
				Id: 319,
			},
			Name: "Sharpedo",
		},
		{
			BaseModel: BaseModel{
				Id: 320,
			},
			Name: "Wailmer",
		},
		{
			BaseModel: BaseModel{
				Id: 321,
			},
			Name: "Wailord",
		},
		{
			BaseModel: BaseModel{
				Id: 322,
			},
			Name: "Numel",
		},
		{
			BaseModel: BaseModel{
				Id: 323,
			},
			Name: "Camerupt",
		},
		{
			BaseModel: BaseModel{
				Id: 324,
			},
			Name: "Torkoal",
		},
		{
			BaseModel: BaseModel{
				Id: 325,
			},
			Name: "Spoink",
		},
		{
			BaseModel: BaseModel{
				Id: 326,
			},
			Name: "Grumpig",
		},
		{
			BaseModel: BaseModel{
				Id: 327,
			},
			Name: "Spinda",
		},
		{
			BaseModel: BaseModel{
				Id: 328,
			},
			Name: "Trapinch",
		},
		{
			BaseModel: BaseModel{
				Id: 329,
			},
			Name: "Vibrava",
		},
		{
			BaseModel: BaseModel{
				Id: 330,
			},
			Name: "Flygon",
		},
		{
			BaseModel: BaseModel{
				Id: 331,
			},
			Name: "Cacnea",
		},
		{
			BaseModel: BaseModel{
				Id: 332,
			},
			Name: "Cacturne",
		},
		{
			BaseModel: BaseModel{
				Id: 333,
			},
			Name: "Swablu",
		},
		{
			BaseModel: BaseModel{
				Id: 334,
			},
			Name: "Altaria",
		},
		{
			BaseModel: BaseModel{
				Id: 335,
			},
			Name: "Zangoose",
		},
		{
			BaseModel: BaseModel{
				Id: 336,
			},
			Name: "Seviper",
		},
		{
			BaseModel: BaseModel{
				Id: 337,
			},
			Name: "Lunatone",
		},
		{
			BaseModel: BaseModel{
				Id: 338,
			},
			Name: "Solrock",
		},
		{
			BaseModel: BaseModel{
				Id: 339,
			},
			Name: "Barboach",
		},
		{
			BaseModel: BaseModel{
				Id: 340,
			},
			Name: "Whiscash",
		},
		{
			BaseModel: BaseModel{
				Id: 341,
			},
			Name: "Corphish",
		},
		{
			BaseModel: BaseModel{
				Id: 342,
			},
			Name: "Crawdaunt",
		},
		{
			BaseModel: BaseModel{
				Id: 343,
			},
			Name: "Baltoy",
		},
		{
			BaseModel: BaseModel{
				Id: 344,
			},
			Name: "Claydol",
		},
		{
			BaseModel: BaseModel{
				Id: 345,
			},
			Name: "Lileep",
		},
		{
			BaseModel: BaseModel{
				Id: 346,
			},
			Name: "Cradily",
		},
		{
			BaseModel: BaseModel{
				Id: 347,
			},
			Name: "Anorith",
		},
		{
			BaseModel: BaseModel{
				Id: 348,
			},
			Name: "Armaldo",
		},
		{
			BaseModel: BaseModel{
				Id: 349,
			},
			Name: "Feebas",
		},
		{
			BaseModel: BaseModel{
				Id: 350,
			},
			Name: "Milotic",
		},
		{
			BaseModel: BaseModel{
				Id: 351,
			},
			Name: "Castform",
		},
		{
			BaseModel: BaseModel{
				Id: 352,
			},
			Name: "Kecleon",
		},
		{
			BaseModel: BaseModel{
				Id: 353,
			},
			Name: "Shuppet",
		},
		{
			BaseModel: BaseModel{
				Id: 354,
			},
			Name: "Banette",
		},
		{
			BaseModel: BaseModel{
				Id: 355,
			},
			Name: "Duskull",
		},
		{
			BaseModel: BaseModel{
				Id: 356,
			},
			Name: "Dusclops",
		},
		{
			BaseModel: BaseModel{
				Id: 357,
			},
			Name: "Tropius",
		},
		{
			BaseModel: BaseModel{
				Id: 358,
			},
			Name: "Chimecho",
		},
		{
			BaseModel: BaseModel{
				Id: 359,
			},
			Name: "Absol",
		},
		{
			BaseModel: BaseModel{
				Id: 360,
			},
			Name: "Wynaut",
		},
		{
			BaseModel: BaseModel{
				Id: 361,
			},
			Name: "Snorunt",
		},
		{
			BaseModel: BaseModel{
				Id: 362,
			},
			Name: "Glalie",
		},
		{
			BaseModel: BaseModel{
				Id: 363,
			},
			Name: "Spheal",
		},
		{
			BaseModel: BaseModel{
				Id: 364,
			},
			Name: "Sealeo",
		},
		{
			BaseModel: BaseModel{
				Id: 365,
			},
			Name: "Walrein",
		},
		{
			BaseModel: BaseModel{
				Id: 366,
			},
			Name: "Clamperl",
		},
		{
			BaseModel: BaseModel{
				Id: 367,
			},
			Name: "Huntail",
		},
		{
			BaseModel: BaseModel{
				Id: 368,
			},
			Name: "Gorebyss",
		},
		{
			BaseModel: BaseModel{
				Id: 369,
			},
			Name: "Relicanth",
		},
		{
			BaseModel: BaseModel{
				Id: 370,
			},
			Name: "Luvdisc",
		},
		{
			BaseModel: BaseModel{
				Id: 371,
			},
			Name: "Bagon",
		},
		{
			BaseModel: BaseModel{
				Id: 372,
			},
			Name: "Shelgon",
		},
		{
			BaseModel: BaseModel{
				Id: 373,
			},
			Name: "Salamence",
		},
		{
			BaseModel: BaseModel{
				Id: 374,
			},
			Name: "Beldum",
		},
		{
			BaseModel: BaseModel{
				Id: 375,
			},
			Name: "Metang",
		},
		{
			BaseModel: BaseModel{
				Id: 376,
			},
			Name: "Metagross",
		},
		{
			BaseModel: BaseModel{
				Id: 377,
			},
			Name: "Regirock",
		},
		{
			BaseModel: BaseModel{
				Id: 378,
			},
			Name: "Regice",
		},
		{
			BaseModel: BaseModel{
				Id: 379,
			},
			Name: "Registeel",
		},
		{
			BaseModel: BaseModel{
				Id: 380,
			},
			Name: "Latias",
		},
		{
			BaseModel: BaseModel{
				Id: 381,
			},
			Name: "Latios",
		},
		{
			BaseModel: BaseModel{
				Id: 382,
			},
			Name: "Kyogre",
		},
		{
			BaseModel: BaseModel{
				Id: 383,
			},
			Name: "Groudon",
		},
		{
			BaseModel: BaseModel{
				Id: 384,
			},
			Name: "Rayquaza",
		},
		{
			BaseModel: BaseModel{
				Id: 385,
			},
			Name: "Jirachi",
		},
		{
			BaseModel: BaseModel{
				Id: 386,
			},
			Name: "Deoxys",
		},
		{
			BaseModel: BaseModel{
				Id: 387,
			},
			Name: "Turtwig",
		},
		{
			BaseModel: BaseModel{
				Id: 388,
			},
			Name: "Grotle",
		},
		{
			BaseModel: BaseModel{
				Id: 389,
			},
			Name: "Torterra",
		},
		{
			BaseModel: BaseModel{
				Id: 390,
			},
			Name: "Chimchar",
		},
		{
			BaseModel: BaseModel{
				Id: 391,
			},
			Name: "Monferno",
		},
		{
			BaseModel: BaseModel{
				Id: 392,
			},
			Name: "Infernape",
		},
		{
			BaseModel: BaseModel{
				Id: 393,
			},
			Name: "Piplup",
		},
		{
			BaseModel: BaseModel{
				Id: 394,
			},
			Name: "Prinplup",
		},
		{
			BaseModel: BaseModel{
				Id: 395,
			},
			Name: "Empoleon",
		},
		{
			BaseModel: BaseModel{
				Id: 396,
			},
			Name: "Starly",
		},
		{
			BaseModel: BaseModel{
				Id: 397,
			},
			Name: "Staravia",
		},
		{
			BaseModel: BaseModel{
				Id: 398,
			},
			Name: "Staraptor",
		},
		{
			BaseModel: BaseModel{
				Id: 399,
			},
			Name: "Bidoof",
		},
		{
			BaseModel: BaseModel{
				Id: 400,
			},
			Name: "Bibarel",
		},
		{
			BaseModel: BaseModel{
				Id: 401,
			},
			Name: "Kricketot",
		},
		{
			BaseModel: BaseModel{
				Id: 402,
			},
			Name: "Kricketune",
		},
		{
			BaseModel: BaseModel{
				Id: 403,
			},
			Name: "Shinx",
		},
		{
			BaseModel: BaseModel{
				Id: 404,
			},
			Name: "Luxio",
		},
		{
			BaseModel: BaseModel{
				Id: 405,
			},
			Name: "Luxray",
		},
		{
			BaseModel: BaseModel{
				Id: 406,
			},
			Name: "Budew",
		},
		{
			BaseModel: BaseModel{
				Id: 407,
			},
			Name: "Roserade",
		},
		{
			BaseModel: BaseModel{
				Id: 408,
			},
			Name: "Cranidos",
		},
		{
			BaseModel: BaseModel{
				Id: 409,
			},
			Name: "Rampardos",
		},
		{
			BaseModel: BaseModel{
				Id: 410,
			},
			Name: "Shieldon",
		},
		{
			BaseModel: BaseModel{
				Id: 411,
			},
			Name: "Bastiodon",
		},
		{
			BaseModel: BaseModel{
				Id: 412,
			},
			Name: "Burmy",
		},
		{
			BaseModel: BaseModel{
				Id: 413,
			},
			Name: "Wormadam",
		},
		{
			BaseModel: BaseModel{
				Id: 414,
			},
			Name: "Mothim",
		},
		{
			BaseModel: BaseModel{
				Id: 415,
			},
			Name: "Combee",
		},
		{
			BaseModel: BaseModel{
				Id: 416,
			},
			Name: "Vespiquen",
		},
		{
			BaseModel: BaseModel{
				Id: 417,
			},
			Name: "Pachirisu",
		},
		{
			BaseModel: BaseModel{
				Id: 418,
			},
			Name: "Buizel",
		},
		{
			BaseModel: BaseModel{
				Id: 419,
			},
			Name: "Floatzel",
		},
		{
			BaseModel: BaseModel{
				Id: 420,
			},
			Name: "Cherubi",
		},
		{
			BaseModel: BaseModel{
				Id: 421,
			},
			Name: "Cherrim",
		},
		{
			BaseModel: BaseModel{
				Id: 422,
			},
			Name: "Shellos",
		},
		{
			BaseModel: BaseModel{
				Id: 423,
			},
			Name: "Gastrodon",
		},
		{
			BaseModel: BaseModel{
				Id: 424,
			},
			Name: "Ambipom",
		},
		{
			BaseModel: BaseModel{
				Id: 425,
			},
			Name: "Drifloon",
		},
		{
			BaseModel: BaseModel{
				Id: 426,
			},
			Name: "Drifblim",
		},
		{
			BaseModel: BaseModel{
				Id: 427,
			},
			Name: "Buneary",
		},
		{
			BaseModel: BaseModel{
				Id: 428,
			},
			Name: "Lopunny",
		},
		{
			BaseModel: BaseModel{
				Id: 429,
			},
			Name: "Mismagius",
		},
		{
			BaseModel: BaseModel{
				Id: 430,
			},
			Name: "Honchkrow",
		},
		{
			BaseModel: BaseModel{
				Id: 431,
			},
			Name: "Glameow",
		},
		{
			BaseModel: BaseModel{
				Id: 432,
			},
			Name: "Purugly",
		},
		{
			BaseModel: BaseModel{
				Id: 433,
			},
			Name: "Chingling",
		},
		{
			BaseModel: BaseModel{
				Id: 434,
			},
			Name: "Stunky",
		},
		{
			BaseModel: BaseModel{
				Id: 435,
			},
			Name: "Skuntank",
		},
		{
			BaseModel: BaseModel{
				Id: 436,
			},
			Name: "Bronzor",
		},
		{
			BaseModel: BaseModel{
				Id: 437,
			},
			Name: "Bronzong",
		},
		{
			BaseModel: BaseModel{
				Id: 438,
			},
			Name: "Bonsly",
		},
		{
			BaseModel: BaseModel{
				Id: 439,
			},
			Name: "Mime Jr.",
		},
		{
			BaseModel: BaseModel{
				Id: 440,
			},
			Name: "Happiny",
		},
		{
			BaseModel: BaseModel{
				Id: 441,
			},
			Name: "Chatot",
		},
		{
			BaseModel: BaseModel{
				Id: 442,
			},
			Name: "Spiritomb",
		},
		{
			BaseModel: BaseModel{
				Id: 443,
			},
			Name: "Gible",
		},
		{
			BaseModel: BaseModel{
				Id: 444,
			},
			Name: "Gabite",
		},
		{
			BaseModel: BaseModel{
				Id: 445,
			},
			Name: "Garchomp",
		},
		{
			BaseModel: BaseModel{
				Id: 446,
			},
			Name: "Munchlax",
		},
		{
			BaseModel: BaseModel{
				Id: 447,
			},
			Name: "Riolu",
		},
		{
			BaseModel: BaseModel{
				Id: 448,
			},
			Name: "Lucario",
		},
		{
			BaseModel: BaseModel{
				Id: 449,
			},
			Name: "Hippopotas",
		},
		{
			BaseModel: BaseModel{
				Id: 450,
			},
			Name: "Hippowdon",
		},
		{
			BaseModel: BaseModel{
				Id: 451,
			},
			Name: "Skorupi",
		},
		{
			BaseModel: BaseModel{
				Id: 452,
			},
			Name: "Drapion",
		},
		{
			BaseModel: BaseModel{
				Id: 453,
			},
			Name: "Croagunk",
		},
		{
			BaseModel: BaseModel{
				Id: 454,
			},
			Name: "Toxicroak",
		},
		{
			BaseModel: BaseModel{
				Id: 455,
			},
			Name: "Carnivine",
		},
		{
			BaseModel: BaseModel{
				Id: 456,
			},
			Name: "Finneon",
		},
		{
			BaseModel: BaseModel{
				Id: 457,
			},
			Name: "Lumineon",
		},
		{
			BaseModel: BaseModel{
				Id: 458,
			},
			Name: "Mantyke",
		},
		{
			BaseModel: BaseModel{
				Id: 459,
			},
			Name: "Snover",
		},
		{
			BaseModel: BaseModel{
				Id: 460,
			},
			Name: "Abomasnow",
		},
		{
			BaseModel: BaseModel{
				Id: 461,
			},
			Name: "Weavile",
		},
		{
			BaseModel: BaseModel{
				Id: 462,
			},
			Name: "Magnezone",
		},
		{
			BaseModel: BaseModel{
				Id: 463,
			},
			Name: "Lickilicky",
		},
		{
			BaseModel: BaseModel{
				Id: 464,
			},
			Name: "Rhyperior",
		},
		{
			BaseModel: BaseModel{
				Id: 465,
			},
			Name: "Tangrowth",
		},
		{
			BaseModel: BaseModel{
				Id: 466,
			},
			Name: "Electivire",
		},
		{
			BaseModel: BaseModel{
				Id: 467,
			},
			Name: "Magmortar",
		},
		{
			BaseModel: BaseModel{
				Id: 468,
			},
			Name: "Togekiss",
		},
		{
			BaseModel: BaseModel{
				Id: 469,
			},
			Name: "Yanmega",
		},
		{
			BaseModel: BaseModel{
				Id: 470,
			},
			Name: "Leafeon",
		},
		{
			BaseModel: BaseModel{
				Id: 471,
			},
			Name: "Glaceon",
		},
		{
			BaseModel: BaseModel{
				Id: 472,
			},
			Name: "Gliscor",
		},
		{
			BaseModel: BaseModel{
				Id: 473,
			},
			Name: "Mamoswine",
		},
		{
			BaseModel: BaseModel{
				Id: 474,
			},
			Name: "Porygon-Z",
		},
		{
			BaseModel: BaseModel{
				Id: 475,
			},
			Name: "Gallade",
		},
		{
			BaseModel: BaseModel{
				Id: 476,
			},
			Name: "Probopass",
		},
		{
			BaseModel: BaseModel{
				Id: 477,
			},
			Name: "Dusknoir",
		},
		{
			BaseModel: BaseModel{
				Id: 478,
			},
			Name: "Froslass",
		},
		{
			BaseModel: BaseModel{
				Id: 479,
			},
			Name: "Rotom",
		},
		{
			BaseModel: BaseModel{
				Id: 480,
			},
			Name: "Uxie",
		},
		{
			BaseModel: BaseModel{
				Id: 481,
			},
			Name: "Mesprit",
		},
		{
			BaseModel: BaseModel{
				Id: 482,
			},
			Name: "Azelf",
		},
		{
			BaseModel: BaseModel{
				Id: 483,
			},
			Name: "Dialga",
		},
		{
			BaseModel: BaseModel{
				Id: 484,
			},
			Name: "Palkia",
		},
		{
			BaseModel: BaseModel{
				Id: 485,
			},
			Name: "Heatran",
		},
		{
			BaseModel: BaseModel{
				Id: 486,
			},
			Name: "Regigigas",
		},
		{
			BaseModel: BaseModel{
				Id: 487,
			},
			Name: "Giratina",
		},
		{
			BaseModel: BaseModel{
				Id: 488,
			},
			Name: "Cresselia",
		},
		{
			BaseModel: BaseModel{
				Id: 489,
			},
			Name: "Phione",
		},
		{
			BaseModel: BaseModel{
				Id: 490,
			},
			Name: "Manaphy",
		},
		{
			BaseModel: BaseModel{
				Id: 491,
			},
			Name: "Darkrai",
		},
		{
			BaseModel: BaseModel{
				Id: 492,
			},
			Name: "Shaymin",
		},
		{
			BaseModel: BaseModel{
				Id: 493,
			},
			Name: "Arceus",
		},
		{
			BaseModel: BaseModel{
				Id: 494,
			},
			Name: "Victini",
		},
		{
			BaseModel: BaseModel{
				Id: 495,
			},
			Name: "Snivy",
		},
		{
			BaseModel: BaseModel{
				Id: 496,
			},
			Name: "Servine",
		},
		{
			BaseModel: BaseModel{
				Id: 497,
			},
			Name: "Serperior",
		},
		{
			BaseModel: BaseModel{
				Id: 498,
			},
			Name: "Tepig",
		},
		{
			BaseModel: BaseModel{
				Id: 499,
			},
			Name: "Pignite",
		},
		{
			BaseModel: BaseModel{
				Id: 500,
			},
			Name: "Emboar",
		},
		{
			BaseModel: BaseModel{
				Id: 501,
			},
			Name: "Oshawott",
		},
		{
			BaseModel: BaseModel{
				Id: 502,
			},
			Name: "Dewott",
		},
		{
			BaseModel: BaseModel{
				Id: 503,
			},
			Name: "Samurott",
		},
		{
			BaseModel: BaseModel{
				Id: 504,
			},
			Name: "Patrat",
		},
		{
			BaseModel: BaseModel{
				Id: 505,
			},
			Name: "Watchog",
		},
		{
			BaseModel: BaseModel{
				Id: 506,
			},
			Name: "Lillipup",
		},
		{
			BaseModel: BaseModel{
				Id: 507,
			},
			Name: "Herdier",
		},
		{
			BaseModel: BaseModel{
				Id: 508,
			},
			Name: "Stoutland",
		},
		{
			BaseModel: BaseModel{
				Id: 509,
			},
			Name: "Purrloin",
		},
		{
			BaseModel: BaseModel{
				Id: 510,
			},
			Name: "Liepard",
		},
		{
			BaseModel: BaseModel{
				Id: 511,
			},
			Name: "Pansage",
		},
		{
			BaseModel: BaseModel{
				Id: 512,
			},
			Name: "Simisage",
		},
		{
			BaseModel: BaseModel{
				Id: 513,
			},
			Name: "Pansear",
		},
		{
			BaseModel: BaseModel{
				Id: 514,
			},
			Name: "Simisear",
		},
		{
			BaseModel: BaseModel{
				Id: 515,
			},
			Name: "Panpour",
		},
		{
			BaseModel: BaseModel{
				Id: 516,
			},
			Name: "Simipour",
		},
		{
			BaseModel: BaseModel{
				Id: 517,
			},
			Name: "Munna",
		},
		{
			BaseModel: BaseModel{
				Id: 518,
			},
			Name: "Musharna",
		},
		{
			BaseModel: BaseModel{
				Id: 519,
			},
			Name: "Pidove",
		},
		{
			BaseModel: BaseModel{
				Id: 520,
			},
			Name: "Tranquill",
		},
		{
			BaseModel: BaseModel{
				Id: 521,
			},
			Name: "Unfezant",
		},
		{
			BaseModel: BaseModel{
				Id: 522,
			},
			Name: "Blitzle",
		},
		{
			BaseModel: BaseModel{
				Id: 523,
			},
			Name: "Zebstrika",
		},
		{
			BaseModel: BaseModel{
				Id: 524,
			},
			Name: "Roggenrola",
		},
		{
			BaseModel: BaseModel{
				Id: 525,
			},
			Name: "Boldore",
		},
		{
			BaseModel: BaseModel{
				Id: 526,
			},
			Name: "Gigalith",
		},
		{
			BaseModel: BaseModel{
				Id: 527,
			},
			Name: "Woobat",
		},
		{
			BaseModel: BaseModel{
				Id: 528,
			},
			Name: "Swoobat",
		},
		{
			BaseModel: BaseModel{
				Id: 529,
			},
			Name: "Drilbur",
		},
		{
			BaseModel: BaseModel{
				Id: 530,
			},
			Name: "Excadrill",
		},
		{
			BaseModel: BaseModel{
				Id: 531,
			},
			Name: "Audino",
		},
		{
			BaseModel: BaseModel{
				Id: 532,
			},
			Name: "Timburr",
		},
		{
			BaseModel: BaseModel{
				Id: 533,
			},
			Name: "Gurdurr",
		},
		{
			BaseModel: BaseModel{
				Id: 534,
			},
			Name: "Conkeldurr",
		},
		{
			BaseModel: BaseModel{
				Id: 535,
			},
			Name: "Tympole",
		},
		{
			BaseModel: BaseModel{
				Id: 536,
			},
			Name: "Palpitoad",
		},
		{
			BaseModel: BaseModel{
				Id: 537,
			},
			Name: "Seismitoad",
		},
		{
			BaseModel: BaseModel{
				Id: 538,
			},
			Name: "Throh",
		},
		{
			BaseModel: BaseModel{
				Id: 539,
			},
			Name: "Sawk",
		},
		{
			BaseModel: BaseModel{
				Id: 540,
			},
			Name: "Sewaddle",
		},
		{
			BaseModel: BaseModel{
				Id: 541,
			},
			Name: "Swadloon",
		},
		{
			BaseModel: BaseModel{
				Id: 542,
			},
			Name: "Leavanny",
		},
		{
			BaseModel: BaseModel{
				Id: 543,
			},
			Name: "Venipede",
		},
		{
			BaseModel: BaseModel{
				Id: 544,
			},
			Name: "Whirlipede",
		},
		{
			BaseModel: BaseModel{
				Id: 545,
			},
			Name: "Scolipede",
		},
		{
			BaseModel: BaseModel{
				Id: 546,
			},
			Name: "Cottonee",
		},
		{
			BaseModel: BaseModel{
				Id: 547,
			},
			Name: "Whimsicott",
		},
		{
			BaseModel: BaseModel{
				Id: 548,
			},
			Name: "Petilil",
		},
		{
			BaseModel: BaseModel{
				Id: 549,
			},
			Name: "Lilligant",
		},
		{
			BaseModel: BaseModel{
				Id: 550,
			},
			Name: "Basculin",
		},
		{
			BaseModel: BaseModel{
				Id: 551,
			},
			Name: "Sandile",
		},
		{
			BaseModel: BaseModel{
				Id: 552,
			},
			Name: "Krokorok",
		},
		{
			BaseModel: BaseModel{
				Id: 553,
			},
			Name: "Krookodile",
		},
		{
			BaseModel: BaseModel{
				Id: 554,
			},
			Name: "Darumaka",
		},
		{
			BaseModel: BaseModel{
				Id: 555,
			},
			Name: "Darmanitan",
		},
		{
			BaseModel: BaseModel{
				Id: 556,
			},
			Name: "Maractus",
		},
		{
			BaseModel: BaseModel{
				Id: 557,
			},
			Name: "Dwebble",
		},
		{
			BaseModel: BaseModel{
				Id: 558,
			},
			Name: "Crustle",
		},
		{
			BaseModel: BaseModel{
				Id: 559,
			},
			Name: "Scraggy",
		},
		{
			BaseModel: BaseModel{
				Id: 560,
			},
			Name: "Scrafty",
		},
		{
			BaseModel: BaseModel{
				Id: 561,
			},
			Name: "Sigilyph",
		},
		{
			BaseModel: BaseModel{
				Id: 562,
			},
			Name: "Yamask",
		},
		{
			BaseModel: BaseModel{
				Id: 563,
			},
			Name: "Cofagrigus",
		},
		{
			BaseModel: BaseModel{
				Id: 564,
			},
			Name: "Tirtouga",
		},
		{
			BaseModel: BaseModel{
				Id: 565,
			},
			Name: "Carracosta",
		},
		{
			BaseModel: BaseModel{
				Id: 566,
			},
			Name: "Archen",
		},
		{
			BaseModel: BaseModel{
				Id: 567,
			},
			Name: "Archeops",
		},
		{
			BaseModel: BaseModel{
				Id: 568,
			},
			Name: "Trubbish",
		},
		{
			BaseModel: BaseModel{
				Id: 569,
			},
			Name: "Garbodor",
		},
		{
			BaseModel: BaseModel{
				Id: 570,
			},
			Name: "Zorua",
		},
		{
			BaseModel: BaseModel{
				Id: 571,
			},
			Name: "Zoroark",
		},
		{
			BaseModel: BaseModel{
				Id: 572,
			},
			Name: "Minccino",
		},
		{
			BaseModel: BaseModel{
				Id: 573,
			},
			Name: "Cinccino",
		},
		{
			BaseModel: BaseModel{
				Id: 574,
			},
			Name: "Gothita",
		},
		{
			BaseModel: BaseModel{
				Id: 575,
			},
			Name: "Gothorita",
		},
		{
			BaseModel: BaseModel{
				Id: 576,
			},
			Name: "Gothitelle",
		},
		{
			BaseModel: BaseModel{
				Id: 577,
			},
			Name: "Solosis",
		},
		{
			BaseModel: BaseModel{
				Id: 578,
			},
			Name: "Duosion",
		},
		{
			BaseModel: BaseModel{
				Id: 579,
			},
			Name: "Reuniclus",
		},
		{
			BaseModel: BaseModel{
				Id: 580,
			},
			Name: "Ducklett",
		},
		{
			BaseModel: BaseModel{
				Id: 581,
			},
			Name: "Swanna",
		},
		{
			BaseModel: BaseModel{
				Id: 582,
			},
			Name: "Vanillite",
		},
		{
			BaseModel: BaseModel{
				Id: 583,
			},
			Name: "Vanillish",
		},
		{
			BaseModel: BaseModel{
				Id: 584,
			},
			Name: "Vanilluxe",
		},
		{
			BaseModel: BaseModel{
				Id: 585,
			},
			Name: "Deerling",
		},
		{
			BaseModel: BaseModel{
				Id: 586,
			},
			Name: "Sawsbuck",
		},
		{
			BaseModel: BaseModel{
				Id: 587,
			},
			Name: "Emolga",
		},
		{
			BaseModel: BaseModel{
				Id: 588,
			},
			Name: "Karrablast",
		},
		{
			BaseModel: BaseModel{
				Id: 589,
			},
			Name: "Escavalier",
		},
		{
			BaseModel: BaseModel{
				Id: 590,
			},
			Name: "Foongus",
		},
		{
			BaseModel: BaseModel{
				Id: 591,
			},
			Name: "Amoonguss",
		},
		{
			BaseModel: BaseModel{
				Id: 592,
			},
			Name: "Frillish",
		},
		{
			BaseModel: BaseModel{
				Id: 593,
			},
			Name: "Jellicent",
		},
		{
			BaseModel: BaseModel{
				Id: 594,
			},
			Name: "Alomomola",
		},
		{
			BaseModel: BaseModel{
				Id: 595,
			},
			Name: "Joltik",
		},
		{
			BaseModel: BaseModel{
				Id: 596,
			},
			Name: "Galvantula",
		},
		{
			BaseModel: BaseModel{
				Id: 597,
			},
			Name: "Ferroseed",
		},
		{
			BaseModel: BaseModel{
				Id: 598,
			},
			Name: "Ferrothorn",
		},
		{
			BaseModel: BaseModel{
				Id: 599,
			},
			Name: "Klink",
		},
		{
			BaseModel: BaseModel{
				Id: 600,
			},
			Name: "Klang",
		},
		{
			BaseModel: BaseModel{
				Id: 601,
			},
			Name: "Klinklang",
		},
		{
			BaseModel: BaseModel{
				Id: 602,
			},
			Name: "Tynamo",
		},
		{
			BaseModel: BaseModel{
				Id: 603,
			},
			Name: "Eelektrik",
		},
		{
			BaseModel: BaseModel{
				Id: 604,
			},
			Name: "Eelektross",
		},
		{
			BaseModel: BaseModel{
				Id: 605,
			},
			Name: "Elgyem",
		},
		{
			BaseModel: BaseModel{
				Id: 606,
			},
			Name: "Beheeyem",
		},
		{
			BaseModel: BaseModel{
				Id: 607,
			},
			Name: "Litwick",
		},
		{
			BaseModel: BaseModel{
				Id: 608,
			},
			Name: "Lampent",
		},
		{
			BaseModel: BaseModel{
				Id: 609,
			},
			Name: "Chandelure",
		},
		{
			BaseModel: BaseModel{
				Id: 610,
			},
			Name: "Axew",
		},
		{
			BaseModel: BaseModel{
				Id: 611,
			},
			Name: "Fraxure",
		},
		{
			BaseModel: BaseModel{
				Id: 612,
			},
			Name: "Haxorus",
		},
		{
			BaseModel: BaseModel{
				Id: 613,
			},
			Name: "Cubchoo",
		},
		{
			BaseModel: BaseModel{
				Id: 614,
			},
			Name: "Beartic",
		},
		{
			BaseModel: BaseModel{
				Id: 615,
			},
			Name: "Cryogonal",
		},
		{
			BaseModel: BaseModel{
				Id: 616,
			},
			Name: "Shelmet",
		},
		{
			BaseModel: BaseModel{
				Id: 617,
			},
			Name: "Accelgor",
		},
		{
			BaseModel: BaseModel{
				Id: 618,
			},
			Name: "Stunfisk",
		},
		{
			BaseModel: BaseModel{
				Id: 619,
			},
			Name: "Mienfoo",
		},
		{
			BaseModel: BaseModel{
				Id: 620,
			},
			Name: "Mienshao",
		},
		{
			BaseModel: BaseModel{
				Id: 621,
			},
			Name: "Druddigon",
		},
		{
			BaseModel: BaseModel{
				Id: 622,
			},
			Name: "Golett",
		},
		{
			BaseModel: BaseModel{
				Id: 623,
			},
			Name: "Golurk",
		},
		{
			BaseModel: BaseModel{
				Id: 624,
			},
			Name: "Pawniard",
		},
		{
			BaseModel: BaseModel{
				Id: 625,
			},
			Name: "Bisharp",
		},
		{
			BaseModel: BaseModel{
				Id: 626,
			},
			Name: "Bouffalant",
		},
		{
			BaseModel: BaseModel{
				Id: 627,
			},
			Name: "Rufflet",
		},
		{
			BaseModel: BaseModel{
				Id: 628,
			},
			Name: "Braviary",
		},
		{
			BaseModel: BaseModel{
				Id: 629,
			},
			Name: "Vullaby",
		},
		{
			BaseModel: BaseModel{
				Id: 630,
			},
			Name: "Mandibuzz",
		},
		{
			BaseModel: BaseModel{
				Id: 631,
			},
			Name: "Heatmor",
		},
		{
			BaseModel: BaseModel{
				Id: 632,
			},
			Name: "Durant",
		},
		{
			BaseModel: BaseModel{
				Id: 633,
			},
			Name: "Deino",
		},
		{
			BaseModel: BaseModel{
				Id: 634,
			},
			Name: "Zweilous",
		},
		{
			BaseModel: BaseModel{
				Id: 635,
			},
			Name: "Hydreigon",
		},
		{
			BaseModel: BaseModel{
				Id: 636,
			},
			Name: "Larvesta",
		},
		{
			BaseModel: BaseModel{
				Id: 637,
			},
			Name: "Volcarona",
		},
		{
			BaseModel: BaseModel{
				Id: 638,
			},
			Name: "Cobalion",
		},
		{
			BaseModel: BaseModel{
				Id: 639,
			},
			Name: "Terrakion",
		},
		{
			BaseModel: BaseModel{
				Id: 640,
			},
			Name: "Virizion",
		},
		{
			BaseModel: BaseModel{
				Id: 641,
			},
			Name: "Tornadus",
		},
		{
			BaseModel: BaseModel{
				Id: 642,
			},
			Name: "Thundurus",
		},
		{
			BaseModel: BaseModel{
				Id: 643,
			},
			Name: "Reshiram",
		},
		{
			BaseModel: BaseModel{
				Id: 644,
			},
			Name: "Zekrom",
		},
		{
			BaseModel: BaseModel{
				Id: 645,
			},
			Name: "Landorus",
		},
		{
			BaseModel: BaseModel{
				Id: 646,
			},
			Name: "Kyurem",
		},
		{
			BaseModel: BaseModel{
				Id: 647,
			},
			Name: "Keldeo",
		},
		{
			BaseModel: BaseModel{
				Id: 648,
			},
			Name: "Meloetta",
		},
		{
			BaseModel: BaseModel{
				Id: 649,
			},
			Name: "Genesect",
		},
		{
			BaseModel: BaseModel{
				Id: 650,
			},
			Name: "Chespin",
		},
		{
			BaseModel: BaseModel{
				Id: 651,
			},
			Name: "Quilladin",
		},
		{
			BaseModel: BaseModel{
				Id: 652,
			},
			Name: "Chesnaught",
		},
		{
			BaseModel: BaseModel{
				Id: 653,
			},
			Name: "Fennekin",
		},
		{
			BaseModel: BaseModel{
				Id: 654,
			},
			Name: "Braixen",
		},
		{
			BaseModel: BaseModel{
				Id: 655,
			},
			Name: "Delphox",
		},
		{
			BaseModel: BaseModel{
				Id: 656,
			},
			Name: "Froakie",
		},
		{
			BaseModel: BaseModel{
				Id: 657,
			},
			Name: "Frogadier",
		},
		{
			BaseModel: BaseModel{
				Id: 658,
			},
			Name: "Greninja",
		},
		{
			BaseModel: BaseModel{
				Id: 659,
			},
			Name: "Bunnelby",
		},
		{
			BaseModel: BaseModel{
				Id: 660,
			},
			Name: "Diggersby",
		},
		{
			BaseModel: BaseModel{
				Id: 661,
			},
			Name: "Fletchling",
		},
		{
			BaseModel: BaseModel{
				Id: 662,
			},
			Name: "Fletchinder",
		},
		{
			BaseModel: BaseModel{
				Id: 663,
			},
			Name: "Talonflame",
		},
		{
			BaseModel: BaseModel{
				Id: 664,
			},
			Name: "Scatterbug",
		},
		{
			BaseModel: BaseModel{
				Id: 665,
			},
			Name: "Spewpa",
		},
		{
			BaseModel: BaseModel{
				Id: 666,
			},
			Name: "Vivillon",
		},
		{
			BaseModel: BaseModel{
				Id: 667,
			},
			Name: "Litleo",
		},
		{
			BaseModel: BaseModel{
				Id: 668,
			},
			Name: "Pyroar",
		},
		{
			BaseModel: BaseModel{
				Id: 669,
			},
			Name: "Flabébé",
		},
		{
			BaseModel: BaseModel{
				Id: 670,
			},
			Name: "Floette",
		},
		{
			BaseModel: BaseModel{
				Id: 671,
			},
			Name: "Florges",
		},
		{
			BaseModel: BaseModel{
				Id: 672,
			},
			Name: "Skiddo",
		},
		{
			BaseModel: BaseModel{
				Id: 673,
			},
			Name: "Gogoat",
		},
		{
			BaseModel: BaseModel{
				Id: 674,
			},
			Name: "Pancham",
		},
		{
			BaseModel: BaseModel{
				Id: 675,
			},
			Name: "Pangoro",
		},
		{
			BaseModel: BaseModel{
				Id: 676,
			},
			Name: "Furfrou",
		},
		{
			BaseModel: BaseModel{
				Id: 677,
			},
			Name: "Espurr",
		},
		{
			BaseModel: BaseModel{
				Id: 678,
			},
			Name: "Meowstic",
		},
		{
			BaseModel: BaseModel{
				Id: 679,
			},
			Name: "Honedge",
		},
		{
			BaseModel: BaseModel{
				Id: 680,
			},
			Name: "Doublade",
		},
		{
			BaseModel: BaseModel{
				Id: 681,
			},
			Name: "Aegislash",
		},
		{
			BaseModel: BaseModel{
				Id: 682,
			},
			Name: "Spritzee",
		},
		{
			BaseModel: BaseModel{
				Id: 683,
			},
			Name: "Aromatisse",
		},
		{
			BaseModel: BaseModel{
				Id: 684,
			},
			Name: "Swirlix",
		},
		{
			BaseModel: BaseModel{
				Id: 685,
			},
			Name: "Slurpuff",
		},
		{
			BaseModel: BaseModel{
				Id: 686,
			},
			Name: "Inkay",
		},
		{
			BaseModel: BaseModel{
				Id: 687,
			},
			Name: "Malamar",
		},
		{
			BaseModel: BaseModel{
				Id: 688,
			},
			Name: "Binacle",
		},
		{
			BaseModel: BaseModel{
				Id: 689,
			},
			Name: "Barbaracle",
		},
		{
			BaseModel: BaseModel{
				Id: 690,
			},
			Name: "Skrelp",
		},
		{
			BaseModel: BaseModel{
				Id: 691,
			},
			Name: "Dragalge",
		},
		{
			BaseModel: BaseModel{
				Id: 692,
			},
			Name: "Clauncher",
		},
		{
			BaseModel: BaseModel{
				Id: 693,
			},
			Name: "Clawitzer",
		},
		{
			BaseModel: BaseModel{
				Id: 694,
			},
			Name: "Helioptile",
		},
		{
			BaseModel: BaseModel{
				Id: 695,
			},
			Name: "Heliolisk",
		},
		{
			BaseModel: BaseModel{
				Id: 696,
			},
			Name: "Tyrunt",
		},
		{
			BaseModel: BaseModel{
				Id: 697,
			},
			Name: "Tyrantrum",
		},
		{
			BaseModel: BaseModel{
				Id: 698,
			},
			Name: "Amaura",
		},
		{
			BaseModel: BaseModel{
				Id: 699,
			},
			Name: "Aurorus",
		},
		{
			BaseModel: BaseModel{
				Id: 700,
			},
			Name: "Sylveon",
		},
		{
			BaseModel: BaseModel{
				Id: 701,
			},
			Name: "Hawlucha",
		},
		{
			BaseModel: BaseModel{
				Id: 702,
			},
			Name: "Dedenne",
		},
		{
			BaseModel: BaseModel{
				Id: 703,
			},
			Name: "Carbink",
		},
		{
			BaseModel: BaseModel{
				Id: 704,
			},
			Name: "Goomy",
		},
		{
			BaseModel: BaseModel{
				Id: 705,
			},
			Name: "Sliggoo",
		},
		{
			BaseModel: BaseModel{
				Id: 706,
			},
			Name: "Goodra",
		},
		{
			BaseModel: BaseModel{
				Id: 707,
			},
			Name: "Klefki",
		},
		{
			BaseModel: BaseModel{
				Id: 708,
			},
			Name: "Phantump",
		},
		{
			BaseModel: BaseModel{
				Id: 709,
			},
			Name: "Trevenant",
		},
		{
			BaseModel: BaseModel{
				Id: 710,
			},
			Name: "Pumpkaboo",
		},
		{
			BaseModel: BaseModel{
				Id: 711,
			},
			Name: "Gourgeist",
		},
		{
			BaseModel: BaseModel{
				Id: 712,
			},
			Name: "Bergmite",
		},
		{
			BaseModel: BaseModel{
				Id: 713,
			},
			Name: "Avalugg",
		},
		{
			BaseModel: BaseModel{
				Id: 714,
			},
			Name: "Noibat",
		},
		{
			BaseModel: BaseModel{
				Id: 715,
			},
			Name: "Noivern",
		},
		{
			BaseModel: BaseModel{
				Id: 716,
			},
			Name: "Xerneas",
		},
		{
			BaseModel: BaseModel{
				Id: 717,
			},
			Name: "Yveltal",
		},
		{
			BaseModel: BaseModel{
				Id: 718,
			},
			Name: "Zygarde",
		},
		{
			BaseModel: BaseModel{
				Id: 719,
			},
			Name: "Diancie",
		},
		{
			BaseModel: BaseModel{
				Id: 720,
			},
			Name: "Hoopa",
		},
		{
			BaseModel: BaseModel{
				Id: 721,
			},
			Name: "Volcanion",
		},
		{
			BaseModel: BaseModel{
				Id: 722,
			},
			Name: "Rowlet",
		},
		{
			BaseModel: BaseModel{
				Id: 723,
			},
			Name: "Dartrix",
		},
		{
			BaseModel: BaseModel{
				Id: 724,
			},
			Name: "Decidueye",
		},
		{
			BaseModel: BaseModel{
				Id: 725,
			},
			Name: "Litten",
		},
		{
			BaseModel: BaseModel{
				Id: 726,
			},
			Name: "Torracat",
		},
		{
			BaseModel: BaseModel{
				Id: 727,
			},
			Name: "Incineroar",
		},
		{
			BaseModel: BaseModel{
				Id: 728,
			},
			Name: "Popplio",
		},
		{
			BaseModel: BaseModel{
				Id: 729,
			},
			Name: "Brionne",
		},
		{
			BaseModel: BaseModel{
				Id: 730,
			},
			Name: "Primarina",
		},
		{
			BaseModel: BaseModel{
				Id: 731,
			},
			Name: "Pikipek",
		},
		{
			BaseModel: BaseModel{
				Id: 732,
			},
			Name: "Trumbeak",
		},
		{
			BaseModel: BaseModel{
				Id: 733,
			},
			Name: "Toucannon",
		},
		{
			BaseModel: BaseModel{
				Id: 734,
			},
			Name: "Yungoos",
		},
		{
			BaseModel: BaseModel{
				Id: 735,
			},
			Name: "Gumshoos",
		},
		{
			BaseModel: BaseModel{
				Id: 736,
			},
			Name: "Grubbin",
		},
		{
			BaseModel: BaseModel{
				Id: 737,
			},
			Name: "Charjabug",
		},
		{
			BaseModel: BaseModel{
				Id: 738,
			},
			Name: "Vikavolt",
		},
		{
			BaseModel: BaseModel{
				Id: 739,
			},
			Name: "Crabrawler",
		},
		{
			BaseModel: BaseModel{
				Id: 740,
			},
			Name: "Crabominable",
		},
		{
			BaseModel: BaseModel{
				Id: 741,
			},
			Name: "Oricorio",
		},
		{
			BaseModel: BaseModel{
				Id: 742,
			},
			Name: "Cutiefly",
		},
		{
			BaseModel: BaseModel{
				Id: 743,
			},
			Name: "Ribombee",
		},
		{
			BaseModel: BaseModel{
				Id: 744,
			},
			Name: "Rockruff",
		},
		{
			BaseModel: BaseModel{
				Id: 745,
			},
			Name: "Lycanroc",
		},
		{
			BaseModel: BaseModel{
				Id: 746,
			},
			Name: "Wishiwashi",
		},
		{
			BaseModel: BaseModel{
				Id: 747,
			},
			Name: "Mareanie",
		},
		{
			BaseModel: BaseModel{
				Id: 748,
			},
			Name: "Toxapex",
		},
		{
			BaseModel: BaseModel{
				Id: 749,
			},
			Name: "Mudbray",
		},
		{
			BaseModel: BaseModel{
				Id: 750,
			},
			Name: "Mudsdale",
		},
		{
			BaseModel: BaseModel{
				Id: 751,
			},
			Name: "Dewpider",
		},
		{
			BaseModel: BaseModel{
				Id: 752,
			},
			Name: "Araquanid",
		},
		{
			BaseModel: BaseModel{
				Id: 753,
			},
			Name: "Fomantis",
		},
		{
			BaseModel: BaseModel{
				Id: 754,
			},
			Name: "Lurantis",
		},
		{
			BaseModel: BaseModel{
				Id: 755,
			},
			Name: "Morelull",
		},
		{
			BaseModel: BaseModel{
				Id: 756,
			},
			Name: "Shiinotic",
		},
		{
			BaseModel: BaseModel{
				Id: 757,
			},
			Name: "Salandit",
		},
		{
			BaseModel: BaseModel{
				Id: 758,
			},
			Name: "Salazzle",
		},
		{
			BaseModel: BaseModel{
				Id: 759,
			},
			Name: "Stufful",
		},
		{
			BaseModel: BaseModel{
				Id: 760,
			},
			Name: "Bewear",
		},
		{
			BaseModel: BaseModel{
				Id: 761,
			},
			Name: "Bounsweet",
		},
		{
			BaseModel: BaseModel{
				Id: 762,
			},
			Name: "Steenee",
		},
		{
			BaseModel: BaseModel{
				Id: 763,
			},
			Name: "Tsareena",
		},
		{
			BaseModel: BaseModel{
				Id: 764,
			},
			Name: "Comfey",
		},
		{
			BaseModel: BaseModel{
				Id: 765,
			},
			Name: "Oranguru",
		},
		{
			BaseModel: BaseModel{
				Id: 766,
			},
			Name: "Passimian",
		},
		{
			BaseModel: BaseModel{
				Id: 767,
			},
			Name: "Wimpod",
		},
		{
			BaseModel: BaseModel{
				Id: 768,
			},
			Name: "Golisopod",
		},
		{
			BaseModel: BaseModel{
				Id: 769,
			},
			Name: "Sandygast",
		},
		{
			BaseModel: BaseModel{
				Id: 770,
			},
			Name: "Palossand",
		},
		{
			BaseModel: BaseModel{
				Id: 771,
			},
			Name: "Pyukumuku",
		},
		{
			BaseModel: BaseModel{
				Id: 772,
			},
			Name: "Type: Null",
		},
		{
			BaseModel: BaseModel{
				Id: 773,
			},
			Name: "Silvally",
		},
		{
			BaseModel: BaseModel{
				Id: 774,
			},
			Name: "Minior",
		},
		{
			BaseModel: BaseModel{
				Id: 775,
			},
			Name: "Komala",
		},
		{
			BaseModel: BaseModel{
				Id: 776,
			},
			Name: "Turtonator",
		},
		{
			BaseModel: BaseModel{
				Id: 777,
			},
			Name: "Togedemaru",
		},
		{
			BaseModel: BaseModel{
				Id: 778,
			},
			Name: "Mimikyu",
		},
		{
			BaseModel: BaseModel{
				Id: 779,
			},
			Name: "Bruxish",
		},
		{
			BaseModel: BaseModel{
				Id: 780,
			},
			Name: "Drampa",
		},
		{
			BaseModel: BaseModel{
				Id: 781,
			},
			Name: "Dhelmise",
		},
		{
			BaseModel: BaseModel{
				Id: 782,
			},
			Name: "Jangmo-o",
		},
		{
			BaseModel: BaseModel{
				Id: 783,
			},
			Name: "Hakamo-o",
		},
		{
			BaseModel: BaseModel{
				Id: 784,
			},
			Name: "Kommo-o",
		},
		{
			BaseModel: BaseModel{
				Id: 785,
			},
			Name: "Tapu Koko",
		},
		{
			BaseModel: BaseModel{
				Id: 786,
			},
			Name: "Tapu Lele",
		},
		{
			BaseModel: BaseModel{
				Id: 787,
			},
			Name: "Tapu Bulu",
		},
		{
			BaseModel: BaseModel{
				Id: 788,
			},
			Name: "Tapu Fini",
		},
		{
			BaseModel: BaseModel{
				Id: 789,
			},
			Name: "Cosmog",
		},
		{
			BaseModel: BaseModel{
				Id: 790,
			},
			Name: "Cosmoem",
		},
		{
			BaseModel: BaseModel{
				Id: 791,
			},
			Name: "Solgaleo",
		},
		{
			BaseModel: BaseModel{
				Id: 792,
			},
			Name: "Lunala",
		},
		{
			BaseModel: BaseModel{
				Id: 793,
			},
			Name: "Nihilego",
		},
		{
			BaseModel: BaseModel{
				Id: 794,
			},
			Name: "Buzzwole",
		},
		{
			BaseModel: BaseModel{
				Id: 795,
			},
			Name: "Pheromosa",
		},
		{
			BaseModel: BaseModel{
				Id: 796,
			},
			Name: "Xurkitree",
		},
		{
			BaseModel: BaseModel{
				Id: 797,
			},
			Name: "Celesteela",
		},
		{
			BaseModel: BaseModel{
				Id: 798,
			},
			Name: "Kartana",
		},
		{
			BaseModel: BaseModel{
				Id: 799,
			},
			Name: "Guzzlord",
		},
		{
			BaseModel: BaseModel{
				Id: 800,
			},
			Name: "Necrozma",
		},
		{
			BaseModel: BaseModel{
				Id: 801,
			},
			Name: "Magearna",
		},
		{
			BaseModel: BaseModel{
				Id: 802,
			},
			Name: "Marshadow",
		},
		{
			BaseModel: BaseModel{
				Id: 803,
			},
			Name: "Poipole",
		},
		{
			BaseModel: BaseModel{
				Id: 804,
			},
			Name: "Naganadel",
		},
		{
			BaseModel: BaseModel{
				Id: 805,
			},
			Name: "Stakataka",
		},
		{
			BaseModel: BaseModel{
				Id: 806,
			},
			Name: "Blacephalon",
		},
		{
			BaseModel: BaseModel{
				Id: 807,
			},
			Name: "Zeraora",
		},
		{
			BaseModel: BaseModel{
				Id: 808,
			},
			Name: "Meltan",
		},
		{
			BaseModel: BaseModel{
				Id: 809,
			},
			Name: "Melmetal",
		},
	}

	for _, poke := range pokemonDefaults {
		err := db.Insert(&poke)
		if err != nil {
			continue
		}
	}

	var pokemon Pokemon
	err := db.Model(&pokemon).First()
	if err != nil {
		return err
	}

	return nil
}
