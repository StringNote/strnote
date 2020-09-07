package main

import (
	"fmt"

	firebase "firebase.google.com/go"
)

func setFirestore(key, value string) error {
	logOutput(fmt.Sprintf("key:%s, value:%s", key, value))
	ctx, cancel := getContext(nil)
	defer cancel()
	app, err := firebase.NewApp(ctx, nil, getOption())
	if err != nil {
		logOutput(err.Error())
		return err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		logOutput(err.Error())
		return err
	}
	defer client.Close()
	// 書き出し
	data := make(map[string]interface{})
	data["data"] = value
	logOutput(fmt.Sprintf("%+v", data))
	_, err = client.Collection("Data").Doc(key).Set(ctx, data)
	if err != nil {
		logOutput(err.Error())
		return err
	}
	logOutput("done")
	return nil
}

func getFirestore(key string) (string, error) {
	logOutput(fmt.Sprintf("key:%s", key))
	ctx, cancel := getContext(nil)
	defer cancel()
	app, err := firebase.NewApp(ctx, nil, getOption())
	if err != nil {
		logOutput(err.Error())
		return "", err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		logOutput(err.Error())
		return "", err
	}
	defer client.Close()
	// 読み出し
	dsnap, err := client.Collection("Data").Doc(key).Get(ctx)
	if err != nil {
		logOutput(err.Error())
		return "", err
	}
	if dsnap.Exists() == false {
		logOutput("return not exist")
		return "", nil
	}
	data := dsnap.Data()
	logOutput(fmt.Sprintf("%+v", data))
	if txt, ok := data["data"].(string); ok {
		logOutput(fmt.Sprintf("return \"%s\"", txt))
		return txt, nil
	}
	logOutput("return empty")
	return "", nil
}

func deleteFirestore(key string) error {
	logOutput(fmt.Sprintf("key:%s", key))
	ctx, cancel := getContext(nil)
	defer cancel()
	app, err := firebase.NewApp(ctx, nil, getOption())
	if err != nil {
		logOutput(err.Error())
		return err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		logOutput(err.Error())
		return err
	}
	defer client.Close()
	// 削除
	_, err = client.Collection("Data").Doc(key).Delete(ctx)
	if err != nil {
		logOutput(err.Error())
		return err
	}
	return err
}
