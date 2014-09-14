package ninja

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	TOKEN_ADD                 = "add"
	TOKEN_ASSIGN              = "assign"
	TOKEN_COLON               = "colon"
	TOKEN_COMMA               = "comma"
	TOKEN_DIV                 = "div"
	TOKEN_DOT                 = "dot"
	TOKEN_EQ                  = "eq"
	TOKEN_FLOORDIV            = "floordiv"
	TOKEN_GT                  = "gt"
	TOKEN_GTEQ                = "gteq"
	TOKEN_LBRACE              = "lbrace"
	TOKEN_LBRACKET            = "lbracket"
	TOKEN_LPAREN              = "lparen"
	TOKEN_LT                  = "lt"
	TOKEN_LTEQ                = "lteq"
	TOKEN_MOD                 = "mod"
	TOKEN_MUL                 = "mul"
	TOKEN_NE                  = "ne"
	TOKEN_PIPE                = "pipe"
	TOKEN_POW                 = "pow"
	TOKEN_RBRACE              = "rbrace"
	TOKEN_RBRACKET            = "rbracket"
	TOKEN_RPAREN              = "rparen"
	TOKEN_SEMICOLON           = "semicolon"
	TOKEN_SUB                 = "sub"
	TOKEN_TILDE               = "tilde"
	TOKEN_WHITESPACE          = "whitespace"
	TOKEN_FLOAT               = "float"
	TOKEN_INTEGER             = "integer"
	TOKEN_NAME                = "name"
	TOKEN_STRING              = "string"
	TOKEN_OPERATOR            = "operator"
	TOKEN_BLOCK_BEGIN         = "block_begin"
	TOKEN_BLOCK_END           = "block_end"
	TOKEN_VARIABLE_BEGIN      = "variable_begin"
	TOKEN_VARIABLE_END        = "variable_end"
	TOKEN_RAW_BEGIN           = "raw_begin"
	TOKEN_RAW_END             = "raw_end"
	TOKEN_COMMENT_BEGIN       = "comment_begin"
	TOKEN_COMMENT_END         = "comment_end"
	TOKEN_COMMENT             = "comment"
	TOKEN_LINESTATEMENT_BEGIN = "linestatement_begin"
	TOKEN_LINESTATEMENT_END   = "linestatement_end"
	TOKEN_LINECOMMENT_BEGIN   = "linecomment_begin"
	TOKEN_LINECOMMENT_END     = "linecomment_end"
	TOKEN_LINECOMMENT         = "linecomment"
	TOKEN_DATA                = "data"
	TOKEN_INITIAL             = "initial"
	TOKEN_EOF                 = "eof"
)

var operatorsMap = map[string]string{
	"+":  TOKEN_ADD,
	"-":  TOKEN_SUB,
	"/":  TOKEN_DIV,
	"//": TOKEN_FLOORDIV,
	"*":  TOKEN_MUL,
	"%":  TOKEN_MOD,
	"**": TOKEN_POW,
	"~":  TOKEN_TILDE,
	"[":  TOKEN_LBRACKET,
	"]":  TOKEN_RBRACKET,
	"(":  TOKEN_LPAREN,
	")":  TOKEN_RPAREN,
	"{":  TOKEN_LBRACE,
	"}":  TOKEN_RBRACE,
	"==": TOKEN_EQ,
	"!=": TOKEN_NE,
	">":  TOKEN_GT,
	">=": TOKEN_GTEQ,
	"<":  TOKEN_LT,
	"<=": TOKEN_LTEQ,
	"=":  TOKEN_ASSIGN,
	".":  TOKEN_DOT,
	":":  TOKEN_COLON,
	"|":  TOKEN_PIPE,
	",":  TOKEN_COMMA,
	";":  TOKEN_SEMICOLON,
}

var operatorsArray = []string{`\/\/`, `\*\*`, `\=\=`, `\!\=`, `\>\=`, `\<\=`, `\+`, `\-`, `\/`, `\*`, `\%`, `\~`, `\[`, `\]`, `\(`, `\)`, `\{`, `\}`, `\>`, `\<`, `\=`, `\.`, `\:`, `\|`, `\,`, `\;`}

var (
	whitespaceRe = regexp.MustCompile(`^\s+`)
	floatRe      = regexp.MustCompile(`^\d+\.\d+`)
	integerRe    = regexp.MustCompile(`^\d+`)
	nameRe       = regexp.MustCompile(`^\b[a-zA-Z_][a-zA-Z0-9_]*\b`)
	stringRe     = regexp.MustCompile(`(?s)^('([^'\\]*(?:\\.[^'\\]*)*)'|"([^"\\]*(?:\\.[^"\\]*)*)")`)
	operatorRe   = regexp.MustCompile(fmt.Sprintf("^(%s)", strings.Join(operatorsArray, "|")))
	newlineRe    = regexp.MustCompile(`(\r\n|\r|\n)`)
)

