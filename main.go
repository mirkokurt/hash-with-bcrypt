package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Sample program that take user and password from command line and create a file with the hashed password using bcrypt. Ideally the output of the hashed output has to be written in a database.
// Then just for testing purpose the program read the file and verify the password. The program itself is useless, it's just a basic showcase of password hashing and verifying in go using bcrypt
func main() {

	user := flag.String("user", "", "username")
	pass := flag.String("pass", "", "password")
	flag.Parse()

	if len(os.Args) < 2 || *user == "" || *pass == "" {
		usage("Please insert user name and password")
	}

	password := []byte(*pass)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	var line = *user + ":" + string(hashedPassword)

	content := []string{line}
	err = createAndWriteFile("hashedPasswordFile.txt", content)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Retrieve the hashed password from file
	psw, err := readFileIfExists("hashedPasswordFile.txt", *user)
	if err != nil {
		fmt.Println(err)
	}

	b := []byte(psw)
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(b, password)
	if err != nil {
		fmt.Println(err)
	}
}

func createAndWriteFile(path string, content []string) error {
	// Open or create the output file
	fo, err := os.Create(path)
	if err != nil {
		return err
	}
	// Close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			return
		}
	}()

	// Create a writer
	writer := bufio.NewWriter(fo)
	defer writer.Flush()

	for _, elem := range content {
		writer.WriteString(elem + "\n")
	}

	return nil
}

func readFileIfExists(path string, user string) (string, error) {

	// Open the input file
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		// Take a line from the file
		line := scanner.Text()
		values := strings.Split(line, ":")
		if values[0] == user {
			return values[1], nil
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return "", err
	}

	return "not_found", nil

}

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command> <password>\n"+
			"       where <username> is the username to be created (not encrypted) \n"+
			"       and <password> is the password you want to encrypt\n",
		errmsg, os.Args[0])
	os.Exit(2)
}
