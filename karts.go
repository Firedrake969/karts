package karts

import (
    "fmt"
    "os"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"

    "github.com/firedrake969/karts/staticfiles"
    "github.com/firedrake969/karts/views"
)

// This runs everything.  Pass it a mapping of strings (url routes)
// to views.View structs and it will serve both your views and
// staticfiles.
func RunKarts(routes map[string]views.View) {
    fmt.Println("Starting...")
    router := httprouter.New()
    
    staticlist := staticfiles.GetStaticfiles()
    for staticfile := range staticlist {
        router.GET(staticlist[staticfile].Servedpath, staticlist[staticfile].Serve)
    }

    for k := range routes {
        route := routes[k]
        router.GET(k, route.HandleGet)
        router.POST(k, route.HandlePost)
    }

    port := os.Args[1]

    log.Fatal(http.ListenAndServe(":" + port, router))
}