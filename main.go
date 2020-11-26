package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/fgunawan1995/julo/wallet/config"
	"github.com/fgunawan1995/julo/wallet/controller"
	"github.com/fgunawan1995/julo/wallet/middleware"
	"github.com/fgunawan1995/julo/wallet/model"
)

func main() {
	gracefulShutdown()

	// Init DB
	db, err := model.ConnectDb(config.GetConfig())
	if err != nil {
		panic(err)
	}
	model.Db = db
	defer db.Close()

	// Init cache
	model.InitCache()

	// Init routes
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/api/v1/init", controller.InitWallet).Methods("POST")
	r.HandleFunc("/api/v1/wallet", middleware.ValidateJWT(controller.GetWallet)).Methods("GET")
	r.HandleFunc("/api/v1/wallet", middleware.ValidateJWT(controller.EnableWallet)).Methods("POST")
	r.HandleFunc("/api/v1/wallet", middleware.ValidateJWT(controller.DisableWallet)).Methods("PATCH")
	r.HandleFunc("/api/v1/wallet/deposits", middleware.ValidateJWT(controller.DepositBalance)).Methods("POST")
	r.HandleFunc("/api/v1/wallet/withdrawals", middleware.ValidateJWT(controller.WithdrawalBalance)).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})
	handler := c.Handler(r)
	fmt.Println("Server started at :3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Shutting down gracefully.")
		err := model.SaveCacheToFile()
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}()
}
