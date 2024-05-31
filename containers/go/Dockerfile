FROM golang:1.21

# 作業ディレクトリを設定
WORKDIR /usr/local/go/src/go-party-box/

# Airツールをインストール
RUN go install github.com/cosmtrek/air@v1.42.0

# go.mod と go.sum をコピーして依存関係をダウンロード
COPY application/go.mod application/go.sum ./
RUN go mod download

# Prisma Clientを事前にキャッシュするために prefetch を実行
RUN go run github.com/steebchen/prisma-client-go prefetch

# アプリケーションのソースコードをコピーしてビルド
COPY application/. /usr/local/go/src/go-party-box/

# Prisma Clientを生成
RUN go run github.com/steebchen/prisma-client-go generate
