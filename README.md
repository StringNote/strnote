
# strnote

URL: [https://strnote.com](https://strnote.com)  
MAIL: [contact@strnote.com](mailto:contact@strnote.com)  

## 概要

* gmail などのアカウントでログインし、短文（ノート）を登録します。
* アカウントでログインすれば、即座にノートを登録できます。
* アカウントひとつにつき４個のノートが登録できます。

ノートはユーザ識別子(UID)と、それを基とした埋め込みタグを持ちます。

* 埋め込みタグはブログやホームページに埋め込んで使用します。
* また[ビューア](http://strnote.com/static/view3.html)に複数の UID を記録することで動向表のような使い方ができます。

## 入会／退会

アカウントで利用者を判別しています。  
そのため、本サービス自身には入会／退会という手順はありません。  
利用を停止したい場合には、ノートをクリアしてください。  
サービスのデータベースからノートの情報を消去します。  
ただしアカウントは同じですので、利用を再開しても以前に振り出されたUIDは変わりません。  

## 利用方法１

ブログやホームページにおいては、払いだされた埋め込みタグを HTML 内に埋め込むことでノートの表示を行うことが可能です。  

``` HTML
    <iframe src="https://strnote.com/api/v1/4B_4CSj0U_nhWa6YCbzB1w"></iframe>
```

HTML に慣れている方は、styleで表示を調整できます。  
以下はトップページで表示されている公式アナウンスです。  

``` HTML
    <iframe src="https://strnote.com/api/v1/4B_4CSj0U_nhWa6YCbzB1w"
   style="border: 0px; width: 750px; height: 32em;"></iframe>
```

## 利用方法２

また JavaScript の知識が必要になりますが、JSON-API を使用するならば、自由に見た目を変更することも可能です。  

API に関しては[参照時 API](API.md) をご覧ください。

## ライセンス

**本ソースコードは、StringNote のセキュリティ検証用に公開するものであり、クローンサイトなどを許可する目的のものではありません。**  
したがって、ライセンスは存在しませんのでご了承ください。  

なお、個々のソースコードの部品に関して、サンプルとしてのご利用は自由です。  

[english version]  
**This source code is for security verification of StringNote, and is not intended to allow clone sites and the like.**  
Therefore, please note that there is no license.  

You are free to use the individual source code components as samples.  

## 利用パッケージ

StringNoteでは以下のパッケージを利用させて頂いております。  

* [Google APP engine](https://github.com/golang/appengine) : Apache License 2.0  
* [Firebase](https://github.com/firebase/firebase-admin-go) : Apache License 2.0  
* [Echo](https://github.com/labstack/echo) : MIT  

（以上）
