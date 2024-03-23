# goのバージョン指定
FROM golang:1.19.2-alpine3.16
# ディレクトリ作成
RUN mkdir /go/src/app
# 作業ディレクトリ変更
WORKDIR /go/src/app
# アプリケーションのソースコードをコピー
COPY app/ .
# 依存関係をインストール
RUN go mod tidy
# main.go実行
CMD ["go", "run", "main.go"]
