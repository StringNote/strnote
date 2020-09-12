package datastore

import "fmt"

var (
	mapcache map[string]string
	// NotMatchKey はキーにマッチする値が見つからなかった場合のエラーです。
	NotMatchKey error
)

func init() {
	NotMatchKey = fmt.Errorf("not match key")
}

func setMapcache(key, value string) error {
	if mapcache == nil {
		mapcache = map[string]string{}
	}
	mapcache[key] = value
	return nil
}

func getMapcache(key string) (string, error) {
	if mapcache == nil {
		return "", NotMatchKey
	}
	if val, ok := mapcache[key]; ok {
		return val, nil
	}
	return "", NotMatchKey
}

func deleteMapcache(key string) error {
	delete(mapcache, key)
	return nil
}
