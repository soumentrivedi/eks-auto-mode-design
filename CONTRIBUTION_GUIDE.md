# Contribution Guide - EKS Auto Mode SGPP Enhancement

## Getting Started

### Prerequisites
- Go 1.19+ installed
- Docker installed
- AWS CLI configured
- kubectl installed
- eksctl installed

### Repository Setup
```bash
# Fork the amazon-vpc-cni-k8s repository
# Clone your fork
git clone https://github.com/soumentrivedi/amazon-vpc-cni-k8s.git
cd amazon-vpc-cni-k8s

# Add upstream remote
git remote add upstream https://github.com/aws/amazon-vpc-cni-k8s.git

# Create feature branch
git checkout -b feature/auto-mode-sgpp-support
```

## Development Workflow

### 1. Code Structure
```
pkg/
├── eniconfig/
│   ├── auto_mode_detector.go      # Auto Mode detection
│   ├── auto_mode_eni.go           # ENI management for Auto Mode
│   ├── auto_mode_sg.go           # Security group management
│   └── eniconfig.go              # Existing ENI configuration
├── ipamd/
│   └── ipamd.go                  # IP address management
└── awsutils/
    └── awsutils.go               # AWS API utilities
```

### 2. Coding Standards
- Follow Go coding standards
- Use meaningful variable and function names
- Add comprehensive comments
- Include error handling
- Write unit tests for all functions

### 3. Testing Requirements
- Unit tests with >90% coverage
- Integration tests for key workflows
- Performance tests for critical paths
- End-to-end tests for complete scenarios

## Implementation Guidelines

### Auto Mode Detection
```go
// Example implementation structure
type AutoModeDetector struct {
    clusterName string
    region      string
    eksClient   *eks.Client
}

func (d *AutoModeDetector) IsAutoMode() (bool, error) {
    // Implementation details
}
```

### ENI Management
```go
// Example implementation structure
type AutoModeENIManager struct {
    detector     *AutoModeDetector
    eniPool      *ENIPool
    sgManager    *SecurityGroupManager
}

func (m *AutoModeENIManager) AllocateENIForPod(podSpec *PodSpec) (*ENI, error) {
    // Implementation details
}
```

### Security Group Management
```go
// Example implementation structure
type SecurityGroupManager struct {
    ec2Client    *ec2.Client
    autoMode     bool
}

func (sgm *SecurityGroupManager) GetSecurityGroupForPod(podSpec *PodSpec) (*SecurityGroup, error) {
    // Implementation details
}
```

## Testing Guidelines

### Unit Tests
```go
func TestAutoModeDetector_IsAutoMode(t *testing.T) {
    // Test cases for Auto Mode detection
}

func TestAutoModeENIManager_AllocateENIForPod(t *testing.T) {
    // Test cases for ENI allocation
}

func TestSecurityGroupManager_GetSecurityGroupForPod(t *testing.T) {
    // Test cases for security group assignment
}
```

### Integration Tests
```go
func TestAutoModeSGPPIntegration(t *testing.T) {
    // Test complete SGPP workflow in Auto Mode
}
```

### Performance Tests
```go
func BenchmarkAutoModeENIAllocation(b *testing.B) {
    // Benchmark ENI allocation performance
}
```

## Pull Request Process

### 1. Before Submitting
- [ ] Run all tests locally
- [ ] Ensure code passes linting
- [ ] Update documentation
- [ ] Add test coverage
- [ ] Review code changes

### 2. Pull Request Template
```markdown
## Description
Brief description of changes

## Related Issues
Fixes #issue_number

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Performance tests pass
- [ ] Manual testing completed

## Documentation
- [ ] README updated
- [ ] API documentation updated
- [ ] User guide updated

## Checklist
- [ ] Code follows project standards
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] Backward compatibility maintained
```

### 3. Review Process
- Address all review comments
- Update implementation based on feedback
- Ensure all tests pass
- Update documentation as needed

## Code Review Guidelines

### For Reviewers
- Focus on code quality and correctness
- Check for security vulnerabilities
- Ensure performance considerations
- Validate test coverage
- Review documentation updates

### For Contributors
- Respond to review comments promptly
- Provide clear explanations for design decisions
- Update code based on feedback
- Ensure all tests pass after changes

## Community Guidelines

### Communication
- Use respectful and professional language
- Provide clear and constructive feedback
- Ask questions when clarification is needed
- Share knowledge and help others

### Issue Reporting
- Use clear and descriptive titles
- Provide detailed reproduction steps
- Include relevant logs and configurations
- Tag issues appropriately

## Resources

### Documentation
- [AWS EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/)
- [VPC CNI Plugin Documentation](https://github.com/aws/amazon-vpc-cni-k8s)
- [Go Documentation](https://golang.org/doc/)

### Tools
- [Go Testing](https://golang.org/pkg/testing/)
- [AWS SDK for Go](https://aws.github.io/aws-sdk-go-v2/)
- [Kubernetes Client Go](https://github.com/kubernetes/client-go)

### Support
- GitHub Issues for bug reports
- GitHub Discussions for questions
- AWS Forums for general questions
- Slack channels for real-time communication

## License

This project is licensed under the Apache License 2.0. By contributing, you agree that your contributions will be licensed under the same license.
