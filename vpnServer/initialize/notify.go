package initialize

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func NotifyBackend(server net.Listener){
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigs
		fmt.Println(fmt.Sprintf("接收到信号量 [%v]", sig))
		done <- true
	}()

	fmt.Println("等待信号量的触发")
	<-done
	fmt.Println("信号量被触发了。。。。。")

	err := server.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("关闭失败 err %v", err))
	}

	fmt.Println("exiting--------")

}
