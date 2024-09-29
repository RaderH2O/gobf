package parser

import "strings"

func ParseBf(input string) []Bf {

    output := make([]Bf, 0)
    // loop := BfLoop{}
    var loop BfLoop = BfLoop{}
    inLoop := false

    for index := 0; index < len(input); index++ {
        v := input[index]
        switch v {
        case '+':
            if inLoop {
                loop.Body = append(loop.Body, BfAdd)
            } else {
                output = append(output, BfAdd)
            }
        case '-':
            if inLoop {
                loop.Body = append(loop.Body, BfReduce)
            } else {
                output = append(output, BfReduce)
            }
        case '>':
            if inLoop {
                loop.Body = append(loop.Body, BfNext)
            } else {
                output = append(output, BfNext)
            }
        case '<':
            if inLoop {
                loop.Body = append(loop.Body, BfPrevious)
            } else {
                output = append(output, BfPrevious)
            }
        case '.':
            if inLoop {
                loop.Body = append(loop.Body, BfPrint)
            } else {
                output = append(output, BfPrint)
            }
        case ',':
            if inLoop {
                loop.Body = append(loop.Body, BfInput)
            } else {
                output = append(output, BfInput)
            }
        case '[':
            if !inLoop {
                inLoop = true
            } else {
                endOfNestedLoop := strings.Index(input[index:], "]")
                loop.Body = append(loop.Body, ParseBf(input[index:index + endOfNestedLoop + 1])...)
                index += endOfNestedLoop
            }

        case ']':
            if inLoop {
                inLoop = false
                output = append(output, loop)
                loop = BfLoop{}
            }
        }
    }

    return output
}
