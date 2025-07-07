package s3filehandler

import (
	tomlread "HOTEL-REGISTRY_API/common/TomlRead"
	"HOTEL-REGISTRY_API/helpers"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Download file from S3 and return as byte slice or stream
func S3FileDownload(pDebug *helpers.HelperStruct, pFileName string) ([]byte, error) {
	pDebug.Log(helpers.Statement, "S3FileDownload (+)")

	var lAwsCred AwsConfig

	// Load config from TOML
	lTomlconfig := tomlread.ReadTomlConfig("toml/AwsCredentials.toml")
	lAwsCred.AccessKey = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["AccessKey"])
	lAwsCred.SecretKey = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["SecretKey"])
	lAwsCred.Region = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["Region"])
	lAwsCred.BucketName = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["BucketName"])
	lAwsCred.BucketFolderName = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["BucketFolderName"])

	// Load AWS config
	lAwsConfig, lErr := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(lAwsCred.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			lAwsCred.AccessKey, lAwsCred.SecretKey, "",
		)),
	)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "S3FD001", lErr)
		return nil, lErr
	}

	s3Client := s3.NewFromConfig(lAwsConfig)

	lObjKey := lAwsCred.BucketFolderName + pFileName

	// Call S3 GetObject
	resp, lErr := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(lAwsCred.BucketName),
		Key:    aws.String(lObjKey),
	})
	if lErr != nil {
		pDebug.Log(helpers.Elog, "S3FD002", lErr)
		return nil, lErr
	}
	defer resp.Body.Close()

	// Read the response body
	buf := new(bytes.Buffer)
	_, lErr = io.Copy(buf, resp.Body)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "S3FD003", lErr)
		return nil, lErr
	}

	pDebug.Log(helpers.Statement, "S3FileDownload (-)")
	return buf.Bytes(), nil
}
