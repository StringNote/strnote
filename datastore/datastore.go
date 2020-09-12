package datastore

import (
	"net/http"

	"github.com/StringNote/strnote/firebase"
	"github.com/StringNote/strnote/util"
)

var fs *firebase.Collection

func init() {
	fs = firebase.NewCollection("Data")
}

// SetValueCache はキーに対する値を記録します。
func SetValueCache(r *http.Request, key, value string) error {
	err := setMemcache(r, key, []byte(value))
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}

	setMapcache(key, value)

	return nil
}

// SetValue はキーに対する値を記録します。
func SetValue(r *http.Request, key, value string) error {
	// Firestore に書き込む
	err := fs.Set(key, value)
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}

	return SetValueCache(r, key, value)
}

// GetValue はキーに対しての値を取得します。
func GetValue(r *http.Request, key string) (string, error) {
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
	value, err = fs.Get(key)
	if err != nil {
		// Firestore にもなかった
		util.LogOutput(err.Error())
		return "", err
	}

	// キャッシュに書き込む
	setMapcache(key, value)
	// logOutput("memcache.Set")
	err = setMemcache(r, key, []byte(value))
	if err != nil {
		util.LogOutput(err.Error())
		return "", err
	}

	return value, err
}

// DeleteValue はキーのデータを削除します。
func DeleteValue(r *http.Request, key string) error {
	deleteMapcache(key)
	deleteMemcache(r, key)
	fs.Delete(key)
	return nil
}
