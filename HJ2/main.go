/**
 * HJ2 计算某字符出现的次数
 *
 * 描述：
 * 		写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字符，然后输出输入字符串中该字符的出现次数。（不区分大小写字母）
 *
 * 输入描述：
 * 		第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字符。
 *
 * 输出描述：
 * 		输出输入字符串中含有该字符的个数。（不区分大小写字母）
 *
 * 示例：
 * 		输入：ABCabc
 * 		     A
 * 		输出：2
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	chmap := make(map[byte]int)

	// Get string
	input.Scan()
	str := input.Text()

	// Get target character
	input.Scan()
	c := input.Text()

	// Caculate
	str = strings.ToLower(str)
	c = strings.ToLower(c)
	for i := 0; i < len(str); i++ {
		chmap[str[i]]++
	}
	fmt.Println(chmap[c[0]])
}
