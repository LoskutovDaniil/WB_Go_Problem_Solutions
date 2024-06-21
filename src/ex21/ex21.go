package main

import "fmt"

type Writer interface {
    WriteString(text string)
}

func PrintToConsole(text string) {
    fmt.Println(text)
}

type ConsoleWriterAdapter struct{}

func (cwa *ConsoleWriterAdapter) WriteString(text string) {
    PrintToConsole(text)
}

func main() {
    adapter := &ConsoleWriterAdapter{}
    adapter.WriteString("Hello, Adapter Pattern!")
}
