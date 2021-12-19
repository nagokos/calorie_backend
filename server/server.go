package server

import (
	"net/http"

	"github.com/nagokos/calorie_backend/config"
	"github.com/nagokos/calorie_backend/controllers"
)

func StartWebServer() error {
	http.HandleFunc("/api/v1/users", controllers.UserCreate)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
