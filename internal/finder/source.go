package finder

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/charlievieth/fastwalk"
	"github.com/mitchellh/go-homedir"
)

type Source struct {
	Path  string `json:"path"`
	Depth uint8  `json:"depth"`
}

func (s *Source) Find() ([]string, error) {
	expanded, err := homedir.Expand(s.Path)

	if err != nil {
		return nil, err
	}

	s.Path = strings.TrimSuffix(expanded, "/")

	if s.Depth == 0 {
		return s.depthZero()
	}

	if s.Depth == 1 {
		return s.depthOne()
	}

	return s.depthGreater()
}

func (s *Source) depthZero() ([]string, error) {
	isDir, err := isPathDir(s.Path)

	if err != nil {
		return nil, err
	}

	if isDir {
		return []string{s.Path}, nil
	}

	return nil, nil
}

func (s *Source) depthOne() ([]string, error) {
	entries, err := os.ReadDir(s.Path)

	if err != nil {
		return nil, err
	}

	paths := []string{s.Path}

	for _, entry := range entries {
		path := filepath.Join(s.Path, entry.Name())
		isDir, err := isPathDir(path)

		if err != nil {
			return nil, err
		}

		if isDir {
			paths = append(paths, path)
		}
	}

	return paths, nil
}

func (s *Source) depthGreater() ([]string, error) {
	paths := make([]string, 0)
	var mu sync.Mutex

	walkFn := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if currentDepth(s.Path, path) > s.Depth {
			return fs.SkipDir
		}

		mu.Lock()
		defer mu.Unlock()

		paths = append(paths, path)

		return err
	}

	err := fastwalk.Walk(
		&fastwalk.Config{Follow: true},
		s.Path,
		walkFn,
	)

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func currentDepth(root, curr string) uint8 {
	relPath, _ := filepath.Rel(root, curr)
	return uint8(len(strings.Split(relPath, "/")))
}

func isPathDir(path string) (bool, error) {
	info, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}
