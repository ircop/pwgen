package pwgen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type pwgen struct {
	words      []string
	config     PwgenConfig
	consonants []string
	vowels     []string
	r          *rand.Rand
}

type PwgenConfig struct {
	upper       bool
	numbers     bool
	numChance   uint
	upperChance uint
	delims      []string
}

func Defaults() PwgenConfig {
	return PwgenConfig{
		upper:       true,
		upperChance: 1,
		numbers:     true,
		numChance:   5,
		delims:      []string{"-", "_", "!", "@", "%", "-", ".", "#"},
	}
}

func New(config PwgenConfig) *pwgen {
	p := pwgen{}
	p.words = getDict()
	p.config = config

	p.consonants = []string{"b","c","d","g","k","m","n","p","r","s","t"}
	p.vowels = []string{"a","e","i","o","u"}
	p.r = rand.New(rand.NewSource(time.Now().Unix()))

	return &p
}

func NewWithDict(config PwgenConfig, words []string) *pwgen {
	p := New(Defaults())
	p.words = words

	return p
}

func (p *pwgen) Generate() string {
	return p.getWord() + p.getDelim() + p.getRandom()
}

func (p *pwgen) getDelim() string {
	if len(p.config.delims) == 0 {
		return ""
	}
	return p.config.delims[p.r.Intn(len(p.config.delims))]
}

func (p *pwgen) getWord() string {
	if len(p.words) == 0 {
		return ""
	}

	slice := strings.Split(p.words[p.r.Intn(len(p.words))], "")
	return p.randomize(slice)
}

func (p *pwgen) randomize(str []string) string {
	if p.config.upper {
		for i := range str {
			if uint(p.r.Intn(10)) < p.config.upperChance {
				str[i] = strings.ToUpper(str[i])
			}
		}
	}

	result := strings.Join(str, "")

	if p.config.numbers {
		if uint(p.r.Intn(10)) < p.config.numChance {
			result += fmt.Sprintf("%d", p.r.Intn(9))
		}
	}

	return result
}

func (p *pwgen) getRandom() string {
	// c+v+c+c+v
	res := []string {
		p.consonants[p.r.Intn(len(p.consonants))],
		p.vowels[p.r.Intn(len(p.vowels))],
		p.consonants[p.r.Intn(len(p.consonants))],
		p.consonants[p.r.Intn(len(p.consonants))],
		p.vowels[p.r.Intn(len(p.vowels))],
	}

	return p.randomize(res)
}
