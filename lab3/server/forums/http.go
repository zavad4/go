package forums

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zavad4/go/tree/main/lab3/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListForums(store, rw)
		} else if r.Method == "POST" {
			handleForumCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleForumCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var f Forum
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateForum(f.ForumName)
	if err == nil {
		tools.WriteJsonOk(rw, &f)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListForums(store *Store, rw http.ResponseWriter) {
	res, err := store.ListForums()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
