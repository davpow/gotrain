/**
Forever
If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
 */
package main

import "fmt"

func main() {
	for {
		fmt.Println("Infinite Loop...")
		break
	}

}
