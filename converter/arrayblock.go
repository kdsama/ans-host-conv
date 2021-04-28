package converter 

// import "bufio"
import (
	"strings"
	// "fmt"
	"errors"
	stringutils "github.com/kdsama/ans-host-conv/utils"
)

type ArrayBlock struct {
	stringArray []string
}

var splitString = "Host "

var (
	errINVALID_INPUT = errors.New("invalid input")
)

func newArrayBlock(input []string)ArrayBlock{
	outputArray := []string{}
	for _,value := range input {
		if value != ""{
			outputArray = append(outputArray,value)
		}
	}
	return ArrayBlock{outputArray}
}


func (ab *ArrayBlock) convertToHost() (string,error){
		toReturn := ""
		for _,value := range ab.stringArray{
			resp,err := NewAnsibleBlock(value)
			if err != nil {
				// fmt.Println(err)
				continue;
			}
			toReturn += resp.GetString() + "\n"
		}
		if len(toReturn) == 0 {
			return "",errINVALID_INPUT
		}
		return toReturn,nil

}
func StringToString(input string) (string,error){

	inputArray := strings.Split(input,splitString)
	// toReturnString := ""

	ab := newArrayBlock(inputArray)
	outputString,err := ab.convertToHost()
	if err != nil {
		return "",err
	}

	return outputString,nil
}


func StringToBuff(input string) ([]byte,error){
	outputString,err := StringToString(input)
	if err != nil {
		return nil, err
	}
	outputByte := stringutils.ConvertStringToBuffer(outputString)
	return outputByte,nil

}



func BuffToString(input []byte) (string,error){
	inputString := string(input[:])
	outputString,err := StringToString(inputString)
	if err != nil {
		return "", err
	}
	
	return outputString,nil

}


func BuffToBuff(input []byte) ([]byte,error){
	inputString := string(input[:])
	outputString,err := StringToString(inputString)
	if err != nil {
		return nil, err
	}
	outputByte := stringutils.ConvertStringToBuffer(outputString)
	return outputByte,nil

}


// Buffer to Buffer 
// Buffer to File 
//  Buffer to String 
// String to String 
// String to Buffer 