package hashutil

import (
	"hash/fnv"
	"encoding/json"
	"fmt"
)

func GetHashInt(obj ...interface{}) uint32 {
	if obj == nil {
		return 0
	}
	contents, _ := json.Marshal(obj)
	h := fnv.New32a()
	h.Write(contents)
	return h.Sum32()
}

func GetHashString(obj ...interface{}) string{
	if obj == nil {
		return ""
	}
	return fmt.Sprintf("%x", GetHashInt(obj))
}