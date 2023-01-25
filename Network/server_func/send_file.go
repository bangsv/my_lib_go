package send_files

import (
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

// func ServeFile() handles requests to serve a particular file. 
// It takes a ResponseWriter (w) and a Request (r) as arguments. 
func ServeFile(w http.ResponseWriter, r *http.Request) {
    // extract the "id" parameter from the URL
    vars := mux.Vars(r)
    id := vars["id"]
    
    // read the file into memory with ioutil.ReadFile 
    file, err := ioutil.ReadFile(id)
    
    // if an error occurs, return a 404 not found response
    if err != nil {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    
    // set the Content-Type of the response based on the file extension
    contentType := mime.TypeByExtension(filepath.Ext(id))
    w.Header().Set("Content-Type", contentType)
    
    // write the file to the ResponseWriter
    w.Write(file)
}
