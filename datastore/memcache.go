package datastore

import (
	"net/http"

	"github.com/StringNote/strnote/util"
	"google.golang.org/appengine/memcache"
)

// キャッシュに書き込む
func setMemcache(r *http.Request, key string, value []byte) error {
	ctx, cancel := util.GetContext(r)
	defer cancel()
	item := &memcache.Item{
		Key:   key,
		Value: value,
	}
	return memcache.Set(ctx, item)
}

// キャッシュから読み込む
func getMemcache(r *http.Request, key string) (string, error) {
	ctx, cancel := util.GetContext(r)
	defer cancel()
	item, err := memcache.Get(ctx, key)
	if item != nil {
		return string(item.Value), err
	}
	return "", err
}

// キャッシュから削除する
func deleteMemcache(r *http.Request, key string) error {
	ctx, cancel := util.GetContext(r)
	defer cancel()
	return memcache.Delete(ctx, key)
}
