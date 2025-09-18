# API Reference - EKS Auto Mode SGPP Enhancement

This directory contains the complete API reference documentation for the EKS Auto Mode Security Groups per Pod enhancement.

## 📋 **API Documentation Structure**

### **Core Components**
- **[Auto Mode Detector API](auto-mode-detector.md)** - Auto Mode detection and cluster information APIs
- **[ENI Manager API](eni-manager.md)** - Enhanced ENI management APIs
- **[Security Group Manager API](security-group-manager.md)** - Security group management APIs
- **[Configuration Manager API](configuration-manager.md)** - Configuration management APIs

### **Integration APIs**
- **[IPAM Integration API](ipam-integration.md)** - IPAM context integration APIs
- **[Pod Management API](pod-management.md)** - Pod lifecycle management APIs

### **Data Structures**
- **[Data Types](data-types.md)** - Core data structures and types
- **[Error Handling](error-handling.md)** - Error codes and handling patterns

## 🚀 **Quick Reference**

### **Auto Mode Detection**
```go
// Check if cluster is in Auto Mode
isAutoMode, err := detector.IsAutoMode(ctx)

// Get cluster information
clusterInfo, err := detector.GetClusterInfo(ctx)
```

### **ENI Management**
```go
// Allocate ENI for pod
eni, err := eniManager.AllocateENIForPod(ctx, podSpec)

// Release ENI for pod
err := eniManager.ReleaseENIForPod(ctx, podSpec)
```

### **Security Group Management**
```go
// Get security group for pod
sg, err := sgManager.GetSecurityGroupForPod(podSpec)

// Attach security group to ENI
err := sgManager.AttachSecurityGroupToENI(ctx, eniID, sgID)
```

## 📚 **Usage Examples**

See the [examples directory](../examples/) for comprehensive usage examples and patterns.

## 🔧 **Configuration**

All APIs support configuration through environment variables. See the [Configuration Manager API](configuration-manager.md) for details.

## 🐛 **Error Handling**

All APIs return structured errors with appropriate error codes. See [Error Handling](error-handling.md) for details.

## 📄 **License**

This API documentation is part of the EKS Auto Mode SGPP enhancement project and is licensed under the Apache License 2.0.
