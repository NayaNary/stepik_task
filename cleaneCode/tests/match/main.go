package main

import (
	"fmt"
	matchcontains "match/matchContains"
)

func main() {
	s := "go is awesome"
	fmt.Println(matchcontains.MatchContains("is", s))
	// true
	fmt.Println(matchcontains.MatchContains("go.*awesome", s))
	// false

	fmt.Println(matchcontains.MatchRegexp("go.*awesome", s))
	// true
	fmt.Println(matchcontains.MatchRegexp("^go", s))
	// true
	fmt.Println(matchcontains.MatchRegexp("awesome$", s))
	// true
}
