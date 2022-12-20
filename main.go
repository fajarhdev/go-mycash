package main

import (
	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/migrate"
	"github.com/fajarhdev/go-mycash/routes"
	_ "github.com/qodrorid/godaemon"
)

func init() {
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
	migrate.Migrate()
}

func main() {
	r := routes.InitRouter()
	r.Run(":3000")
}
