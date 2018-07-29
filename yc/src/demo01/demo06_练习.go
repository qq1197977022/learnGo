package demo01

import (
	"fmt"
	"time"
)

func Work01() {
	daySec := 60 * 60 * 24
	nowSecond := int(time.Now().Unix())
	/*根据nowSecond计算当前 年-月-日
	 */
	day := nowSecond / daySec
	remainder := nowSecond % (day * daySec)
	fmt.Println(day, remainder)
	fmt.Println(day / 12)
	fmt.Println(day / 365)

}
