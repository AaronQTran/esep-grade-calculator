package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 59, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 70, Assignment)
	gc.AddGrade("exam", 70, Exam)
	gc.AddGrade("essay", 70, Essay)

	if gc.GetFinalGrade() != "C" {
		t.Error("Expected C")
	}
}
func TestGetGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 60, Assignment)
	gc.AddGrade("exam", 60, Exam)
	gc.AddGrade("essay", 60, Essay)

	if gc.GetFinalGrade() != "D" {
		t.Error("Expected D")
	}
}
func TestNoGrades(t *testing.T) {
	gc := NewGradeCalculator()
	if gc.GetFinalGrade() != "F" {
		t.Error("Expected F")
	}
}
func TestGradeType(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Error("Wrong grade type")
	}
}
