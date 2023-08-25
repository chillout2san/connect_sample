# 開発に必要なライブラリをインストールする
.PHONY: install
install:
	./tool/install.sh

# フォーマットを行う
.PHONY: fmt
fmt:
	gofmt -w ./