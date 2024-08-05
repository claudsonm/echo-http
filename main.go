package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type RequestInfo struct {
    Method        string              `json:"method"`
    URL           string              `json:"url"`
    Headers       map[string][]string `json:"headers"`
    Body          string              `json:"body"`
    Host          string              `json:"host"`
    ContentLength int64               `json:"content_length"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    var body string
    if r.Body != nil {
        bodyBytes, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Unable to read request body", http.StatusInternalServerError)
            return
        }
        body = string(bodyBytes)
    }

    reqInfo := RequestInfo{
        Method:        r.Method,
        URL:           r.URL.String(),
        Headers:       r.Header,
        Body:          body,
        Host:          r.Host,
        ContentLength: r.ContentLength,
    }

    reqInfoJSON, err := json.MarshalIndent(reqInfo, "", "  ")
    if err != nil {
        http.Error(w, "Unable to marshal request info to JSON", http.StatusInternalServerError)
        return
    }

    fmt.Println(string(reqInfoJSON))

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(reqInfoJSON)
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Starting server on :80")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatalf("ListenAndServe failed: %v", err)
    }
}

