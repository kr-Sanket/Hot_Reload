package loghub

import (
	"fmt"
	"net/http"
)

func StartServer(hub *Hub) {

	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		ch := hub.Subscribe()

		for {
			msg := <-ch

			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dashboard.html")
	})

	go http.ListenAndServe(":8090", nil)
}