package ecr

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	log "github.com/sirupsen/logrus"
)

// Decorate the ecr.ECR struct so that new methods can be added
type awsECR struct {
	*ecr.ECR
}

func NewECR(sess *session.Session) *awsECR {
	return &awsECR{ecr.New(sess)}
}

// Add a new method of HandleError onto the *ecr.ECR -> *awsECR
func (m *awsECR) HandleError(err error, cmd string) {
	if aerr, ok := err.(awserr.Error); ok {
		errorLogger := log.WithFields(log.Fields{
			"service": "ecr",
			"command": cmd,
			"status_code": aerr.Code(),
		})
		switch aerr.Code() {
		case ecr.ErrCodeInvalidParameterException:
		case ecr.ErrCodeRepositoryNotFoundException:
			errorLogger.Warn(aerr.Error())
		case ecr.ErrCodeServerException:
		default:
			errorLogger.Fatal(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.WithFields(
			log.Fields{
				"service": "ecr",
				"command": cmd,
				"status_code": 101,
			}).Fatal(err.Error())
	}
}
