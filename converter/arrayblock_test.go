package converter 

import (
	"testing"
)


func TestBuffToBuff(t *testing.T){
	input:=`Host search-ui 
	HostName  XXXXXXXX
	User test
	Host search-server 
	HostName  XXXXXXXX
	User test
	`
	want:="[search-ui]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n[search-server]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n"
	// want := []byte(input)
	got,_ := BuffToBuff([]byte(input))
	if (string(got[:]) != want){
		
			t.Errorf("Expected %v , but got %v",want,string(got[:]))
		
	}

}



func TestBuffToString(t *testing.T){
	input:=`Host search-ui 
	HostName  XXXXXXXX
	User test
	Host search-server 
	HostName  XXXXXXXX
	User test
	`
	want:="[search-ui]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n[search-server]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n"
	// want := []byte(input)
	got,_ := BuffToString([]byte(input))
	if (got!= want){
			t.Errorf("Expected %v , but got %v",want,string(got[:]))
	}

}

func TestStringToBuff(t *testing.T){
	input:=`Host search-ui 
	HostName  XXXXXXXX
	User test
	Host search-server 
	HostName  XXXXXXXX
	User test
	`
	want:="[search-ui]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n[search-server]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n"
	// want := []byte(input)
	got,_ := StringToBuff(input)
	if (string(got[:]) != want){
		
			t.Errorf("Expected %v , but got %v",want,string(got[:]))
		
	}

}


func TestStringtoString(t *testing.T){
	t.Run("Get FormattedBlock from InputString",TestAnsibleInputString)
	t.Run("Get Errors  from InputString",TestAnsibleInputForError)
}

func TestAnsibleInputForError(t *testing.T){

	testcases := []struct{
		name string
		input string
		want error
	}{
		{
			name:"Should Empty Input Array",
			input:`asdasd
			asd
			as
			d
			as
			d
			a
			sd`,
			want:errINVALID_INPUT,
	},
		}
		for _,obj := range testcases{
			t.Run(obj.name,func (t *testing.T){
				_,got := StringToString(obj.input)
				assertErrorInputString(t,got,obj.want)
			})
	}

}


func TestAnsibleInputString(t *testing.T){

	testcases := []struct{
		name string
		input string
		want string
	}{
		{name:"All Blocks to be Returned",
		input:`Host search-ui 
		HostName  XXXXXXXX
		User test
		Host search-server 
		HostName  XXXXXXXX
		User test
		`,
		want:"[search-ui]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n[search-server]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n",
		},{
				name:"Second block has missing HostName , Only First Block will appear",
				input:`Host search-ui 
				HostName  XXXXXXXX
				User test
				Host search-server 
				XXXXXXXX
				User test`,
				want:"[search-ui]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n",
		},{
				name:"First Block has missing User , Only Second  Block will appear",
				input:`Host search-ui 
				HostName  XXXXXXXX
				test
				Host search-server 
				HostName  XXXXXXXX
				User test`,
				want:"[search-server]\nXXXXXXXX    ansible_connection=ssh    ansible_user=test\n",
		},
		{
			name:"Should Return Empty String",
			input:`asdasd
			asd
			as
			d
			as
			d
			a
			sd`,
			want:"",
	},
		}
		for _,obj := range testcases{
			t.Run(obj.name,func (t *testing.T){
				got,_ := StringToString(obj.input)
				assertValidInputString(t,obj.input,got,obj.want)
			})
	}

}
func assertErrorInputString(t *testing.T,got error, want error){

	if got != want {
		t.Errorf("Expected ERROR ::::: %v , but got  ERROR ::::: %v",want,got)
	}


}

func assertValidInputString(t *testing.T, input string ,got string, want string){

	if got != want {
		t.Errorf("For input\n,\n%s\n\nExpected \n\n %s , but got \n\n %s",input,want,got)
	}


}