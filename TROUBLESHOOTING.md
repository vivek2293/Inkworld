# Pre-Commit Hooks Troubleshooting
This section contains common issues and their solutions for the pre-commit hook.

1. Go-imports
   - **Issue**: The `go-imports` hook fails with an error message about import order or formatting.
   - **Solution**: Run `goimports -w .` to fix the import order and formatting issues in your Go files.

2. Go-sec-mod
    - **Issue**: The `go-sec-mod` hook fails with an error message about security issues in your Go code.
    - **Solution**: Review the reported issues and fix them in your code. You can also run `gosec ./...` to see the issues in detail.