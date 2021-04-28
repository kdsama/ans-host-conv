package converter

import (
	"errors"
	"fmt"
	stringutils "github.com/kdsama/ans-host-conv/utils"
	"strings"
)

var (
	ErrINCOMPLETE_DATA    = errors.New("incomplete data . expects hostName,host")
	ErrHOST_INVALID       = errors.New("host is invalid")
	ErrHOSTNAME_NOT_FOUND = errors.New("hostName not found")
	ErrUSER_NOT_FOUND     = errors.New("user is not found")
)

var (
	BLOCK_SIZE     = 3
	HOSTNAME       = "hostname"
	USER           = "user"
	VARIABLE_ARRAY = []string{HOSTNAME, USER}
)

type AnsibleBlock struct {
	Input    string
	HostName string
	Host     string
	User     string
}

func (a AnsibleBlock) GetString() string {
	return fmt.Sprintf("[%s]\n%s    ansible_connection=ssh    ansible_user=%s", a.Host, a.HostName, a.User)
}

func NewAnsibleBlock(input string) (AnsibleBlock, error) {

	ac := SplitBlockIntoArray(input)

	return convert(ac, input)
}

func convert(inputArray []string, input string) (AnsibleBlock, error) {
	var host string
	var hostName string
	var user string
	if len(inputArray) < BLOCK_SIZE {
		return AnsibleBlock{}, ErrINCOMPLETE_DATA
	}
	firstArray := strings.Fields(inputArray[0])
	
	if len(firstArray) > 1 {
		return AnsibleBlock{}, ErrHOST_INVALID
	} else {
		if len(inputArray[0]) == 0{
			return AnsibleBlock{}, ErrHOST_INVALID
		}
		host = inputArray[0]
	}

	secondArray := strings.Fields(inputArray[1])
	if len(secondArray) !=2 {
		return AnsibleBlock{},ErrINCOMPLETE_DATA
	}
	if stringutils.GetIndexOfStringElement(VARIABLE_ARRAY, secondArray[0]) == -1 {
		return AnsibleBlock{}, ErrINCOMPLETE_DATA
	} else {

		if strings.ToLower(secondArray[0]) == HOSTNAME {
			hostName = secondArray[1]
		} else if strings.ToLower(secondArray[0]) == USER {
			user = secondArray[1]
		}
	}

	thirdArray := strings.Fields(inputArray[2])
	if len(thirdArray) !=2 {
		return AnsibleBlock{},ErrINCOMPLETE_DATA
	}
	if stringutils.GetIndexOfStringElement(VARIABLE_ARRAY, thirdArray[0]) == -1 {
		return AnsibleBlock{}, ErrINCOMPLETE_DATA
	} else {

		if strings.ToLower(thirdArray[0]) == HOSTNAME {
			hostName = thirdArray[1]
		} else if strings.ToLower(thirdArray[0]) == USER {
			user = thirdArray[1]
		}
	}
	if hostName == "" {
		return AnsibleBlock{}, ErrHOSTNAME_NOT_FOUND
	}
	if user == "" {
		return AnsibleBlock{}, ErrUSER_NOT_FOUND
	}
	
	return AnsibleBlock{Input: input, Host: host, HostName: hostName, User: user}, nil
}

func (a *AnsibleBlock) ReadInput() string {
	return a.Input
}

func SplitBlockIntoArray(input string) []string {
	toReturn := strings.Split(input, "\n")
	index := 0
	for index < len(toReturn) {
		if len(toReturn[index]) == 0 {
			toReturn = stringutils.RemoveIndexFromStringArray(toReturn, index)
			continue
		}
		toReturn[index] = strings.TrimSpace(toReturn[index])
		index++
	}
	return toReturn
}

