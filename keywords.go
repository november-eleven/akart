package main

import (
	"strings"
)

// Keyword define a possible auto-complete result
type Keyword string

// GetKey define the keyword's key used in the radix tree
func (k Keyword) GetKey() []byte {
	return []byte(strings.ToLower(k.GetValue()))
}

// GetValue define the keyword's value used in the radix tree
func (k Keyword) GetValue() string {
	return string(k)
}

// Keywords define a list of keywords
var Keywords = []Keyword{
	"Pandora",
	"Pinterest",
	"Paypal",
	"Pg&e",
	"Project free tv Priceline",
	"Press democrat",
	"Progressive",
	"Project runway",
	"Proactive",
	"Programming",
	"Progeria",
	"Progesterone",
	"Progenex",
	"Procurable",
	"Processor",
	"Proud",
	"Print",
	"Prank",
	"Bowl",
	"Owl",
	"River",
	"Phone",
	"Kayak",
	"Stamps",
	"Reprobe",
}
