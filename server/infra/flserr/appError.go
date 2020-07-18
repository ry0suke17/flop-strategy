// Package flserr はアプリエラーを表すパッケージ。
package flserr

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/yneee/flop-strategy/infra/flspath"
)

type appError struct {
	cause      error
	annotation string
	file       string
	line       int
}

func (e *appError) Error() string {
	base := fmt.Sprintf("%s at %s line %d", e.cause, flspath.RelativePathWith(e.file), e.line)
	if e.annotation == "" {
		return base
	}
	return e.annotation + ": " + base
}

func (e *appError) Cause() error { return e.cause }

// Wrap は cause をラップした error 値を生成する。
func Wrap(cause error) error {
	return wrap(cause, "")
}

// Wrapf は追加の状況説明を付与して cause をラップした error 値を生成する。
func Wrapf(cause error, annotation string, a ...interface{}) error {
	return wrap(cause, fmt.Sprintf(annotation, a...))
}

// New は errors.New と同等の関数で、 error 値を生成する
func New(message string) error {
	return wrap(errors.New(message), "")
}

// Errorf は fmt.Errorf と同等の関数で、 error 値を生成する
func Errorf(format string, a ...interface{}) error {
	return wrap(fmt.Errorf(format, a...), "")
}

// Cause はラップされた error 値を返す。
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

func wrap(cause error, annotation string) error {
	if cause == nil {
		cause = Errorf("[BUG] got cause nil, but want not nil, annotation=%s", annotation)
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return cause
	}
	return &appError{cause, annotation, file, line}
}
