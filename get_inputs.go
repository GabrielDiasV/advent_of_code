package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func makeRequest(day, year int) string {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	session := os.Getenv("session")

	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error creating cookie jar: ", err)
		os.Exit(1)
	}

	client := &http.Client{
		Jar: jar,
	}

	urlString := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		os.Exit(1)
	}

	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		os.Exit(1)
	}
	jar.SetCookies(u, []*http.Cookie{cookie})

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		os.Exit(1)
	}

	return string(body)
}

func saveFile(day, year int, inputString string) error {

	dir := fmt.Sprintf("%d/input_files", year)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	filePath := fmt.Sprintf("%s/input_day%d.txt", dir, day)

	if err := os.WriteFile(filePath, []byte(inputString), 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func main() {
	var year, day int
	var option int
	var inputString string

	fmt.Print("Enter the year of reference to import the inputs of the challanges: ")
	fmt.Scanln(&year)
	fmt.Print(
		"Do you want to import all the inputs or the inputs from a specific date (Please select the number associated with the option)?",
		"\n[1] All the Inputs\n[2] Inputs from a specific day\n",
		"Option Selected: ",
	)
	fmt.Scanln(&option)

	if option == 1 {
		for day := 1; day <= 25; day++ {
			fmt.Println("Processing input from day ", day)
			inputString = makeRequest(day, year)
			saveFile(day, year, inputString)
		}
	}

	if option == 2 {
		fmt.Print("Insert the day: ")
		fmt.Scanln(&day)
		fmt.Println("Processing input from day ", day)
		inputString = makeRequest(day, year)
		saveFile(day, year, inputString)
	}
}
