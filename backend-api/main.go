package main

import (
	"os"

	"./infrastructure"
)

func main() {
	// export API_SERVER_PORT=:8888
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
