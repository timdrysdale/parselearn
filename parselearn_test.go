package parselearn

import (
	"testing"
)

// Example receipt (anonymised)

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

func TestSubmissionField(t *testing.T) {

	sub := Submission{}
	processSubmission("There is no student submission text data for this assignment.", &sub)
	assertEqual(t, sub.SubmissionField, "There is no student submission text data for this assignment.")
}

func TestComments(t *testing.T) {

	sub := Submission{}
	processSubmission("There are no student comments for this assignment", &sub)
	assertEqual(t, sub.SubmissionField, "There are no student comments for this assignment")
}

func TestOriginalFilename(t *testing.T) {

	sub := Submission{}
	processOriginalFilename("Original filename: OnlineExam-Bxxxxxx.pdf", &sub)
	assertEqual(t, sub.OriginalFilename, "OnlineExam-Bxxxxxx.pdf")
}

func TestFilename(t *testing.T) {

	sub := Submission{}
	processFilename("Filename: Practice Exam Drop Box_sxxxxxxx_attempt_yyyy-mm-dd-hh-mm-ss_OnlineExam-Bxxxxxx.pdf", &sub)
	assertEqual(t, sub.Filename, "Practice Exam Drop Box_sxxxxxxx_attempt_yyyy-mm-dd-hh-mm-ss_OnlineExam-Bxxxxxx.pdf")
}

func TestParseFile(t *testing.T) {

	sub, err := parseLearnReceipt("./test/receipt2.txt")
	if err != nil {
		t.Error(err)
	}

	assertEqual(t, sub.FirstName, "John")
	assertEqual(t, sub.LastName, "Smith")
	assertEqual(t, sub.Matriculation, "s00000000")
	assertEqual(t, sub.Assignment, "Some Exam Or Other")
	assertEqual(t, sub.DateSubmitted, "Tuesday, dd April yyyy hh:mm:ss o'clock BST")
	assertEqual(t, sub.SubmissionField, "There is no student submission text data for this assignment.")
	assertEqual(t, sub.Comments, "There are no student comments for this assignment.")
	assertEqual(t, sub.OriginalFilename, "ENGI1234-Bxxxxxx.pdf")
	assertEqual(t, sub.Filename, "Practice Exam Drop Box_sxxxxxxx_attempt_yyyy-mm-dd-hh-mm-ss_ENGI1234-Bxxxxxx.pdf")

}
