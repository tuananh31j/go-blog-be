package common

import (
	"fmt"

	"nta-blog/internal/lib/logger"
)

func AppRecover() {
	if err := recover(); err != nil {
		logger.ZeroLog.Debug().Msg(fmt.Sprintf("Side effect is faild! >>>>>%v", err))
	}
}
