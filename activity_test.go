package databasequery

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"io/ioutil"
	"testing"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	
	}
	return activityMetadata
}

func TestCreate(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

//debugging

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	
	//setup attr for mysql
	fmt.Println("===============================")
	fmt.Println("Unit Test ===> MySQL Connection")
	fmt.Println("===============================")
	fmt.Println("")
	tc.SetInput("driverName", "mysql")
	tc.SetInput("datasourceName", "sushant:kumbhar.6242@tcp(cvadb.cpvs8qpluubi.us-east-1.rds.amazonaws.com:3306)/CVAMysqlDB")
	//tc.SetInput("query", "select * from user_details where mobile='4' ")
	//tc.SetInput("query", "insert into user_details (mobile) values (8) ")
	tc.SetInput("query", "update user_details set mobile='55', password='test' where mobile='5' ")
	fmt.Println("==========Set Input Done=====================")


/* 	
	//setup attr for sqlite3
	fmt.Println("===============================")
	fmt.Println("Unit Test ===> SQLITE3 Connection")
	fmt.Println("===============================")
	fmt.Println("")
	tc.SetInput("driverName", "sqlite3")
	tc.SetInput("datasourceName", "akashdb")
	tc.SetInput("query", "select * from person")
*/

/* 
	//setup attr for postgres
	fmt.Println("===============================")
	fmt.Println("Unit Test ===> POSTGRES Connection")
	fmt.Println("===============================")
	fmt.Println("")
	tc.SetInput("driverName", "postgres")
	tc.SetInput("datasourceName", "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable")
	tc.SetInput("query", "select * from company")
 */
 fmt.Println("=============Eval Started==================")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Println("result: ", result)

	if result == nil {
		t.Fail()
	}

}
