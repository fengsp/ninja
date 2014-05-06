package ninja

type Parser struct {
	stream *TokenStream
}

func NewParser(source string) *Parser {
	lexer := NewLexer()
	stream := lexer.tokenize(source)
	parser := &Parser{stream: stream}
	return parser
}

func (parser *Parser) subParse() {
	c := parser.tokenize()

	body := []string{}
	dataBuffer := []interface{}{}

	for token := range c {
		if token.tp == "data" {
			append(dataBuffer, TemplateDataNode{token.value, token.lineno, false})
		} else if token.tp == "variable_begin" {
			append(dataBuffer, parser.parseTuple())
		} else if token.tp == "block_begin" {
			pass
		} else {
			panic("internal parsing error")
		}
	}
	return body
}

// Parse the template into a node
func (parser *Parser) parse() {
	result := NewTemplateNode(parser.subParse(), 1)
	return result
}
