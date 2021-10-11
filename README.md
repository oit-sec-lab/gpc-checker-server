# dnt-verify-server
GPCChecker(https://github.com/oit-sec-lab/GPCChecker) のサーバープログラムです。
<br>
現在外部公開用のリソースで公開していないので、ローカルで利用してもらう事が前提となっています。

## 起動手順
本リポジトリをクローン後、リポジトリ内にdb用の.envを作成しdockerのビルドを行ってください。
<br>
httpsで利用する場合、後述するようにdocker-compose.ymlを変更する必要があります。

```
# .env
DB_NAME=(任意のDB_NAME)
DB_USER=(任意のユーザーネーム)
DB_PASS=(任意のパスワード)
TZ=Asia/Tokyo
```

```
docker-compose build
docker-compose up -d
```

serverのシェルに入った後、go.modの設定を行います。

```
docker-compose exec server /bin/sh
go mod init
go mod tidy
```

正常終了を確認後、server.goを起動します。
```
go run server.go
```

## 利用手順
ローカルで利用する場合、GPCChecker(https://github.com/oit-sec-lab/GPCChecker) のcontent.jsを一行目のURLを
<br>
**localhost:8080/sites** に変更する必要があります。

```
// content.js

let URL = "http://localhost:8080/sites"
```

httpsで確認する場合、docker-compose.ymlの9行目を任意のドメインに変更し、etc/hostsファイルにdockerを起動しているマシンのIPに対して任意のドメインを紐づけてください。

```
# docker-compose.yml

(省略)
DOMAINS: "(任意のドメイン) -> http://server:8080/"
```
```
# /etc/hosts (winであればC:\Windows\System32\drivers\etc\hosts)

(dockerを起動しているマシンのIP) (任意のドメイン)

(例)
192.168.0.5 test.cyanos
```
```
// content.js

let URL = "(任意のドメイン)"
```
その後、GPCCheckerをchromeの拡張として読み込ませると利用が可能です。
<br>
詳しくはGPCChecker(https://github.com/oit-sec-lab/GPCChecker) のREADMEを参照してください。
