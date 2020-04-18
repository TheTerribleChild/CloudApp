package timeutil

import(
	"time"
)

type TimeUtil struct{
	GetTime func() (time.Time, error)
}

func (instance *TimeUtil) GetTime() {currTime time.Time, err error}{
	err = nil
	if instance.GetTime == nil {
		currTime = time.Now()
		return
	}else{
		currTime, err = instance.GetTime()
	}
}