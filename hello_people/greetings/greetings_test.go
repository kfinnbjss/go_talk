package greetings

import "testing"

type greetingTest struct {
	testName         string
	name             string
	language         string
	expectedGreeting string
	wantErr          bool
}

func Test_Greet(t *testing.T) {
	tests := []greetingTest{
		{
			testName:         "English",
			name:             "john",
			language:         "english",
			expectedGreeting: "hello john!",
			wantErr:          false,
		},
		{
			testName:         "French",
			name:             "john",
			language:         "french",
			expectedGreeting: "bonjour john!",
			wantErr:          false,
		},
		{
			testName:         "Spanish",
			name:             "john",
			language:         "spanish",
			expectedGreeting: "hola john!",
			wantErr:          false,
		},
		{
			testName:         "German",
			name:             "john",
			language:         "german",
			expectedGreeting: "guten tag john!",
			wantErr:          false,
		},
		{
			testName:         "Hawaiian",
			name:             "john",
			language:         "hawaiian",
			expectedGreeting: "aloha e john!",
			wantErr:          false,
		},
		{
			testName:         "Swahili",
			name:             "john",
			language:         "swahili",
			expectedGreeting: "",
			wantErr:          true,
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			actualGreeting, err := Greet(test.name, test.language)
			gotErr := err != nil

			if test.wantErr && !gotErr {
				tt.Error("expected an error but got none")
			}
			if !test.wantErr && gotErr {
				tt.Errorf("got unexpected error %v", err)
			}
			if actualGreeting != test.expectedGreeting {
				tt.Errorf("wrong output. Expected %s, got %s", test.expectedGreeting, actualGreeting)
			}
		},
		)
	}

}
