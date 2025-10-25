package parser

import (
	"strconv"
)

// func main() {
// 	// startNode := BuildTree(["2", "-", "3", "/", "3", "+", "(", "(", "3", "+", "2", ")", "*", "5", "^", "2", ")", "+", "2"]

// 	startNode := BuildTree([]string{"2", "-", "3", "/", "3", "+", "(", "(", "3", "+", "2", ")", "*", "5", "^", "2", ")", "+", "2"})
// 	fmt.Println("value = ", EvaluateTree(startNode, 0))
// }

type Node struct {
	value string
	left  *Node
	right *Node
}

func CreateNode(op string, l *Node, r *Node) *Node {
	return &Node{
		value: op,
		left:  l,
		right: r,
	}
}

func applyOperation(nodes []*Node, operators []string) ([]*Node, []string) {
	op := operators[len(operators)-1]
	operators = operators[:len(operators)-1]

	rightNode := nodes[len(nodes)-1]
	nodes = nodes[:len(nodes)-1]

	leftNode := nodes[len(nodes)-1]
	nodes = nodes[:len(nodes)-1]

	newNode := CreateNode(op, leftNode, rightNode)
	nodes = append(nodes, newNode)

	return nodes, operators
}

func BuildTree(expressionToken []string) *Node {
	var operators []string
	var nodes []*Node

	precedenceMap := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}

	for i := 0; i < len(expressionToken); i++ {
		currentToken := expressionToken[i]

		switch currentToken {
		case "(":
			operators = append(operators, currentToken)
		case ")":
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				nodes, operators = applyOperation(nodes, operators)
			}
			operators = operators[:len(operators)-1]
		case "+", "-", "*", "/", "^":
			currentTokenPrecedence := precedenceMap[currentToken]

			for len(operators) > 0 {
				top := operators[len(operators)-1]
				topOperatorPrecedence := precedenceMap[top]

				if top == "(" || topOperatorPrecedence < currentTokenPrecedence {
					break
				}

				nodes, operators = applyOperation(nodes, operators)
			}
			operators = append(operators, currentToken)
		default:
			node := CreateNode(currentToken, nil, nil)
			nodes = append(nodes, node)
		}
	}

	for len(operators) > 0 {
		nodes, operators = applyOperation(nodes, operators)
	}

	if len(nodes) == 0 {
		return nil
	}

	return nodes[0]
}

// func printExpression(node *Node) string {
// 	if node == nil {
// 		return ""
// 	}
// 	if node.left == nil && node.right == nil {
// 		return node.value
// 	}
// 	leftExpr := printExpression(node.left)
// 	rightExpr := printExpression(node.right)
// 	return "(" + leftExpr + " " + node.value + " " + rightExpr + ")"
// }

func EvaluateTree(rootNode *Node, subExpressionValue int) int {
	if rootNode.left == nil && rootNode.right == nil {
		val, _ := strconv.Atoi(rootNode.value)
		return val
	}

	leftValue := EvaluateTree(rootNode.left, subExpressionValue)
	rightValue := EvaluateTree(rootNode.right, subExpressionValue)

	operator := rootNode.value
	// subExpressionValue := 0

	switch operator {
	case "+":
		subExpressionValue += leftValue + rightValue
	case "-":
		subExpressionValue += leftValue - rightValue
	case "*":
		subExpressionValue += leftValue * rightValue
	case "/":
		subExpressionValue += leftValue / rightValue
	case "^":
		product := 1
		for i := 1; i <= rightValue; i++ {
			product *= leftValue
		}
		subExpressionValue += product
	}
	return subExpressionValue
}

// func applyOperation(nodes *[]*Node, operators *[]string) {
// 	op := operators[len(operators) - 1]
// 	*operators = (*operators)[:len(operators) - 1]

// 	leftNode := (*nodes)[len(nodes) - 1]
// 	*nodes = (*nodes)[:len(nodes) - 1]

// 	rightNode := (*nodes)[len(nodes) - 1]
// 	*nodes = (*nodes)[:len(nodes) - 1]

// 	combinedNewNode := CreateNode(op, leftNode, rightNode)
// 	*nodes = append(*nodes, combinedNewNode)
// }
