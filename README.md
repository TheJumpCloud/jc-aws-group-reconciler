# AWS Group Checker

Set the following environment variables:
```
	JUMPCLOUD_API_KEY - a jumpcloud api key
	JUMPCLOUD_APPLICATION_IDS - the application ids bound to AWS
	AWS_REGION - the default region for your AWS instance (where the groups exist)
	AWS_ACCESS_KEY_ID - the access key for the user used to access AWS
	AWS_SECRET_ACCESS_KEY - the secret access key for the user used to access AWS
	AWS_ID_STORE_ID - the AWS ID Store id (can be found under the IAM Identity Center settings in AWS)
  (optional) AWS_SESSION_TOKEN - set this if your AWS instance is using SSO to authenticate
```

Clone this repository and run:
```
go run main.go
```
