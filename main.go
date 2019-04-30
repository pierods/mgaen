package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var helpFlag = flag.Bool("help", false, "")

func main() {
	flag.Parse()
	if *helpFlag {
		fmt.Println("Usage: mgaen clearfile encryptedfile")
		return
	}
	inFile := ""
	outFile := ""
	switch len(os.Args) {
	case 1:
		fmt.Print("Enter clear file name/path: ")
		consoleReader := bufio.NewReader(os.Stdin)
		inFile, _ = consoleReader.ReadString('\n')
		inFile = strings.Replace(inFile, "\n", "", -1)

		fmt.Print("Enter encrypted file name/path: ")
		outFile, _ = consoleReader.ReadString('\n')
		outFile = strings.Replace(outFile, "\n", "", -1)

	case 2:
		inFile = os.Args[1]
		fmt.Println("Using clear file " + inFile)

		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter encrypted file name/path: ")
		outFile, _ = consoleReader.ReadString('\n')
		outFile = strings.Replace(outFile, "\n", "", -1)

	case 3:
		inFile = os.Args[1]
		fmt.Println("Using clear file " + inFile)

		outFile := os.Args[2]
		fmt.Println("Creating encrypted file " + outFile)
	}

}
