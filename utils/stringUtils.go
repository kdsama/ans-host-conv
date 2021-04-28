package utils

import ( 
	"strings"
)
func RemoveIndexFromStringArray(s []string, index int) []string {
	s = append(s[:index], s[index+1:]...)
	return s
}
func GetIndexOfStringElement(s []string, val string) int {
	for k, v := range s {
		if strings.EqualFold(val,v){
			return k
		}
	}
	return -1
}

func ConvertStringToBuffer(s string) []byte {
	return []byte(s)
}