var ignoredTokens = map[string]bool{
	TOKEN_COMMENT_BEGIN:     true,
	TOKEN_COMMENT:           true,
	TOKEN_COMMENT_END:       true,
	TOKEN_WHITESPACE:        true,
	TOKEN_LINECOMMENT_BEGIN: true,
	TOKEN_LINECOMMENT_END:   true,
	TOKEN_LINECOMMENT:       true,
}

var ignoreIfEmpty = map[string]bool{
	TOKEN_WHITESPACE:  true,
	TOKEN_DATA:        true,
	TOKEN_COMMENT:     true,
	TOKEN_LINECOMMENT: true,
}

func compile(x string) *regexp.Regexp {
	x = `(?ms)^` + x
	r := regexp.MustCompile(x)
	return r
}

type Token struct {
	lineno int
	tp     string
	value  interface{}
}

type StateToken struct {
	regex    *regexp.Regexp
	tokens   []string
	newState string
}

type TokenStream struct {
	iter    chan *Token
	current *Token
}

func (stream *TokenStream) next() {
    rv = stream.current
    if stream.current.tp != TOKEN_EOF {
        stream.current, ok <- stream.iter
        if (!ok) {
            stream.close()
        }
    }
    return rv
}

type Lexer struct {
	rules map[string][]*StateToken
}

func NewLexer() *Lexer {
	lexer := new(Lexer)

	tagRules := []*StateToken{
		&StateToken{whitespaceRe, []string{TOKEN_WHITESPACE}, "nil"},
		&StateToken{floatRe, []string{TOKEN_FLOAT}, "nil"},
		&StateToken{integerRe, []string{TOKEN_INTEGER}, "nil"},
		&StateToken{nameRe, []string{TOKEN_NAME}, "nil"},
		&StateToken{stringRe, []string{TOKEN_STRING}, "nil"},
		&StateToken{operatorRe, []string{TOKEN_OPERATOR}, "nil"},
	}

	//lstripRe := `^[ \t]*`
	//noLstripRe := `+`
	//blockPrefixRe := fmt.Sprintf(`%s{%|{%\+?`, lstripRe)

	lexer.rules = make(map[string][]*StateToken)

	rootTagRules := map[string]string{
		"comment":  "{#",
		"block":    "{%",
		"variable": "{{",
	}
	regexArray := []string{}
	regexArray = append(regexArray, `(?P<raw_begin>(?:{%)\s*raw\s*(?:%}))`)
	for n, r := range rootTagRules {
		regex := fmt.Sprintf(`(?P<%s_begin>%s)`, n, r)
		regexArray = append(regexArray, regex)
	}
	lexer.rules["root"] = []*StateToken{
		&StateToken{
			compile(fmt.Sprintf(`(.*?)(?:%s)`, strings.Join(regexArray, `|`))),
			[]string{TOKEN_DATA, "#bygroup"},
			"#bygroup",
		},
		&StateToken{
			compile(".+"),
			[]string{TOKEN_DATA},
			"nil",
		},
	}

	lexer.rules[TOKEN_COMMENT_BEGIN] = []*StateToken{
		&StateToken{
			compile(`(.*?)((?:#}))`),
			[]string{TOKEN_COMMENT, TOKEN_COMMENT_END},
			"#pop",
		},
		&StateToken{
			compile("(.)"),
			[]string{"Failure: Missing end of comment tag"},
			"nil",
		},
	}

	lexer.rules[TOKEN_BLOCK_BEGIN] = append([]*StateToken{
		&StateToken{
			compile(`(?:%})`),
			[]string{TOKEN_BLOCK_END},
			"#pop",
		},
	}, tagRules...)

	lexer.rules[TOKEN_VARIABLE_BEGIN] = append([]*StateToken{
		&StateToken{
			compile(`}}`),
			[]string{TOKEN_VARIABLE_END},
			"#pop",
		},
	}, tagRules...)

	lexer.rules[TOKEN_RAW_BEGIN] = []*StateToken{
		&StateToken{
			compile(`(.*?)((?:{%)\s*endraw\s*(?:%}))`),
			[]string{TOKEN_DATA, TOKEN_RAW_END},
			"#pop",
		},
		&StateToken{
			compile("(.)"),
			[]string{"Failure: Missing end of raw directive"},
			"nil",
		},
	}

	lexer.rules[TOKEN_LINESTATEMENT_BEGIN] = append([]*StateToken{
		&StateToken{
			compile(`\s*(\n|$)`),
			[]string{TOKEN_LINESTATEMENT_END},
			"#pop",
		},
	}, tagRules...)

	lexer.rules[TOKEN_LINECOMMENT_BEGIN] = []*StateToken{
		&StateToken{
			compile(`(.*?)()(?:\n|$)`),
			[]string{TOKEN_LINECOMMENT, TOKEN_LINECOMMENT_END},
			"#pop",
		},
	}

	return lexer
}

