// Package fnombre contains functions for generating random, human-friendly names.
package fnombre

import (
	"bufio"
	_ "embed"
	"strings"
)

//go:embed data/adjective
var _adjective string
var adjectives = parseGenderedWordList(_adjective)

//go:embed data/adverb
var _adverb string
var adverbs = parseNongenderedWordList(_adverb)

//go:embed data/noun
var _noun string
var nouns = parseGenderedWordList(_noun)

//go:embed data/verb
var _verb string
var verbs = parseGenderedWordList(_verb)

type Gender int
const (
	GenderNeutral   Gender = 0
	GenderFeminine  Gender = 1
	GenderMasculine Gender = 2
)

type Plurality int
const (
	PluralitySingular Plurality = 0
	PluralityPlural   Plurality = 1
)

type Word struct {
	Text   string
	Gender Gender
	Plurality Plurality
}

// Dictionary is a collection of words.
type Dictionary struct {
	adjectives []Word
	adverbs    []Word
	nouns      []Word
	verbs      []Word
}

// NewDictionary creates a new dictionary.
func NewDictionary() *Dictionary {
	// TODO: allow for custom dictionary
	return &Dictionary{
		adjectives: adjectives,
		adverbs:    adverbs,
		nouns:      nouns,
		verbs:      verbs,
	}
}

// LengthAdjective returns the number of adjectives in the dictionary.
func (d *Dictionary) LengthAdjective() int {
	return len(d.adjectives)
}

// LengthAdverb returns the number of adverbs in the dictionary.
func (d *Dictionary) LengthAdverb() int {
	return len(d.adverbs)
}

// LengthNoun returns the number of nouns in the dictionary.
func (d *Dictionary) LengthNoun() int {
	return len(d.nouns)
}

// LengthVerb returns the number of verbs in the dictionary.
func (d *Dictionary) LengthVerb() int {
	return len(d.verbs)
}

func parseGenderedWordList(data string) []Word {
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)
	var words []Word
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			word := parts[0]
			gender := toGender(parts[1])
			plurality := toPlurality(parts[2])
			words = append(words, Word{Text: word, Gender: gender, Plurality: plurality})
		}
	}
	return words
}

func parseNongenderedWordList(data string) []Word {
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)
	var words []Word
	for scanner.Scan() {
		line := scanner.Text()
		word := line // Since each line is now just the word itself
		// Use the fixed values for gender and plurality
		words = append(words, Word{Text: word, Gender: GenderNeutral, Plurality: PluralitySingular})
	}
	return words
}


// FIXME: Naive AF
func toGender(s string) Gender {
	switch s {
	case "1":
		return GenderFeminine
	case "2":
		return GenderMasculine
	default:
		return GenderNeutral
	}
}

func toPlurality(s string) Plurality {
	switch s {
	case "1":
		return PluralityPlural
	default:
		return PluralitySingular
	}
}

// Matches checks if the adjective matches the noun's gender and number.
func (w *Word) Matches(other Word) bool {
	// If the noun is gender-neutral (2), choose the feminine form by default.
	if other.Gender == GenderNeutral {
		other.Gender = GenderFeminine // Assume feminine for words like 'persona'
	}

	// Match if both have the same gender or if one is neutral (2).
	if (w.Gender == other.Gender || w.Gender == GenderNeutral || other.Gender == GenderNeutral) &&
		w.Plurality == other.Plurality {
		return true
	}
	return false
}
