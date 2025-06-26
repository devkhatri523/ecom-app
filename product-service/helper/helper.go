package helper

import (
	"strconv"
	"strings"
)

func BuildInCluaseQuery(ids []int32) string {
	strs := make([]string, len(ids))
	for i, id := range ids {
		strs[i] = strconv.Itoa(int(id))
	}
	return strings.Join(strs, ",")
}
