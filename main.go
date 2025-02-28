package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64)
// }

func main() {
	// const port uint16 = 3000 // one way to create variables -> now, using 'Load' Method
	config.Load()
	router := router.Create() // another way to create variables

	fmt.Printf("Escutando na porta %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)) // concatenate strings

}
