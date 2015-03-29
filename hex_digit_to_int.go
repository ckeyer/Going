/*
* 文件名:  字符进制转换
* 创建日期:  2015.3.17
* 作者:  ckeyer
* 功能:  实现16进制字符转为10进制
**/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(hex_digit_to_int('2'))
}

func hex_digit_to_int(in byte) (out int) {
	fmt.Println('A' | 0x60)
	return (int)(in ^ 0x30)
}
