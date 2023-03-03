package syscmd

import (
	"context"
	"log"
	"os/exec"

	"github.com/gorilla/websocket"
)

func ShellWithContext(ctx context.Context, ws *websocket.Conn) error {
	sc := &socketConn{ws: ws}
	name := shellEnv()

	cmd := exec.CommandContext(ctx, name)
	cmd.Stdin = sc
	cmd.Stdout = sc
	cmd.Stderr = sc

	err := cmd.Run()
	log.Printf("cmd 执行结果：%v", err)

	return err
}
