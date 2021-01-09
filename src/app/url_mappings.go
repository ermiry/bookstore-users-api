package app

import (
	"../controllers"
)

func mapURLs() {
	router.GET("/ping", controllers.Ping)
}
