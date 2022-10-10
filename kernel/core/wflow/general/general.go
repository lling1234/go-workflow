package general

var (
	UserId0   uint64 = 101
	UserName0        = "admin"
	UserId1   uint64 = 61769798
	UserName1        = "李四"
	UserId2   uint64 = 381496
	UserName2        = "旅人"
	UserId3   uint64 = 11025
	UserName3        = "xiaoming"
	UserId4   uint64 = 568898
	UserName4        = "王翠花"
	UserId5   uint64 = 489564
	UserName5        = "李秋香"
	UserId6   uint64 = 6418616
	UserName6        = "张三"

	TargetId uint64 = 1727882

	MyUserId = UserId4
	// MyUserId = UserId2
)

// 通过两重循环过滤重复元素
func RemoveRepByLoop(source []int64) []int64 {
	var result []int64 // 存放结果
	for i := range source {
		flag := true
		for j := range result {
			if source[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, source[i])
		}
	}
	return result
}
