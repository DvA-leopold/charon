package methods

import (
	"fmt"
	"net/http"
)

func LoadByURL(out http.ResponseWriter, req *http.Request) {
	fmt.Println("load files " + req.URL.Path)
}
