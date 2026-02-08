package server

import (
	"fmt"
	"net/http"
	"time"
)

func SseHandler(downloaderChan chan map[string]int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		w.Header().Set("Access-Control-Allow-Origin", "*")

		clientGone := r.Context().Done()

		rc := http.NewResponseController(w)
		t := time.NewTicker(time.Second)
		defer t.Stop()

		for {
			select {
			case <-clientGone:
				fmt.Println("Client disconnected")
				return
			case event, ok := <-downloaderChan:
				if !ok {
					return
				}
				_, err := fmt.Fprintf(w, "%v\n\n", event)
				if err != nil {
					return
				}
				fmt.Fprintf(w, "data: %v\n\n", event)
				err = rc.Flush()
				if err != nil {
					return
				}
			}
		}
	}
}
