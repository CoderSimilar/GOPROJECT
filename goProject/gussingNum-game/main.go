package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	// "os"
	// "strconv"
	// "strings"
	"time"
)

func main() {
	maxNum := 100	//猜数字的范围是0 ~ 100
	rand.Seed(time.Now().Unix())//设置随机数种子
	secretNumber := rand.Intn(maxNum)//生成随机数
	// fmt.Println(secretNumber)
	fmt.Println("Hello,Welcome to play Gussing Number Game!")
	fmt.Println("Please input your guess:")
	guess := 0
	
	guessNum := 1
	for
	{
		// input, err := reader.ReadString('\n')//读入标准输入的换行符
		// if err != nil {
		// 	fmt.Println("An error occured while reading input. Please try again!")
		// 	continue
		// }
		// input = strings.Trim(input, "\r\n")//去除标准输入的换行符
		// fmt.Scanf("%d", &reader)//从标准输入读取一个数字
		// guess, err := strconv.Atoi(reader)//将读入的字符串转换成整数
		// if err != nil {
		// 	fmt.Println("An error occured while reading input. Please try again!")
		// 	continue
		// }
		fmt.Scanf("%d", &guess)//从标准输入读取一个数字
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than secret number, Please try again.")
			guessNum += 1
		}else if guess < secretNumber {
			fmt.Println("Your guess is smaller than secret number, Please try again.")
			guessNum += 1
		}else {
			fmt.Println("Correct, you Legend! Use times is", guessNum)
			break
		}

	}
}