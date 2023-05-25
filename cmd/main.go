package main

import "go_restfulapi/commons"

func main() {
	router := commons.Router()
	router.Run(":8080")

}
