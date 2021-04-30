
##### Version 0.1.0
**The package helps you to convert ssh config files to ansible inventories.**

input and output support for string and bytearray is present.

###### func StringToString ( input string )  ( string , error )
###### func StringToBuff ( input string ) ( []byte, error )
###### func BuffToString ( input []byte ) (string, error )
###### func BuffToBuff ( input []byte ) ( []byte , error )


Currently, these properties of ssh config files are supported 
**1. Hostname**
**2. Host**
**3. User**


Example1 
```go
package main 

import (
	"fmt"
	"ans-host-conv/converter"
)
func main() {
	input := `Host testing-server1 
	HostName  XXXXXXXX
	User test
	Host testing-server2 
	HostName  XXXXXXXX
	User test`
	
	ac, err := converter.StringToString(input)
	// Takes string  as input and gives string as output 
	
	if err != nil {
			fmt.Println(err)
		}
	fmt.Println(ac)
//	 Output is below 
//     [testing-server1]
//     XXXXXXXX    ansible_connection=ssh    ansible_user=test
//     [testing-server2]
//     XXXXXXXX    ansible_connection=ssh    ansible_user=test
}
```

** If any one of these entities are not available (Host/HostName/User) the whole ssh config block will be skipped. **



