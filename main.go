package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yuki-maruyama/actions-with-cache/handlers"
	"github.com/yuki-maruyama/actions-with-cache/logger"
	"github.com/yuki-maruyama/actions-with-cache/metrics"

	_ "github.com/aws/aws-sdk-go-v2/service/amplify"
	_ "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	_ "github.com/aws/aws-sdk-go-v2/service/appconfig"
	_ "github.com/aws/aws-sdk-go-v2/service/appconfigdata"
	_ "github.com/aws/aws-sdk-go-v2/service/appmesh"
	_ "github.com/aws/aws-sdk-go-v2/service/apprunner"
	_ "github.com/aws/aws-sdk-go-v2/service/athena"
	_ "github.com/aws/aws-sdk-go-v2/service/batch"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := handlers.NewHandler(logger.Logger, metrics.HTTPRequests)

	r := mux.NewRouter()
	r.HandleFunc("/health", handler.HealthHandler).Methods("GET")
	r.HandleFunc("/hello", handler.HelloHandler).Methods("GET")
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")

	logger.Logger.WithField("port", port).Info("Server starting")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
