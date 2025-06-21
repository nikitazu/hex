package main

import (
    "flag"
    "fmt"
    "io"
    "os"
)

type CliOptions struct {
    Version   bool
    LineSize  int
    BlockSize int
    Args      []string
}

func ParseCliOptions() *CliOptions {
    opts := &CliOptions{}
    flag.Usage = PrintHelp
    flag.BoolVar(&opts.Version, "v", false, "Prints the app version (shorthand).")
    flag.BoolVar(&opts.Version, "version", false, "Prints the app version.")
    flag.IntVar(&opts.LineSize, "l", 32, "Number of bytes on one line (shorthand).")
    flag.IntVar(&opts.LineSize, "line-size", 32, "Number of bytes on one line.")
    flag.IntVar(&opts.BlockSize, "b", 4, "Number of bytes in one block (shorthand).")
    flag.IntVar(&opts.BlockSize, "block-size", 4, "Number of bytes in one block.")
    flag.Parse()
    opts.Args = flag.Args()
    return opts
}

func PrintHelp() {
    fmt.Printf("This software reads file names from the command line\n")
    fmt.Printf("and prints their byte content in hexadecimal format.\n\n")
    fmt.Printf("USAGE: hex file.png\n")
    fmt.Printf("       hex file1.png file2.txt file3.jpg\n")
    fmt.Printf("\n")
    fmt.Printf("OPTIONS\n\n")
    fmt.Printf("  -h    Prints the help.\n")
    fmt.Printf("  -help\n        Prints the help.\n")
    fmt.Printf("\n")
    flag.PrintDefaults()
}

func main() {
    opts := ParseCliOptions()

    if opts.Version {
        fmt.Fprintf(os.Stderr, "hex 0.1.0\n")
    }

    if opts.LineSize < 16 || opts.LineSize > 256 {
        fmt.Fprintf(os.Stderr, "Error: option `line-size` must be within range (16, 256).\n\n")
        os.Exit(1)
    }

    if opts.BlockSize < 0 || opts.BlockSize > opts.LineSize {
        fmt.Fprintf(os.Stderr, "Error: option `block-size` must be within range (0, line-size).\n\n")
        os.Exit(1)
    }

    if len(opts.Args) == 0 {
        fmt.Fprintf(os.Stderr, "Error: no files specified.\n\n")
        os.Exit(1)
    }

    countProcessedFiles := 0
    countFailedFiles := 0

    for i := 0; i < len(opts.Args); i++ {
        path := opts.Args[i]
        processingErrorHappened := false

        if i > 0 {
            fmt.Println()
        }

        fmt.Printf("FILE: %v\n", path)

        file, err := os.Open(path)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error: failed opening file\n  %v\n", err)
            processingErrorHappened = true
            countFailedFiles++
        }

        if processingErrorHappened {
            continue
        }

        defer file.Close()

        bufferSize := opts.LineSize
        buffer := make([]byte, bufferSize)

        for {
            bytesRead, err := file.Read(buffer)
            if err != nil {
                if err == io.EOF {
                    processingErrorHappened = true
                    countFailedFiles++
                    break
                }
                fmt.Fprintf(os.Stderr, "Error: failed reading file\n  %v\n", err)
                processingErrorHappened = true
                countFailedFiles++
                break
            }

            for filePos := 0; filePos < bytesRead; filePos++ {
                byteValue := buffer[filePos]
                fmt.Printf("%02X", byteValue)
                if opts.BlockSize > 0 && (filePos%opts.BlockSize) == (opts.BlockSize-1) {
                    fmt.Printf(" ")
                }
            }
            fmt.Println()
        }

        if processingErrorHappened {
            continue
        }

        countProcessedFiles++
    }

    if countFailedFiles > 0 {
        os.Exit(1)
    }
}
