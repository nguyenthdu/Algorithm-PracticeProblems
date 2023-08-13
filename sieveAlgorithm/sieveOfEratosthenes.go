package main

import "fmt"

//TODO: Hàm sieveOfEratosthenes: Đây là hàm thực hiện thuật toán Eratosthenes.
//Trong hàm này, chúng ta tạo một mảng isPrime để đánh dấu các số nguyên tố.
//Ban đầu, tất cả các phần tử trong mảng đều được đánh dấu là true,
//ngụ ý rằng chúng ta giả định mọi số đều là số nguyên tố.

func sieveOfEratosthenes(limit int) []int {
	// Create an array to mark prime numbers
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	p := 2
	// Start the loop, iterate until the square of p exceeds the limit
	for p*p <= limit {
		// If p is prime
		if isPrime[p] {
			// Mark multiples of p (except p) as non-prime
			for i := p * p; i <= limit; i += p {
				isPrime[i] = false
			}
		}
		p++
	}
	// Create a list containing the last primes
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var n int
	fmt.Print("Enter the natural number n:")
	_, err := fmt.Scanf("%d", &n)
	if err != nil || n <= 0 {
		fmt.Println("Please enter a natural number > 0!!!")
	} else {
		primeNumbers := sieveOfEratosthenes(n - 1)
		fmt.Println("Prime nubmers < n:", primeNumbers)
	}

}

//TODO:Vòng Lặp Sàng Eratosthenes: Chúng ta sử dụng vòng lặp để thực hiện sàng Eratosthenes.
//Lặp qua tất cả các số từ 2 đến căn bậc hai của giới hạn (limit). Khi p là số nguyên tố,
//chúng ta đánh dấu tất cả các bội số của p (ngoại trừ p) trong mảng isPrime là false.
//TODO: Tạo Danh Sách Số Nguyên Tố: Sau khi hoàn thành vòng lặp sàng Eratosthenes,
//chúng ta tạo danh sách primes chứa các số nguyên tố cuối cùng bằng cách lặp qua mảng isPrime và thêm các phần tử có giá trị true vào danh sách.
