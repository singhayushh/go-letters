package main

import "Goletters/backend"

var server = backend.Server{}

func main() {
	server.Run("8080")
}
