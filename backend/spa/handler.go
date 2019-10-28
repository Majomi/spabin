package spa

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/shurcooL/httpfs/vfsutil"
)

// Handler serves the Index.html if no other static file is found.
type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := vfsutil.Stat(assets, r.URL.Path)
	if os.IsNotExist(err) {
		f, err := vfsutil.ReadFile(assets, "index.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic("no index.html found")
		}
		render.HTML(w, r, string(f))
		return
	}
	http.FileServer(assets).ServeHTTP(w, r)
}
