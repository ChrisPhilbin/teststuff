package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"time"
	// "github.com/gocron"
)

// func createnew() string {
// 	t := time.Now()
// 	dirname := t.Format("20060102")
// 	os.Mkdir(dirname, os.FileMode(0666))
// 	return dirname
// }

func main() {
		url := "https://radblast.wunderground.com/cgi-bin/radar/WUNIDS_map?station=DIX&brand=wui&num=1&delay=15&type=N0R&frame=0&scale=1.000&noclutter=0&showstorms=0&mapx=400&mapy=240&centerx=400&centery=240&transx=0&transy=0&showlabels=1&severe=0&rainsnow=0&lightning=0&smooth=0&rand=25369994&lat=40.18000031&lon=-75.54000092&label=Royersford"
		ticker := time.NewTicker(10 * time.Second)
        for range ticker.C {
			t := time.Now()
			timestamp := t.Format("20060102150405")

			response, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
		
			file, err := os.Create("radar"+timestamp+".jpg")
			if err != nil {
				fmt.Println(err)
			}
		
			_, err = io.Copy(file, response.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Success! Saved file")
			}
			
			file.Close()
			}
}