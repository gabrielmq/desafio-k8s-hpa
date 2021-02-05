package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func greeting(msg string) string {
	return fmt.Sprintf("<b>%s</b>", msg)
}

func loop() float64 {
	x := 0.0001
	for i := 0; i <= 100000000; i++ {
		x = x + math.Sqrt(x)
	}
	return x
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		loop()
		fmt.Fprintf(w, greeting("Code.education Rocks!"))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
