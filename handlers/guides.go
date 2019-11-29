package handlers

import (
	"encoding/json"
	"log"
	"loosidAPI/db"
	"loosidAPI/generated"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetGuides - handler for GET /guides
func (si ServerWrapper) GetGuides(ctx echo.Context) error {
	var guides generated.Guides

	err := db.ReadAllGuides(&guides)

	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	out, err := json.Marshal(guides)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	json.Unmarshal(out, &guides)

	if len(guides) == 0 {
		return ctx.JSON(http.StatusNotFound, generated.Error{0, "No guides found"})
	}

	return ctx.JSON(http.StatusOK, guides)
}

// AddGuide - handler for POST /guides/guide/
func (si ServerWrapper) AddGuide(ctx echo.Context) error {
	guide := new(generated.Guide)

	if err := ctx.Bind(guide); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	db.InsertGuide(guide)

	return ctx.JSON(http.StatusOK, nil)
}

// AddGuides - handler for POST /guides
func (si ServerWrapper) AddGuides(ctx echo.Context) error {
	var guides = new(generated.Guides)

	if err := ctx.Bind(guides); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err500)
	}

	for _, guide := range *guides {
		db.InsertGuide(&guide)
	}

	return ctx.JSON(http.StatusOK, nil)
}
