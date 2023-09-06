package main

import (
	"fmt"
	"strings"
)

// kiểm tra số ký tự trong 1 chuỗi có thể tạo thành chuỗi kia hay không

func canConstruct(ransomNote string, magazine string) bool {

	for i := 0; i < len(ransomNote); i++ {
		if !strings.Contains(magazine, string(ransomNote[i])) {
			return false
		}
		magazine = strings.Replace(magazine, string(ransomNote[i]), "", 1)

	}
	return true

}

// using hash map
func canConstruct1(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("abc", "aabb"))
	fmt.Println(canConstruct1("abc", "aabb"))
}
