package timeutil

import(
	"time"
)

type TimeUtil struct{
	GetTimeFunc func() (time.Time, error)
}

func (instance *TimeUtil) GetTime() (currTime time.Time) {
	if instance.GetTimeFunc == nil {
		currTime = time.Now()
		return
	}else{
		currTime, _ = instance.GetTimeFunc()
	}
	return
}

func (instance *TimeUtil) GetTimeUnix() int64 {
	return instance.GetTime().Unix()
}