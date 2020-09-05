package main

import (
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// verifyUser はIDトークンを検証した結果を返します。
func verifyUser(w http.ResponseWriter, reqToken string) (*auth.Token, error) {
	ctx, cancel := getContext(nil)
	defer cancel()
	app, err := firebase.NewApp(ctx, nil, getOption())
	if err != nil {
		logOutput(err.Error())
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return nil, err
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		logOutput(err.Error())
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return nil, err
	}
	token, err := auth.VerifyIDToken(ctx, reqToken)
	if err != nil {
		logOutput(err.Error())
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return nil, err
	}
	return token, nil
}
