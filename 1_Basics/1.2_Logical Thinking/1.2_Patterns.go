package main

import (
	"fmt"
)

func main() {
	var n, i, j, k, l int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)

	// *****
	// *****
	// *****
	// *****
	// *****
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// *
	// **
	// ***
	// ****
	// *****
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print("*")
			if i == j {
				break
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// 1
	// 12
	// 123
	// 1234
	// 12345
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print(j + 1)
			if i == j {
				break
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// 1
	// 22
	// 333
	// 4444
	// 55555
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print(i + 1)
			if i == j {
				break
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// *****
	// ****
	// ***
	// **
	// *
	for i = n; i > 0; i-- {
		for j = 1; j <= n; j++ {
			fmt.Print("*")
			if i == j {
				break
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// 12345
	// 1234
	// 123
	// 12
	// 1
	for i = n; i > 0; i-- {
		for j = 1; j <= n; j++ {
			fmt.Print(j)
			if i == j {
				break
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	//     *
	//    ***
	//   *****
	//  *******
	// *********
	for i = 0; i < n; i++ {
		for j = 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < (2*i)+1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// *********
	//  *******
	//   *****
	//    ***
	//     *
	for i = n - 1; i >= 0; i-- {
		for j = 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < (2*i)+1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	//     *
	//    ***
	//   *****
	//  *******
	// *********
	// *********
	//  *******
	//   *****
	//    ***
	//     *
	for i = 0; i < n; i++ {
		for j = 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < (2*i)+1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i = n - 1; i >= 0; i-- {
		for j = 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < (2*i)+1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// *
	// **
	// ***
	// ****
	// *****
	// ****
	// ***
	// **
	// *
	for i = 0; i <= n; i++ {
		for j = 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i = 1; i <= n; i++ {
		for j = 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// 1
	// 10
	// 101
	// 1010
	// 10101
	for i = 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(j % 2)
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// 1        1
	//  2      2
	//   3    3
	//    4  4
	//     55
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if i == j {
				fmt.Print(j)
			} else {
				fmt.Print(" ")
			}
		}
		for j = n; j > 0; j-- {
			if i == j {
				fmt.Print(j)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("\n------------")

	// 1        1
	// 12      21
	// 123    321
	// 1234  4321
	// 1234554321
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if j <= i {
				fmt.Print(j)
			} else {
				fmt.Print(" ")
			}
		}
		for j = n; j > 0; j-- {
			if j <= i {
				fmt.Print(j)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("\n------------")

	// 1
	// 2 3
	// 4 5 6
	// 7 8 9 10
	// 11 12 13 14 15
	count := 1
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print(count, " ")
			count++
			if i == j {
				break
			}
		}
		fmt.Println()
	}
	fmt.Println("\n------------")

	// A
	// AB
	// ABC
	// ABCD
	// ABCDE
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print(string(byte(j + 65)))
			if i == j {
				break
			}
		}
		fmt.Println()
	}
	fmt.Println("\n------------")

	// ABCDE
	// ABCD
	// ABC
	// AB
	// A
	for i = n - 1; i >= 0; i-- {
		for j = 0; j < n; j++ {
			fmt.Print(string(byte(j + 65)))
			if i == j {
				break
			}
		}
		fmt.Println()
	}
	fmt.Println("\n------------")

	// A
	// BB
	// CCC
	// DDDD
	// EEEEE
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Print(string(byte(i + 65)))
			if i == j {
				break
			}
		}
		fmt.Println()
	}
	fmt.Println("\n------------")

	// 	   A
	//    ABA
	//   ABCBA
	//  ABCDCBA
	// Tricky
	for i = 1; i <= n; i++ {
		for j = 1; j < n-i+1; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Print(string(byte(k + 64)))
		}
		for l = 1; l < i; l++ {
			fmt.Print(string(byte(i - l + 64)))
		}
		fmt.Println()
	}
	fmt.Println("\n------------")

	// E
	// DE
	// CDE
	// BCDE
	// ABCDE
	// Tricky
	for i = n; i > 0; i-- {
		for j = i; j <= n; j++ {
			fmt.Print(string(byte(j + 64)))
		}
		fmt.Println()
	}
	fmt.Println("\n------------")

	// **********
	// ****  ****
	// ***    ***
	// **      **
	// *        *
	// *        *
	// **      **
	// ***    ***
	// ****  ****
	// **********
	for i = 1; i <= n; i++ {
		for j = n; j > 0; j-- {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		for j = 1; j <= n; j++ {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if j <= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		for j = n - 1; j >= 0; j-- {
			if j <= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// *        *
	// **      **
	// ***    ***
	// ****  ****
	// **********
	// ****  ****
	// ***    ***
	// **      **
	// *        *
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if j <= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		for j = n; j > 0; j-- {
			if j <= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	for i = 1; i < n; i++ {
		for j = n - 1; j >= 0; j-- {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		for j = 0; j < n; j++ {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// *****
	// *   *
	// *   *
	// *   *
	// *****
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if i == 0 || j == 0 || j == n-1 || i == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n------------")

	// Inner Reducing Pattern
	// [4 4 4 4 4 4 4]
	// [4 3 3 3 3 3 4]
	// [4 3 2 2 2 3 4]
	// [4 3 2 1 2 3 4]
	// [4 3 2 2 2 3 4]
	// [4 3 3 3 3 3 4]
	// [4 4 4 4 4 4 4]
	size := (2 * n) - 1
	front := 0
	back := size - 1
	count = n
	// Create array with 0's inplace
	arr := make([][]int, size)
	for i = 0; i < size; i++ {
		arr[i] = make([]int, size)
	}
	// Populate values
	for count > 0 {
		for i = front; i <= back; i++ {
			for j = front; j <= back; j++ {
				arr[i][j] = count
			}
		}
		count--
		front++
		back--
	}
	// Print Array
	for i = 0; i < size; i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("\n------------")
}
