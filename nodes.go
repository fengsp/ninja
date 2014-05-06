package ninja

type Node struct {
	fields []string
	values map[string]string
	lineno int
}

func (node *Node) iterFields() string {
	for _, name := range node.fields {
		return name
	}
	return "end"
}

type StmtNode struct {
	Node
}

type HelperNode struct {
	Node
}

type ExprNode struct {
	Node
}

type BinExprNode struct {
	ExprNode // left right
}

type UnaryExprNode struct {
	ExprNode // node
}

type NameNode struct {
	ExprNode
}

type LiteralNode struct {
	ExprNode
}

type ConstNode struct {
	LiteralNode
}

type TemplateDataNode struct {
	LiteralNode
}

type TupleNode struct {
	LiteralNode
}

type ListNode struct {
	LiteralNode
}

type DictNode struct {
	LiteralNode
}

type PairNode struct {
	HelperNode
}

type KeywordNode struct {
	HelperNode
}

type CondExprNode struct {
	ExprNode
}

type FilterNode struct {
	ExprNode
}

type CallNode struct {
	ExprNode
}

func NewTemplateNode(body string, lineno int) {
	node := new(Node)
	node.fields = []string{"body"}
	node.values["body"] = body
	node.lineno = lineno
	// node.attributes["environment"] = environment
}

func NewOutputNode(nodes string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"nodes"}
	node.values["nodes"] = nodes
	node.lineno = lineno
}

func NewExtendsNode(template string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"template"}
	node.values["template"] = template
	node.lineno = lineno
}

func NewForNode(target string, iter string, body string, else_ string, test string, recursive string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"target", "iter", "body", "else_", "test", "recursive"}
	node.values["target"] = target
	node.values["iter"] = iter
	node.values["body"] = body
	node.values["else_"] = else_
	node.values["test"] = test
	node.values["recursive"] = recursive
	node.lineno = lineno
}

func NewIfNode(test string, body string, else_ string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"test", "body", "else_"}
	node.values["test"] = test
	node.values["body"] = body
	node.values["else_"] = else_
	node.lineno = lineno
}

func NewMacroNode(name string, args string, defaults string, body string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"name", "args", "defaults", "body"}
	node.values["name"] = name
	node.values["args"] = args
	node.values["defaults"] = defaults
	node.values["body"] = body
	node.lineno = lineno
}

func NewCallBlockNode(call string, args string, defaults string, body string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"call", "args", "defaults", "body"}
	node.values["call"] = call
	node.values["args"] = args
	node.values["defaults"] = defaults
	node.values["body"] = body
	node.lineno = lineno
}

func NewFilterBlockNode(body string, filter string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"body", "filter"}
	node.values["body"] = body
	node.values["filter"] = filter
	node.lineno = lineno
}

func NewBlockNode(name string, body string, scoped string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"name", "body", "scoped"}
	node.values["name"] = name
	node.values["body"] = body
	node.values["scoped"] = scoped
	node.lineno = lineno
}

func NewIncludeNode(template string, with_context string, ignore_missing string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"template", "with_context", "ignore_missing"}
	node.values["template"] = template
	node.values["with_context"] = with_context
	node.values["ignore_missing"] = ignore_missing
	node.lineno = lineno
}

func NewImportNode(template string, target string, with_context string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"template", "target", "with_context"}
	node.values["template"] = template
	node.values["target"] = target
	node.values["with_context"] = with_context
	node.lineno = lineno
}

func NewFromImportNode(template string, names string, with_context string, lineno int) {
	node := new(StmtNode)
	node.fields = []string{"template", "names", "with_context"}
	node.values["template"] = template
	node.values["names"] = names
	node.values["with_context"] = with_context
	node.lineno = lineno
}

func NewExprStmtNodeNode(node string, lineno int) {
	n := new(StmtNode)
	n.fields = []string{"node"}
	n.values["node"] = node
	n.lineno = lineno
}

func NewAssignNode(target string, node string, lineno int) {
	n := new(StmtNode)
	n.fields = []string{"target", "node"}
	n.values["target"] = target
	n.values["node"] = node
	n.lineno = lineno
}

func NewNameNode(name string, ctx string, lineno int) {
	node := new(NameNode)
	node.fields = []string{"name", "ctx"}
	node.values["name"] = name
	node.values["ctx"] = ctx
	node.lineno = lineno
}

func NewConstNode(value string, lineno int) {
	node := new(ConstNode)
	node.fields = []string{"value"}
	node.values["value"] = value
	node.lineno = lineno
}

func NewTemplateDataNode(data string, lineno int) {
	node := new(TemplateDataNode)
	node.fields = []string{"data"}
	node.values["data"] = data
	node.lineno = lineno
}

func NewTupleNode(items string, ctx string, lineno int) {
	node := new(TupleNode)
	node.fields = []string{"items", "ctx"}
	node.values["items"] = items
	node.values["ctx"] = ctx
	node.lineno = lineno
}

func NewListNode(items string, lineno int) {
	node := new(ListNode)
	node.fields = []string{"items"}
	node.values["items"] = items
	node.lineno = lineno
}

func NewDictNode(items string, lineno int) {
	node := new(DictNode)
	node.fields = []string{"items"}
	node.values["items"] = items
	node.lineno = lineno
}

func NewPairNode(key string, value string, lineno int) {
	node := new(PairNode)
	node.fields = []string{"key", "value"}
	node.values["key"] = key
	node.values["value"] = value
	node.lineno = lineno
}

func NewKeywordNode(key string, value string, lineno int) {
	node := new(KeywordNode)
	node.fields = []string{"key", "value"}
	node.values["key"] = key
	node.values["value"] = value
	node.lineno = lineno
}

func NewCondExprNode(test string, expr1 string, expr2 string, lineno int) {
	node := new(CondExprNode)
	node.fields = []string{"test", "expr1", "expr2"}
	node.values["test"] = test
	node.values["expr1"] = expr1
	node.values["expr2"] = expr2
	node.lineno = lineno
}

func NewFilterNode(node string, name string, args string, kwargs string, dyn_args string, dyn_kwargs string) {
	n := new(FilterNode)
	n.fields = []string{"node", "name", "args", "kwargs", "dyn_args", "dyn_kwargs"}
	n.values["node"] = node
	n.values["name"] = name
	n.values["args"] = args
	n.values["kwargs"] = kwargs
	n.values["dyn_args"] = dyn_args
	n.values["dyn_kwargs"] = dyn_kwargs
}

func NewCallNode(node string, args string, kwargs string, dyn_args string, dyn_kwargs string) {
	n := new(FilterNode)
	n.fields = []string{"node", "args", "kwargs", "dyn_args", "dyn_kwargs"}
	n.values["node"] = node
	n.values["args"] = args
	n.values["kwargs"] = kwargs
	n.values["dyn_args"] = dyn_args
	n.values["dyn_kwargs"] = dyn_kwargs
}
