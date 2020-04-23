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

/*
//Name: First Last (sxxxxxxx)
func processName(line string, sub *Submission) {
	return errors.New("Not Implemented")
}

//Assignment: Practice Exam Drop Box
func processAssignment(line string, sub *Submission) {
	return errors.New("Not Implemented")
}

//Date Submitted: Monday, dd April yyyy hh:mm:ss o'clock BST
func processDateSubmitted(line string, sub *Submission) {
	return errors.New("Not Implemented")
}

//Submission Field:
//There is no student submission text data for this assignment.
func processSubmission(line string, sub *Submission) {
	return errors.New("Not Implemented")
}

//Comments:
//There are no student comments for this assignment
func processComments(line string, sub *Submission) {
	return errors.New("Not Implemented")
}

//Files:
//	Original filename: OnlineExam-Bxxxxxx.pdf
//	Filename: Practice Exam Drop Box_sxxxxxxx_attempt_yyyy-mm-dd-hh-mm-ss_OnlineExam-Bxxxxxx.pdf
func processOriginalFilename(line string, sub *Submission) {
	return errors.New("Not Implemented")
}
func processFilename(line string, sub *Submission) {
	return errors.New("Not Implemented")
}

func writeSubmissionsToCSV(subs []Submission, outpath string) error {
	// wrap the marshalling library in case we need converters etc later
	file, err := os.OpenFile(inputPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	return gocsv.MarshalFile(&subs, outpath)
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
