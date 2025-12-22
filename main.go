package main

import (
	"log"
	"os"
	_ "time/tzdata"

	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/diamond-planner/diamond-planner/dp"
	_ "github.com/diamond-planner/diamond-planner/migrations"
	"github.com/subosito/gotenv"
)

// / Loads environment.
//   - Locally: from .env file
//   - Production: from container env
func init() {
	_ = gotenv.Load()

	if os.Getenv("APPLICATION_CONTEXT") == "" {
		log.Fatal("APPLICATION_CONTEXT not set, error loading environment variables")
	}

	if os.Getenv("BSM_API_KEY") == "" || os.Getenv("BSM_API_HOST") == "" {
		log.Fatal("BSM API key or host not set, error loading environment variables")
	}
}

func main() {
	client := bsm.NewAPIClient()
	app := dp.NewDiamondPlanner(client)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
