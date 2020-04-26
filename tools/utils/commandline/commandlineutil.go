package commandlineutil

import (
    "bufio"
    "fmt"
    "golang.org/x/crypto/ssh/terminal"
    "os"
    "strconv"
    "strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadUserInput() (str string, err error) {
    str, err = reader.ReadString('\n')
    if err != nil {
        return
    }
    str = strings.TrimSuffix(str, "\n")
    return
}

func ReadUserInputOfLength(minLength uint, maxLength uint) (str string, err error) {
    for {
        str, err = ReadUserInput()
        if err != nil {
            return
        }
        strLen := uint(len(str))
        if minLength <= strLen && strLen <= maxLength {
            return
        }
        fmt.Printf("Input length must be between %d and %d:", minLength, maxLength)
    }
}

func ReadInteger() (num int, err error) {
    str, err := ReadUserInput()
    if err != nil {
        return
    }
    num, err = strconv.Atoi(str)
    return
}

func ReadIntegerBetween(min int, max int) (num int, err error) {
    num, err = ReadInteger()
    if err != nil {
        return
    }
    if num < min || num > max {
        err = fmt.Errorf("%d is out of range", num)
        return
    }
    return
}

func ReadUserPassword() string {
    bytePassword, _ := terminal.ReadPassword(0)
    return string(bytePassword)
}