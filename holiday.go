package holiday

import (
	"strconv"
	"time"
)

func IsHoliday(date string) (bool, error) {
	var (
		t time.Time
		di int
		err error
	)
	if t, err = time.Parse("20060102", date); err != nil {
		return false, err
	}
	if di, err  = strconv.Atoi(date); err != nil {
		return false, err
	}
	if _, ok := OffWorkDayData[di]; ok {  // 判断是否是节假日
		return true, nil
	} else if _, ok := OnWorkDayData[di]; ok { // 判断是否是调休
		return false, nil
	} else if t.Weekday() == 0 || t.Weekday() == 6 {
		return true, nil
	} else {
		return false, nil
	}
}
