package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		insertElement(nums2[i], nums1, m)
		m++
	}

}
func insertElement(x int, arr []int, m int) {
	existing := false // if in the array there is an element  greater than x
	for k := 0; k < m; k++ {
		if arr[k] > x {
			existing = true //had found
			for i := m - 1; i >= k; i-- {
				arr[i+1] = arr[i] //arr[i+1]: move the following  elements k
			}
			arr[k] = x // assign x to arr[k]
			break
		}

	}
	if !existing {
		arr[m] = x //if  there is no element  in the array  greater than x, insert x into the current position
	}
}

//cách 2: dùng 2 con trỏ
func merge1(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] > nums2[j] {
			insertElement(nums2[j], nums1, m)
			j++
			m++
		} else {
			i++
		}
	}
	for j < n {
		insertElement(nums2[j], nums1, m)
		j++
		m++
	}
}

//cách 3: dùng Complexity
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k-- //giai thich: boi vi k la vi tri cuoi cung cua mang nums1, nen sau khi gan gia tri cho nums1[k] thi ta phai giam k di 1 don vi
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--

	}
}

func main() {
	arr := []int{2, 4, 5, 0, 0}
	arr2 := []int{1, 2}
	merge(arr, 3, arr2, 2)
	fmt.Println("DONE")
}
