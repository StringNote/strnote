package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"google.golang.org/appengine"
)

// ログ出力
func logOutput(mes string) {
	called := GetCallerPC().String()
	log.Output(2, called+": "+mes)
}

// CalledPC は呼び出し元を示す識別コードです。
type CalledPC uintptr

// String はstringerです。
func (c CalledPC) String() string {
	pc := uintptr(c)
	fpc := runtime.FuncForPC(pc)
	n, l := fpc.FileLine(pc)
	return fmt.Sprintf("%s (%s:%d)", fpc.Name(), n, l)
}

// GetCallerPC はそれを呼び出した関数の呼び出し元の PC を取得します。
func GetCallerPC() CalledPC {
	pc, _, _, _ := runtime.Caller(2)
	return CalledPC(pc)
}

// getContext はキャンセルContextを新規生成します
func getContext(r *http.Request) (context.Context, context.CancelFunc) {
	var ctx context.Context
	if r != nil {
		ctx = appengine.NewContext(r)
	} else {
		ctx = context.Background()
	}
	return context.WithCancel(ctx)
}

func strcat(str string, size int) string {
	rn := []rune(str)
	len := len(rn)
	if len > size {
		len = size
	}
	return string(rn[:len])
}

// func portStr() string {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	log.Printf("port %s", port)
// 	return port
// }

const ymdFMT = "20060102150405" // 日付書式（golangの決まり）

func time2ymdhms(t time.Time) string {
	return t.Format(ymdFMT)
}
func ymdhms2time(ymdhms string) (time.Time, error) {
	return time.Parse(ymdFMT, ymdhms)
}
func ymdhms2GMT(ymdhms string) (string, error) {
	utc, err := ymdhms2time(ymdhms)
	if err != nil {
		return "", err
	}
	headtime := utc.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	return headtime, nil
}
