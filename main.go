package main

import (
	"fmt"

	"pweb-api.abdanhafidz.com/config"
	"pweb-api.abdanhafidz.com/router"
)

func main() {
	fmt.Println("Server started on ", config.TCP_ADDRESS, ", port :", config.HOST_PORT)
	router.StartService()

}
