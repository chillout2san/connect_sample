# 開発に必要なライブラリをインストールする
.PHONY: install
install:
	./tool/install.sh

# サーバを立ち上げる
.PHONY: serve
serve:
	go run ./cmd/main.go

# Go のフォーマットを行う
.PHONY: fmt
fmt:
	gofmt -w ./

# buf のリントを行う
.PHONY: lintbuf
lintbuf:
	buf lint

# buf の生成を行う
.PHONY: genbuf
genbuf:
	buf generate
