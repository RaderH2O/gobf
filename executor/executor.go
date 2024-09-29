package executor

import (
    "fmt"
    "raderh2o/brainfuck_interpreter/parser"
)


func ExecuteBf(bfTokens []parser.Bf, current *int, vals []uint8) []uint8 {
    for _, operation := range bfTokens {
        switch v := operation.(type) {
        case parser.Operation:
            switch v {
            case parser.BfAdd:
                if vals[*current] >= 255 {
                    vals[*current] = 0
                } else {
                    vals[*current]++
                }
            case parser.BfReduce:
                if vals[*current] <= 0 {
                    vals[*current] = 255
                } else {
                    vals[*current]--
                }
            case parser.BfNext:
                if cap(vals) >= *current+1 {
                    *current++
                    vals = append(vals, uint8(0))
                }
            case parser.BfPrevious:
                if *current <= 0 {
                    vals = append([]uint8{0}, vals...)
                    *current = 0
                } else {
                    *current--
                }
            case parser.BfPrint:
                fmt.Print(string(vals[*current]))
            case parser.BfInput:
                var i string = ""
                fmt.Scan(&i)
                vals[*current] = uint8(i[0])
            }
        case parser.BfLoop:
            endLoop := false
            for !endLoop {
                vals = ExecuteBf(v.Body, current, vals)
                if vals[*current] == 0 {
                    endLoop = true
                }
            }
        }
    }
    return vals
}
