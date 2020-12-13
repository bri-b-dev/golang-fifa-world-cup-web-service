package handlers

import (
	"github.com/gin-gonic/gin"
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler godoc
// @Summary returns an empty body status code
// @Success 204
// @Router / [get]
func RootHandler(c *gin.Context) {
	rootHandler(c.Writer, c.Request)
}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners godoc
// @Summary returns winners from the list
// @Produce json
// @Success 200
// @Router /winners [get]
func ListWinners(c *gin.Context) {
	listWinners(c.Writer, c.Request)
}

func listWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	year := req.URL.Query().Get("year")

	if year == "" {
		winners, err := data.ListAllJSON()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Write(winners)
	} else {
		filteredWinners, err := data.ListAllByYear(year)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		res.Write(filteredWinners)
	}
}

// AddNewWinner godoc
// @Summary adds new winner to the list
// @Produce json
// @Success 201
// @Router /winners [post]
func AddNewWinner(c *gin.Context) {
	addNewWinner(c.Writer, c.Request)
}

func addNewWinner(res http.ResponseWriter, req *http.Request) {

	accessToken := req.Header.Get("X-ACCESS-TOKEN")
	isTokenValid := data.IsAccessTokenValid(accessToken)

	if !isTokenValid {
		res.WriteHeader(http.StatusUnauthorized)
	} else {
		if err := data.AddNewWinner(req.Body); err != nil {
			res.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		res.WriteHeader(http.StatusCreated)
	}
}

// WinnersHandler godoc
// @Summary is the dispatcher for all /winners URL
// @Produce json
// @Success 201
// @Router /winners [any]
func WinnersHandler(c *gin.Context) {
	winnersHandler(c.Writer, c.Request)
}

func winnersHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		listWinners(res, req)
	case http.MethodPost:
		addNewWinner(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
