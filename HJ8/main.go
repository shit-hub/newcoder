/**
 * HJ8 合并表记录
 *
 * 描述：
 * 		数据表记录包含表索引index和数值value（int范围的正整数），请对表索引相同的记录进行合并，
 * 		即将相同索引的数值进行求和运算，输出按照index值升序进行输出。
 *
 * 输入描述：
 * 		先输入键值对的个数n（1 <= n <= 500）
 * 		接下来n行每行输入成对的index和value值，以空格隔开
 *
 * 输出描述：
 * 		输出合并后的键值对（多行）
 *
 * 示例：
 * 		输入：4
 * 			 0 1
 * 			 0 2
 * 		     1 2
 * 			 3 4
 * 		输出：0 3
 * 			 1 2
 * 			 3 4
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var key []int

	map1 := make(map[int]int)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var key int
		var value int
		fmt.Scan(&key, &value)
		map1[key] += value
	}

	for k := range map1 {
		key = append(key, k)
	}
	sort.Ints(key)

	for _, v := range key {
		fmt.Println(v, map1[v])
	}

}
