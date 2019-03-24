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
	Upper       bool
	Numbers     bool
	UpperChance uint
	Delims      []string
}

func Defaults() PwgenConfig {
	return PwgenConfig{
		Upper:       true,
		UpperChance: 1,
		Numbers:     true,
		Delims:      []string{"-", "_", "!", "@", "%", "-", ".", "#"},
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
	p := New(config)
	p.words = words

	return p
}

func (p *pwgen) Generate() string {
	word := p.getWord()
	delim := p.getDelim()
	random := p.getRandom()

	if p.config.Numbers {
		if p.r.Intn(9) < 5 {
			word = p.addNum(word)
		} else {
			random = p.addNum(random)
		}
	}

	return word + delim + random;
	//return p.getWord() + p.getDelim() + p.getRandom()
}

func (p *pwgen) addNum(s string) string {
	if p.r.Intn(9) < 5 {
		return fmt.Sprintf("%s%d", s, p.r.Intn(9));
	} else {
		return fmt.Sprintf("%d%s", p.r.Intn(9), s);
	}
}

func (p *pwgen) getDelim() string {
	if len(p.config.Delims) == 0 {
		return ""
	}
	return p.config.Delims[p.r.Intn(len(p.config.Delims))]
}

func (p *pwgen) getWord() string {
	if len(p.words) == 0 {
		return ""
	}

	slice := strings.Split(p.words[p.r.Intn(len(p.words))], "")
	return p.randomize(slice)
}

func (p *pwgen) randomize(str []string) string {
	if p.config.Upper {
		for i := range str {
			if uint(p.r.Intn(10)) < p.config.UpperChance {
				str[i] = strings.ToUpper(str[i])
			}
		}
	}

	result := strings.Join(str, "")

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
