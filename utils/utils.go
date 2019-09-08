package utils

import (
	"net"
)

// 返回map中key对应的value，key不存在返回默认值
func GetMapValue(m map[string]string, key, defaultValue string) string {
	if value, ok := m[key]; ok {
		return value
	}
	return defaultValue
}

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}
	return "", nil
}