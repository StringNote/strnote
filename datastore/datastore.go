package datastore

import (
	"net/http"

	"github.com/StringNote/strnote/firebase"
	"github.com/StringNote/strnote/util"
)

// DataStore はデータを管理します。
type DataStore struct {
	name string
	fs   *firebase.Collection
}

// NewDataStore は DataStore を生成します。
func NewDataStore(name string) *DataStore {
	return &DataStore{name: name + "/", fs: firebase.NewCollection(name)}
}

// SetValueCache はキーに対する値を記録します。
func (d *DataStore) SetValueCache(r *http.Request, key, value string) error {
	cachekey := d.name + key
	err := setMemcache(r, cachekey, []byte(value))
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}

	setMapcache(cachekey, value)

	return nil
}

// SetValue はキーに対する値を記録します。
func (d *DataStore) SetValue(r *http.Request, key, value string) error {
	// Firestore に書き込む
	err := d.fs.Set(key, value)
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}

	return d.SetValueCache(r, key, value)
}

// Keys はキーの一覧を返します。
func (d *DataStore) Keys(r *http.Request) []string {
	return d.fs.Keys()
}

// GetValue はキーに対しての値を取得します。
func (d *DataStore) GetValue(r *http.Request, key string) (string, error) {
	cachekey := d.name + key
	value, err := getMapcache(cachekey)
	if value != "" {
		// キャッシュにあったので返す
		return value, err
	}

	// キャッシュを検索
	// logOutput("memcache.Get")
	value, err = getMemcache(r, cachekey)
	if value != "" {
		// キャッシュにあったので返す
		setMapcache(cachekey, value)
		return value, err
	}

	// キャッシュにないので Firestore から取得する
	// logOutput("getFirestore")
	value, err = d.fs.Get(key)
	if err != nil {
		// Firestore にもなかった
		util.LogOutput(err.Error())
		return "", err
	}

	// キャッシュに書き込む
	setMapcache(cachekey, value)
	// logOutput("memcache.Set")
	err = setMemcache(r, cachekey, []byte(value))
	if err != nil {
		util.LogOutput(err.Error())
		return "", err
	}

	return value, err
}

// DeleteValue はキーのデータを削除します。
func (d *DataStore) DeleteValue(r *http.Request, key string) error {
	cachekey := d.name + key
	deleteMapcache(cachekey)
	deleteMemcache(r, cachekey)
	d.fs.Delete(key)
	return nil
}

// Keys はキーの一覧を返します。
func (d *DataStore) Keys() []string {
	return d.fs.Keys()
}