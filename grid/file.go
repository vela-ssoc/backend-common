package grid

import (
	"database/sql"
	"io"
	"io/fs"
	"mime"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type File interface {
	fs.File
	fs.FileInfo
	Checksum() string
	ContentType() string
	ContentLength() string
	Attachment() string
}

// file 文件信息表（MySQL）
// CREATE TABLE `grid_file`
// (
//
//	`id`         BIGINT AUTO_INCREMENT COMMENT '文件 ID',
//	`name`       TEXT                                 NOT NULL COMMENT '文件名',
//	`size`       BIGINT     DEFAULT 0                 NOT NULL COMMENT '文件大小',
//	`sha1`       CHAR(40)                             NOT NULL COMMENT '文件 SHA1',
//	`burst`      INT        DEFAULT 0                 NOT NULL COMMENT '分片大小（单位：bytes，要和 grid_part.data 配合使用）',
//	`done`       TINYINT(1) DEFAULT 0                 NOT NULL COMMENT '是否上传完毕',
//	`created_at` DATETIME   DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间（一般代表上传开始时间）',
//	`updated_at` DATETIME   DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新时间（一般代表上传结束时间）',
//	CONSTRAINT grid_file_pk PRIMARY KEY (`id`)
//
// ) COMMENT '文件信息表';
type file struct {
	ID        int64     `json:"id,string"` // BIGINT 自增
	Filename  string    `json:"name"`
	Filesize  int64     `json:"size"`
	SHA1      string    `json:"sha1"`
	Burst     int       `json:"burst"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	db     *sql.DB // 数据库连接
	serial int64   // 分片序号
	buffer []byte  // 缓存
	eof    bool    // 是否读完了
}

// Checksum 文件校验码
func (fl *file) Checksum() string {
	return fl.SHA1
}

func (fl *file) ContentType() string {
	ct := mime.TypeByExtension(filepath.Ext(fl.Filename))
	// 当 Content-Type 是 text/html 时，就算设置了 Content-Length 长度，
	// 浏览器下载的时候也不会正常显示进度条，改一下 Content-Type 即可。
	if ct == "" || strings.HasPrefix(ct, "text/html") {
		ct = "application/octet-stream"
	}

	return ct
}

func (fl *file) ContentLength() string {
	return strconv.FormatInt(fl.Filesize, 10)
}

func (fl *file) Attachment() string {
	pam := map[string]string{"filename": fl.Filename}

	return mime.FormatMediaType("attachment", pam)
}

func (fl *file) Close() error               { return nil }
func (fl *file) Name() string               { return fl.Filename }
func (fl *file) Size() int64                { return fl.Filesize }
func (fl *file) Mode() fs.FileMode          { return 0o444 }
func (fl *file) ModTime() time.Time         { return fl.UpdatedAt }
func (fl *file) IsDir() bool                { return false }
func (fl *file) Sys() any                   { return nil }
func (fl *file) Stat() (fs.FileInfo, error) { return fl, nil }

func (fl *file) Read(p []byte) (int, error) {
	if fl.db == nil {
		return 0, fs.ErrInvalid
	}
	if fl.eof {
		return 0, io.EOF
	}

	var n int
	psz := len(p)
	for !fl.eof && psz > n {
		if len(fl.buffer) == 0 {
			if err := fl.readPart(); err != nil {
				break
			}
		}

		i := copy(p[n:], fl.buffer)
		fl.buffer = fl.buffer[i:]
		n += i
	}
	if n > 0 {
		return n, nil
	}

	return n, io.EOF
}

// readPart 读取数据分片
func (fl *file) readPart() error {
	queryPart := "SELECT `data` FROM grid_part WHERE file_id = ? AND `serial` = ?"
	var pt part
	if err := fl.db.QueryRow(queryPart, fl.ID, fl.serial).
		Scan(&pt.Data); err != nil {
		fl.eof = true
		return io.EOF
	}

	fl.serial++
	fl.buffer = pt.Data

	return nil
}

// part 文件分片（MySQL）
// CREATE TABLE `grid_part`
// (
//
//	`id`      BIGINT AUTO_INCREMENT COMMENT '内容分片 ID',
//	`file_id` BIGINT           NOT NULL COMMENT '所属文件 ID',
//	`serial`  BIGINT DEFAULT 0 NOT NULL COMMENT '分片序号 (0-n)',
//	`data`    BLOB             NOT NULL COMMENT '分片内容',
//	CONSTRAINT grid_part_pk PRIMARY KEY (`id`),
//	CONSTRAINT grid_part_pk2 UNIQUE (`file_id`, `serial`)
//
// ) COMMENT '文件分片';
type part struct {
	ID     int64  // 分片 ID （无业务意义）
	FileID int64  // 所属文件 ID
	Serial int64  // 分片序号 （0-n）
	Data   []byte // 分片内容
}
