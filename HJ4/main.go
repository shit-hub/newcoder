/**
 * HJ4 字符串分隔
 *
 * 描述：
 * 		输入一个字符串，请按长度为8拆分每隔输入字符串并进行输出；
 * 		长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。
 *
 * 输入描述：
 * 		连续输入字符串(每个字符串长度小于等于100)
 *
 * 输出描述：
 * 		依次输出所有分割后的长度为8的新字符串
 *
 * 示例：
 * 		输入：abc
 * 		输出：abc00000
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		panic("error")
	}

	str := []rune(scanner.Text())
	m, n := len(str)/8, len(str)%8
	for i := 0; i < m; i++ {
		fmt.Println(string(str[i*8 : i*8+8]))
	}

	if n != 0 {
		fmt.Print(string(str[len(str)-n:]))
		for j := 0; j < 8-n; j++ {
			fmt.Print(0)
		}
		fmt.Println()
	}
}
