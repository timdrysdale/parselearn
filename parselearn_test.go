package parselearn

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestProcessName(t *testing.T) {

	sub := Submission{}
	processName("Name: First Last (sxxxxxxx)", &sub)
	assertEqual(t, sub.FirstName, "First")
	assertEqual(t, sub.LastName, "Last")
	assertEqual(t, sub.Matriculation, "sxxxxxxx")
}

func TestProcessAssignment(t *testing.T) {

	sub := Submission{}
	processAssignment("Assignment: Practice Exam Drop Box", &sub)
	assertEqual(t, sub.Assignment, "Practice Exam Drop Box")
}

func TestDateSubmitted(t *testing.T) {

	sub := Submission{}
	processDateSubmitted("Date Submitted: Monday, dd April yyyy hh:mm:ss o'clock BST", &sub)
	assertEqual(t, sub.DateSubmitted, "Monday, dd April yyyy hh:mm:ss o'clock BST")
}

/*
//Date Submitted: Monday, dd April yyyy hh:mm:ss o'clock BST
func processDateSubmitted(line string, sub *Submission) {
	return errors.New("Not Implemented")
}

//Submission Field:
//There is no student submission text data for this assignment.
func processSubmission(line string, sub *Submission) {

}

//Comments:
//There are no student comments for this assignment
func processComments(line string, sub *Submission) {

}

//Files:
//	Original filename: OnlineExam-Bxxxxxx.pdf
//	Filename: Practice Exam Drop Box_sxxxxxxx_attempt_yyyy-mm-dd-hh-mm-ss_OnlineExam-Bxxxxxx.pdf
func processOriginalFilename(line string, sub *Submission) {

}
func processFilename(line string, sub *Submission) {

}
*/

//Name: First Last (sxxxxxxx)
//Assignment: Practice Exam Drop Box
//Date Submitted: Monday, dd April yyyy hh:mm:ss o'clock BST
//Current Mark: Needs Marking
//
//Submission Field:
//There is no student submission text data for this assignment.
//
//Comments:
//There are no student comments for this assignment.
//
//Files:
//	Original filename: OnlineExam-Bxxxxxx.pdf
//	Filename: Practice Exam Drop Box_sxxxxxxx_attempt_yyyy-mm-dd-hh-mm-ss_OnlineExam-Bxxxxxx.pdf
