package main

import (
	"net/http"
)

func main() {
	// This would be the way to create the FileServer without the built-in function
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open("../public" + r.URL.Path)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
			}

			defer f.Close()
			var contentType string

			switch {
			case strings.HasSuffix(r.URL.Path, "css"):
				contentType = "text/css"
			case strings.HasSuffix(r.URL.Path, "html"):
				contentType = "text/html"
			case strings.HasSuffix(r.URL.Path, "png"):
				contentType = "image/png"
			}

			w.Header().Add("Content-Type", contentType)
			io.Copy(w, f)
		})

		http.ListenAndServe(":8090", nil)
	*/

	// However, all of this can be done much easier
	http.ListenAndServe(":8080", http.FileServer(http.Dir("public")))
}
