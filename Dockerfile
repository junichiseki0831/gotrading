FROM golang:1.15.2-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /gotrading
# ワーキングディレクトリの設定
WORKDIR /gotrading
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /gotrading