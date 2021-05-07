package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

// Manager holds a list of the crypto and stocks we are watching
type Manager struct {
	Watching map[string]*Board
	Cache    *redis.Client
	Context  context.Context
	sync.RWMutex
}

// NewManager stores all the information about the current stocks being watched and
// listens for api requests on 8080
func NewManager(port string, cache *redis.Client, context context.Context) *Manager {
	m := &Manager{
		Watching: make(map[string]*Board),
		Cache:    cache,
		Context:  context,
	}

	r := mux.NewRouter()
	r.HandleFunc("/tickerboard", m.AddBoard).Methods("POST")
	r.HandleFunc("/tickerboard/{id}", m.DeleteBoard).Methods("DELETE")
	r.HandleFunc("/tickerboard", m.GetBoards).Methods("GET")

	srv := &http.Server{
		Addr:         "localhost:" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	logger.Infof("Starting api server on %s...", port)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	return m
}

// BoardRequest represents the json coming in from the request
type BoardRequest struct {
	Items      []string `json:"items"`
	Token      string   `json:"token"`
	Name       string   `json:"name"`
	Header     string   `json:"header"`
	Nickname   bool     `json:"set_nickname"`
	Crypto     bool     `json:"crypto"`
	Color      bool     `json:"set_color"`
	Percentage bool     `json:"percentage"`
	Arrows     bool     `json:"arrows"`
	Frequency  int      `json:"frequency"`
}

// AddBoard adds a new board to the list of what to watch
func (m *Manager) AddBoard(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a ticker")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading body: %v", err)
		return
	}

	// unmarshal into struct
	var boardReq BoardRequest
	if err := json.Unmarshal(body, &boardReq); err != nil {
		logger.Errorf("Error unmarshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error unmarshalling: %v", err)
		return
	}

	// ensure token is set
	if boardReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: token required")
		return
	}

	// ensure name is set
	if boardReq.Name == "" {
		logger.Error("Board Name required")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: name required")
		return
	}

	// add stock or crypto ticker
	if boardReq.Crypto {

		// check if already existing
		if _, ok := m.Watching[strings.ToUpper(boardReq.Name)]; ok {
			logger.Error("Error: board already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		crypto := NewCryptoBoard(boardReq.Items, boardReq.Token, boardReq.Name, boardReq.Header, boardReq.Nickname, boardReq.Color, boardReq.Percentage, boardReq.Arrows, boardReq.Frequency, m.Cache, m.Context)
		m.addBoard(crypto)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(crypto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	// check if already existing
	if _, ok := m.Watching[strings.ToUpper(boardReq.Name)]; ok {
		logger.Error("Error: board already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	stock := NewStockBoard(boardReq.Items, boardReq.Token, boardReq.Name, boardReq.Header, boardReq.Nickname, boardReq.Color, boardReq.Percentage, boardReq.Arrows, boardReq.Frequency)
	m.addBoard(stock)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stock)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addBoard(b *Board) {
	m.Watching[b.Name] = b
}

// DeleteBoard removes a board
func (m *Manager) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a board")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.Watching[id]; !ok {
		logger.Error("Error: no ticker found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: ticker not found")
		return
	}
	// send shutdown sign
	m.Watching[id].Shutdown()

	// remove from cache
	delete(m.Watching, id)

	logger.Infof("Deleted board %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetBoards returns a list of what the manager is watching
func (m *Manager) GetBoards(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.Watching); err != nil {
		logger.Errorf("Error serving request: %v", err)
		fmt.Fprintf(w, "Error: %v", err)
	}
}
