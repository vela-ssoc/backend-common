// Package grid 名字灵感源自 MongoDB 的 GridFS，其实现的功能也类似于
// MongoDB GridFS：将文件分片存储到 MySQL 中，实现了 fs.FS fs.File
// fs.FileInfo 几个接口适用性更强也更利于扩展。
//
// 其实将文件存储到数据库，不是个好的做法：
// 1. 文件上传下载都会占用大量的数据库 IO，与业务 SQL 抢资源
// 2. 但是这些文件是需要 manager broker 节点之间共享的，所以也不能放在
// 某个节点的本地磁盘中，如果放在节点本地磁盘中，需要考虑节点间文件共享
// 方案
// 3. 比较好的办法就是部署一个专门的文件管理服务，比如 minio，那么项目
// 依赖的服务增加，部署时增加复杂度，申请的 IT 资源也更多，不是很符合这
// 个项目简单的特点
// 综合考虑：我们项目中的文件不会很多，文件一般也不大，基本都在几兆到几
// 十兆，为了减少数据库的 IO，还增加了 CDN 功能（即：grid.NewCDN），每
// 个节点都会将文件缓存到本地磁盘，以后的下载就不走数据库了
package grid
