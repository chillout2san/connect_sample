package tool

// ダブルクォーテーションで囲まれた文字列を検知して install しているので、
// 足さないこと
import (
	_ "connectrpc.com/connect/cmd/protoc-gen-connect-go"
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
