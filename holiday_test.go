package holiday

import (
	"fmt"
	"testing"
)

func TestIsHoliday(t *testing.T) {
	var (
		date = "20211009"
		expected = true
		err error
		out bool
	)
	out, err = IsHoliday(date)
	if err != nil {
		fmt.Errorf("err:%s", err.Error())
	}
	if out != expected  {
		t.Error("不符合预期")
	}
}
