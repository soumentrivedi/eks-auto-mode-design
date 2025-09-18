# Pod Management API

The Pod Management API provides high-level APIs for managing pod lifecycle with Auto Mode SGPP support.

## 📋 **API Overview**

The Pod Management API is responsible for:
- Pod creation and deletion workflows
- Security group assignment and validation
- ENI lifecycle management
- Integration with Kubernetes pod events

## 🔧 **Core APIs**

### **NewPodManager**

Creates a new Pod Manager instance.

```go
func NewPodManager(clusterName, region string) (*PodManager, error)
```

**Parameters:**
- `clusterName` (string): EKS cluster name
- `region` (string): AWS region where the cluster is located

**Returns:**
- `*PodManager`: Manager instance
- `error`: Error if initialization fails

### **CreatePod**

Creates a pod with Auto Mode SGPP support.

```go
func (m *PodManager) CreatePod(ctx context.Context, podSpec *PodSpec) (*PodNetworkInfo, error)
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `podSpec` (*PodSpec): Pod specification

**Returns:**
- `*PodNetworkInfo`: Pod network information
- `error`: Error if creation fails

### **DeletePod**

Deletes a pod and cleans up associated resources.

```go
func (m *PodManager) DeletePod(ctx context.Context, podSpec *PodSpec) error
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `podSpec` (*PodSpec): Pod specification

**Returns:**
- `error`: Error if deletion fails

### **UpdatePodSecurityGroups**

Updates security groups for an existing pod.

```go
func (m *PodManager) UpdatePodSecurityGroups(ctx context.Context, podSpec *PodSpec, newSecurityGroups []string) error
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `podSpec` (*PodSpec): Pod specification
- `newSecurityGroups` ([]string): New security group IDs

**Returns:**
- `error`: Error if update fails

## 📊 **Data Structures**

### **PodManager**

Main manager for pod lifecycle management.

```go
type PodManager struct {
    detector           *AutoModeDetector
    eniManager         *AutoModeENIManager
    securityGroupManager *SecurityGroupManager
    configManager      *ConfigurationManager
    k8sClient          kubernetes.Interface
}
```

### **PodEvent**

Represents a pod lifecycle event.

```go
type PodEvent struct {
    Type      PodEventType `json:"type"`
    PodSpec   *PodSpec     `json:"podSpec"`
    Timestamp time.Time    `json:"timestamp"`
    Error     error        `json:"error,omitempty"`
}
```

### **PodEventType**

Type of pod event.

```go
type PodEventType string

const (
    PodEventCreated   PodEventType = "created"
    PodEventDeleted   PodEventType = "deleted"
    PodEventUpdated   PodEventType = "updated"
    PodEventFailed    PodEventType = "failed"
)
```

## ⚙️ **Pod Lifecycle Management**

The Pod Management API handles the complete pod lifecycle:

### **Pod Creation**
1. Validate pod specification
2. Check if cluster is in Auto Mode
3. Parse security group annotations
4. Allocate ENI with security groups
5. Update pod network configuration
6. Monitor pod status

### **Pod Deletion**
1. Identify pod resources
2. Release allocated ENI
3. Clean up security group associations
4. Update resource pools
5. Log cleanup completion

### **Pod Updates**
1. Detect configuration changes
2. Validate new security groups
3. Update ENI security groups
4. Maintain pod connectivity

## 🔄 **Event Handling**

The Pod Manager integrates with Kubernetes events:

### **Pod Events**
- Watches for pod creation, deletion, and updates
- Processes events asynchronously
- Maintains event history for debugging

### **Error Handling**
- Captures and logs all errors
- Provides detailed error information
- Implements retry mechanisms

## 🐛 **Error Handling**

Common error scenarios:
- **Invalid Pod Specification**: Returns validation error
- **Auto Mode Detection Failure**: Falls back to standard mode
- **ENI Allocation Failure**: Returns allocation error
- **Security Group Validation Failure**: Returns validation error
- **Kubernetes API Errors**: Returns wrapped K8s errors

## 📝 **Examples**

### **Basic Pod Creation**

```go
package main

import (
    "context"
    "log"
    
    "github.com/aws/amazon-vpc-cni-k8s/pkg/automode"
)

func main() {
    // Create pod manager
    manager, err := automode.NewPodManager("my-cluster", "us-west-2")
    if err != nil {
        log.Fatal(err)
    }
    
    // Create pod with security groups
    podSpec := &automode.PodSpec{
        Name:      "my-pod",
        Namespace: "default",
        Annotations: map[string]string{
            "vpc.amazonaws.com/security-groups": "sg-12345678,sg-87654321",
        },
    }
    
    ctx := context.Background()
    networkInfo, err := manager.CreatePod(ctx, podSpec)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Created pod %s with ENI %s", networkInfo.PodName, networkInfo.ENIID)
}
```

### **Pod Deletion**

```go
// Delete pod and cleanup resources
err := manager.DeletePod(ctx, podSpec)
if err != nil {
    log.Printf("Failed to delete pod: %v", err)
    return err
}

log.Printf("Successfully deleted pod %s", podSpec.Name)
```

### **Security Group Update**

```go
// Update pod security groups
newSecurityGroups := []string{"sg-11111111", "sg-22222222"}
err := manager.UpdatePodSecurityGroups(ctx, podSpec, newSecurityGroups)
if err != nil {
    log.Printf("Failed to update security groups: %v", err)
    return err
}

log.Printf("Updated security groups for pod %s", podSpec.Name)
```

### **Kubernetes Pod YAML**

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
    ports:
    - containerPort: 80
```

## 🔗 **Related APIs**

- [Auto Mode Detector API](auto-mode-detector.md)
- [ENI Manager API](eni-manager.md)
- [Security Group Manager API](security-group-manager.md)
- [IPAM Integration API](ipam-integration.md)
