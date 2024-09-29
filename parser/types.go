package parser

type Operation uint8

const (
    BfAdd Operation = iota
    BfReduce
    BfNext
    BfPrevious
    BfInput
    BfPrint
)

func (Operation) bf() {
    panic("bf called")
}

func (op Operation) String() string {
    switch op {
    case BfAdd:
        return "BfAdd"
    case BfReduce:
        return "BfReduce"
    case BfNext:
        return "BfNext"
    case BfPrevious:
        return "BfPrevious"
    case BfPrint:
        return "BfPrint"
    case BfInput:
        return "BfInput"
    }
    return ""
}

type Bf interface {
    String() string
    bf()
}

type BfLoop struct {
    Body []Bf
}

func (BfLoop) bf() {
    panic("bf called")
}

func (loop BfLoop) String() string {
    output := "BfLoop { "

    for _, op := range loop.Body {
        output += op.String()
        output += " "
    }

    output += "}"
    return string(output)
}

