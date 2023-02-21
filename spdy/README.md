# SPDY

A multiplexed stream library.

尚未实现 `流量控制`、`stream 并发数控制`、`stream 读写超时控制`

- [Example](https://github.com/dfcfw/spdy-example)

## 帧格式

```text
0                   1                   2                   3
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|      Flag     |                  Stream ID                    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|   Stream ID   |          Data Length          |     Data      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Data                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

Flag: `uint8`

Stream ID: `uint32`

Data Length: `uint16`

Data: 变长，由 `Data Length` 决定

### SYN - 新建连接

SYN 为变长帧，代表新建虚拟连接

### FIN - 结束连接

FIN 为虚拟连接的最后一帧，收到 FIN 则代表对方已经断开了虚拟连接，

FIN 帧为定长帧（7 bytes），只能包含 `Flag` `Stream ID` `Data Length` 信息，且 `Data Length` 填充为 `0`

### DAT - 数据报文

## 参考链接

[spdystream](https://github.com/moby/spdystream)

[yamux](https://github.com/hashicorp/yamux)

[smux](https://github.com/xtaci/smux)

[muxado](https://github.com/inconshreveable/muxado)

[multiplex](https://github.com/whyrusleeping/go-smux-multiplex)
