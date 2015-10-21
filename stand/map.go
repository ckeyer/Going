// bug version: mapref.go
package main

import (
	"fmt"
)

func main() {
	foo := make(map[string]map[string]map[string]float32)
	foo_s12 := map[string]float32{"s2": 0.1}
	foo_s1 := map[string]map[string]float32{"s1": foo_s12}
	foo["x1"] = foo_s1
	foo_s22 := map[string]float32{"s2": 0.5}
	foo_s2 := map[string]map[string]float32{"s1": foo_s22}
	foo["x2"] = foo_s2

	x3 := make(map[string]map[string]float32)
	for _, v := range foo {
		for sk, sv := range v {
			if _, ok := x3[sk]; ok {
				for tmpk, tmpv := range sv {
					if _, ok := x3[sk][tmpk]; ok {
						x3[sk][tmpk] += tmpv
					} else {
						x3[sk][tmpk] = tmpv
					}
				}
			} else {
				x3[sk] = sv // 注意这里，map的赋值是个引用行为！
			}
		}
	}
	fmt.Printf("foo=%v\n", foo)
	fmt.Printf("x3=%v\n", x3)
}
