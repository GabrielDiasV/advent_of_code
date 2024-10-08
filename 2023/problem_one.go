package main

import (
	"fmt"
	"os"
	"time"
	"errors"
	"net/http"
)

func make_request(){
	resp, err := http.Get("https://adventofcode.com/2023/day/1/input")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
}

func main(){
	
}