package golib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ParseJson(f string, v any) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(v); err != nil {
		return fmt.Errorf("parse %s: %s", f, err.Error())
	}
	return nil
}

func ParseIni(f string) (map[string]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	kvs := map[string]string{}
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line[0] == '#' {
			continue
		}
		a := strings.SplitN(line, "=", 2)
		if len(a) < 2 {
			return nil, fmt.Errorf("parse %s:%d: unparsable", f, lineNum)
		}
		k := strings.TrimSpace(a[0])
		if k == "" {
			return nil, fmt.Errorf("parse %s:%d: empty key", f, lineNum)
		}
		kvs[k] = strings.TrimSpace(a[1])
	}
	return kvs, nil
}
