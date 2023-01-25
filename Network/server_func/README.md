ServeFile
===
This function takes in a ResponseWriter and a Request and tries to read the 
given file by trying to get the id from the mux.Vars() 
function and if detected it reads the file using ioutil.ReadFile() function. 
If no file could be found it renders the 404 not found error using the http.Error() function.
If a valid file is found it set the content type using the mime.TypeByExtension() and filepath.Ext() 
functions and then writes it to the ResponseWriter using the w.Write() function
