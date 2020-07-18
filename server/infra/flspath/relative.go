package flspath

import (
	"path/filepath"
	"runtime"
	"strings"
)

// RelativePath は呼び出し元のパスとプロジェクトルートパスからはじめて異なるところのパスを返す。
func RelativePath() string {
	_, callFile, _, ok := runtime.Caller(1)
	if !ok {
		return callFile
	}

	return relativePath(callFile)
}

// RelativePathWith は引数のパスとプロジェクトルートパスからはじめて異なるところからのパスを返す。
func RelativePathWith(name string) string {
	return relativePath(name)
}

// ProjectRootPath はプロジェクトのルートパスを返す。
func ProjectRootPath() string {
	_, this, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	return filepath.Dir(filepath.Dir(filepath.Dir(this)))
}

func relativePath(name string) string {
	root := ProjectRootPath()
	if len(root) == 0 {
		return name
	}

	names := strings.Split(name, "/")
	roots := strings.Split(root, "/")
	idx := diffIdx(names, roots)
	if idx == 0 {
		return name
	}

	return strings.Join(names[idx:], "/")
}

// diffIdx は二つのスライスの要素が初めて異なる添字を返す
func diffIdx(s1, s2 []string) int {
	n1, n2 := len(s1), len(s2)
	var i int
	for ; i < n1 && i < n2; i++ {
		if s1[i] != s2[i] {
			break
		}
	}

	return i
}
