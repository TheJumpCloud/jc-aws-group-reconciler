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

## Building the Tool

### Option 1: Building a binary (recommended)
The project includes a Makefile to simplify the build process:

1. To build for your current platform:

```
make build
```

2. To build for multiple platforms (Windows, macOS, and Linux):

```
make build-all
```

3. To clean up build artifacts:

```
make clean
```

### Option 2: Running from source
Clone this repository and run:

```
go run .
```

### Distributing the binary
After building the binary for the required platforms using `make build-all`, you can distribute the appropriate binary file to users.
The binary is standalone and doesn't require Go to be installed on the target machine.

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
