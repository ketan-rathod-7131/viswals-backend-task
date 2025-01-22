package redis

import "fmt"

func GetKey(prefix string, key interface{}) string {
	return fmt.Sprintf("%v:%v", prefix, key)
}
