FROM golang:1.15.5

ENV GO111MODULE on

RUN apt-get update && apt-get install vim -y
RUN go get -u github.com/cosmtrek/air \
  && go get -u github.com/swaggo/swag/cmd/swag \
  && go get -u golang.org/x/lint/golint \
  && go get github.com/rubenv/sql-migrate/...

CMD air