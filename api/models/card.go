package models

import (
	"api/db"
	"api/filter"
	"net/url"

	"database/sql"
)

// Card defines the properties of a card in Spellslingers
type Card struct {
	ID          int64         `db:"id"`
	Name        string        `db:"name"`
	Set         Set           `db:"set"`
	Cost        int           `db:"cost"`
	Power       sql.NullInt32 `db:"power"`
	Health      sql.NullInt32 `db:"health"`
	Ability     string        `db:"ability"`
	Chance      sql.NullInt32 `db:"chance"`
	Rarity      Rarity        `db:"rarity"`
	CardType    CardType      `db:"card_type"`
	SubType     []SubType     `db:"-"`
	Artist      Artist        `db:"artist"`
	Colour      []Colour      `db:"-"`
	Legendary   bool          `db:"legendary"`
	Collectible bool          `db:"collectible"`
	Material    Material      `db:"material"`
	Mox         Material      `db:"mox"`
}

// CardFilter foo
type CardFilter struct {
	Name       string `filter-alias:"C" filter-column-name:"name"`
	SetID      int    `filter-alias:"C" filter-column-name:"set_id"`
	Cost       int    `filter-alias:"C" filter-column-name:"cost"`
	Power      int    `filter-alias:"C" filter-column-name:"power"`
	Health     int    `filter-alias:"C" filter-column-name:"health"`
	Ability    string `filter-alias:"C" filter-column-name:"ability"`
	Chance     int    `filter-alias:"C" filter-column-name:"chance"`
	RarityID   int    `filter-alias:"C" filter-column-name:"rarity_id"`
	CardTypeID int    `filter-alias:"C" filter-column-name:"card_type_id"`
	// SubType     []SubType `filter-alias:"C" filter-column-name:"sub_type_id"`
	ArtistID int `filter-alias:"C" filter-column-name:"artist_id"`
	// Colour      []Colour `filter-alias:"C" filter-column-name:"colour_id"`
	Legendary   bool `filter-alias:"C" filter-column-name:"legendary"`
	Collectible bool `filter-alias:"C" filter-column-name:"collectible"`
	MaterialID  int  `filter-alias:"C" filter-column-name:"material_id"`
	MoxID       int  `filter-alias:"C" filter-column-name:"mox_id"`
}

// AllCards enumerates all artists, using filters specified in the request
func AllCards(queryParams url.Values) ([]Card, error) {
	cardSql := `
		SELECT
			C.id,
			C.name,
			S.id AS "set.id",
			S.name AS "set.name",
			C.cost,
			C.power,
			C.health,
			C.ability,
			C.chance,
			R.id AS "rarity.id",
			R.name AS "rarity.name",
			CT.id AS "card_type.id",
			CT.name AS "card_type.name",
			A.id AS "artist.id",
			A.name AS "artist.name",
			C.legendary,
			C.collectible,
			M1.id AS "material.id",
			M1.name AS "material.name",
			M2.id AS "mox.id",
			M2.name AS "mox.name"
		FROM public.cards C
		INNER JOIN public.sets S ON C.set_id = S.id
		INNER JOIN public.rarities R ON C.rarity_id = R.id
		INNER JOIN public.card_types CT ON C.card_type_id = CT.id
		INNER JOIN public.artists A ON C.artist_id = A.id
		INNER JOIN public.materials M1 ON C.material_id = M1.id
		INNER JOIN public.materials M2 ON C.mox_id = M2.id
	`

	colourSql := `
		SELECT
			CO.id,
			CO.name
		FROM public.cards CA
		INNER JOIN public.card_colours CC ON CA.id = CC.card_id
		INNER JOIN public.colours CO ON CO.id = CC.colour_id
		WHERE CA.id = $1
	`

	subTypeSql := `
		SELECT
			ST.id,
			ST.name
		FROM public.cards CA
		INNER JOIN public.card_sub_types CST ON CA.id = CST.card_id
		INNER JOIN public.sub_types ST ON ST.id = CST.sub_type_id
		WHERE CA.id = $1
	`

	whereClause, args, err := filter.BuildWhereClauseFromQuery(queryParams, CardFilter{})
	if err != nil {
		return nil, err
	}

	var allCards []Card
	err = db.GlobalSqlxDB.Select(&allCards, cardSql+whereClause, args...)
	if err != nil {
		return nil, err
	}

	for i, card := range allCards {
		var colours []Colour
		err = db.GlobalSqlxDB.Select(&colours, colourSql, card.ID)
		if err != nil {
			return nil, err
		} else {
			allCards[i].Colour = make([]Colour, len(colours))
			copy(allCards[i].Colour, colours)
		}

		var subTypes []SubType
		err = db.GlobalSqlxDB.Select(&subTypes, subTypeSql, card.ID)
		if err != nil {
			return nil, err
		} else {
			allCards[i].SubType = make([]SubType, len(subTypes))
			copy(allCards[i].SubType, subTypes)
		}
	}

	return allCards, nil
}
