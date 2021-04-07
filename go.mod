module hex-go

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/bitly/go-simplejson v0.5.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/iancoleman/strcase v0.1.2
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.9
)

replace hex-go/internal/datasource => ./internal/datasource

replace hex-go/internal/usecase => ./internal/usecase

replace hex-go/internal/port => ./internal/port

replace hex-go/internal/domain => ./internal/domain

replace hex-go/internal/repository => ./internal/repository

replace hex-go/internal/router => ./internal/router

replace hex-go/internal/model => ./internal/model

replace hex-go/http/controller => ./http/controller

replace hex-go/http/response => ./http/response
