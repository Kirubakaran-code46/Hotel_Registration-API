package s3filehandler

import (
	"HOTEL-REGISTRY_API/common"
	tomlread "HOTEL-REGISTRY_API/common/TomlRead"
	"HOTEL-REGISTRY_API/helpers"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsConfig struct {
	AccessKey        string
	SecretKey        string
	Region           string
	BucketName       string
	BucketFolderName string
}

// --------------------------------------------------------------------
// UPLOAD FILE INTO S3 BUCKET
// --------------------------------------------------------------------

func S3FileUpload(pDebug *helpers.HelperStruct, pReq *http.Request, pFileKey string) (string, error) {
	pDebug.Log(helpers.Statement, "S3FileUpload (+)")

	if pFileKey == "" {
		return "", nil
	}
	lFileName := ""

	// Read FormFile
	lFile, handler, lErr := pReq.FormFile(pFileKey)
	if lErr == http.ErrMissingFile {

		// File is optional, skip silently or log
		pDebug.Log(helpers.Elog, "S3FU001", lErr)
		return lFileName, nil

	} else if lErr != nil {
		pDebug.Log(helpers.Elog, "S3FU002", lErr)
		return lFileName, lErr
	}

	defer lFile.Close()

	//Generate DocId (or) FileName
	lDocId := common.GenerateDocID()
	ext := filepath.Ext(handler.Filename)
	lFileName = fmt.Sprintf("%s%s", lDocId, ext)

	// Load AWS config from TOML
	var lAwsCred AwsConfig

	lTomlconfig := tomlread.ReadTomlConfig("../AwsCredentials.toml")
	lAwsCred.AccessKey = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["AccessKey"])
	lAwsCred.SecretKey = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["SecretKey"])
	lAwsCred.Region = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["Region"])
	lAwsCred.BucketName = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["BucketName"])
	lAwsCred.BucketFolderName = fmt.Sprintf("%v", lTomlconfig.(map[string]interface{})["BucketFolderName"])

	// Read File Content
	lFileBytes := make([]byte, handler.Size)
	_, lErr = lFile.Read(lFileBytes)
	if lErr != nil {
		pDebug.Log(helpers.Elog, "S3FU003", lErr)
		return lFileName, lErr
	}

	lAwsConfig, lErr := config.LoadDefaultConfig(context.TODO(), config.WithRegion(lAwsCred.Region), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(lAwsCred.AccessKey, lAwsCred.SecretKey, "")))
	if lErr != nil {
		pDebug.Log(helpers.Elog, "S3FU004", lErr)
		return lFileName, lErr
	}

	s3Client := s3.NewFromConfig(lAwsConfig)

	lObjKey := lAwsCred.BucketFolderName + lFileName

	// Upload S3 Bucket
	_, lErr = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(lAwsCred.BucketName),
		Key:         aws.String(lObjKey),
		Body:        bytes.NewReader(lFileBytes),
		ContentType: aws.String(handler.Header.Get("Content-Type")),
	})
	if lErr != nil {
		pDebug.Log(helpers.Elog, "S3FU005", lErr)
		return lFileName, lErr
	}

	pDebug.Log(helpers.Statement, "S3FileUpload (-)")
	return lFileName, nil
}
