package main

import (
	"github.com/cherry-game/cherry/logger"
	"github.com/cherry-game/cherry/timer"
	"time"
)

func main() {
	cherryLogger.DefaultSet()

	t := cherryTimer.NewTimer(func() {
		//time.Sleep(time.Second * 3)
		cherryLogger.Infof("execute func....")
	}, time.Second*1, 3)

	cherryTimer.AddTimer(t)

	go func() {
		for {
			cherryTimer.Cron()
		}
	}()

	time.Sleep(time.Second * 10000)
}
