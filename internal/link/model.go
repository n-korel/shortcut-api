package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
		Hash: RandStringRunes(7),
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(7)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}