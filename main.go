package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type NFTFloor struct {
	Floor float64 `json:"floor"`
}

func serverWs(w http.ResponseWriter, r *http.Request) {
	wallet := r.PathValue("wallet")
	slog.Debug(wallet)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("upgrader.Upgrade:", "error", err.Error())
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				slog.Error("ReadMessage:", "error", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			floor, err := GetNFTCollectionFloor(wallet)
			if err != nil {
				slog.Error("GetNFTCollectionFloor error", "wallet", "error", wallet, err)
			}
			slog.Debug("GetNFTCollectionFloor", "floor", floor)
			jsondata, err := json.Marshal(NFTFloor{floor})
			if err != nil {
				slog.Error("json.Marshal", "error", err)
			}
			err = conn.WriteMessage(websocket.TextMessage, jsondata)
			if err != nil {
				slog.Error("WriteMessage", "error", err)
			}
		}

	}
}

func main() {
	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)

	http.HandleFunc("/{wallet}", serverWs)

	defaultPort := ":8080"

	if envPort := os.Getenv("PORT"); envPort != "" {
		defaultPort = ":" + envPort
	}

	slog.Debug("defaultPort", "defaultPort", defaultPort)

	err := http.ListenAndServe(defaultPort, nil)
	if err != nil {
		panic(err)
	}
}
