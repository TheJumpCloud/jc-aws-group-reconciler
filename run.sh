#!/usr/bin/env bash
set -euo pipefail

# Helper script to run the jc-aws-group-reconciler with environment variables

# Set your environment variables here
export JUMPCLOUD_API_KEY="your-api-key"
export JUMPCLOUD_APPLICATION_IDS="app-id-1,app-id-2"
export JUMPCLOUD_ORG_ID="your-org-id"
export AWS_REGION="us-west-2"
export AWS_ACCESS_KEY_ID="your-access-key"
export AWS_SECRET_ACCESS_KEY="your-secret-key"
export AWS_ID_STORE_ID="your-id-store"
# Uncomment if needed:
# export AWS_SESSION_TOKEN="your-session-token"

# Run the reconciler
./jc-aws-group-reconciler
