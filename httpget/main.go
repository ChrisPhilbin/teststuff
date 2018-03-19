package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    )

func main() {
	response, err := http.Get("http://www.google.com")
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        defer response.Body.Close()
        pagecontent, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Printf("%s\n", string(pagecontent))
    }
}