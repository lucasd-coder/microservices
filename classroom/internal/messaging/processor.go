package messaging

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	graphqlmodel "github.com/lucasd-coder/classroom/internal/graphql/model"
	"github.com/lucasd-coder/classroom/internal/interfaces"
	"github.com/lucasd-coder/classroom/internal/model"
	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/lucasd-coder/classroom/internal/service"
)

type PurchasesProcessor struct {
	StudentsService    interfaces.StudentsService
	CoursesService     interfaces.CoursesService
	EnrollmentsService interfaces.EnrollmentsService
}

func NewPurchasesProcessor(
	studentsService *service.StudentsService,
	coursesService *service.CoursesService,
	enrollmentsService *service.EnrollmentsService,
) *PurchasesProcessor {
	return &PurchasesProcessor{
		StudentsService:    studentsService,
		CoursesService:     coursesService,
		EnrollmentsService: enrollmentsService,
	}
}

func (c *PurchasesProcessor) PurchaseCreated(msg *kafka.Message) error {
	var payload model.PurchaseCreatedPayload

	err := json.Unmarshal(msg.Value, &payload)
	if err != nil {
		return err
	}

	student, err := c.createStudent(payload.Customer.AuthUserId)
	if err != nil {
		return err
	}

	course, err := c.createCourse(payload.Product.Slug, payload.Product.Title)
	if err != nil {
		return err
	}

	enrollment, err := c.EnrollmentsService.CreateEnrollment(course.ID, student.ID)
	if err != nil {
		return err
	}

	logger.Log.Infof("Created successfully PurchaseId: { %s }", enrollment.ID)
	return err
}

func (c *PurchasesProcessor) createStudent(authUserId string) (*graphqlmodel.User, error) {
	if getStudent, err := c.StudentsService.GetStudentByAuthUserId(authUserId); err != nil {

		student, err := c.StudentsService.CreateStudent(authUserId)
		if err != nil {
			return &graphqlmodel.User{}, err
		}
		return student, nil
	} else {
		return getStudent, nil
	}
}

func (c *PurchasesProcessor) createCourse(slug, title string) (*graphqlmodel.Course, error) {
	if getCourse, err := c.CoursesService.GetCourseBySlug(slug); err != nil {
		course, err := c.CoursesService.CreateCourse(title)
		if err != nil {
			return &graphqlmodel.Course{}, err
		}
		return course, nil
	} else {
		return getCourse, nil
	}
}
