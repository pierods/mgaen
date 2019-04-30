package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/pierods/mgaen/encrypt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"syscall"
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
		consoleReader := bufio.NewScanner(os.Stdin)
		consoleReader.Scan()
		inFile = consoleReader.Text()

		fmt.Print("Enter encrypted file name/path: ")
		consoleReader.Scan()
		outFile = consoleReader.Text()

	case 2:
		inFile = os.Args[1]
		fmt.Println("Using clear file " + inFile)

		consoleReader := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter encrypted file name/path: ")
		consoleReader.Scan()
		outFile = consoleReader.Text()

	case 3:
		inFile = os.Args[1]
		fmt.Println("Using clear file " + inFile)
		outFile = os.Args[2]
		fmt.Println("Creating encrypted file " + outFile)
	}

	password := getPassword()

	inData, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	encryptedData, err := encrypt.Seal(inData, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	err = ioutil.WriteFile(outFile, encryptedData, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if len(password) == 0 {
		fmt.Println("Empty password. Exiting.")
		os.Exit(2)
	}
	fmt.Print("\nConfirm password: ")
	passwordConfirm, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if bytes.Compare(password, passwordConfirm) != 0 {
		fmt.Println("\nPasswords don't match. Exiting")
		os.Exit(3)
	}
	fmt.Println()
	return password
}
