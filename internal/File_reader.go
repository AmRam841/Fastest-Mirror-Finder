package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// readLinesFromFile reads the file line by line and prints each line as-is
func readLinesFromFile(path string) error {
    fileHandle, err := os.Open(path)
    if err != nil {
        return err
    }
    defer fileHandle.Close()
    scanner := bufio.NewReader(fileHandle)
    for {
        textLine, err := scanner.ReadString('\n')
        if err == io.EOF {
            if len(textLine) != 0 {
                fmt.Print(textLine) // Print last line if not empty
            }
            break
        }
        if err != nil {
            return fmt.Errorf("error reading from file: %w", err)
        }
        fmt.Print(textLine)
    }
    return nil
}