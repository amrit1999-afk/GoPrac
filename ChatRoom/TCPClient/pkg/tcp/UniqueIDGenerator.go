package tcp

import (
	"math/rand"
	"strconv"
)

func UniqueIDGenerator() string {
	uniqueID := ""

	for i := 1; i <= 10; i++ {
		choice := rand.Intn(3-1+1) + 1

		switch choice {
		case 1:
			{
				randVal := rand.Intn(90-65+1) + 65
				uniqueID += strconv.Itoa(randVal)
			}
		case 2:
			{
				randVal := rand.Intn(122-97+1) + 97
				uniqueID += strconv.Itoa(randVal)
			}
		default:
			{
				randVal := rand.Intn(57-48+1) + 48
				uniqueID += strconv.Itoa(randVal)
			}
		}
	}
	return uniqueID
}
