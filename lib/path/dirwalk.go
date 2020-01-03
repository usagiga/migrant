package path

import (
	"io/ioutil"
	"path/filepath"
)

// Get all files recursive in specific dir.
// https://qiita.com/tanksuzuki/items/7866768c36e13f09eedb
func Dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}