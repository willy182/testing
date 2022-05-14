package main

import (
	"fmt"

	"github.com/willy182/testing/alami"
	"github.com/willy182/testing/garasidotid"
	"github.com/willy182/testing/okadoc"
	"github.com/willy182/testing/ruangguru"
	"github.com/willy182/testing/tiketdotcom"
	"github.com/willy182/testing/ula"
)

func exampleFunction(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}

	return sum
}

func main() {
	// ula test hackerearth 1
	str := "occurrences"
	countOccurrences := ula.GetNumberLetter(str)
	fmt.Println("countOccurrences", countOccurrences)
	fmt.Println("============================ ula test hackerearth 1")
	fmt.Println("")

	// ula test hackerearth 1
	str = "CamelCase"
	camelCase := ula.CamelToSnake(str)
	fmt.Println("camelCase", camelCase)
	fmt.Println("============================ ula test hackerearth 1")
	fmt.Println("")

	// tiket test codility 1
	x := []int{1, 2, 3, 1, 2}
	y := []int{2, 4, 6, 5, 10}
	fractions := tiketdotcom.Fractions(x, y)
	fmt.Println("fractions", fractions)
	fmt.Println("============================ tiket test codility 1")
	fmt.Println("")

	// tiket test codility 2
	inputan := []int{1, 3, 1, 2}
	pohon := tiketdotcom.Pohon(inputan)
	fmt.Println("pohon", pohon)
	fmt.Println("============================ tiket test codility 2")
	fmt.Println("")

	// ruang guru test live coding
	nilaiPangkat := ruangguru.Pangkat(2, 0)
	fmt.Println("nilaiPangkat", nilaiPangkat)
	fmt.Println("============================ ruang guru test live coding")
	fmt.Println("")

	// okadoc testDome findMaxSum
	list := []int{5, 9, 7, 11}
	fmt.Println(list, okadoc.FindMaxSum(list))
	fmt.Println("============================ okadoc testDome findMaxSum")
	fmt.Println("")

	// okadoc testDome fillChannel
	expensiveFunction := func() int { return exampleFunction(2000000000) }
	cheapFunction := func() int { return exampleFunction(100000000) }
	ch := okadoc.FillChannel(expensiveFunction, cheapFunction)

	if ch != nil {
		for i := 0; i < 2; i++ {
			fmt.Printf("Result: %d\n", <-ch)
		}
	}
	fmt.Println("============================ okadoc testDome fillChannel")
	fmt.Println("")

	// garasi.id hacker rank Is Prime
	var number int64 = 929
	if garasidotid.IsPrimeByNative(number) == 1 {
		fmt.Printf("%d is primes\n", number)
	} else {
		fmt.Printf("%d is not primes\n", number)
	}
	fmt.Println("============================ garasi hackerank IsPrime")
	fmt.Println("")

	// alami
	inputFactorial := 5
	factorial := alami.Factorial(inputFactorial)
	fmt.Printf("Factorial %d = %d \n", inputFactorial, factorial)
	fmt.Println("============================ alami Factorial")
	fmt.Println("")

	numericReverse := 5793
	numReverse := alami.NumericReverse(numericReverse)
	fmt.Printf("Factorial %d = %d \n", numericReverse, numReverse)
	fmt.Println("============================ alami NumericReverse")
	fmt.Println("")

	fmt.Println("============================ testing")
	f := float64(5.2)
	fmt.Println(f)
	fmt.Println(int(f))
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
