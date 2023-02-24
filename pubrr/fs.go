package pubrr

import (
	"sort"
	"strings"
	"time"
)

type FS struct {
	Abs   string `json:"abs"`
	Files Files  `json:"files"`
}

type File struct {
	Path  string    `json:"path"`
	Name  string    `json:"name"`
	Size  int64     `json:"size"`
	Mtime time.Time `json:"mtime"`
	Dir   bool      `json:"dir"`
	Mode  string    `json:"mode"`
}

type Files []*File

func (fs Files) NameAsc() Files {
	sort.Slice(fs, func(i, j int) bool { return strings.Compare(fs[i].Name, fs[j].Name) < 0 })
	return fs
}

func (fs Files) NameDesc() Files {
	sort.Slice(fs, func(i, j int) bool { return strings.Compare(fs[i].Name, fs[j].Name) > 0 })
	return fs
}

func (fs Files) SizeAsc() Files {
	sort.Slice(fs, func(i, j int) bool { return fs[i].Size < fs[j].Size })
	return fs
}

func (fs Files) SizeDesc() Files {
	sort.Slice(fs, func(i, j int) bool { return fs[i].Size > fs[j].Size })
	return fs
}

func (fs Files) Human() Files {
	sort.Slice(fs, func(i, j int) bool {
		ifs, jfs := fs[i], fs[j]
		if ifs.Dir && !jfs.Dir {
			return true
		} else if !ifs.Dir && jfs.Dir {
			return false
		}

		return strings.Compare(ifs.Name, jfs.Name) < 0
	})
	return fs
}
