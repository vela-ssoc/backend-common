package pubody

import (
	"sort"
	"strings"
	"time"
)

// Folder 目录列表
type Folder struct {
	Abs       string    `json:"abs"`       // 当前绝对路径
	Separator string    `json:"separator"` // 当前路径分隔符
	Items     FileItems `json:"items"`     // 文件列表
}

type FileItem struct {
	Path  string    `json:"path"`
	Name  string    `json:"name"`
	Size  int64     `json:"size"`
	Mtime time.Time `json:"mtime"`
	Dir   bool      `json:"dir"`
	Ext   string    `json:"ext"` // 扩展名
	Mode  string    `json:"mode"`
}

type FileItems []*FileItem

func (fs FileItems) NameAsc() FileItems {
	sort.Slice(fs, func(i, j int) bool { return strings.Compare(fs[i].Name, fs[j].Name) < 0 })
	return fs
}

func (fs FileItems) NameDesc() FileItems {
	sort.Slice(fs, func(i, j int) bool { return strings.Compare(fs[i].Name, fs[j].Name) > 0 })
	return fs
}

func (fs FileItems) SizeAsc() FileItems {
	sort.Slice(fs, func(i, j int) bool { return fs[i].Size < fs[j].Size })
	return fs
}

func (fs FileItems) SizeDesc() FileItems {
	sort.Slice(fs, func(i, j int) bool { return fs[i].Size > fs[j].Size })
	return fs
}

func (fs FileItems) Human() FileItems {
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
