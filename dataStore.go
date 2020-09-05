package main

import (
	"net/http"
)

// setValue はキーに対する値を記録します。
func setValue(r *http.Request, key, value string) error {
	// Firestore に書き込む
	// logOutput("setFirestore")
	err := setFirestore(key, value)
	if err != nil {
		logOutput(err.Error())
		return err
	}

	// logOutput("memcache.Set")
	err = setMemcache(r, key, []byte(value))
	if err != nil {
		logOutput(err.Error())
		return err
	}

	setMapcache(key, value)

	return nil
}

// getValue はキーに対しての値を取得します。
func getValue(r *http.Request, key string) (string, error) {
	value, err := getMapcache(key)
	if value != "" {
		// キャッシュにあったので返す
		return value, err
	}

	// キャッシュを検索
	// logOutput("memcache.Get")
	value, err = getMemcache(r, key)
	if value != "" {
		// キャッシュにあったので返す
		setMapcache(key, value)
		return value, err
	}

	// キャッシュにないので Firestore から取得する
	// logOutput("getFirestore")
	value, err = getFirestore(key)
	if err != nil {
		// Firestore にもなかった
		logOutput(err.Error())
		return "", err
	}

	// キャッシュに書き込む
	setMapcache(key, value)
	// logOutput("memcache.Set")
	err = setMemcache(r, key, []byte(value))
	if err != nil {
		logOutput(err.Error())
		return "", err
	}

	return value, err
}
