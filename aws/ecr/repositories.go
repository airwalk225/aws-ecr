package ecr

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecr"
	"os"
	"text/tabwriter"
	log "github.com/sirupsen/logrus"
)

func GetRepositories(svc *awsECR, flt *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error) {
	result, err := svc.DescribeRepositories(flt)

	return result, err
}

func getImages(ctx aws.Context, svc *awsECR, repo string) *ecr.ListImagesOutput {
	input := &ecr.ListImagesInput{
		RepositoryName: aws.String(repo),
	}

	result, err := svc.ListImagesWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			log.Print(aerr.Error())
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Print(err.Error())
		}
	}

	if //noinspection GoNilness
	len(result.ImageIds) < 1 {
		log.Printf("No images found in repo: %s", repo)
	}

	return result
}

func DisplayRepos(res *ecr.DescribeRepositoriesOutput) {
	if len(res.Repositories) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
		_, _ = fmt.Fprintln(w, "Repository\tARN\t")
		for _, repo := range res.Repositories {
			_, _ = fmt.Fprintf(w, "%s\t%s\t\n", *repo.RepositoryName, *repo.RepositoryArn)
		}
		_ = w.Flush()
	}
}

func DescribeRepoInDepth(res *ecr.DescribeRepositoriesOutput) {
	if len(res.Repositories) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
		for _, repo := range res.Repositories {
			//fmt.Printf("%+v\n", repo)
			_, _ = fmt.Fprintf(w, "%s\t%s\t\n", "CreatedAt", *repo.CreatedAt)
			_, _ = fmt.Fprintf(w, "%s\t%s\t\n", "ImageScanningConfiguration:", repo.ImageScanningConfiguration)
			// TODO: *repo.ImageTagMutability is a string, but has json in it, make this look nice
			_, _ = fmt.Fprintf(w, "%s\t%s\t\n", "ImageTagMutability.ScanOnPush", *repo.ImageTagMutability)
			_, _ = fmt.Fprintf(w, "%s\t%s\t\n", "RegistryId:", *repo.RegistryId)
			_, _ = fmt.Fprintf(w, "%s\t%s\t\n", "RepositoryArn:", *repo.RepositoryArn)
			_, _ = fmt.Fprintf(w, "%s\t%s\t\n", "RepositoryName:", *repo.RepositoryName)
			_, _ = fmt.Fprintf(w, "%s\t%s\t\n", "RepositoryUri:", *repo.RepositoryUri)
		}
		_ = w.Flush()
	}
}