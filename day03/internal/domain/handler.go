package domain

import (
	"encoding/json"
	"html/template"
	"net/http"
	"search/internal/db"
	"search/internal/types"
	"strconv"
	"strings"
)

type Server struct {
	store *db.ElasticStore
}

func NewServer(store *db.ElasticStore) *Server {
	return &Server{store: store}
}

type Data struct {
	Places       []types.Place
	Total        int
	IsPrevious   bool
	IsNext       bool
	PreviousPage int
	NextPage     int
	LastPage     int
}

func (server *Server) GetPlacesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := server.PrepareData(r)
	if err.Message != "" {
		http.Error(w, err.Message, http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/template/template.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

type JsonResponse struct {
	Name     string        `json:"name"`
	Total    int           `json:"total"`
	Places   []types.Place `json:"places"`
	PrevPage int           `json:"prev_page"`
	NextPage int           `json:"next_page"`
	LastPage int           `json:"last_page"`
}

type CustomError struct {
	Message string `json:"error"`
}

func (server *Server) GetPlacesJSONHandler(w http.ResponseWriter, r *http.Request) {
	data, custom := server.PrepareData(r)
	if custom.Message != "" {
		if strings.Contains(custom.Message, "Invalid 'page' value") {
			WriteJSON(w, CustomError{custom.Message})
		} else {
			http.Error(w, custom.Message, http.StatusInternalServerError)
		}
		return
	}

	response := JsonResponse{
		Name:     "Places",
		Total:    data.Total,
		Places:   data.Places,
		PrevPage: data.PreviousPage,
		NextPage: data.NextPage,
		LastPage: data.LastPage,
	}

	WriteJSON(w, response)
}

func WriteJSON(w http.ResponseWriter, response interface{}) {
	jsonResp, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Error encoding JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResp)
}

func (server *Server) PrepareData(r *http.Request) (Data, CustomError) {
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		return Data{}, CustomError{"Invalid 'page' value " + pageParam}
	}

	limit := 10
	offset := (page - 1) * limit

	places, total, err := server.store.GetPlaces(limit, offset)
	if err != nil {
		return Data{}, CustomError{"Error fetching places: " + err.Error()}
	}

	answer := Data{
		Places: places,
		Total:  total,
	}

	if offset > 0 {
		answer.IsPrevious = true
		answer.PreviousPage = page - 1
	}

	if offset+limit < total {
		answer.IsNext = true
		answer.NextPage = page + 1
	}

	answer.LastPage = (total + limit - 1) / limit

	if page > answer.LastPage {
		return Data{}, CustomError{"Invalid 'page' value " + pageParam}
	}

	return answer, CustomError{}
}
