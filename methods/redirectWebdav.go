package methods

import (
	"fmt"
	"net/http"
)

func RedirectWebdav(out http.ResponseWriter, req *http.Request) {
	var newIP = "79.137.175.159:8080" // take from database
	var newRedirectURI = "http://" + newIP + "/charon" + req.URL.Path[7:]
	fmt.Println("new URI: " + newRedirectURI)
	http.Redirect(out, req, newRedirectURI, 307)
}
