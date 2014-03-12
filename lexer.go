package ninja

import (
	"fmt"
	"regexp"
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

var operators = map[string]string{
	"+":  TOKEN_ADD,
	"-":  TOKEN_SUB,
	"/":  TOKEN_DIV,
	"//": TOKEN_FLOORDIV,
	"*":  TOKEN_MUL,
	"%%":  TOKEN_MOD,
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

func compile(x string) *regexp.Regexp {
	x = `(?ms)^` + x
	r, _ := regexp.Compile(x)
	return r
}

type Token struct {
	lineno int
	tp     string
	value  string
}

type StateToken struct {
	regex    *regexp.Regexp
	tokens   []string
	newState string
}

type TokenStream struct {
}

type Lexer struct {
	rules map[string][]*StateToken
}

func NewLexer() *Lexer {
	lexer := new(Lexer)

	whitespaceRe, _ := regexp.Compile(`\s+`)
	floatRe, _ := regexp.Compile(`(?<!\.)\d+\.\d+`)
	integerRe, _ := regexp.Compile(`\d+`)
	nameRe, _ := regexp.Compile(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`)
	stringRe, _ := regexp.Compile(`(?s)('([^'\\]*(?:\\.[^'\\]*)*)'|"([^"\\]*(?:\\.[^"\\]*)*)")`)
	keyArray := make([]string, 0)
	for k, _ := range operators {
		keyArray = append(keyArray, k)
	}
	operatorRe, _ := regexp.Compile(fmt.Sprintf("(%s)", strings.Join(keyArray, "|")))

	tagRules := []*StateToken{
		&StateToken{whitespaceRe, []string{TOKEN_WHITESPACE}, "nil"},
		&StateToken{floatRe, []string{TOKEN_FLOAT}, "nil"},
		&StateToken{integerRe, []string{TOKEN_INTEGER}, "nil"},
		&StateToken{nameRe, []string{TOKEN_NAME}, "nil"},
		&StateToken{stringRe, []string{TOKEN_STRING}, "nil"},
		&StateToken{operatorRe, []string{TOKEN_OPERATOR}, "nil"},
	}

	lstripRe := `^[ \t]*`
	noLstripRe := `+`
	blockPrefixRe := fmt.Sprintf(`%s{%%(?!%s)|{%%\+?`, lstripRe, noLstripRe)

	lexer.rules = make(map[string][]*StateToken)

	rootTagRules := map[string]string{
		"comment":  "{#",
		"block":    "{%",
		"variable": "{{",
	}
	regexArray := []string{}
	regexArray = append(regexArray, fmt.Sprintf(`(?P<raw_begin>(?:\s*{%%\-|%s)\s*raw\s*(?:\-%%}\s*|%%}))`, blockPrefixRe))
	for n, r := range rootTagRules {
		regex := fmt.Sprintf(`(?P<%s_begin>\s*%s\-)`, n, r)
		regexArray = append(regexArray, regex)
	}
    fmt.Println(fmt.Sprintf(`(.*?)(?:%s)`, strings.Join(regexArray, `|`)))
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
			compile(`(.*?)((?:\-#}\s*|#}))`),
			[]string{TOKEN_COMMENT, TOKEN_COMMENT_END},
			"#pop",
		},
		&StateToken{
			compile("(.)"),
			[]string{"Failure"},
			"nil",
		},
	}

	lexer.rules[TOKEN_BLOCK_BEGIN] = append([]*StateToken{
		&StateToken{
			compile(`(?:\-%%}\s*|%%})`),
			[]string{TOKEN_BLOCK_END},
			"#pop",
		},
	}, tagRules...)

	lexer.rules[TOKEN_VARIABLE_BEGIN] = append([]*StateToken{
		&StateToken{
			compile(`\-}}\s*|}}`),
			[]string{TOKEN_VARIABLE_END},
			"#pop",
		},
	}, tagRules...)

	lexer.rules[TOKEN_RAW_BEGIN] = []*StateToken{
		&StateToken{
			compile(fmt.Sprintf(`(.*?)((?:\s*{%%\-|%s)\s*endraw\s*(?:\-%%}\s*|%%}))`, blockPrefixRe)),
			[]string{TOKEN_DATA, TOKEN_RAW_END},
			"#pop",
		},
		&StateToken{
			compile("(.)"),
			[]string{"Failure"},
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
			compile(`(.*?)()(?=\n|$)`),
			[]string{TOKEN_LINECOMMENT, TOKEN_LINECOMMENT_END},
			"#pop",
		},
	}

    return lexer
}

func (lexer *Lexer) tokeniter(source string) chan *Token {
	c := make(chan *Token)

	go func() {
		//lines := strings.Split(source, "\n")
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
							for key, value := range subMap {
								c <- &Token{lineno, key, value}
								lineno += strings.Count(value, "\n")
								break
							}
						} else {
							data := regex.FindStringSubmatch(source)[idx+1]
							c <- &Token{lineno, token, data}
							lineno += strings.Count(data, "\n")
						}
					}
				} else {
					token := tokens[0]
					data := regex.FindStringSubmatch(source)[0]
					if token == "operator" {
						if data == "{" {
							balancingStack = append(balancingStack, "}")
						} else if data == "(" {
							balancingStack = append(balancingStack, ")")
						} else if data == "[" {
							balancingStack = append(balancingStack, "]")
						} else if data == "}" || data == ")" || data == "]" {
							balancingStack = balancingStack[:len(balancingStack)-1]
						}
					}
					c <- &Token{lineno, token, data}
					lineno += strings.Count(data, "\n")
				}

				pos2 := index[1]

				if newState != "nil" {
					if newState == "#pop" {
						stack = stack[:len(stack)-1]
					} else if newState == "#bygroup" {
						subMap := FindStringSubmatchMap(regex, source)
						for key, _ := range subMap {
							stack = append(stack, key)
							break
						}
					} else {
						stack = append(stack, newState)
					}
					stateTokens = lexer.rules[stack[len(stack)-1]]
				}

				pos = pos2
				source = source[pos:]
				breaked = true
				break
			}
			if !breaked {
				break
				//panic("weird")
			}
		}
		close(c)
	}()

	return c
}

func (lexer *Lexer) tokenize(source string) chan *Token {
	stream := lexer.tokeniter(source)
    return stream
}
