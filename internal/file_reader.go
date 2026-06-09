package file_reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)


func ReadLinesFromFile(path string) (string, error) {
    fileHandle, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer fileHandle.Close()

    scanner := bufio.NewReader(fileHandle)
    var builder strings.Builder  

    replacer := strings.NewReplacer("#", "", "$", "") 

    for {
        textLine, err := scanner.ReadString('\n')

        if strings.Contains(textLine, "#Server") {
            cleaned := replacer.Replace(textLine)
            builder.WriteString(cleaned)
        }

        
        if err == io.EOF {
            break  
        }
        if err != nil {
            return "", fmt.Errorf("error reading from file: %w", err)
        }
    }
    println(builder.String())
    return builder.String(), nil  
}