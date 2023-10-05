# AWS Group Checker

This tool compares user groups in JumpCloud to user groups in AWS. It identifies and reconciles
a couple of things:
1. Any extra groups in AWS that are not bound to an AWS application in JumpCloud.
2. Any users that exist in AWS groups that are not a member of the bound JumpCloud group.

## Setup

### AWS Permissions and Setup
Perform the following steps in the AWS console:
1. Go to Identity and Access Management (IAM) Dashboard
2. Enable access to the `Identity Store` service
3. Enable and attach the following actions to the entity calling the script:
```
ListGroupMemberships
ListGroups
ListUsers
DescribeUser
```

Additionally, gather the following information for the AWS environment:
1. Default AWS region (where the groups exist)
2. The access key for the user used to access AWS
3. The secret access key for the user used to access AWS
4. The AWS ID store id. This can be found under the IAM Identity Center setting in AWS
4. _Optional_ The AWS session token. Use this if your AWS instance uses SSO

### JumpCloud Permissions and Setup
Gather the following information:
1. The JumpCloud api key for the environment
2. The application IDs for all AWS applications bound to JumpCloud user groups


## Using the Tool
Using the information gathered above, set the following environment variables:
```
JUMPCLOUD_API_KEY - a jumpcloud api key
JUMPCLOUD_APPLICATION_IDS - the application ids bound to AWS
AWS_REGION - the default region for your AWS instance (where the groups exist)
AWS_ACCESS_KEY_ID - the access key for the user used to access AWS
AWS_SECRET_ACCESS_KEY - the secret access key for the user used to access AWS
AWS_ID_STORE_ID - the AWS ID Store id (can be found under the IAM Identity Center settings in AWS)
(optional) AWS_SESSION_TOKEN - set this if your AWS instance uses SSO to authenticate
```

Clone this repository and run:
```
go run .
```
