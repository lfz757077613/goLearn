package utils

import (
	"bufio"
	"io"
	"net"
	"os"
	"strings"
)

// 返回map中key对应的value，key不存在返回默认值
func GetMapValue(m map[string]string, key, defaultValue string) string {
	if value, ok := m[key]; ok {
		return value
	}
	return defaultValue
}

func GetLocalIp() string {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return ""
}

func FileLineProcess(path string, f func(line string) error) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		line = strings.Trim(line, "\n")
		if err == io.EOF {
			if line != "" {
				return f(line)
			}
		}
		if err := f(line); err != nil {
			return nil
		}
	}
}
