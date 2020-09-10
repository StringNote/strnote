
# strnote

## 概要

StringNote は Google などにより認証されたユーザが、短文(以下、ノート)を登録し、登録時に払いだされたユーザ識別子(以下、UID)によって第三者が登録されたノートを参照できるものとするサービスです。  

## 利用方法

ブログやホームページにおいては、UID と同時に払いだされた埋め込みタグを HTML 内に埋め込むことでノートの表示を行うことが可能です。  
JavaScript の知識が必要になりますが、JSON-API を使用するならば、自由に見た目を変更することも可能です。  

## サービス

URL: [https://strnote.com](https://strnote.com)  
MAIL: [contact@strnote.com](mailto:contact@strnote.com)  

## ビューア

static/view.html および view2.html は、グループとしたいメンバーのUIDを登録し、API を用いて一覧表示を行うサンプルとなっています。  
コード内に記述してありますように MIT ライセンスでの参考実装ですので、改造したものを公開することなどは自由です。  

## ライセンス

本ソースコードは、StringNote のセキュリティ検証用に公開するものであり、クローンサイトなどを許可する目的のものではありません。したがって、ライセンスは存在しませんのでご了承ください。  

なお、個々のソースコードの部品に関しては著作権は存在しないという見解です。  
よって、サンプルとしてのご利用は自由です。  

[english version]  
This source code is for security verification of StringNote, and is not intended to allow clone sites and the like. Therefore, please note that there is no license.  

In addition, for the individual components of the source code is the view that the copyright does not exist.  
Therefore, you are free to use them as samples.  

## 利用パッケージ

StringNoteでは以下のパッケージを利用させて頂いております。  

* [Google APP engine](https://github.com/golang/appengine) : Apache License 2.0  
* [Firebase](https://github.com/firebase/firebase-admin-go) : Apache License 2.0  
* [Echo](https://github.com/labstack/echo) : MIT  

以上
