# Go Hexagonal Architecture
ヘキサゴナルアーキテクチャ勉強用。

## 立ち上げ
```
$ docker-compose build
$ docker-compose up -d // airを実行してホットリロードしたい場合は -d オプション無し
$ docker-compose exec app bash
$ make init
$ make migrate
$ make seed
$ go run main.go
```

アクセス
```
http://localhost:8080/v1/articles?offset=0&limit=10
```

### ボリューム削除
```
$ docker-compose down --volumes
```

## MySQL
```
$ docker-compose exec db bash
$ mysql -u root -p // root
```
