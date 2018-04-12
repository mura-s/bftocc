package main

type Token rune

const (
	Add       = Token('+')
	Sub       = Token('-')
	Right     = Token('>')
	Left      = Token('<')
	Read      = Token(',')
	Write     = Token('.')
	BeginLoop = Token('[')
	EndLoop   = Token(']')
)

type Tokenizer struct{}

func (t *Tokenizer) Tokenize(input string) []Token {
	tokens := make([]Token, 0)

	for _, c := range input {
		t := Token(c)
		switch t {
		case Add, Sub, Right, Left, Read, Write, BeginLoop, EndLoop:
			tokens = append(tokens, t)
		default:
			// do nothing
		}
	}

	return tokens
}
