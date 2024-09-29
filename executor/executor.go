package executor

import (
    "fmt"
    "raderh2o/brainfuck_interpreter/parser"
)


func ExecuteBf(bfTokens []parser.Bf, current *int, vals []uint8) []uint8 {
    values := vals
    for _, operation := range bfTokens {
        switch v := operation.(type) {
        case parser.Operation:
            switch v {
            case parser.BfAdd:
                if values[*current] >= 255 {
                    values[*current] = 0
                } else {
                    values[*current]++
                }
            case parser.BfReduce:
                if values[*current] <= 0 {
                    values[*current] = 255
                } else {
                    values[*current]--
                }
            case parser.BfNext:
                if cap(values) >= *current+1 {
                    *current++
                    values = append(values, uint8(0))
                }
            case parser.BfPrevious:
                if *current <= 0 {
                    values = append([]uint8{0}, values...)
                    *current = 0
                } else {
                    *current--
                }
            case parser.BfPrint:
                fmt.Print(string(values[*current]))
            case parser.BfInput:
                var i string = ""
                fmt.Scan(&i)
                values[*current] = uint8(i[0])
            }
        case parser.BfLoop:
            endLoop := false
            for !endLoop {
                values = ExecuteBf(v.Body, current, values)
                if values[*current] == 0 {
                    endLoop = true
                }
            }
        }
    }
    return values
}
