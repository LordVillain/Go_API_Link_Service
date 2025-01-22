package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Url: url,
		Hash: RandStringRunes(6),
	}
}

var letterRunes = []rune("abcdefghijkmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")

func RandStringRunes(n int) string {
		str := make([]rune, n)
		for i := range str {
			str[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(str)
}