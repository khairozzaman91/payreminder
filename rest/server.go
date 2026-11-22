package rest

import (
	"fmt"
	"net/http"
	"payreminder/config"
	userhandler "payreminder/rest/handler/user"
	"payreminder/rest/middleware"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Server(cnf *config.Config) {

	manager := middleware.NewManager()
	manager.Use(
		middleware.Logger,
		middleware.Recovery,
		middleware.CORSMiddleware,
	)

	mux := http.NewServeMux()
	// Swagger UI
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	mux.Handle("POST /users", http.HandlerFunc(userhandler.CreateUser))

	wrappedMux := manager.WrapMux(mux)

	fmt.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", wrappedMux)

	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
}
