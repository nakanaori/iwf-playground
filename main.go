package main

import "iwf-playground/server"

func main() {
	router := server.GetRouter()
	router.Run(":8803")
}
