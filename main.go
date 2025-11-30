/*
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-27
 * @fileoverview User enters 5 rows and receives correct number triangle.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many rows would you like? ")
	rowStr, _ := reader.ReadString('\n')
	rowStr = strings.TrimSpace(rowStr)
	rows, _ := strconv.Atoi(rowStr)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
	fmt.Println("\nDone.")
}