func (lexer *Lexer) tokeniter(source string) chan *Token {
	c := make(chan *Token)

	go func() {
		lines := strings.Split(source, "\n")
		if lines[len(lines)-1] == "" {
			lines = lines[0 : len(lines)-1]
		}
		source = strings.Join(lines, "\n")
		pos := 0
		lineno := 1
		stack := []string{"root"}
		//state := "root"
		// stack = append(stack, "en")
		stateTokens := lexer.rules[stack[len(stack)-1]]
		//sourceLength := len(source)

		balancingStack := make([]string, 0)

		for {
			breaked := false
			for _, stateToken := range stateTokens {
				regex, tokens, newState := stateToken.regex, stateToken.tokens, stateToken.newState
				m := regex.MatchString(source)
				index := regex.FindStringIndex(source)
				if m == false {
					continue
				}
				if len(balancingStack) > 0 && (tokens[0] == "variable_end" || tokens[0] == "block_end" || tokens[0] == "linestatement_end") {
					continue
				}

				if len(tokens) > 1 {
					for idx, token := range tokens {
						if token == "#bygroup" {
							subMap := FindStringSubmatchMap(regex, source)
							if len(subMap) <= 0 {
								panic("Can't resolve token, no group matched")
							}
							for key, value := range subMap {
								c <- &Token{lineno, key, value}
								lineno += strings.Count(value, "\n")
								break
							}
						} else {
							data := regex.FindStringSubmatch(source)[idx+1]
							if data != "" || !ignoreIfEmpty[token] {
								c <- &Token{lineno, token, data}
							}
							lineno += strings.Count(data, "\n")
						}
					}
				} else {
					token := tokens[0]
					if strings.HasPrefix(token, "Failure") {
						panic(token)
					}
					data := regex.FindString(source)
					if token == "operator" {
						if data == "{" {
							balancingStack = append(balancingStack, "}")
						} else if data == "(" {
							balancingStack = append(balancingStack, ")")
						} else if data == "[" {
							balancingStack = append(balancingStack, "]")
						} else if data == "}" || data == ")" || data == "]" {
							if len(balancingStack) <= 0 {
								panic("unexpected '" + data + "'")
							}
							expectedOp := balancingStack[len(balancingStack)-1]
							if expectedOp != data {
								panic(fmt.Sprintf("unexpected '%s', expected '%s'", data, expectedOp))
							}
							balancingStack = balancingStack[:len(balancingStack)-1]
						}
					}
					if data != "" || !ignoreIfEmpty[token] {
						c <- &Token{lineno, token, data}
					}
					lineno += strings.Count(data, "\n")
				}

				pos2 := index[1]

				if newState != "nil" {
					if newState == "#pop" {
						stack = stack[:len(stack)-1]
					} else if newState == "#bygroup" {
						subMap := FindStringSubmatchMap(regex, source)
						if len(subMap) <= 0 {
							panic("Can't resolve new state, no group matched")
						}
						for key, _ := range subMap {
							stack = append(stack, key)
							break
						}
					} else {
						stack = append(stack, newState)
					}
					stateTokens = lexer.rules[stack[len(stack)-1]]
				} else if pos2 == pos {
					panic("empty string yielded and without stack change")
				}

				pos = pos2
				source = source[pos:]
				breaked = true
				break
			}
			if !breaked {
				if len(source) <= 0 {
					break
				}
				panic(fmt.Sprintf("unexpected char %s at %d", source[0], lineno))
			}
		}
		close(c)
	}()

	return c
}

func (lexer *Lexer) wrap(stream chan *Token) chan *Token {
	c := make(chan *Token)
	var wrappedValue interface{}

	go func() {
		for token := range stream {
			lineno := token.lineno
			tokenType := token.tp
			value := token.value.(string)
			wrappedValue = token.value
			if ignoredTokens[tokenType] {
				continue
			} else if tokenType == "raw_begin" || tokenType == "raw_end" {
				continue
			} else if tokenType == "data" {
				wrappedValue = newlineRe.ReplaceAllString(value, "\n")
			} else if tokenType == "keyword" {
				tokenType = value
			} else if tokenType == "name" {
				wrappedValue = string(value)
			} else if tokenType == "string" {
				wrappedValue = string(value)
			} else if tokenType == "integer" {
				wrappedValue, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}
			} else if tokenType == "float" {
				wrappedValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					panic(err)
				}
			} else if tokenType == "operator" {
				tokenType = operatorsMap[value]
			}
			c <- &Token{lineno, tokenType, wrappedValue}
		}
		close(c)
	}()

	return c
}

func (lexer *Lexer) tokenize(source string) *TokenStream {
	stream := lexer.tokeniter(source)
	return &TokenStream{lexer.wrap(stream), &Token{1, TOKEN_INITIAL, ""}}
}
