package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/appengine"
)

const (
	uidPRE        = "u."                 // ユーザのキーのプリフィックス
	confPRE       = "c."                 // 設定のキーのプリフィックス
	tokenHeader   = "X-Requested-With"   // トークンのヘッダ
	optHeader     = "X-Requested-Option" // オプションのヘッダ
	maxOptNum     = 4                    // オプションの数
	maxNameLength = 64                   // ニックネームの最大長
	maxMesLength  = 512                  // メッセージの最大長
)

var (
	proccesing map[string]bool
	mu         sync.Mutex
)

func main() {
	proccesing = map[string]bool{}
	echoInst := echo.New()
	// CORS対応 (XMLHttpRequest で GET したときに、単純リクエストなのになぜかクロスサイトアクセスが必要)
	echoInst.Use(middleware.CORS())

	// トップページ
	echoInst.GET("/", topPage)
	// 取得 API
	echoInst.GET("/api/v1/:uid", getAPI)
	echoInst.GET("/api/v2/:uid", getAPI2)
	// 更新時刻一覧 API
	echoInst.POST("/api/v2/list", listAPI2)
	// 更新 API
	echoInst.POST("/api/v1/user", updateAPI)

	http.Handle("/", echoInst)
	appengine.Main()
}

// メインページ
func topPage(c echo.Context) error {
	// // TODO: CORS調査
	// w := c.Response()
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	type SignParam struct {
		Acc int
	}
	param := SignParam{
		Acc: addAccess(c.Request()),
	}
	return templateRender(http.StatusOK, "sign", param, c)
}

// 取得 API
func getAPI(c echo.Context) error {
	uid := c.Param("uid")
	note, err := getNote(c, uid)
	if err != nil {
		logOutput(err.Error())
	}
	// 応答
	type GetParam struct {
		ID   string
		Name string
		Mes  []string
		UTC  string
	}
	param := GetParam{
		ID:   uid,
		Name: note.Name,
		Mes:  strings.Split(note.Mes, "\n"),
		UTC:  note.UTC,
	}
	// 最終更新時刻
	// gmtstr, err := ymdhms2GMT(note.UTC)
	// w := c.Response()
	// w.Header().Set("Last-Modified", gmtstr)
	// TODO: 何故か更新結果の読み出しが不味い。原理がわからないのでお蔵入り
	// 多分Googleのキャッシュ制御かなにかが更新されていないと思ってしまう感じ？
	return templateRender(http.StatusOK, "note", param, c)
}

// 取得 API
func getAPI2(c echo.Context) error {
	uid := c.Param("uid")
	note, err := getNote(c, uid)
	if err != nil {
		logOutput(err.Error())
	}
	// 応答
	return c.JSON(http.StatusOK, note)
}

// 更新時刻一覧 API
func listAPI2(c echo.Context) error {
	mu.Lock()
	// リクエストしてきたリモート
	requestIP := c.RealIP()
	if b, ok := proccesing[requestIP]; ok && b {
		mu.Unlock()
		// 処理中にまた来た
		// 不正なリクエストとしてエラー
		mes := "processing"
		logOutput(mes)
		return c.String(http.StatusBadRequest, mes)
	}
	// → 処理中
	proccesing[requestIP] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		// → 処理終了
		proccesing[requestIP] = false
		mu.Unlock()
	}()

	utcs := map[string]string{}

	ids := []string{}
	idsjson := c.FormValue("ids")
	err := json.Unmarshal([]byte(idsjson), &ids)
	if err != nil {
		// 応答
		return c.JSON(http.StatusBadRequest, utcs)
	}

	for _, uid := range ids {
		if note, err := getNote(c, uid); err == nil {
			utcs[uid] = note.UTC
		}
	}

	// 応答
	return c.JSON(http.StatusOK, utcs)
}

// 更新 API
func updateAPI(c echo.Context) error {
	mu.Lock()
	// リクエストしてきたリモート
	requestIP := c.RealIP()
	if b, ok := proccesing[requestIP]; ok && b {
		mu.Unlock()
		// 処理中にまた来た
		// 不正なリクエストとしてエラー
		mes := "processing"
		logOutput(mes)
		return c.String(http.StatusBadRequest, mes)
	}
	// → 処理中
	proccesing[requestIP] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		// → 処理終了
		proccesing[requestIP] = false
		mu.Unlock()
	}()

	// 独自ヘッダの確認
	reqToken := c.Request().Header.Get(tokenHeader)
	if reqToken == "" {
		// 不正なリクエスト
		mes := "need header X-Requested-With"
		logOutput(mes)
		return c.String(http.StatusBadRequest, mes)
	}

	// トークンの認可
	resToken, err := verifyUser(c.Response().Writer, reqToken)
	if err != nil {
		// 認可失敗
		logOutput(err.Error())
		// 連打防止対策 (総当たり攻撃)としてレスポンスに10秒のペナルティ
		// ペナルティ中に同じリモートから再POSTされるとエラーで返している
		time.Sleep(time.Duration(10) * time.Second)
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// 独自ヘッダの確認
	opt := ""
	optToken := c.Request().Header.Get(optHeader)
	if optToken != "" {
		type Opt struct {
			Num int `json:"num"`
		}
		o := Opt{}
		bad := false
		if err = json.Unmarshal([]byte(optToken), &o); err == nil {
			if 1 < o.Num && o.Num < maxOptNum {
				opt = strconv.Itoa(o.Num)
			} else {
				bad = true
			}
		} else {
			bad = true
		}
		if bad {
			// 不正なリクエスト
			mes := "bad header " + optHeader
			logOutput(mes)
			return c.String(http.StatusBadRequest, mes)
		}
	}

	// ノートを保存
	uid := convertUID(c.Request(), opt+resToken.UID)
	name := strcat(c.FormValue("name"), maxNameLength)
	mes := strcat(c.FormValue("mes"), maxMesLength)
	setNote(c, uid, name, mes)

	// レスポンス
	type RetParam struct {
		UID  string `json:"uid"`
		Name string `json:"name"`
		Mes  string `json:"mes"`
	}
	note, _ := getNote(c, uid)
	param := RetParam{
		UID:  uid,
		Name: note.Name,
		Mes:  note.Mes,
	}
	return c.JSON(http.StatusOK, param)
}