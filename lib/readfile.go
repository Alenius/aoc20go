package lib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

// returns map of named subgroups in a regex on format subgroupname to regexpmatch
func FindNamedSubgroupInString(s string, re *regexp.Regexp, names []string) (map[string]string, error) {
	matches := map[string]string{}
	for _, name := range names {
		ix := re.SubexpIndex(name)
		if ix < 0 {
			return nil, fmt.Errorf("could not find any subgroup in regexp with name: %v", name)
		}
		match := re.FindStringSubmatch(s)
		if match == nil {
			return nil, fmt.Errorf("no match for in regexp: %v", name)
		}

		subexpix := re.SubexpIndex(name)
		if subexpix < 0 {
			return nil, fmt.Errorf("no match on subgroup with name: %v", name)
		}

		submatch := match[subexpix]
		matches[name] = submatch
	}

	return matches, nil
}
