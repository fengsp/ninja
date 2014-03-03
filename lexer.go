package ninja

import (
    "regexp"
    "fmt"
    "strings"
)

const (
    TOKEN_ADD = "add"
    TOKEN_ASSIGN = "assign"
    TOKEN_COLON = "colon"
    TOKEN_COMMA = "comma"
    TOKEN_DIV = "div"
    TOKEN_DOT = "dot"
    TOKEN_EQ = "eq"
    TOKEN_FLOORDIV = "floordiv"
    TOKEN_GT = "gt"
    TOKEN_GTEQ = "gteq"
    TOKEN_LBRACE = "lbrace"
    TOKEN_LBRACKET = "lbracket"
    TOKEN_LPAREN = "lparen"
    TOKEN_LT = "lt"
    TOKEN_LTEQ = "lteq"
    TOKEN_MOD = "mod"
    TOKEN_MUL = "mul"
    TOKEN_NE = "ne"
    TOKEN_PIPE = "pipe"
    TOKEN_POW = "pow"
    TOKEN_RBRACE = "rbrace"
    TOKEN_RBRACKET = "rbracket"
    TOKEN_RPAREN = "rparen"
    TOKEN_SEMICOLON = "semicolon"
    TOKEN_SUB = "sub"
    TOKEN_TILDE = "tilde"
    TOKEN_WHITESPACE = "whitespace"
    TOKEN_FLOAT = "float"
    TOKEN_INTEGER = "integer"
    TOKEN_NAME = "name"
    TOKEN_STRING = "string"
    TOKEN_OPERATOR = "operator"
    TOKEN_BLOCK_BEGIN = "block_begin"
    TOKEN_BLOCK_END = "block_end"
    TOKEN_VARIABLE_BEGIN = "variable_begin"
    TOKEN_VARIABLE_END = "variable_end"
    TOKEN_RAW_BEGIN = "raw_begin"
    TOKEN_RAW_END = "raw_end"
    TOKEN_COMMENT_BEGIN = "comment_begin"
    TOKEN_COMMENT_END = "comment_end"
    TOKEN_COMMENT = "comment"
    TOKEN_LINESTATEMENT_BEGIN = "linestatement_begin"
    TOKEN_LINESTATEMENT_END = "linestatement_end"
    TOKEN_LINECOMMENT_BEGIN = "linecomment_begin"
    TOKEN_LINECOMMENT_END = "linecomment_end"
    TOKEN_LINECOMMENT = "linecomment"
    TOKEN_DATA = "data"
    TOKEN_INITIAL = "initial"
    TOKEN_EOF = "eof"
)

func compile(x string) *regexp.Regexp {
    x = "(?ms)" + x
    r, _ := regexp.Compile(x)
    return r
}

type Token struct {
    lineno int
    type
    value string
}

type StateToken struct {
    regex *regexp.Regexp
    tokens []string
    newState string
}

type TokenStream struct {
}

type Lexer struct {
    rules map[string][]*StateToken
}

func NewLexer() *Lexer {
    lexer := new(Lexer)

    lstripRe := "^[ \t]*"
    noLstripRe = "+"
    blockPrefixRe := fmt.Sprintf("%s{%(?!%s)|{%\+?", lstripRe, noLstripRe)

    lexer.rules = make(map[string][]*StateToken)

    rootTagRules := map(string)string {
        "comment": "{#",
        "block": "{%",
        "variable": "{{"
    }
    regexArray := []string{}
    append(regexArray, fmt.Sprintf("(?P<raw_begin>(?:\s*{%\-|%s)\s*raw\s*(?:\-%}\s*|%}))", blockPrefixRe))
    for n, r := range rootTagRules {
        regex := fmt.Sprintf("(?P<%s_begin>\s*%s\-|%s)", n, r, prefixRe.get(n, r))
        append(regexArray, regex)
    }
    lexer.rules["root"] = []*StateToken {
        &StateToken{
            compile(fmt.Sprintf("(.*?)(?:%s)", strings.Join(rules, "|"))),
            []string{TOKEN_DATA, "#bygroup"},
            "#bygroup"
        },
        &StateToken{
            compile(".+"),
            []string{TOKEN_DATA},
            nil
        }
    }

    lexer.rules[TOKEN_COMMENT_BEGIN] = []*StateToken {
        &StateToken {
            compile("(.*?)((?:\-#}\s*|#}))"),
            []string{TOKEN_COMMENT, TOKEN_COMMENT_END},
            "#pop"
        },
        &StateToken {
            compile("(.)"),
            []string{"Failure"},
            nil
        }
    }

    lexer.rules[TOKEN_BLOCK_BEGIN] = append([]*StateToken {
        &StateToken {
            compile("(?:\-%}\s*|%})"),
            []string{TOKEN_BLOCK_END},
            "#pop"
        }
    }, tagRules)

    lexer.rules[TOKEN_VARIABLE_BEGIN = append([]*StateToken {
        &StateToken {
            compile("\-}}\s*|}}"),
            []string{TOKEN_VARIABLE_END},
            "#pop"
        }
    }, tagRules)

    lexer.rules[TOKEN_RAW_BEGIN] = []*StateToken {
        &StateToken {
            compile(fmt.Sprintf("(.*?)((?:\s*{%\-|%s)\s*endraw\s*(?:\-%}\s*|%}))", blockPrefixRe)),
            []string{TOKEN_DATA, TOKEN_RAW_END},
            "#pop"
        },
        &StateToken {
            compile("(.)"),
            []string{"Failure"},
            nil
        }
    }

    lexer.rules[TOKEN_LINESTATEMENT_BEGIN] = append([]*StateToken {
        &StateToken {
            compile("\s*(\n|$)"),
            []string{TOKEN_LINESTATEMENT_END},
            "#pop"
        }s
    }, tagRules...)

    lexer.rules[TOKEN_LINECOMMENT_BEGIN] = []*StateToken {
        &StateToken {
            compile("(.*?)()(?=\n|$)"),
            []string{TOKEN_LINECOMMENT, TOKEN_LINECOMMENT_END},
            "#pop"
        }
    }
}

func (lexer *Lexer) tokeniter(source string) {
    lines := source.Split("\n")
    pos := 0
    lineno := 1
    stack := []string{"root"}
    state := "root"
    // stack = append(stack, "en")
    stateTokens := lexer.rules[stack[-1]]
    sourceLength := len(source)
    
    balancingStack := []string

    for {
        for regex, tokens, newState := range stateTokens {
            m := 
}

func (lexer *Lexer) tokenize(source string) {
    stream := lexer.tokeniter(source)
}
