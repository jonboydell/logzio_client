package logshipping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	logzio_client "github.com/logzio/logzio_terraform_client"
)

const (
	createS3BucketServiceMethod string = http.MethodPost
	createS3BucketMethodSuccess int    = http.StatusOK
)

// type FieldError struct {
// 	Field   string
// 	Message string
// }

// func (e FieldError) Error() string {
// 	return fmt.Sprintf("%v: %v", e.Field, e.Message)
// }

func validateCreateS3BucketRequest(bucket S3Bucket) error {

	if len(bucket.AccessKey) == 0 {
		return fmt.Errorf("accessKey must be set")
	}
	if len(bucket.SecretKey) == 0 {
		return fmt.Errorf("secretKey must be set")
	}
	if len(bucket.Bucket) == 0 {
		return fmt.Errorf("bucket must be set")
	}
	if len(bucket.LogsType) == 0 {
		return fmt.Errorf("logsType must be set")
	}

	return nil
}

func buildCreateS3BucketRequest(bucket S3Bucket) map[string]interface{} {
	return map[string]interface{}{
		fldAwsAccessKey:                bucket.AccessKey,
		fldAwsSecretKey:                bucket.SecretKey,
		fldAwsArn:                      bucket.Arn,
		fldAwsBucket:                   bucket.Bucket,
		fldAwsPrefix:                   bucket.Prefix,
		fldAwsActive:                   bucket.Active,
		fldAwsAddS3ObjectKeyAsLogField: bucket.AddS3ObjectKeyAsLogField,
		fldAwsRegion:                   bucket.Region,
		fldLogsType:                    bucket.LogsType,
	}
}

func (c *LogShippingClient) buildCreateAPIRequest(apiToken string, jsonObject map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseURL := c.BaseUrl
	req, err := http.NewRequest(createS3BucketServiceMethod, fmt.Sprintf(createS3BucketServiceURL, baseURL), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// CreateBucket will return the created bucket if successful, an error otherwise
func (c *LogShippingClient) CreateBucket(bucket S3Bucket) (*S3Bucket, error) {
	err := validateCreateS3BucketRequest(bucket)
	if err != nil {
		return nil, err
	}

	createBucket := buildCreateS3BucketRequest(bucket)
	req, _ := c.buildCreateAPIRequest(c.ApiToken, createBucket)
	jsonResponse, err := logzio_client.CreateHttpRequest(req)
	if err != nil {
		return nil, err
	}

	retVal := jsonS3BucketToS3Bucket(jsonResponse)

	return &retVal, nil
}
