package webserver

import (
	"fmt"
	"net/http"
)

func StartWebServer() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Host)
	http.ServeFile(res, req, "./webserver/index.html")
}
