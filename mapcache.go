package main

import "fmt"

var (
	mapcache    map[string]string
	notMatchKey error
)

func init() {
	notMatchKey = fmt.Errorf("not match key")
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
		return "", notMatchKey
	}
	if val, ok := mapcache[key]; ok {
		return val, nil
	}
	return "", notMatchKey
}

func deleteMapcache(key string) error {
	delete(mapcache, key)
	return nil
}
