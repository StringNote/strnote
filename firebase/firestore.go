package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/StringNote/strnote/util"
)

// Collection はfirestoreのコレクションを表します。
type Collection struct {
	collection string
}

// NewCollection は Collecction のインスタンスを生成します。
func NewCollection(collection string) *Collection {
	return &Collection{collection: collection}
}

// Set は Collection に値を保存します。
func (f *Collection) Set(key, value string) error {
	util.LogOutput(fmt.Sprintf("key:%s, value:%s", key, value))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app, err := firebase.NewApp(ctx, nil, GetOption())
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}
	defer client.Close()
	// 書き出し
	data := make(map[string]interface{})
	data["data"] = value
	util.LogOutput(fmt.Sprintf("%+v", data))
	_, err = client.Collection(f.collection).Doc(key).Set(ctx, data)
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}
	util.LogOutput("done")
	return nil
}

// Get は Collection でキーから値を検索します。
func (f *Collection) Get(key string) (string, error) {
	util.LogOutput(fmt.Sprintf("key:%s", key))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app, err := firebase.NewApp(ctx, nil, GetOption())
	if err != nil {
		util.LogOutput(err.Error())
		return "", err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		util.LogOutput(err.Error())
		return "", err
	}
	defer client.Close()
	// 読み出し
	dsnap, err := client.Collection(f.collection).Doc(key).Get(ctx)
	if err != nil {
		util.LogOutput(err.Error())
		return "", err
	}
	if dsnap.Exists() == false {
		util.LogOutput("return not exist")
		return "", nil
	}
	data := dsnap.Data()
	util.LogOutput(fmt.Sprintf("%+v", data))
	if txt, ok := data["data"].(string); ok {
		util.LogOutput(fmt.Sprintf("return \"%s\"", txt))
		return txt, nil
	}
	util.LogOutput("return empty")
	return "", nil
}

// Delete は Collection からキーを削除します。
func (f *Collection) Delete(key string) error {
	util.LogOutput(fmt.Sprintf("key:%s", key))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app, err := firebase.NewApp(ctx, nil, GetOption())
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}
	defer client.Close()
	// 削除
	_, err = client.Collection(f.collection).Doc(key).Delete(ctx)
	if err != nil {
		util.LogOutput(err.Error())
		return err
	}
	return err
}
