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

    body := []interface{}{}
	dataBuffer := []interface{}{}

	for parser.stream {
		if token.tp == "data" {
			append(dataBuffer, TemplateDataNode{token.value, token.lineno, false})
            parser.stream.next()
		} else if token.tp == "variable_begin" {
            parser.stream.next()
			append(dataBuffer, parser.parseTuple())
            parser.stream.expect("variable_end")
		} else if token.tp == "block_begin" {
			if len(dataBuffer) > 0 {
                lineno := dataBuffer[0].lineno
                append(body, NewOutputNode(dataBuffer, lineno))
                dataBuffer = dataBuffer[:0]
            }
            parser.stream.next()
            rv := parser.parseStatement()
            append(body, rv...)
            parser.stream.expect('block_end')
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
