package C

import (
	"bufio"
	"os"
	"strings"
)

type fileInfo struct {
	lines []string
}

func Set() fileInfo {
	return fileInfo{}
}

func (f *fileInfo) Load(Filename string) error {
	if strings.HasSuffix(Filename, ".set") {
		data, err := os.ReadFile(Filename)
		if err != nil {
			return err
		}
		newValue := bufio.NewScanner(strings.NewReader(string(data)))
		for newValue.Scan() {
			f.lines = append(f.lines, newValue.Text())
		}
	}
	return nil
}

func (f *fileInfo) GetValue(value string) string {
	for _, i := range f.lines {
		i = strings.Replace(i, " ", "", -1)
		if i != "" {
			line := strings.Split(i, "=")
			if line[0] == value {
				return line[1]
			}
		}
	}
	return ""
}
