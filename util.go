package leetcode

import (
	"errors"
	"fmt"
)

// intSlice 切片类型结果校验
func intSlice(in, res []int) error {
	// "无解" 校验
	if len(in) == 0 && len(res) != 0 {
		return errors.New("原题有解")
	}
	if len(in) != 0 && len(res) == 0 {
		return errors.New("原题无解")
	}

	/* 不做长度校验 因为要提取不存在的解 */

	// 结果校验字典
	var mp = make(map[int]struct{})

	// 结果存入字典
	for _, j := range res {
		mp[j] = struct{}{}
	}

	// 取出结果 取不到就报错
	for _, i := range in {
		if _, ok := mp[i]; !ok {
			return fmt.Errorf("错误解 %d", i)
		}
		delete(mp, i)
	}

	// mp 中剩余的解就是未答出的解
	for k := range mp {
		return fmt.Errorf("未答出解 %d", k)
	}
	return nil
}
