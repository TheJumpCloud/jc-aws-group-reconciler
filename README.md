# AWS Group Reconciler

This tool compares user groups in JumpCloud to user groups in AWS Identity Center. It identifies and reconciles:

1. Extra groups in AWS that are not bound to an AWS application in JumpCloud
2. Users in AWS groups who are not members of the corresponding JumpCloud group

## Quick Start

### 1. Download the Tool

Pre-built binaries are available on the [Releases page](https://github.com/TheJumpCloud/jc-aws-group-reconciler/releases).

```bash
# Example for macOS arm64 (replace vX.Y.Z with the latest version)
curl -LO https://github.com/TheJumpCloud/jc-aws-group-reconciler/releases/download/vX.Y.Z/jc-aws-group-reconciler-X.Y.Z-macos-arm64.zip
unzip jc-aws-group-reconciler-X.Y.Z-macos-arm64.zip
chmod +x jc-aws-group-reconciler-X.Y.Z-macos-arm64
```

### 2. Configure Environment Variables

Set the required environment variables:

```bash
export JUMPCLOUD_API_KEY="your-jumpcloud-api-key"
export JUMPCLOUD_APPLICATION_IDS="app-id-1,app-id-2"
export AWS_REGION="your-aws-region"
export AWS_ACCESS_KEY_ID="your-access-key"
export AWS_SECRET_ACCESS_KEY="your-secret-key"
export AWS_ID_STORE_ID="your-id-store-id"
# Optional: export AWS_SESSION_TOKEN="your-session-token"
```

### 3. Run the tool
```bash
./jc-aws-group-reconciler-X.Y.Z-macos-arm64
```

## Prerequisites

### JumpCloud Requirements

1. JumpCloud API Key with admin privileges
2. Application IDs for all AWS applications bound to JumpCloud user groups

### AWS Requirements

#### Required Permissions
Your AWS credentials need the following permissions:

* ListGroupMemberships
* ListGroups
* ListUsers
* DescribeUser

#### Setup Steps
1. Go to Identity and Access Management (IAM) Dashboard
2. Enable access to the Identity Store service
3. Create a policy with the required permissions
4. Attach the policy to your IAM user/role

#### Required Information
1. AWS Region (where your Identity Center is configured)
2. AWS Access Key ID and Secret Access Key
3. AWS ID Store ID (found in IAM Identity Center settings)
4. Session Token (only if you're using AWS SSO)

## Usage Guide

### Environment Variables

| Variable                 | Description                                | Required |
|--------------------------|--------------------------------------------|----------|
| JUMPCLOUD_API_KEY        | JumpCloud admin API key                    | Yes      |
| JUMPCLOUD_APPLICATION_IDS| Comma-separated list of JumpCloud application IDs | Yes      |
| AWS_REGION               | AWS region for Identity Center             | Yes      |
| AWS_ACCESS_KEY_ID        | AWS access key ID                          | Yes      |
| AWS_SECRET_ACCESS_KEY    | AWS secret access key                      | Yes      |
| AWS_ID_STORE_ID          | AWS Identity Store ID                      | Yes      |
| AWS_SESSION_TOKEN        | AWS session token (for SSO)                | No       |

#### Helper Script

For convenience, you can create a helper script to set environment variables. A sample `run.sh` is provided in the repository.

## Verification

You can verify the integrity of downloaded binaries using the SHA256SUMS.txt file included in each release:

```bash
sha256sum -c SHA256SUMS.txt
```

## Building From Source

### Prerequisites
* Go 1.18 or newer
* Git

### Option 1: Using Make (Recommended)

```bash
# Clone the repository
git clone https://github.com/TheJumpCloud/jc-aws-group-reconciler.git
cd jc-aws-group-reconciler

# Build for your platform
make build

# Build for all platforms
make build-all

# Clean build artifacts
make clean
```

### Option 2: Using Go Directly

```bash
# Clone the repository
git clone https://github.com/TheJumpCloud/jc-aws-group-reconciler.git
cd jc-aws-group-reconciler

# Run directly
go run .

# Or build manually
go build -o jc-aws-group-reconciler .
```

## For Contributors

### Releasing New Versions

1. Develop and test your changes
2. Create a PR and merge changes to the main branch
3. Check the existing tags to establish the next version
4. Run the 'Create Tag' GitHub Actions workflow using the next version as input
5. The 'Release' GitHub Actions workflow will automatically build binaries and create a release
6. Binaries can be downloaded from the GitHub Releases page
