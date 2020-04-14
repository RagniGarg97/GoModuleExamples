package main

import (
    "net/http"
    "my_project/logger"
)

func main() {
    l := new(logger.Logger)

    http.Handle("/valar-morghulis", loggerware(l, http.HandlerFunc(handle)))
    http.ListenAndServe(":5600", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("valar dohaeris"))
}

func loggerware(l *logger.Logger, next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        requestPath :=  r.URL
        l.LogInfo("Request Made To Server Two: ", requestPath)
    })
}
