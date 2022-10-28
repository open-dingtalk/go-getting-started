package main

import (
	"flag"
	"fmt"
	"github.com/open-dingtalk/go-getting-started/pkg/config"
	"github.com/open-dingtalk/go-getting-started/pkg/logger"
	"github.com/open-dingtalk/go-getting-started/pkg/router"
	"time"
)

/**
 * @Author linya.jj
 * @Date 2022/9/21 5:00 PM
 */

var (
	flagConfigFilepath string
	flagPort           int
	flagStaticFilePath string
)

func main() {
	var err error

	logger.ActiveInitLogger()

	flag.StringVar(&flagConfigFilepath, "c", "", "file path to config file")
	flag.IntVar(&flagPort, "p", 80, "port to listen")
	flag.StringVar(&flagStaticFilePath, "s", "", "file path to static file")
	flag.Parse()

	if flagConfigFilepath == "" {
		panic("use -c to specify config file.")
	}

	if flagStaticFilePath == "" {
		panic("use -s to specify static file.")
	}

	err = config.ActiveInitConfigWithFile(flagConfigFilepath)
	if err != nil {
		fmt.Println("应用启动失败。\n错误原因：", err.Error())
		return
	}

	port := 80
	if flagPort != 0 {
		port = flagPort
	}

	go func() {
		err = router.ActiveInitHttpRouter(port, flagStaticFilePath)
		if err != nil {
			fmt.Println("应用启动失败。\n错误原因：", err.Error())
			return
		}
	}()

	go func() {
		time.Sleep(time.Second)
		if err == nil {
			fmt.Println("恭喜!您已成功启动酷应用服务。\n可返回开发者后台继续操作。")
		}
	}()

	select {}
}
