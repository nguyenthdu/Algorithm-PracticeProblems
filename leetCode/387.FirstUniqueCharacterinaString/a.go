package main

func firstUniqChar(s string) int {
	charUnique := make(map[rune]int)
	for _, c := range s {
		charUnique[c]++
	}
	for i, c := range s {
		if charUnique[c] == 1 {
			return i
		}
	}
	//do phuc tap O(n)
	return -1
}
func firstUniqChar2(s string) int {
	var charUnique [26]int//mang 26 phan tu vi tieng anh co 26 chu cai
	for _, c := range s {
		charUnique[c-'a']++//c-'a' de chuyen tu char sang int, vi du 'a'-'a'=0, 'b'-'a'=1
	}
	for i, c := range s {
		if charUnique[c-'a'] == 1 {//neu chi xuat hien 1 lan thi tra ve vi tri
			return i
		}
	}
	//do phuc tap O(n)
	return -1
}