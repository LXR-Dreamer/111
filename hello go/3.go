package main

import (
	"fmt"
	"time"
)

var x, y, z, i, j int
var a, b, c int
var t1 int

func main() {

	fmt.Println("What year were you born ?")

	fmt.Scanf("%d\n", &x)
Start1:
	fmt.Println("What month were you born [MM]?")
	fmt.Scanf("%d\n", &y)
	if y > 12 {
		fmt.Println("请输入正确的月份")
		goto Start1
	}
Start2:
	fmt.Println("What day were you born [DD]?")
	fmt.Scanf("%d\n", &z)
	if y == 1 || y == 3 || y == 5 || y == 7 || y == 8 || y == 10 || y == 12 {
		if z > 31 {
			fmt.Println("请输入正确的日期")
			goto Start2
		}
	}
	if y == 4 || y == 6 || y == 9 || y == 11 {
		if z > 30 {
			fmt.Println("请输入正确的日期")
			goto Start2
		}
	}
	if y == 2 {
		if a%4 == 0 && a%100 != 0 || a%400 == 0 {
			if z > 29 {
				fmt.Println("请输入正确的日期")
				goto Start2
			}
		} else {
			if z > 28 {
				fmt.Println("请输入正确的日期")
				goto Start2
			}
		}
	}
	fmt.Println("It looks like you were born on", z, "/", y, "/", x)
	a := int(time.Now().Year())
	b := int(time.Now().Month())
	c := int(time.Now().Day())
	var sum int
	if y < b {
		sum := c - z
		for j = y; j < b; j++ {
			if j == 1 {
				sum += 31
			} else if j == 2 {
				if a%4 == 0 && a%100 != 0 || a%400 == 0 {
					sum += 29
				} else {
					sum += 28
				}
			} else if j == 3 {
				sum += 31
			} else if j == 4 {
				sum += 30
			} else if j == 5 {
				sum += 31
			} else if j == 6 {
				sum += 30
			} else if j == 7 {
				sum += 31
			} else if j == 8 {
				sum += 31
			} else if j == 9 {
				sum += 30
			} else if j == 10 {
				sum += 31
			} else if j == 11 {
				sum += 30
			} else if j == 12 {
				sum += 31
			}
		}
		fmt.Println("It looks like your birthday has passed ", sum, "days")
		fmt.Println("Hope you're looking forward to the next one!")
	} else if y == b && z > c {
		sum = z - c
		fmt.Println("It looks like your birthday has passed ", sum, "days")
		fmt.Println("Hope you're looking forward to the next one!")
	} else if y == b && z < c {
		sum = c - z
		fmt.Println("It looks like there are ", sum, " days from your birthday.")
		fmt.Println("Hope you're looking forward to it!")
	} else if y > b {
		sum = z - c
		for j = b; j < y; j++ {
			if j == 1 {
				sum += 31
			} else if j == 2 {
				sum += 29
			} else if j == 3 {
				sum += 31
			} else if j == 4 {
				sum += 30
			} else if j == 5 {
				sum += 31
			} else if j == 6 {
				sum += 30
			} else if j == 7 {
				sum += 31
			} else if j == 8 {
				sum += 31
			} else if j == 9 {
				sum += 30
			} else if j == 10 {
				sum += 31
			} else if j == 11 {
				sum += 30
			} else if j == 12 {
				sum += 31
			}
			fmt.Println("It looks like there are ", sum, " days from your birthday.")
			fmt.Println("Hope you're looking forward to it!")
		}
	} else if y == b && z == c {
		fmt.Println("Aha,today is your birthday!")
		fmt.Println("Happy birthday to you!")
	}

}
