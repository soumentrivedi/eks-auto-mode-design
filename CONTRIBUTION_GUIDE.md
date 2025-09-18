# Contribution Guide - EKS Auto Mode SGPP Enhancement

This guide follows AWS contribution guidelines for the EKS Auto Mode Security Groups per Pod enhancement.

## 📋 **AWS Contribution Process**

### **1. AWS CLA and Code of Conduct**
- Sign the AWS Contributor License Agreement (CLA)
- Follow the AWS Code of Conduct
- Ensure all contributions align with AWS standards

### **2. Issue Creation**
- Create detailed GitHub issues for all changes
- Use appropriate labels and milestones
- Include design documents for significant changes

### **3. Pull Request Process**
- Create feature branches from main
- Follow AWS coding standards
- Include comprehensive tests
- Update documentation
- Request reviews from AWS maintainers

## 🚀 **Getting Started**

### **Prerequisites**
- Go 1.19+ installed
- Docker installed
- AWS CLI configured
- kubectl installed
- eksctl installed

### **Repository Setup**
```bash
# Fork the amazon-vpc-cni-k8s repository
git clone https://github.com/soumentrivedi/amazon-vpc-cni-k8s.git
cd amazon-vpc-cni-k8s

# Add upstream remote
git remote add upstream https://github.com/aws/amazon-vpc-cni-k8s.git

# Create feature branch
git checkout -b feature/auto-mode-sgpp-support
```

## 🔧 **Development Workflow**

### **1. Code Structure**
```
pkg/
├── automode/
│   ├── detector.go              # Auto Mode detection
│   ├── eni_manager.go           # ENI management for Auto Mode
│   ├── security_group_manager.go # Security group management
│   └── config.go                # Configuration management
├── ipamd/
│   ├── ipamd.go                 # IP address management
│   └── auto_mode_integration.go # Auto Mode integration
└── awsutils/
    └── awsutils.go              # AWS API utilities
```

### **2. AWS Coding Standards**
- Follow Go coding standards and AWS style guide
- Use meaningful variable and function names
- Add comprehensive comments and documentation
- Include proper error handling
- Write unit tests for all functions
- Follow AWS security best practices

### **3. Testing Requirements**
- Unit tests with >90% coverage
- Integration tests for AWS API calls
- End-to-end tests with real EKS clusters
- Performance tests for ENI allocation
- Security tests for security group validation

### **4. Documentation Requirements**
- Update API documentation
- Add user guides for new features
- Include troubleshooting information
- Update architecture diagrams
- Add configuration examples

## 📝 **Pull Request Guidelines**

### **PR Title Format**
```
feat: Add Auto Mode SGPP support
fix: Resolve ENI allocation timeout issue
docs: Update API documentation
test: Add integration tests for Auto Mode
```

### **PR Description Template**
```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Manual testing completed

## Checklist
- [ ] Code follows AWS style guide
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] Tests added/updated
```

### **Review Process**
1. **Self Review**: Review your own code before submitting
2. **Automated Checks**: Ensure all CI/CD checks pass
3. **AWS Maintainer Review**: Request review from AWS maintainers
4. **Testing**: Verify changes work in AWS environments
5. **Approval**: Wait for AWS maintainer approval

## 🔍 **Code Review Checklist**

### **Code Quality**
- [ ] Code follows AWS Go style guide
- [ ] Functions are well-documented
- [ ] Error handling is comprehensive
- [ ] No hardcoded values
- [ ] Proper logging implemented

### **Security**
- [ ] No sensitive data in code
- [ ] Proper input validation
- [ ] Security group validation implemented
- [ ] AWS IAM permissions follow least privilege
- [ ] No security vulnerabilities

### **Performance**
- [ ] Efficient resource usage
- [ ] Proper caching implemented
- [ ] No memory leaks
- [ ] Optimized AWS API calls
- [ ] Performance tests included

### **Testing**
- [ ] Unit tests cover new code
- [ ] Integration tests included
- [ ] Edge cases tested
- [ ] Error scenarios tested
- [ ] Performance benchmarks included

## 🚨 **AWS-Specific Requirements**

### **AWS API Usage**
- Use AWS SDK v2 for all API calls
- Implement proper retry logic with exponential backoff
- Handle AWS service limits and throttling
- Use appropriate AWS regions and endpoints

### **Security Considerations**
- Validate all security group IDs
- Implement proper IAM role assumptions
- Use AWS KMS for sensitive data
- Follow AWS security best practices

### **Monitoring and Logging**
- Use AWS CloudWatch for logging
- Implement proper metrics collection
- Add AWS X-Ray tracing where appropriate
- Follow AWS observability guidelines

## 📚 **Resources**

### **AWS Documentation**
- [AWS VPC CNI Plugin](https://github.com/aws/amazon-vpc-cni-k8s)
- [EKS User Guide](https://docs.aws.amazon.com/eks/)
- [AWS Go SDK](https://aws.github.io/aws-sdk-go-v2/)
- [AWS Security Best Practices](https://aws.amazon.com/security/security-resources/)

### **Contribution Resources**
- [AWS Contributor Guide](https://aws.github.io/aws-contributor-guide/)
- [AWS Code of Conduct](https://aws.github.io/code-of-conduct/)
- [AWS CLA](https://aws.github.io/cla/)

## 🤝 **Getting Help**

### **Community Support**
- GitHub Discussions for questions
- AWS re:Post for AWS-specific questions
- Slack channels for real-time help

### **AWS Maintainers**
- Tag AWS maintainers in PRs
- Use appropriate labels for issues
- Follow AWS communication guidelines

## 📄 **License**

This project is licensed under the Apache License 2.0. By contributing, you agree that your contributions will be licensed under the same license.

---

**Note**: This contribution guide follows AWS standards and guidelines. Please ensure all contributions align with AWS policies and best practices.
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
