package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ParseRESP(input string) (interface{}, error) {
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
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		count, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}
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
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}
		if length == -1 {
			return nil, nil
		}
		buf := make([]byte, length+2)
		_, err = reader.Read(buf)
		if err != nil {
			return nil, err
		}
		return string(buf[:length]), nil

	case '+':
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return strings.TrimSpace(line), nil

	case ':':
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return strconv.Atoi(strings.TrimSpace(line))

	case '-':
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Redis error: %s", strings.TrimSpace(line))

	default:
		return nil, fmt.Errorf("Unknown prefix: %c", prefix)
	}
}
