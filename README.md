# EKS Auto Mode Security Groups per Pod - Design Repository

This repository contains the complete design documentation and implementation plans for enabling Security Groups per Pod (SGPP) support in Amazon EKS Auto Mode clusters.

## 🎯 **Project Overview**

This project addresses the current limitation where EKS Auto Mode does not support Security Groups per Pod, enabling fine-grained network security controls in Auto Mode clusters.

## 📋 **Repository Contents**

### **Design Documents**
- **[DESIGN_DOCUMENT.md](DESIGN_DOCUMENT.md)** - Complete technical design specification
- **[IMPLEMENTATION_PLAN.md](IMPLEMENTATION_PLAN.md)** - Detailed implementation roadmap
- **[PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)** - Executive summary and project overview

### **Documentation**
- **[CONTRIBUTION_GUIDE.md](CONTRIBUTION_GUIDE.md)** - Guidelines for contributing to the project
- **[docs/](docs/)** - Comprehensive user guides and API documentation
- **[prototypes/](prototypes/)** - Working prototype implementations
- **[tests/](tests/)** - Test cases and validation scenarios

## 🏗️ **Architecture Overview**

The solution implements:
- **Auto Mode Detection Service** - Detects EKS clusters running in Auto Mode
- **Enhanced ENI Manager** - Pool-based ENI allocation for Auto Mode
- **Security Group Manager** - Annotation-based security group specification
- **Configuration Management** - Environment variable configuration
- **IPAM Integration** - Seamless integration with existing VPC CNI plugin

## 🚀 **Quick Start**

### **Prerequisites**
- EKS cluster running in Auto Mode
- VPC CNI plugin version with Auto Mode SGPP support
- Appropriate AWS IAM permissions

### **Configuration**
```bash
export CLUSTER_NAME="your-auto-mode-cluster"
export AWS_REGION="us-west-2"
export AUTO_MODE_SGPP_ENABLED="true"
```

### **Pod Security Group Assignment**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
  namespace: default
  annotations:
    vpc.amazonaws.com/security-groups: "sg-12345678"
spec:
  containers:
  - name: app
    image: nginx:latest
```

## 📊 **Key Features**

- ✅ **Pod-level Security Groups**: Assign specific security groups to pods
- ✅ **Multiple Security Groups**: Support up to 5 security groups per pod
- ✅ **Auto Mode Detection**: Automatic detection of Auto Mode clusters
- ✅ **ENI Pool Management**: Efficient resource utilization
- ✅ **Configuration Management**: Flexible environment-based configuration
- ✅ **Comprehensive Testing**: Thorough test coverage
- ✅ **Production Ready**: Performance optimized and thoroughly tested

## 🔗 **Related Repositories**

- **[amazon-vpc-cni-k8s](https://github.com/soumentrivedi/amazon-vpc-cni-k8s)** - Enhanced VPC CNI plugin with Auto Mode SGPP support
- **[eksctl](https://github.com/soumentrivedi/eksctl)** - Enhanced eksctl with Auto Mode support
- **[eks-anywhere](https://github.com/soumentrivedi/eks-anywhere)** - Enhanced EKS Anywhere with Auto Mode support
- **[aws-k8s-tester](https://github.com/soumentrivedi/aws-k8s-tester)** - Enhanced testing framework

## 📚 **Documentation**

- **[Getting Started Guide](docs/user-guide/getting-started.md)** - Complete setup and usage guide
- **[Usage Examples](docs/examples/basic-usage.md)** - Real-world examples and patterns
- **[Architecture Overview](docs/architecture/overview.md)** - Detailed architecture documentation
- **[API Reference](docs/api/README.md)** - Complete API documentation

## 🤝 **Contributing**

This project follows AWS contribution guidelines. Please see the [CONTRIBUTION_GUIDE.md](CONTRIBUTION_GUIDE.md) for detailed guidelines.

### **Contribution Process**
1. Fork the repository
2. Create a feature branch
3. Implement changes following the design specifications
4. Add comprehensive tests
5. Update documentation
6. Submit a pull request

## 📄 **License**

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

## 🎉 **Status**

**DESIGN COMPLETE** ✅

This repository contains the complete design and implementation plan for enabling Security Groups per Pod in EKS Auto Mode clusters. The solution follows AWS architecture patterns and standards.

---

**Note**: This implementation is based on the original AWS VPC CNI plugin and has been enhanced with Auto Mode SGPP support. The enhanced implementation is available in the related repositories listed above.