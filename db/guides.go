package db

import (
	"log"
	"loosidAPI/generated"
)

func initGuide() {
	createDB := `
	CREATE TABLE IF NOT EXISTS guides (
		GuideID SERIAL PRIMARY KEY,
		GuideName TEXT UNIQUE NOT NULL,
		GuideDescription TEXT NOT NULL
	  );
	`

	_, err := DbConnection.Exec(createDB)
	if err != nil {
		panic(err)
	}
}

// InsertGuide - function to insert data
func InsertGuide(data *generated.Guide) {
	sqlQuery := `
	INSERT INTO guides 
		(GuideName, GuideDescription)
		VALUES ($1, $2)
	`

	_, err := DbConnection.Exec(sqlQuery, data.GuideName, data.GuideDescription)

	if err != nil {
		log.Println(err)
	}
}

// ReadAllGuides - function to read all guides
func ReadAllGuides(guides *generated.Guides) error {
	sqlQuery := `SELECT * FROM guides`

	rows, err := DbConnection.Query(sqlQuery)

	defer rows.Close()

	for rows.Next() {
		guide := generated.Guide{}
		err = rows.Scan(
			&guide.GuideID,
			&guide.GuideName,
			&guide.GuideDescription,
		)
		if err != nil {
			return err
		}
		*guides = append(*guides, guide)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
