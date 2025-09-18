# Security Group Manager API

The Security Group Manager provides APIs for managing security groups in EKS Auto Mode clusters with Security Groups per Pod support.

## 📋 **API Overview**

The Security Group Manager is responsible for:
- Annotation-based security group specification
- Security group validation and caching
- Multiple security group support (up to 5 per pod)
- Integration with AWS EC2 security groups

## 🔧 **Core APIs**

### **NewSecurityGroupManager**

Creates a new Security Group Manager instance.

```go
func NewSecurityGroupManager(clusterName, region string) (*SecurityGroupManager, error)
```

**Parameters:**
- `clusterName` (string): EKS cluster name
- `region` (string): AWS region where the cluster is located

**Returns:**
- `*SecurityGroupManager`: Manager instance
- `error`: Error if initialization fails

### **GetSecurityGroupForPod**

Retrieves security groups for a pod based on annotations and labels.

```go
func (m *SecurityGroupManager) GetSecurityGroupForPod(podSpec *PodSpec) ([]string, error)
```

**Parameters:**
- `podSpec` (*PodSpec): Pod specification

**Returns:**
- `[]string`: List of security group IDs
- `error`: Error if retrieval fails

### **ValidateSecurityGroup**

Validates that security groups exist and are accessible.

```go
func (m *SecurityGroupManager) ValidateSecurityGroup(ctx context.Context, sgID string) error
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `sgID` (string): Security group ID to validate

**Returns:**
- `error`: Error if validation fails

### **AttachSecurityGroupToENI**

Attaches security groups to an ENI.

```go
func (m *SecurityGroupManager) AttachSecurityGroupToENI(ctx context.Context, eniID string, sgIDs []string) error
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `eniID` (string): ENI ID
- `sgIDs` ([]string): List of security group IDs

**Returns:**
- `error`: Error if attachment fails

## 📊 **Data Structures**

### **SecurityGroupInfo**

Contains information about a security group.

```go
type SecurityGroupInfo struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Description string            `json:"description"`
    VPCID       string            `json:"vpcId"`
    Tags        map[string]string `json:"tags"`
}
```

## ⚙️ **Configuration**

The Security Group Manager uses the following environment variables:

- `MAX_SECURITY_GROUPS_PER_POD`: Maximum security groups per pod (default: 5)
- `SECURITY_GROUP_CACHE_TTL`: Cache TTL for security group info (default: 5m)
- `SECURITY_GROUP_VALIDATION_TIMEOUT`: Timeout for validation (default: 10s)

## 🔄 **Caching**

The Security Group Manager implements caching for improved performance:
- Security group information is cached for 5 minutes by default
- Cache TTL can be configured via environment variables
- Cache is automatically refreshed when expired

## 🏷️ **Annotation Support**

The manager supports the following pod annotations:

- `vpc.amazonaws.com/security-groups`: Comma-separated list of security group IDs
- `vpc.amazonaws.com/security-group-ids`: Alternative annotation for security group IDs

### **Label Support**

The manager also supports pod labels:

- `vpc.amazonaws.com/security-groups`: Comma-separated list of security group IDs

## 🐛 **Error Handling**

Common error scenarios:
- **Invalid Security Groups**: Returns error for non-existent security groups
- **Too Many Security Groups**: Returns error if pod exceeds maximum limit
- **AWS API Errors**: Returns wrapped AWS SDK errors
- **Permission Denied**: Returns error if security groups are not accessible

## 📝 **Examples**

### **Basic Security Group Assignment**

```go
package main

import (
    "context"
    "log"
    
    "github.com/aws/amazon-vpc-cni-k8s/pkg/automode"
)

func main() {
    // Create security group manager
    manager, err := automode.NewSecurityGroupManager("my-cluster", "us-west-2")
    if err != nil {
        log.Fatal(err)
    }
    
    // Get security groups for pod
    podSpec := &automode.PodSpec{
        Name:      "my-pod",
        Namespace: "default",
        Annotations: map[string]string{
            "vpc.amazonaws.com/security-groups": "sg-12345678,sg-87654321",
        },
    }
    
    sgIDs, err := manager.GetSecurityGroupForPod(podSpec)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Pod security groups: %v", sgIDs)
}
```

### **Security Group Validation**

```go
// Validate security groups before allocation
ctx := context.Background()
for _, sgID := range sgIDs {
    if err := manager.ValidateSecurityGroup(ctx, sgID); err != nil {
        log.Printf("Invalid security group %s: %v", sgID, err)
        return err
    }
}
```

### **Pod YAML Example**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
  namespace: default
  annotations:
    vpc.amazonaws.com/security-groups: "sg-12345678,sg-87654321"
spec:
  containers:
  - name: app
    image: nginx:latest
```

## 🔗 **Related APIs**

- [Auto Mode Detector API](auto-mode-detector.md)
- [ENI Manager API](eni-manager.md)
- [Configuration Manager API](configuration-manager.md)
