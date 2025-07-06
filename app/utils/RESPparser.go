package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseRESP(input string) (interface{}, error) {
	reader := bufio.NewReader(strings.NewReader(input))
	return parse(reader)
}

func parse(reader *bufio.Reader) (interface{}, error) {
	prefix, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch prefix {
	case '*':
		line, _ := reader.ReadString('\n')
		count, _ := strconv.Atoi(strings.TrimSpace(line))
		if count == -1 {
			return nil, nil
		}
		items := make([]interface{}, 0, count)
		for i := 0; i < count; i++ {
			item, err := parse(reader)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
		return items, nil

	case '$':
		line, _ := reader.ReadString('\n')
		length, _ := strconv.Atoi(strings.TrimSpace(line))
		if length == -1 {
			return nil, nil
		}
		buf := make([]byte, length+2)
		_, err := reader.Read(buf)
		if err != nil {
			return nil, err
		}
		return string(buf[:length]), nil

	case '+':
		line, _ := reader.ReadString('\n')
		return strings.TrimSpace(line), nil

	case ':': 
		line, _ := reader.ReadString('\n')
		return strconv.Atoi(strings.TrimSpace(line))

	case '-': 
		line, _ := reader.ReadString('\n')
		return nil, fmt.Errorf("Redis error: %s", strings.TrimSpace(line))

	default:
		return nil, fmt.Errorf("Unknown prefix: %c", prefix)
	}
}
