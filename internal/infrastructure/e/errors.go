package e

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

var (
	// ErrNilHTTPClient は指定されたHTTPクライアントがnilであることを意味します。
	ErrNilHTTPClient = errors.New("指定されたHTTPクライアントがnilです")
	// ErrInvalidBaseURL は指定されたベースURLが無効であることを意味します。
	ErrInvalidBaseURL = errors.New("指定されたベースURLが無効です")
	// ErrGenerateRequestHeaders はリクエストヘッダーの生成に失敗したことを意味します。
	// このエラーが発生した場合、APIキーとAPIシークレットを確認する必要があります。
	ErrGenerateRequestHeaders = errors.New("リクエストヘッダーの生成に失敗しました")
	// ErrNoCredentials は指定された資格情報がnilであることを意味します。
	ErrNoCredentials = errors.New("指定された資格情報がnilです")
)

// WithPrefixError はパッケージプレフィックスを付けたエラーを返します。
func WithPrefixError(err error) error {
	const prefix = "coincheck"
	return errors.New(prefix + ": " + err.Error())
}

func Log(err error) {
	const depth = 10
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:]) // skip printStackTrace + its caller
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		fmt.Printf("at %s:%d (%s)\n", frame.File, frame.Line, frame.Function)
		if !more {
			log.Fatalln(err.Error())
		}
	}
}
