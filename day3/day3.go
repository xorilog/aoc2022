package day3

import (
	"fmt"
	"strings"
	"unicode"
)

var (
	AlphabetPriority = make(map[string]int, 52)
	AlphabetIsSet    = false
)

type Compartment struct {
	content []string
}

type Rucksack struct {
	one Compartment
	two Compartment
}

type AllRucksack struct {
	Data []*Rucksack
}

// NewRucksack takes a string and fill it in the two compartments of a Rucksack returned as a pointer.
func NewRucksack(s string) *Rucksack {
	split := len(s) / 2

	return &Rucksack{
		one: Compartment{content: []string{s[:split]}},
		two: Compartment{content: []string{s[split:]}},
	}
}

// FillRucksacks takes a slice of string as input and return a slice of Rucksack pointers
func FillRucksacks(input []string) []*Rucksack {
	rucksacks := make([]*Rucksack, len(input))
	for k, v := range input {
		rucksacks[k] = NewRucksack(v)
	}
	return rucksacks
}

// Content applies on a Rucksack object and return the assembled compartments as a single string
func (r *Rucksack) Content() string {
	return fmt.Sprintf("%s%s",
		r.One(),
		r.Two(),
	)
}

// One applies on a Rucksack object and return the content of the first compartment as a string
func (r *Rucksack) One() string {
	return strings.Join(r.one.content, "")
}

// Two applies on a Rucksack object and return the content of the second compartment as a string
func (r *Rucksack) Two() string {
	return strings.Join(r.two.content, "")
}

// Intersection applies on a Rucksack object and return a slice of string which contain unique intersection of
// `items` contained in each compartment
func (r *Rucksack) Intersection() []string {
	var (
		cleaned []string
		keys    = make(map[string]bool)
	)
	for _, e := range HashGeneric([]rune(r.One()), []rune(r.Two())) {
		es := string(e)
		if _, v := keys[es]; !v {
			keys[es] = true
			cleaned = append(cleaned, es)
		}
	}

	return cleaned
}

// AllRuckSacks takes a slice of string as input and return an AllRucksack pointer.
func AllRuckSacks(input []string) *AllRucksack {
	return &AllRucksack{Data: FillRucksacks(input)}
}

// Group appllies on a AllRucksack object, takes a group number and a windowSize int and return a slice of Rucksack
// pointers corresponding to the group with the defined window given.
func (ar *AllRucksack) Group(groupNumber, windowSize int) []*Rucksack {
	var (
		windowStart int
	)
	windowStart = groupNumber * windowSize

	return ar.Data[windowStart : windowStart+windowSize]
}

// GroupCount appllies on a AllRucksack object, takes a windowSize int and return the amount of groups we have for that
// windowSize as int.
func (ar *AllRucksack) GroupCount(windowSize int) int {
	// fmt.Println(len(ar.Data) % windowSize)
	count := (len(ar.Data) - 1) / windowSize
	return count
}

// Intersection applies on a AllRucksack object, takes a groupNumber as int and return a string which contain the unique
// common characters to all members of the group `items`.
func (ar *AllRucksack) Intersection(groupNumber int) string {
	var (
		content, cleaned []string
		toCompare        string
	)

	for _, r := range ar.Group(groupNumber, 3) {
		content = append(content, r.Content())
	}

	for k, _ := range content {
		if k < 1 {
			toCompare = content[k]
			continue
		}
		var keys = make(map[string]bool)
		for _, e := range HashGeneric([]rune(toCompare), []rune(content[k])) {
			es := string(e)
			if _, v := keys[es]; !v {
				keys[es] = true
				cleaned = append(cleaned, es)
			}
		}
		toCompare = strings.Join(cleaned, "")
		cleaned = cleaned[0:0]
	}

	return toCompare
}

// SumAllGroupsIntersectionPriorities applies on an AllRucksack object, it returns the item priorities for groups summed
// as int.
func (ar *AllRucksack) SumAllGroupsIntersectionPriorities() int {
	var intSlice []int
	for i := 0; i <= ar.GroupCount(3); i++ {
		intSlice = append(intSlice, i)
	}
	return ar.SumGroupIntersectionPriorities(intSlice)
}

// SumGroupIntersectionPriorities applies on an AllRucksack object, it takes a slice of groupNumbers as int and return
// the item priorities summed for concerned groups as int.
func (ar *AllRucksack) SumGroupIntersectionPriorities(g []int) int {
	var summedPriorities int
	for _, v := range g {
		summedPriorities = summedPriorities + Priority(ar.Intersection(v))
	}

	return summedPriorities
}

// SumIntersectionPriorities applies on an AllRucksack object and return an int of the item priority summed of all
// Rucksack objects.
func (ar *AllRucksack) SumIntersectionPriorities() int {
	var summedPriorities int
	for _, v := range ar.Data {
		for _, vv := range v.Intersection() {
			summedPriorities = summedPriorities + Priority(vv)
		}
	}

	return summedPriorities
}

// HashGeneric takes two slices of comparable types T and return the intersection of those two comparable slices as a
// slice of T. It uses go 1.18 generics.
func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

// prepareAlphabetPriority generate globally available variable containing the priority mapping of each Lower & Upper
// cased Alphabet character if not set and return a boolean true when done in order to avoid generation at each calls.
func prepareAlphabetPriority() bool {
	if !AlphabetIsSet {
		priority := 0
		for r := 'a'; r <= 'z'; r++ {
			priority++
			AlphabetPriority[string(r)] = priority
		}
		for r := unicode.ToUpper('a'); r <= unicode.ToUpper('z'); r++ {
			priority++
			AlphabetPriority[string(r)] = priority
		}
	}
	return true
}

// Priority takes a string and returns it's associated priority as an int value based on AlphabetPriority mapping.
func Priority(s string) int {
	AlphabetIsSet = prepareAlphabetPriority()
	return AlphabetPriority[s]
}
