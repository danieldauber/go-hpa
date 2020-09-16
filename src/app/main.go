package main

import (
	"fmt"
	"math"
	"net/http"
)

func loop() float64 {
	x := 0.0001
	for index := 0; index < 100000000; index++ {
		x = x + math.Sqrt(x)
	}
	return x
}

func greeting(text string) string {
	return fmt.Sprintf("<b>%s</b>", text)
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%e\n", loop())
		fmt.Fprintf(w, greeting("Code.education Rocks! com CD"))
	})
	http.ListenAndServe(":80", nil)
}