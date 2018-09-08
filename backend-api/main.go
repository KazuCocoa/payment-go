package main

import (
	"os"
	"payment-go/backend-api/infrastructure"
)

func main() {
	// export API_SERVER_PORT=:8888
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
