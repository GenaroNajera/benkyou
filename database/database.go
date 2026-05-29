package database

import (
	"database/sql"
	"log"
)

type Kanji struct {
	Db *sql.DB
}

type Data struct {
	ID      int
	Kanji   string
	Meaning string
	Onyomi  string
	Kunyomi string
}

func (k *Kanji) SelectAll() []Data {
	rows, err := k.Db.Query(`SELECT * FROM kyouiku1;`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var row Data
	var data []Data

	for rows.Next() {
		err = rows.Scan(&row.ID, &row.Kanji, &row.Meaning, &row.Onyomi, &row.Kunyomi)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, row)
	}

	return data
}
