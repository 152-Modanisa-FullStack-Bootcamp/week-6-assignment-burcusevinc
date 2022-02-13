package main

import (
	config2 "go.mod/config"
	handler2 "go.mod/handler"
	"go.mod/repository"
	service2 "go.mod/service"
	"net/http"
)

func main() {
	//it returns the config struct
	config := config2.NewConfig()

	store := repository.NewUserLocalStorage()

	service := service2.NewService(store, config.InitialBalance, config.MinumumBalance)

	handler := handler2.NewWalletHandler(service)

	http.HandleFunc("/", handler.WalletMethods)
	//listen port:
	http.ListenAndServe(":8080", nil)

}
