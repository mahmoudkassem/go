package main

import "fmt"

//encoding : a -> 1, b -> 2, ... , z -> 26
func main() {
    count := countDigitSequenceDecodings("111") // aaa, ak, ka
    fmt.Printf("%v -> %v\n", "111", count)

    count = countDigitSequenceDecodings("1234") // abcd, lcd, awd
    fmt.Printf("%v -> %v\n", "1234", count)

    count = countDigitSequenceDecodings("102") // jb
    fmt.Printf("%v -> %v\n", "102", count)
}

func countDigitSequenceDecodings(code string) (count int) {
    if len(code) == 0 || len(code) == 1 {
        return 1
    }

    if code[0] > '0'{
        count = countDigitSequenceDecodings(code[1 : ])
    }

    if code[0] == '1' || code[0] == '2' && code[1] < '7' {
        count += countDigitSequenceDecodings(code[2 : ])
    }

    return
}
