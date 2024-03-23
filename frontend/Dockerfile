# node.jsのバージョン指定
FROM node:18.17

# 作業ディレクトリ変更
WORKDIR /src/app

# ホストのpackage.jsonをコンテナにコピー
COPY ./package.json ./

# コンテナに必要なパッケージをインストール
RUN npm install

# プロジェクトのソースコードをコピー
COPY . .

# アプリケーションの起動コマンド
CMD ["npm", "run", "dev"]