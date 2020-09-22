// Package flsflag は Go の flag パッケージがサポートするコマンドライン引数のパースに加えて、環境変数のパースをサポートする。
package flsflag

import (
	"flag"
	"os"
	"strings"
)

func envName(prefix, sep, argument string) string {
	return strings.Join(
		[]string{
			prefix,
			strings.ToUpper(argument),
		},
		sep,
	)
}

// Parse は Go の flag.Parse に加えて、環境変数のパースを行う。
// コマンドライン引数、環境変数の順に優先される。
func Parse(prefix, sep string) {
	flag.VisitAll(func(f *flag.Flag) {
		if s := os.Getenv(envName(prefix, sep, f.Name)); s != "" {
			err := f.Value.Set(s)
			if err != nil {
				panic(err)
			}
		}
	})
	flag.Parse()
}
