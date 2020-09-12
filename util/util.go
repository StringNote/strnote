package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"google.golang.org/appengine"
)

// LogOutput はログを出力します。
func LogOutput(mes string) {
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

// GetContext はキャンセルContextを新規生成します
func GetContext(r *http.Request) (context.Context, context.CancelFunc) {
	var ctx context.Context
	if r != nil {
		ctx = appengine.NewContext(r)
	} else {
		ctx = context.Background()
	}
	return context.WithCancel(ctx)
}

// Strcat は文字列を切り詰めます。
func Strcat(str string, size int) string {
	rn := []rune(str)
	len := len(rn)
	if len > size {
		len = size
	}
	return string(rn[:len])
}

const ymdFMT = "20060102150405" // 日付書式（golangの決まり）

// Time2ymdhms はTimeをyyyymmddhhmmss形式文字列に変換します。
func Time2ymdhms(t time.Time) string {
	return t.Format(ymdFMT)
}

// Ymdhms2time はyyyymmddhhmmss形式文字列をTimeに変換します。
func Ymdhms2time(ymdhms string) (time.Time, error) {
	return time.Parse(ymdFMT, ymdhms)
}

// Ymdhms2GMT はyyyymmddhhmmss形式文字列をHTTPでのGMT表記に変換します。
func Ymdhms2GMT(ymdhms string) (string, error) {
	utc, err := Ymdhms2time(ymdhms)
	if err != nil {
		return "", err
	}
	headtime := utc.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	return headtime, nil
}

func slice(str string, slen int) []string {
	ret := []string{}
	ls := len(str)
	start := 0
	for start < ls {
		end := start + slen
		if ls < end {
			end = ls
		}
		ret = append(ret, str[start:end])
		start += slen
	}
	return ret
}
func splitEnd(str string, size int) []string {
	ret := []string{}
	maxlen := len(str)
	start := maxlen % size
	if start != 0 {
		ret = append(ret, str[:start])
	}
	if start < maxlen {
		ret = append(ret, slice(str[start:], size)...)
	}
	return ret
}

// Comma は数字をカンマ区切り文字列に変換します。
func Comma(num int) string {
	return strings.Join(splitEnd(strconv.Itoa(num), 3), ",")
}
