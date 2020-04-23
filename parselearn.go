package parselearn

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type Submission struct {
	FirstName          string  `csv:"FirstName"`
	LastName           string  `csv:"LastName"`
	Matriculation      string  `csv:"Matriculation"`
	Assignment         string  `csv:"Assignment"`
	DateSubmitted      string  `csv:"DateSubmitted"`
	SubmissionField    string  `csv:"SubmissionField"`
	Comments           string  `csv:"Comments"`
	OriginalFilename   string  `csv:"OriginalFilename"`
	Filename           string  `csv:"Filename"`
	ExamNumber         string  `csv:"ExamNumber"`
	MatriculationError string  `csv:"MatriculationError"`
	ExamNumberError    string  `csv:"ExamNumberError"`
	FiletypeError      string  `csv:"FiletypeError"`
	FilenameError      string  `csv:"FilenameError"`
	NumberOfPages      string  `csv:"NumberOfPages"`
	FilesizeMB         float64 `csv:"FilesizeMB`
	NumberOfFiles      int     `csv:"NumberOfFiles"`
}

func parseLearnReceipt(inputPath string) (Submission, error) {

	submission := Submission{}

	file, err := os.Open(inputPath)

	if err != nil {
		return submission, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

SCAN:
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		switch {
		case strings.HasPrefix(line, "Name:"):
			processName(line, sub)
		case strings.HasPrefix(line, "Assignment:"):
			processAssignment(line, sub)
		case strings.HasPrefix(line, "Date Submitted:"):
			processDateSubmitted(line, sub)
		case strings.HasPrefix(line, "Submission Field:"):
			processSubmission(scanner.Text(), sub)
		case strings.HasPrefix(line, "Comments:"):
			processComments(scanner.Text(), sub)
		case strings.HasPrefix(line, "Files:"):
			break SCAN
		default:
			continue
		}
	}

	// now read in the files ....
	// TODO figure out nested csv so we can record multiple files
	// meanwhile for safety, count the number of original files

	submission.NumberOfFiles = 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		switch {
		case strings.HasPrefix(line, "Original filename:"):
			processOriginalFilename(line, sub)
			submission.NumberOfFiles++
		case strings.HasPrefix(line, "Filename:"):
			processFilename(line, sub)
		default:
			continue
		}

	}

	return submission, scanner.Err()
}

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
