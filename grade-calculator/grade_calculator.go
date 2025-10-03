package esepunittests

type GradeCalculator struct {
	assignments []Grade
	isPassing bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		isPassing: true,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		gc.isPassing = false
		return "D"
	}

	gc.isPassing = false
	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.assignments = append(gc.assignments, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_sum := 0 
	assignment_count := 0 
	exam_sum := 0
	exam_count := 0
	essay_sum := 0
	essay_count := 0

	for _, assignment := range gc.assignments {
		switch assignment.Type {
			case Assignment: {
				assignment_sum += assignment.Grade
				assignment_count += 1
			}
			case Exam: {
				exam_sum += assignment.Grade
				exam_count += 1
			}
			case Essay: {
				essay_sum += assignment.Grade
				essay_count += 1
			}
		}
	}
	
	weighted_grade := 
		float64(assignment_sum/assignment_count)*0.5 +
		float64(exam_sum/exam_count)*0.35 +
		float64(essay_sum/essay_count)*0.15 


	return int(weighted_grade)
}