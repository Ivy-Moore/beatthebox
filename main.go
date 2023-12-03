package main

import (
	"fmt"
	"math/rand"
)

func main() {
	wins := 0 
	losses := 0
	for i:= 0; i < 1000000; i++ {
		result := simulation()
		if (result == 0){
			losses++
		} else {
			wins++
		}

		// fmt.Println(result)
	}

	fmt.Println("WINS", wins)
	fmt.Println("LOSSES", losses)
}

func simulation() int {
	//11 - J 12 - Q 13 - K 14 - A
	cardset := [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	
	var card [52]int 
	index := 0
	indexofcard := 0
	for i:= 0; i < 4; i++ {
		for j:= 0; j < len(cardset); j++ {
			card[indexofcard] = cardset[index]
			index++
			indexofcard++
		}
		index = 0
	}

	rand.Shuffle(len(card), func(i, j int) {
		card[i], card[j] = card[j], card[i]
	})


	var line []int

	for i:= 0; i < len(card); i++ {
		if (i == 9){
			break
		} else {
			line = append(line, card[i])
			card[i] = -1
		} 
	}

	indexOfLine := 9

	for i:= 9; i < len(card); i++{
		if(indexOfLine != 0) {
			randIndex := rand.Intn(indexOfLine)
			num := line[randIndex]
			if(num == 2){
				if(num == card[i]){
					line = remove(line, randIndex)
					indexOfLine--
					card[i] = -1
				} else {
					line[randIndex] = card[i]
					card[i] = -1
				}
			}
			if(num == 14){
				if(num == card[i]){
					line = remove(line, randIndex)
					indexOfLine--
					card[i] = -1
				} else {
					line[randIndex] = card[i]
					card[i] = -1
				}
			}
			if(num == 8){
				randChoice := rand.Intn(2)
				if(num == card[i]){
					line = remove(line, randIndex)
					indexOfLine--
					card[i] = -1
				} else if(randChoice == 0 && num < card[i]){
					line[randIndex] = card[i]
					card[i] = -1
				}else if(randChoice == 1 && num > card[i]){
					line[randIndex] = card[i]
					card[i] = -1
				} else {
					line[randIndex] = card[i]
					card[i] = -1
				}
			}
			if(num == 3 || num == 4 || num == 5 || num == 6 || num == 7){
				if(num == card[i]){
					line = remove(line, randIndex)
					indexOfLine--
					card[i] = -1
				} else if (num < card[i]){
					line[randIndex] = card[i]
					card[i] = -1
				} else {
					line = remove(line, randIndex)
					indexOfLine--
					card[i] = -1
				}
			}
			if(num == 9 || num == 10 || num ==11 || num == 12 || num == 13){
				if(num == card[i]){
					line = remove(line, randIndex)
					indexOfLine--
					card[i] = -1
				} else if (num > card[i]){
					line[randIndex] = card[i]
					card[i] = -1
				} else {
					line = remove(line, randIndex)
					indexOfLine--
					card[i] = -1
				}
			}
		}else {
			break
		}

	}

	if(card[len(card)-1] == -1){
		return 1
	} else {
		return 0
	}

}

func remove(line []int, i int) []int {
    line[i] = line[len(line)-1]
    return line[:len(line)-1]
}