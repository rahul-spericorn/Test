package handlers

import (
	"encoding/json"
	"log"
	"loosidAPI/db"
	"loosidAPI/generated"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetListings - handler for GET /listings
func (si ServerWrapper) GetListings(ctx echo.Context) error {
	var listings generated.Listings

	err := db.ReadAllListings(&listings)

	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	out, err := json.Marshal(listings)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	json.Unmarshal(out, &listings)

	if len(listings) == 0 {
		return ctx.JSON(http.StatusNotFound, generated.Error{0, "No listings found"})
	}

	return ctx.JSON(http.StatusOK, listings)
}

// AddListings - handler for POST /listings
func (si ServerWrapper) AddListings(ctx echo.Context) error {
	var listings = new(generated.Listings)

	if err := ctx.Bind(listings); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	for _, listing := range *listings {
		db.InsertListing(&listing)
	}

	return ctx.JSON(http.StatusOK, nil)
}

// AddListing - handler for POST /listings/listing
func (si ServerWrapper) AddListing(ctx echo.Context) error {
	listing := new(generated.Listing)

	if err := ctx.Bind(listing); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	db.InsertListing(listing)

	return ctx.JSON(http.StatusOK, nil)
}

// GetListing - handler for GET /listings/listing/{ListingID}
func (si ServerWrapper) GetListing(ctx echo.Context, listingID string) error {
	var listing generated.Listing

	err := db.GetListingByID(&listing, listingID)

	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	out, err := json.Marshal(listing)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	json.Unmarshal(out, &listing)

	if listing.ListingID == "" {
		return ctx.JSON(http.StatusNotFound, generated.Error{0, "Listing not found"})
	}

	return ctx.JSON(http.StatusOK, listing)
}

// SearchListings - handler for GET /listings/search
func (si ServerWrapper) SearchListings(ctx echo.Context, params generated.SearchListingsParams) error {
	var listings generated.Listings

	err := db.SearchListings(&listings, params)

	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	out, err := json.Marshal(listings)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	json.Unmarshal(out, &listings)

	if len(listings) == 0 {
		return ctx.JSON(http.StatusNotFound, generated.Error{0, "No listings found"})
	}

	return ctx.JSON(http.StatusOK, listings)
}
