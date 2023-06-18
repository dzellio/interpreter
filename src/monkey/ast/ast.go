package ast

type Node interface {
    TokenLiteral() string
}

type Statement interfce {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

type Program struct {
    Statements []statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiterlal()
    } else {
        return ""
    }

}