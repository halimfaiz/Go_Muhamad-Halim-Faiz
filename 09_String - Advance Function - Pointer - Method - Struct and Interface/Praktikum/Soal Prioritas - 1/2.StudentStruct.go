package main

import "fmt"

type student struct {
	name  []string
	score []int
}

func (avg *student) avg() int {
	total := 0
	for _, v := range avg.score {
		total += v
	}
	mean := total / len(avg.score)
	return mean
}
func (a *student) min() (string, int) {
	min := a.score[0]
	minName := a.name[0]
	for i, v := range a.score {
		if min > v || min == 0 {
			min = v
			minName = a.name[i]
		}
	}
	return minName, min
}

func (b *student) max() (string, int) {
	max := b.score[0]
	maxName := b.name[0]
	for i, v := range b.score {
		if max < v {
			max = v
			maxName = b.name[i]
		}
	}
	return maxName, max
}

func main() {
	student := student{
		name:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		score: []int{80, 75, 70, 75, 60},
	}
	minN, minS := student.min()
	maxN, maxS := student.max()
	fmt.Println("Average Score : ", student.avg())
	fmt.Printf("Min Score of Students : %s (%d)\n", minN, minS)
	fmt.Printf("Max Score of Students : %s (%d)\n", maxN, maxS)
}
