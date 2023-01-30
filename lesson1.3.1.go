/*Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:

text = Eyjafjallajokull, width = 6 → Eyjafj...
Если строка не превышает указанной длины, менять ее не следует:

text = hello, width = 6 → hello
Гарантируется, что в исходной строке text используются только однобайтовые символы без пробелов, а длина width строго больше 0.*/
package main

import (
	"fmt"
)

func main() {
	var text, res string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	runes := []rune(text)
	end := "..."

	if len(runes) > width {
		res = string(runes[0:width]) + end
	} else {
		res = string(runes)
	}

	fmt.Println(res)
}
