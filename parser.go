package ninja

type Parser struct {
	source string
    lexer *Lexer
}

func NewParser(source string) *Parser {
	lexer := NewLexer()
	parser := &Parser{source: source, lexer: lexer}
	return parser
}

func (parser *Parser) tokenize() chan *Token {
	stream := parser.lexer.tokenize(parser.source)
	return stream
}

func (parser *Parser) subParse() {
    
}

// Parse the template into a node
func (parser *Parser) parse() {
	
}
