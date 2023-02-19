package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("first argument must be a valid integer"))
	}

	aocSession, ok := os.LookupEnv("AOC_SESSION")
	if !ok {
		panic(fmt.Errorf("env var AOC_SESSION is not set"))
	}
	err = downloadInput(day, aocSession)
	if err != nil {
		panic(err)
	}
}

func downloadInput(day int, sessionCookie string) error {

	os.MkdirAll("inputs", os.ModePerm)
	flags := os.O_WRONLY | os.O_CREATE
	file, err := os.OpenFile(fmt.Sprintf("inputs/day%d", day), flags, 0666)
	defer file.Close()

	if os.IsExist(err) {
		return fmt.Errorf("file '%s' already exists", fmt.Sprintf("/inputs/day%d", day))
	} else if err != nil {
		return err
	}

	client := new(http.Client)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day), nil)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", sessionCookie
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
