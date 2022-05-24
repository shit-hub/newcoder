/**
 * HJ1 字符串最后一个单词的长度
 *
 * 描述:
 * 		计算字符串最后一个单词的长度，单词以空格隔开，字符串长度小于5000。（注：字符串末尾不以空格为结尾）
 *
 * 输入描述：
 * 		输入一行，代表要计算的字符串，非空，长度小于5000。
 *
 * 输出描述：
 * 		输出一个整数，表示输入字符串最后一个单词的长度。
 *
 * 示例：
 * 		输入：hello newcoder
 * 		输出：8
 * 		说明：最后一个单词为newcoder，长度为8.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	r := []rune(a)
	var n int
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] == '\n' {
			continue
		} else if r[i] != ' ' {
			n++
		} else {
			break
		}
	}
	fmt.Printf("%d\n", n)
}
