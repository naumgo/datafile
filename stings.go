//пакет предназначен для чтения данных из файла

package datafile

import (
	"bufio"
	"os"
)

//GetStrings читает каждую строку из файла

func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		line := scaner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scaner.Err() != nil {
		return nil, scaner.Err()
	}
	return lines, nil
}
