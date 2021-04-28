package converter
import (
	"testing"
)


func TestAnsibleConverter(t *testing.T) {

	t.Run("Get AnsibleFormattedLine from InputBlock",TestAnsibleOutputBlocks)
	t.Run("Get AnsibleFormattedLine from InputBlock Errors",TestAnsibleOutputBlockErrors)



}
func TestAnsibleOutputBlockErrors(t *testing.T){

	testcases := []struct{
		name string
		expected error
		input string
	}{
		{
			name:"Without Hosts",
			expected:ErrHOST_INVALID,
			input:` 
			HostName  XXXXXXXX
			User test`,
		},
		{
			name:"Without HostName",
			expected:ErrINCOMPLETE_DATA,
			input:`search-ui 
			Hostname  
			User test`,
		},
		{
			name:"Without Name",
			expected:ErrINCOMPLETE_DATA,
			input:`search-ui 
			HostName  XXXXXXXX
			user `,
		},
		{
			name:"Lesser Rows",
			expected:ErrINCOMPLETE_DATA,
			input:`HostName  XXXXXXXX
			User test`,
		},
	}

	for _,testcase := range testcases {
		t.Run(testcase.name,func (t *testing.T){
			_,err := NewAnsibleBlock(testcase.input)
			assertErrors(t,err,testcase.expected)
			// t.Error(field)
		})

	}

}



func TestAnsibleOutputBlocks(t *testing.T){
	input :=  `search-ui 
	HostName  XXXXXXXX
	User test`
	ac,_ := NewAnsibleBlock(input)
	want := "[search-ui]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test"
	got := ac.GetString()
	assertInput(t,got,want)
}

func assertErrors(t testing.TB,got error,want error){
	if got != want {
		t.Errorf("Expected Error %v   , but got %v\n",want,got)
	}
}


func assertInput(t testing.TB,got string , want string ){
	t.Helper()
	if got != want {
		t.Errorf("Got\n %s but wanted\n %s",got ,want)

	}
}