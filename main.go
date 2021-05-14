package main

import (
	"fmt"

	"github.com/willy182/testing/ruangguru"
	"github.com/willy182/testing/tiketdotcom"
	"github.com/willy182/testing/ula"
)

func main() {
	// ula test hackerearth 1
	str := "occurrences"
	countOccurrences := ula.Ula1(str)
	fmt.Println("countOccurrences", countOccurrences)
	fmt.Println("============================ ula test hackerearth 1")

	// ula test hackerearth 1
	str = "CamelCase"
	camelCase := ula.Ula2(str)
	fmt.Println("camelCase", camelCase)
	fmt.Println("============================ ula test hackerearth 1")

	// tiket test codility 1
	x := []int{1, 2, 3, 1, 2}
	y := []int{2, 4, 6, 5, 10}
	fractions := tiketdotcom.TiketTest1(x, y)
	fmt.Println("fractions", fractions)
	fmt.Println("============================ tiket test codility 1")

	// tiket test codility 2
	inputan := []int{1, 3, 1, 2}
	pohon := tiketdotcom.Solution(inputan)
	fmt.Println("pohon", pohon)
	fmt.Println("============================ tiket test codility 2")

	// ruang guru test live coding
	nilaiPangkat := ruangguru.Pangkat(2, 0)
	fmt.Println("nilaiPangkat", nilaiPangkat)
	fmt.Println("============================ ruang guru test live coding")

	// testing time
	// now := time.Now()
	// now1 := now.AddDate(0, 0, 1)
	// fmt.Println(now)
	// fmt.Println(now1)
	// fmt.Println(now.Before(now1))
	// fmt.Println(now.After(now1))
	// fmt.Println("======")
	// fmt.Println(now.Format(time.RFC3339))
}
