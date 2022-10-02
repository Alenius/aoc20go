package lib

import (
	"bufio"
	"fmt"
	"os"
)

// fn takes in the line and does something with it
func ReadFileByLine(day string, fn func(string)) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Open(fmt.Sprintf(`%s/%s/%s.txt`, dir, day, day))
	defer file.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}

	return nil
}
