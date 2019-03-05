package todo

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// List todo file
func List() error {
	fs, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}
	for _, f := range fs {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".td" {
			fmt.Println(f.Name())
		}
	}
	return nil
}

// New todo file
func New(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println(name, "Created")
	return nil
}

// Show todo list of todo name file
func Show(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println("List of", name)
	scanner := bufio.NewScanner(f)
	for i := 1; scanner.Scan(); i++ {
		fmt.Printf("%d) %s\n", i, scanner.Text())
	}
	return scanner.Err()
}

// Add item to todo name file
func Add(name string, item string) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintln(f, item)
	fmt.Printf("%s <= %s\n", name, item)
	return nil
}

// Delete item number n from todo name file
func Delete(name string, n string) error {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	lines := bytes.Split(b, []byte("\n"))
	itemNO, err := strconv.Atoi(n)
	if err != nil {
		return err
	}

	// slice itemNO (index itemNO-1) out of slice lines
	lines = append(lines[:itemNO-1], lines[itemNO:]...)

	err = ioutil.WriteFile(name, bytes.Join(lines, []byte("\n")), 0)
	if err != nil {
		return err
	}
	fmt.Printf("%s <= %s) Deleted\n", name, name)
	return nil
}

// Edit item number n of todo name file with newItem
func Edit(name string, n string, newItem string) error {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	lines := bytes.Split(b, []byte("\n"))
	itemNO, err := strconv.Atoi(n)
	if err != nil {
		return err
	}
	lines[itemNO-1] = []byte(newItem)

	err = ioutil.WriteFile(name, bytes.Join(lines, []byte("\n")), 0)
	if err != nil {
		return err
	}
	fmt.Printf("%s <= %s) %s\n", name, n, newItem)
	return nil
}

//Print pyramid
func Pyramid(input string) error {
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("It's not a number")
	} else {
		for i := 1; i <= num; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

	return nil
}
