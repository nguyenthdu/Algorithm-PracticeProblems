package main

import (
	"bytes"
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := ""
	s2 := ""
	for _, val := range word1 {
		s1 += val
	}
	for _, val := range word2 {
		s2 += val
	}
	return s1 == s2
}
func arrayStringsAreEqual1(word1 []string, word2 []string) bool {
	var s1, s2 strings.Builder
	for _, val := range word1 {
		s1.WriteString(val)
	}
	for _, val := range word2 {
		s2.WriteString(val)
	}
	return s1.String() == s2.String()
}
func arrayStringsAreEqual2(word1 []string, word2 []string) bool {
	var s1, s2 bytes.Buffer
	for _, val := range word1 {
		s1.Write([]byte(val))
	}
	for _, val := range word2 {
		s2.Write([]byte(val))
	}
	return s1.String() == s2.String()

}


