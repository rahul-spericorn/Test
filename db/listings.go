package db

import (
	"database/sql"
	"fmt"
	"log"
	"loosidAPI/generated"
)

func initListing() {
	createDB := `
	CREATE TABLE IF NOT EXISTS listings (
		ListingID SERIAL PRIMARY KEY,
		ListingName TEXT NOT NULL,
		ListingDescription TEXT,
		GuideName TEXT REFERENCES guides(GuideName),
		ListingCategory TEXT,
		Brands TEXT,
		Cost TEXT,
		Email TEXT,
		EndsAt TEXT,
		ExternalImageUrl TEXT,
		ExternalUrl1 TEXT,
		ListingTags TEXT,
		LocationLabel TEXT,
		Phone TEXT,
		PremiumLevel TEXT,
		StartsAt TEXT,
		VenueName TEXT,
		Address1 TEXT,
		Address2 TEXT,
		City TEXT NOT NULL,
		Country TEXT,
		Latitude TEXT,
		Longitude TEXT,
		State TEXT NOT NULL,
		Zip TEXT NOT NULL
	  );
	`

	_, err := DbConnection.Exec(createDB)
	if err != nil {
		panic(err)
	}
}

// InsertListing - function to insert data
func InsertListing(data *generated.Listing) {
	sqlQuery := `
	INSERT INTO listings (
			GuideName, ListingCategory, Brands, Cost, Email, EndsAt,
			ExternalImageUrl, ExternalUrl1, ListingDescription,
			ListingName, ListingTags, LocationLabel, Phone,
			PremiumLevel, StartsAt, VenueName, Address1, Address2,
			City, Country, Latitude, Longitude, State, Zip
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24
		)
	`

	_, err := DbConnection.Exec(
		sqlQuery, data.GuideName, data.ListingCategory, data.Brands, data.Cost, data.Email,
		data.EndsAt, data.ExternalImageUrl, data.ExternalUrl1, data.ListingDescription,
		data.ListingName, data.ListingTags, data.LocationLabel, data.Phone, data.PremiumLevel,
		data.StartsAt, data.VenueName, data.Address1, data.Address2, data.City, data.Country,
		data.Latitude, data.Longitude, data.State, data.Zip,
	)

	if err != nil {
		log.Println(err)
	}
}

// ReadAllListings - function to read all listings
func ReadAllListings(listings *generated.Listings) error {
	sqlQuery := `SELECT * FROM listings`

	rows, err := DbConnection.Query(sqlQuery)

	defer rows.Close()

	for rows.Next() {
		listing := generated.Listing{}
		err = scanListingRows(rows, &listing)
		if err != nil {
			return err
		}
		*listings = append(*listings, listing)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

// GetListingByID - function to read all listings
func GetListingByID(listing *generated.Listing, listingID string) error {
	sqlQuery := fmt.Sprintf("SELECT * FROM listings WHERE listings.ListingID=%s", listingID)

	rows, err := DbConnection.Query(sqlQuery)

	defer rows.Close()

	for rows.Next() {
		err = scanListingRows(rows, listing)
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

// SearchListings - function to search listings table
func SearchListings(listings *generated.Listings, params generated.SearchListingsParams) error {
	sqlQuery := `SELECT * FROM listings`

	if params.GuideName != nil {
		fmt.Println(*params.GuideName)
		// sqlQuery = sqlQuery + " WHERE listings.GuideName=" + *params.GuideName
	}
	if params.Offset != nil {
		sqlQuery = fmt.Sprintf(sqlQuery+" OFFSET %s", *params.Offset)
	}
	if params.Keywords != nil {
		fmt.Println(*params.Keywords)
	}
	if params.Limit != nil {
		sqlQuery = fmt.Sprintf(sqlQuery+" LIMIT %s", *params.Limit)
	}

	rows, err := DbConnection.Query(sqlQuery)

	defer rows.Close()

	for rows.Next() {
		listing := generated.Listing{}
		err = scanListingRows(rows, &listing)
		if err != nil {
			return err
		}
		*listings = append(*listings, listing)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func scanListingRows(rows *sql.Rows, listing *generated.Listing) error {
	err := rows.Scan(
		&listing.ListingID,
		&listing.ListingName,
		&listing.ListingDescription,
		&listing.GuideName,
		&listing.ListingCategory,
		&listing.Brands,
		&listing.Cost,
		&listing.Email,
		&listing.EndsAt,
		&listing.ExternalImageUrl,
		&listing.ExternalUrl1,
		&listing.ListingTags,
		&listing.LocationLabel,
		&listing.Phone,
		&listing.PremiumLevel,
		&listing.StartsAt,
		&listing.VenueName,
		&listing.Address1,
		&listing.Address2,
		&listing.City,
		&listing.Country,
		&listing.Latitude,
		&listing.Longitude,
		&listing.State,
		&listing.Zip,
	)
	if err != nil {
		return err
	}

	return nil
}
