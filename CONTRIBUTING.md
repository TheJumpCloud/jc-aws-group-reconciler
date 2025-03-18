# Contributing to JumpCloud AWS Group Reconciler

Thank you for your interest in contributing to the JumpCloud AWS Group Reconciler. This document provides
guidelines and workflows to help you contribute effectively.

## Code of Conduct

By participating in this project, you agree to act in the best interest of the community of users of this
tool.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally
3. **Add the upstream repository** as a remote:
   ```
   git remote add upstream https://github.com/TheJumpCloud/jc-aws-group-reconciler.git
   ```
4. **Create a branch** for your changes

## Development Environment

1. Ensure you have Go 1.23 or higher installed
2. Install dependencies: `go mod download`
3. Build the project: `make build`

## Making Changes

1. Create a focused branch addressing a specific issue
2. Write clear, commented code following Go best practices
3. Include tests for new functionality
4. Ensure all tests pass: `go test ./...`
5. Run `go fmt` and `go vet` to ensure code quality

## Submitting Changes

1. Push changes to your fork
2. Submit a pull request against the `main` branch of the original repository
3. Include a clear description of the changes and their purpose
4. Reference any related issues using GitHub's #issue syntax

## Pull Request Process

1. Update documentation for any changed functionality
2. Add or update tests as necessary
3. Ensure CI checks pass on your pull request
4. A maintainer will review your pull request and may request changes
5. Once approved, a maintainer will merge your changes

## Release Process

The maintainers handle the release process using the GitHub Actions workflow. When your changes are merged
to main, they'll be included in the next release.

## Questions?

If you have questions or need assistance, please open an issue on GitHub.

Thank you for contributing to the JumpCloud AWS Group Reconciler!
