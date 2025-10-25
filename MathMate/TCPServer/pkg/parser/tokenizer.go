package parser

import (
	"fmt"
	"strconv"
)

func EvaluateExpression(exp string) (int, error) {
	expressionTokenArray, err := TokenizeEquation(exp)

	if err != nil {
		return 0, fmt.Errorf("Invalid Expression")
	}

	startNode := BuildTree(expressionTokenArray)
	expValue := EvaluateTree(startNode, 0)

	return expValue, nil
}

func removeSpaces(exp string) string {
	expression := ""

	for i := 0; i < len(exp); i++ {
		if string(exp[i]) == " " {
			continue
		}

		expression += string(exp[i])
	}

	return expression
}

func ValidToken(currentToken string, checkOperator bool) (string, bool) {
	validTokenMap := map[string]string{
		"*": "*",
		"/": "/",
		"-": "-",
		"+": "+",
		"^": "^",
		")": ")",
		"(": "(",
	}

	token, validTokenFlag := validTokenMap[currentToken]

	if checkOperator {
		if validTokenFlag && token != "(" && token != ")" {
			return token, true
		} else {
			return token, false
		}
	} else {
		_, err := strconv.Atoi(currentToken)

		if validTokenFlag {
			return token, validTokenFlag
		}

		if err == nil {
			return "Number", true
		} else {
			return "", false
		}
	}
}

func TokenizeEquation(exp string) ([]string, error) {
	if exp == "" {
		return nil, fmt.Errorf("invalid expression")
	}

	expression := removeSpaces(exp)
	expressionLength := len(expression)
	tokenAtBeginning, _ := ValidToken(string(expression[0]), true)
	tokenAtEnd, validOperatorAtEndFlag := ValidToken(string(expression[expressionLength-1]), true)
	invalidTokensAtStartAndEnd := tokenAtBeginning == "*" || tokenAtBeginning == "/" || tokenAtBeginning == ")" || tokenAtEnd == "(" || validOperatorAtEndFlag

	if invalidTokensAtStartAndEnd {
		return nil, fmt.Errorf("invalid expression")
	}

	bracketStack := make([]string, 0)
	expressionTokens := make([]string, 0)
	number := ""
	numberCount := 0
	consecutiveOperatorCount := 0
	nextToken := ""

	for i := 0; i < expressionLength; i++ {
		currentToken := string(expression[i])
		if currentToken == " " {
			continue
		}

		validToken, isCurrentTokenvalid := ValidToken(currentToken, false)

		//First if
		if !isCurrentTokenvalid {
			return nil, fmt.Errorf("invalid expression")
		}

		switch validToken {
		case "(":
			{
				if i-1 >= 0 {
					previousTokenType, _ := ValidToken(string(expression[i-1]), false)

					if previousTokenType == "Number" {
						expressionTokens = append(expressionTokens, "*")
					}
				}

				if i+1 < expressionLength {
					nextToken, _ = ValidToken(string(expression[i+1]), false)
					validNextToken := nextToken == "+" || nextToken == "-" || nextToken == "(" || nextToken == "Number"

					if validNextToken {
						expressionTokens = append(expressionTokens, currentToken)
						bracketStack = append(bracketStack, currentToken)
					} else {
						return nil, fmt.Errorf("invalid expression")
					}
				}
			}
		case ")":
			{
				if len(bracketStack) == 0 {
					return nil, fmt.Errorf("invalid expression")
				} else {
					bracketStack = bracketStack[:len(bracketStack)-1]
				}

				previousToken := expressionTokens[len(expressionTokens)-1]
				_, isPreviousTokenAnOperator := ValidToken(previousToken, true)

				if isPreviousTokenAnOperator || previousToken == "(" {
					return nil, fmt.Errorf("invalid expression")
				} else {
					expressionTokens = append(expressionTokens, currentToken)
				}

				if i+1 < expressionLength {
					nextTokenType, _ := ValidToken(string(expression[i+1]), false)

					if nextTokenType == "Number" {
						expressionTokens = append(expressionTokens, "*")
					}
				}
			}

		//could have used default instead, but too late to change now
		case "Number":
			{
				number += currentToken

				if i+1 < expressionLength {
					nextToken = string(expression[i+1])
					nextTokenType, _ := ValidToken(nextToken, false)

					if nextTokenType != "Number" {
						expressionTokens = append(expressionTokens, number)
						number = ""
					}
				} else {
					expressionTokens = append(expressionTokens, number)
				}
				numberCount++
			}
		case "*", "/", "^":
			{
				prevToken := string(expression[i-1])
				prevTokenType, _ := ValidToken(prevToken, false)
				nextToken = string(expression[i+1])
				nextTokenType, _ := ValidToken(nextToken, false)

				if (prevToken == ")" || prevTokenType == "Number") && (nextToken == "(" || nextToken == "+" || nextToken == "-" || nextTokenType == "Number") {
					expressionTokens = append(expressionTokens, currentToken)
				} else {
					// fmt.Println("Next token = ", nextToken)
					return nil, fmt.Errorf("invalid expression")
				}
			}
		case "+", "-":
			{
				consecutiveOperatorCount++
				nextToken = string(expression[i+1])

				if (nextToken == "+" || nextToken == "-") && consecutiveOperatorCount == 2 {
					return nil, fmt.Errorf("invalid expression")
				} else if (currentToken == "+" && nextToken == "+" || currentToken == "+" && nextToken == "-") && consecutiveOperatorCount == 1 {
					continue
				} else if currentToken == "-" && nextToken == "+" && consecutiveOperatorCount == 1 {
					expressionTokens = append(expressionTokens, currentToken)
					i++
					continue
				} else if currentToken == "-" && nextToken == "-" && consecutiveOperatorCount == 1 {
					expressionTokens = append(expressionTokens, "+")
					i++
					continue
				} else if nextToken != "-" && nextToken != "+" {
					consecutiveOperatorCount = 0
				}

				expressionTokens = append(expressionTokens, currentToken)
			}
		}
	}

	if len(bracketStack) > 0 || numberCount == 0 {
		return nil, fmt.Errorf("invalid expression")
	} else {
		return expressionTokens, nil
	}
}
