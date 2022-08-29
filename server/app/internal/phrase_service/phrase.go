package phraseservice

import (
	"math/rand"
	"time"
)

func GetStringInWarcraft3() string {
	listCitat := []string{
		"My life for Ner'zhul",
		"I wish only to serve.",
		"This is the hour of the Scourge!",
		"Death shall cleanse the world!",
		"Let blood drown the weak!",
	}
	rand.Seed(time.Now().Unix())

	return listCitat[rand.Intn(len(listCitat))]
}

func GetHelloWorld() string {
	return "Hello World!!!"
}
