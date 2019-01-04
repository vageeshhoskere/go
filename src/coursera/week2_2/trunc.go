package main

import (
	"fmt"
	"strings"
)

func main() {
	var ip float64
	var response string
	for {
		fmt.Println("Enter floating point number")
		n, err := fmt.Scan(&ip)
		if err == nil && n == 1 {
			fmt.Println("Truncated value is ", int64(ip))
		} else {
			fmt.Println("Invalid input.")
		}
		fmt.Println("Do you want to try once again? (y/n)")
		_, err = fmt.Scan(&response)
		response = strings.TrimSpace(response)
		response = strings.ToLower(response)
		if err != nil || response != "y" {
			break
		}
	}
}
