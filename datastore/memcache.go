package datastore

import (
	"context"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

// キャッシュに書き込む
func setMemcache(r *http.Request, key string, value []byte) error {
	ctx, cancel := context.WithCancel(appengine.NewContext(r))
	defer cancel()
	item := &memcache.Item{
		Key:   key,
		Value: value,
	}
	return memcache.Set(ctx, item)
}

// キャッシュから読み込む
func getMemcache(r *http.Request, key string) (string, error) {
	ctx, cancel := context.WithCancel(appengine.NewContext(r))
	defer cancel()
	item, err := memcache.Get(ctx, key)
	if item != nil {
		return string(item.Value), err
	}
	return "", err
}

// キャッシュから削除する
func deleteMemcache(r *http.Request, key string) error {
	ctx, cancel := context.WithCancel(appengine.NewContext(r))
	defer cancel()
	return memcache.Delete(ctx, key)
}
