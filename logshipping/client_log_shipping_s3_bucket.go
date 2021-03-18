package logshipping

const (
	fldAwsAccessKey                string = "accessKey"
	fldAwsSecretKey                string = "secretKey"
	fldAwsArn                      string = "arn"
	fldAwsBucket                   string = "bucket"
	fldAwsPrefix                   string = "prefix"
	fldAwsActive                   string = "active"
	fldAwsAddS3ObjectKeyAsLogField string = "addS3ObjectKeyAsLogField"
	fldAwsRegion                   string = "region"

	s3BucketServiceURLBase string = serviceEndpoint + "/s3-buckets"
)

// S3Bucket Represents an S3 Bucket resource from logzio API
type S3Bucket struct {
	ID                       int64
	AccessKey                string
	SecretKey                string
	Arn                      string
	Bucket                   string
	Prefix                   string
	Active                   bool
	AddS3ObjectKeyAsLogField bool
	Region                   string
	LogsType                 string
}

func jsonS3BucketToS3Bucket(jsonS3Bucket map[string]interface{}) S3Bucket {
	bucket := S3Bucket{
		Id:                       int64(jsonS3Bucket[fldId].(float64)),
		AccessKey:                jsonS3Bucket[fldAwsAccessKey].(string),
		SecretKey:                jsonS3Bucket[fldAwsSecretKey].(string),
		Arn:                      jsonS3Bucket[fldAwsArn].(string),
		Bucket:                   jsonS3Bucket[fldAwsBucket].(string),
		Prefix:                   jsonS3Bucket[fldAwsPrefix].(string),
		Active:                   jsonS3Bucket[fldAwsActive].(bool),
		AddS3ObjectKeyAsLogField: jsonS3Bucket[fldAwsAddS3ObjectKeyAsLogField].(bool),
		Region:                   jsonS3Bucket[fldAwsRegion].(string),
		LogsType:                 jsonS3Bucket[fldLogsType].(string),
	}
	return bucket
}
