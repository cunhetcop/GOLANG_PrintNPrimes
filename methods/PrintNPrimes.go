package methods

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func FindNPrimes() {
	n, err := getValidNumber()
	if err != nil {
		fmt.Println(err)
		return
	}

	printNPrimes(n)
}

func getValidNumber() (n int, err error) {
	//dừng chương trình sau 5 giây nếu không có phản hồi
	color.Yellow("CHƯƠNG TRÌNH TÌM N SỐ NGUYÊN TỐ ĐẦU TIÊN")
	color.Green("Nhập một số nguyên dương: ")
	time.AfterFunc(5*time.Second, func() {
		color.Red("Chương trình đã tự động kết thúc sau 5 giây không có phản hồi")
		panic("timeout")
	})

	//bắt đầu nhập từ bàn phím
	var input string
	count := 0
	for count < 3 {
		//nhập khoảng trắng
		_, err = fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("Nhập sai %d lần \n", count+1)
			count++
			continue
		}
		//nhập sai
		n, err = strconv.Atoi(input)
		if err != nil || n <= 0 {
			fmt.Printf("Nhập sai %d lần \n", count+1)
			count++
			continue
		}

		return n, nil
	}
	color.Red("Bạn đã nhập sai quá 3 lần. Game over.")
	return
}

func printNPrimes(n int) {
	count := 0
	primeChan := make(chan int)
	go func() {
		for i := 2; count < n; i++ {
			if isPrime(i) {
				count++
				go func(prime int) {
					primeChan <- prime
				}(i)
			}
		}
	}()

	
	printChan := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			prime := <-primeChan
			printChan <- prime
		}
		close(primeChan)
		close(printChan)
	}()
	
	go func() {
		for i := 1; i <= n; i++ {
			prime := <-printChan
			fmt.Printf("Số nguyên tố thứ #%d: %d\n", i, prime)
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
}


func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
