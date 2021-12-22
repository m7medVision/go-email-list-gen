package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println(`

	███▄ ▄███▓ ▄▄▄      ▄▄▄██▀▀▀██░ ██  ▄████▄   ▄████▄  
	▓██▒▀█▀ ██▒▒████▄      ▒██  ▓██░ ██▒▒██▀ ▀█  ▒██▀ ▀█  
	▓██    ▓██░▒██  ▀█▄    ░██  ▒██▀▀██░▒▓█    ▄ ▒▓█    ▄ 
	▒██    ▒██ ░██▄▄▄▄██▓██▄██▓ ░▓█ ░██ ▒▓▓▄ ▄██▒▒▓▓▄ ▄██▒
	▒██▒   ░██▒ ▓█   ▓██▒▓███▒  ░▓█▒░██▓▒ ▓███▀ ░▒ ▓███▀ ░
	░ ▒░   ░  ░ ▒▒   ▓▒█░▒▓▒▒░   ▒ ░░▒░▒░ ░▒ ▒  ░░ ░▒ ▒  ░
	░  ░      ░  ▒   ▒▒ ░▒ ░▒░   ▒ ░▒░ ░  ░  ▒     ░  ▒   
	░      ░     ░   ▒   ░ ░ ░   ░  ░░ ░░        ░        
		   ░         ░  ░░   ░   ░  ░  ░░ ░      ░ ░         
	`)
	fmt.Printf("\n Number of Emails do you: ")
	var numbers int
	fmt.Scanf("%d", &numbers)
	file, _ := os.Open("names.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	list := []string{}
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	listofdomains := []string{}
	listofdomains = append(listofdomains, "gmail.com")
	listofdomains = append(listofdomains, "yahoo.com")
	listofdomains = append(listofdomains, "hotmail.com")
	listofdomains = append(listofdomains, "aol.com")
	listofdomains = append(listofdomains, "outlook.com")
	listofdomains = append(listofdomains, "live.com")
	listofdomains = append(listofdomains, "msn.com")
	listofdomains = append(listofdomains, "outlook.sa")
	files, _ := os.OpenFile("emails.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for i := 0; i < numbers; i++ {
		ss := listofdomains[rand.Intn(len(listofdomains))]
		name1 := list[rand.Intn(len(list))]
		name2 := list[rand.Intn(len(list))]
		numberq := rand.Int() % 1000
		fullemail := ((strings.ToLower(name1) + strings.ToLower(name2) + fmt.Sprintf("%d", numberq)) + "@" + ss)
		files.WriteString(fullemail + "\n")
		fmt.Print("\rCount: ", i)
	}
}
