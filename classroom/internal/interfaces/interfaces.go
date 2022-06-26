package interfaces

import "github.com/lucasd-coder/classroom/internal/graphql/model"

type (
	CoursesRepository interface {
		ListAllCourses() []*model.Course
		GetCourseById(id string) *model.Course
		GetCourseBySlug(slug string) *model.Course
		Create(curso model.Course) (*model.Course, error)
	}

	EnrollmentsRepository interface {
		GetByCourseAndStudentId(courseId, studentId string) *model.Enrollment
		ListAllEnrollments() []*model.Enrollment
		ListEnrollmentsByStudent(studentId string) []*model.Enrollment
		Create(enrollment model.Enrollment) (*model.Enrollment, error)
		FindEnrollmentByID(id string) *model.Enrollment
	}

	StudentsRepository interface {
		ListAllStudents() []*model.User
		GetStudentByAuthUserId(authUserId string) *model.User
		GetStudentById(id string) *model.User
		Create(user model.User) (*model.User, error)
	}

	CoursesService interface {
		ListAllCourses() []*model.Course
		GetCourseById(id string) (*model.Course, error)
		GetCourseBySlug(slug string) (*model.Course, error)
		CreateCourse(title string) (*model.Course, error)
	}

	EnrollmentsService interface {
		GetByCourseAndStudentId(courseId, studentId string) (*model.Enrollment, error)
		ListAllEnrollments() []*model.Enrollment
		ListEnrollmentsByStudent(studentId string) []*model.Enrollment
		CreateEnrollment(courseId, studentId string) (*model.Enrollment, error)
		FindEnrollmentByID(id string) (*model.Enrollment, error)
	}

	StudentsService interface {
		ListAllStudents() []*model.User
		GetStudentByAuthUserId(authUserId string) (*model.User, error)
		GetStudentById(id string) (*model.User, error)
		CreateStudent(authUserId string) (*model.User, error)
	}
)
