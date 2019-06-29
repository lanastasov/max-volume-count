package main

import (
	"fmt"
)

func main() {
	var val1 = 6.75
	var val2 = 6.78
	var val3 = 6.8
	var val4 = 6.98

	for i := 1; i <= 50; i++ {
		fmt.Println("<tr>")
		fmt.Println("    <td>", float64(i)/10, "</td>")
		fmt.Println("    <td>", val1, "--", float64(0.113)*float64(i), "--", val1*float64(i), "</td>")
		fmt.Println("    <td>", val2, "--", float64(0.113)*float64(i), "--", val2*float64(i), "</td>")
		fmt.Println("    <td>", val3, "--", float64(0.113)*float64(i), "--", val3*float64(i), "</td>")
		fmt.Println("    <td>", val4, "--", float64(0.113)*float64(i), "--", val4*float64(i), "</td>")
		fmt.Println("</tr>")
	}
}
