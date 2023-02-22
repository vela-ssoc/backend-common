# grid

## SQL 导入

### 文件信息表

```sql
CREATE TABLE `grid_file`
(
    `id`         BIGINT AUTO_INCREMENT COMMENT '文件 ID',
    `name`       TEXT                               NOT NULL COMMENT '文件名',
    `size`       BIGINT   DEFAULT 0                 NOT NULL COMMENT '文件大小',
    `sha1`       CHAR(40)                           NOT NULL COMMENT '文件 SHA1',
    `burst`      INT      DEFAULT 0                 NOT NULL COMMENT '分片大小（单位：bytes，要和 grid_part.data 配合使用）',
    `done`       TINYINT(1) DEFAULT 0 NOT NULL COMMENT '是否上传完毕',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间（一般代表上传开始时间）',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新时间（一般代表上传结束时间）',
    CONSTRAINT grid_file_pk PRIMARY KEY (`id`)
) COMMENT '文件信息表';
```

### 文件分片表

```sql
CREATE TABLE `grid_part`
(
    `id`      BIGINT AUTO_INCREMENT COMMENT '内容分片 ID',
    `file_id` BIGINT           NOT NULL COMMENT '所属文件 ID',
    `serial`  BIGINT DEFAULT 0 NOT NULL COMMENT '分片序号 (0-n)',
    `data`    BLOB             NOT NULL COMMENT '分片内容',
    CONSTRAINT grid_part_pk PRIMARY KEY (`id`),
    CONSTRAINT grid_part_pk2 UNIQUE (`file_id`, `serial`)
) COMMENT '文件分片';
```

## demo

```go
package main

import (
	"database/sql"
	"os"

	"github.com/dfcfw/grid"
)

func main() {
	db := connDB()
	
	// 创建文件系统
	// gfs := grid.NewFS(db)

	// CDN 缓存目录要先创建好
	dir := "/tmp/grid"
	_ = os.RemoveAll(dir)
	_ = os.MkdirAll(dir, 0644)

	// 创建一个带 CDN 缓存的文件管理：
	// dir 为缓存目录，当文件大于 10240 bytes 才会缓存到磁盘
	gfs := grid.NewCDN(db, dir, 10240)
	file, err := gfs.OpenID(1)
	// file 已经实现了 fs.File fs.FileInfo 接口
	
	// TODO 业务逻辑
}

// connDB 连接数据库
func connDB() *sql.DB {
	// TODO 实现连接数据库
	return nil
}
```
