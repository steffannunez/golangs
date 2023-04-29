package main

import (
	"fmt"
	"net/http"

	"github.com/steffannunez/golangs/pkg/handlers/handlers.go"
)

const puerto = ":8080"

func main() {
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			n, err := fmt.Fprintf(w, "Holiwis")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(fmt.Sprintf("Bytes: %d", n))
		})

		_ = http.ListenAndServe(":8080", nil)*/

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Inicializando aplicación en el puerto %s", puerto))
	_ = http.ListenAndServe(puerto, nil)
}
