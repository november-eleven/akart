package main

import (
	"strings"

	art "github.com/november-eleven/akart/art"
)

// MaxSuggestions define the maximum suggestions requested
const MaxSuggestions = 4

// Engine will offer some suggestions based on the "auto-complete"
type Engine struct {
	t art.Tree
}

// NewEngine returns a Engine
func NewEngine() Engine {

	e := Engine{art.New()}

	for _, keyword := range Keywords {
		e.t.Insert(keyword.GetKey(), keyword.GetValue())
	}

	return e
}

// Find will lookup in the radix tree in order to provide some suggestions
func (e Engine) Find(search string) []string {

	i := 0
	l := []string{}

	e.t.ForEachPrefix(art.Key(strings.ToLower(search)), func(node art.Node) bool {

		if node.Kind() != art.Leaf {
			return true
		}

		switch v := node.Value().(type) {
		case string:
			l = append(l, v)
			i++
		default:
		}

		return i < MaxSuggestions

	})

	return l

}
