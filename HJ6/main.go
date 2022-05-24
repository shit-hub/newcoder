/**
 * HJ6 质数因子
 *
 * 描述：
 * 		功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（重复的也要列举）（如180的质因子为2 2 3 3 5 ）
 *
 * 输入描述：
 * 		输入一个整数。
 *
 * 输出描述：
 * 		按照从小到大的顺序输出它的所有质数的因子，以空格隔开。
 *
 * 示例：
 * 		输入：180
 * 		输出：2 2 3 3 5
 */

package main

import "fmt"

func main() {
	var input int64
	fmt.Scan(&input)

	for i := int64(2); i*i <= input; i++ {
		for input%i == 0 {
			fmt.Printf("%d ", i)
			input /= i
		}
	}
	if input >= 2 {
		fmt.Println(input)
	}
}
