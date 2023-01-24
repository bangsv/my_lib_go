package send_files

import (
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	file, err := ioutil.ReadFile(id)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	contentType := mime.TypeByExtension(filepath.Ext(id))
	w.Header().Set("Content-Type", contentType)
	w.Write(file)
}
