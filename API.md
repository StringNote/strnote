# 参照時 API

作成日：2020/09/11  

StringNote で登録したノートを参照するには、大きく分けて二種類の呼び出し方法があります。  

* 埋め込みタグで表示する HTML-API
* JavaScript で表示する JSON-API

の二種類です。  
以下はその解説になります。  

## GET <https://strnote.com/api/v1/{{.uid}}>

埋め込みタグで表示する HTML-API です。  
UID に対応したノートの内容を編集したページを返します。  

iframeタグで埋め込むことで、静的なページ内にノートを表示させることができます。  
埋め込みタグの例としてはトップページで "your note embedded:" として表示されています。  

``` HTML
    <iframe src="https://strnote.com/api/v1/4B_4CSj0U_nhWa6YCbzB1w"
   style="border: 0px; width: 750px; height: 32em;"></iframe>
```

## GET <https://strnote.com/api/v2/{{.uid}}>

JavaScript で表示する JSON-API です。  
UID に対応したノートの内容を JSON で返します。  

* name：更新時に名付けた自分の名前です。
* mes：ノートの本文です。
* utc：ノートを更新した時に自動的に付与された更新時刻(UTC)です。

```json
{"name":"StringNote","mes":"こんにちは、StringNote の進捗や障害などの通知ノートです。\n　\nこのようにノートをブログなどに埋め込むことで、リアルタイムな情報を発信することが可能になります。\n　\n負荷が上がった時に安定稼働できるかどうか、その調整が済むまでは早期アクセス期間とします。\n　\n【作業】　\n\n現在\n・画面調整\n\n次\n・ドキュメント\n\n将来\n・広告の適用\n　\n--------\n【注意】\nTwitter で複数のアカウントをご利用の場合。\nWEB版でのTwitterで現在選択しているアカウントでのログインになります。\n【注意】\n現在、トップページの調整中です。\n一時的に操作などに支障をきたす場合があるかもしれませんが、ご了承ください。","utc":"20200908050817"}
```

## POST <https://strnote.com/api/v2/list>

"ids" というフォームデータで UID の配列を JSON 形式で送信します。  
戻り値として、UID をキーとした更新時刻(YYYYMMDDHHMMSS)を返します。  
更新時刻は UTC 時刻ですので日本時刻などに変換が必要です。  

```json
{"4B_4CSj0U_nhWa6YCbzB1w":"20200908050817"}
```

## JSON-API 利用ライブラリ

JSON として返される応答はそのままiframe表示しても、おそらくは意味が分かりません。  
ですので、ご利用の際は API 呼び出しをラップした [node.js](https://strnote.com/static/notes.js) を使用してください。  
note.js は MIT ライセンスですので修正はご自由に。  

[view3.html](https://strnote.com/static/view3.html) は、このライブラリにより作られています。  

### ライブラリ読み込み

```js
    <script type="text/javascript" src="https://strnote.com/static/notes.js"></script>
```

### インスタンス作成

```js
    /** notes define */
    var notes = new Notes();
```

### ノートの反映

notes.Add(UID, コールバック関数) としてノートを取得した時の処理を記述します。  
戻り値はノートのインスタンスです。  

```js
        /** refresh note */
        let note = notes.Add(id, function(note) {
            /** display code */
            nameSpan.innerText = note.name + " (" + note.date + ")";
            noteSpan.innerText = note.mes;
        });
```

### ノートの即時更新

note.Refresh() を呼び出すと、即時に取得～表示更新処理が呼ばれます。  
イベントリスナなどで使用する場合には bind() する必要があります。  

```js
        relBtn.addEventListener("click", note.Refresh.bind(note));
```

### 定期チェックの設定

指定のマイクロ秒ごとに notes.Add() した全てのノートの表示を更新します。  

```js
        /** periodic refresh */
        let interval = 5 * 60 * 1000;
        notes.setInterval(interval);
```

（以上）
