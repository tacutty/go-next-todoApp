# Goのバージョン指定（1.23を使用）
FROM golang:1.23

# 必要なパッケージをインストールし、Airをインストール
RUN apt-get update && apt-get install -y curl && \
    go install github.com/air-verse/air@latest

# ディレクトリ作成
RUN mkdir -p /go/src/app

# 作業ディレクトリ変更
WORKDIR /go/src/app


# アプリケーションのソースコードをコピー
COPY app/ .

# 依存関係をインストール
RUN go mod tidy

# Airを使ってアプリケーションをホットリロード
CMD ["air", "-c", ".air.toml"]
