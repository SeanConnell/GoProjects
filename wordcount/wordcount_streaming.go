package main

import (
    "sort"
    "strings"
    "bufio"
    "os"
    "io"
    "fmt"
)

func main() {
    args := os.Args
    if len(args) != 2 { 
        fmt.Println("Expected only one arg: absolute path to file to wordcount")
        os.Exit(1)
    }
    filename := args[1]
    fmt.Println("Counting words in file:", filename)

    file, err := os.Open(filename)
    if err != nil { panic(err) }

    r := bufio.NewReader(file)
    counts := make(map[string]int)

    //iterate file one line at a time
    for {
        if line, prefix, err := r.ReadLine(); err != nil {
            break
        } else if !prefix {
            s := strings.Fields(string(line))
            for _ , v := range s {
                counts[v] = counts[v] + 1
            }
        }
        if err == io.EOF {
            err = nil
        }
    }


    //sort keys to output words deterministically 
    keys := make([]string, len(counts))
    i := 0
    for k, _ := range counts {
        keys[i] = k
        i++
    }
    sort.Strings(keys)

    //display output of word counts
    for _, v := range keys{
        fmt.Println(v,counts[v])
    }
}
