package grid

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"io"
	"io/fs"
	"strconv"
	"time"
)

type FS interface {
	fs.FS
	OpenID(int64) (File, error)
	Remove(int64) error
	Write(io.Reader, string) (File, error)
}

func NewFS(db *sql.DB) FS {
	return &gridFS{db: db, burst: 60 * 1024}
}

type gridFS struct {
	db    *sql.DB // 数据库连接
	burst int     // 60K
}

// Open implement fs.FS
func (gfs *gridFS) Open(id string) (fs.File, error) {
	fid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fs.ErrInvalid
	}

	return gfs.OpenID(fid)
}

func (gfs *gridFS) OpenID(id int64) (File, error) {
	rawSQL := "SELECT id, `name`, size, sha1, burst, done, created_at, updated_at FROM grid_file WHERE id = ?"
	row := gfs.db.QueryRow(rawSQL, id)
	fl := &file{db: gfs.db}
	if err := row.Scan(&fl.ID, &fl.Filename, &fl.Filesize, &fl.SHA1, &fl.Burst,
		&fl.Done, &fl.CreatedAt, &fl.UpdatedAt); err != nil || !fl.Done {
		return nil, fs.ErrNotExist
	}

	return fl, nil
}

func (gfs *gridFS) Remove(id int64) error {
	tx, err := gfs.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer tx.Rollback()

	deleteFile := "DELETE FROM grid_file WHERE id = ?"
	if _, err = tx.Exec(deleteFile, id); err == nil {
		deletePart := "DELETE FROM grid_part WHERE file_id = ?"
		if _, err = tx.Exec(deletePart, id); err == nil {
			return tx.Commit()
		}
	}

	return err
}

func (gfs *gridFS) Write(r io.Reader, name string) (File, error) {
	// 开启事务
	opt := &sql.TxOptions{Isolation: sql.LevelReadCommitted}
	tx, err := gfs.db.BeginTx(context.Background(), opt)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer tx.Rollback()

	burst := gfs.burst
	createdAt := time.Now()
	insertFile := "INSERT INTO grid_file(`name`, sha1, burst, created_at) VALUE (?, ?, ?, ?)"
	ret, err := tx.Exec(insertFile, name, "", burst, createdAt)
	if err != nil {
		return nil, err
	}
	fileID, err := ret.LastInsertId() // 拿到插入的 ID
	if err != nil {
		return nil, err
	}

	insertPart := "INSERT INTO grid_part (file_id, `serial`, `data`) VALUE (?, ?, ?)"
	buf := make([]byte, burst)

	checksum := sha1.New()
	tr := io.TeeReader(r, checksum)

	var n, serial int
	var filesize int64
	for {
		if n, err = tr.Read(buf); err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		if _, err = tx.Exec(insertPart, fileID, serial, buf[:n]); err != nil {
			break
		}
		serial++
		filesize += int64(n)
	}

	if err == nil {
		sum := hex.EncodeToString(checksum.Sum(nil))
		updatedAt := time.Now()
		updateFile := "UPDATE grid_file SET size = ?, sha1 = ?, done = ?, updated_at = ? WHERE id = ?"
		if _, err = tx.Exec(updateFile, filesize, sum, true, updatedAt, fileID); err == nil {
			if err = tx.Commit(); err == nil {
				fl := &file{
					ID:        fileID,
					Filename:  name,
					Filesize:  filesize,
					SHA1:      sum,
					Burst:     burst,
					Done:      true,
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
					db:        gfs.db,
				}

				return fl, nil
			}
		}
	}

	return nil, err
}
