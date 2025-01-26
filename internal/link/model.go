package link

import (
	"go/adv-demo/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(6)
}

var letterRunes = []rune("abcdefghijkmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")

func RandStringRunes(n int) string {
		str := make([]rune, n)
		for i := range str {
			str[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(str)
}

