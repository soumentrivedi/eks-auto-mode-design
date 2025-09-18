# ENI Manager API

The Enhanced ENI Manager provides APIs for managing Elastic Network Interfaces (ENIs) in EKS Auto Mode clusters with Security Groups per Pod support.

## 📋 **API Overview**

The Enhanced ENI Manager is responsible for:
- Pool-based ENI allocation for Auto Mode clusters
- Security group assignment to ENIs
- ENI lifecycle management
- Resource optimization and cleanup

## 🔧 **Core APIs**

### **NewAutoModeENIManager**

Creates a new Enhanced ENI Manager instance.

```go
func NewAutoModeENIManager(clusterName, region string) (*AutoModeENIManager, error)
```

**Parameters:**
- `clusterName` (string): EKS cluster name
- `region` (string): AWS region where the cluster is located

**Returns:**
- `*AutoModeENIManager`: Manager instance
- `error`: Error if initialization fails

### **AllocateENIForPod**

Allocates an ENI for a pod with specified security groups.

```go
func (m *AutoModeENIManager) AllocateENIForPod(ctx context.Context, podSpec *PodSpec) (*ENIInfo, error)
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `podSpec` (*PodSpec): Pod specification with security group requirements

**Returns:**
- `*ENIInfo`: ENI information
- `error`: Error if allocation fails

### **ReleaseENIForPod**

Releases an ENI when a pod is terminated.

```go
func (m *AutoModeENIManager) ReleaseENIForPod(ctx context.Context, podSpec *PodSpec) error
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `podSpec` (*PodSpec): Pod specification

**Returns:**
- `error`: Error if release fails

## 📊 **Data Structures**

### **ENIInfo**

Contains information about an allocated ENI.

```go
type ENIInfo struct {
    ENIID           string   `json:"eniId"`
    SecurityGroups  []string `json:"securityGroups"`
    SubnetID        string   `json:"subnetId"`
    PrivateIP       string   `json:"privateIp"`
    AllocationTime  time.Time `json:"allocationTime"`
}
```

### **PodSpec**

Pod specification with security group requirements.

```go
type PodSpec struct {
    Name            string            `json:"name"`
    Namespace       string            `json:"namespace"`
    SecurityGroups  []string          `json:"securityGroups"`
    Annotations     map[string]string `json:"annotations"`
    Labels          map[string]string `json:"labels"`
}
```

## ⚙️ **Configuration**

The ENI Manager uses the following environment variables:

- `ENI_POOL_SIZE`: Maximum number of ENIs in the pool (default: 10)
- `ENI_ALLOCATION_TIMEOUT`: Timeout for ENI allocation (default: 30s)
- `ENI_CLEANUP_INTERVAL`: Interval for ENI cleanup (default: 5m)

## 🔄 **Pool Management**

The ENI Manager implements a pool-based allocation strategy:
- Pre-allocates ENIs for faster pod startup
- Maintains a pool of available ENIs
- Automatically scales the pool based on demand
- Cleans up unused ENIs periodically

## 🐛 **Error Handling**

Common error scenarios:
- **Pool Exhaustion**: Returns error when ENI pool is empty
- **AWS API Errors**: Returns wrapped AWS SDK errors
- **Invalid Security Groups**: Returns error for invalid security group IDs
- **Allocation Timeout**: Returns timeout error if allocation takes too long

## 📝 **Examples**

### **Basic ENI Allocation**

```go
package main

import (
    "context"
    "log"
    
    "github.com/aws/amazon-vpc-cni-k8s/pkg/automode"
)

func main() {
    // Create ENI manager
    manager, err := automode.NewAutoModeENIManager("my-cluster", "us-west-2")
    if err != nil {
        log.Fatal(err)
    }
    
    // Allocate ENI for pod
    podSpec := &automode.PodSpec{
        Name:           "my-pod",
        Namespace:      "default",
        SecurityGroups: []string{"sg-12345678"},
    }
    
    ctx := context.Background()
    eniInfo, err := manager.AllocateENIForPod(ctx, podSpec)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Allocated ENI %s with IP %s", eniInfo.ENIID, eniInfo.PrivateIP)
}
```

## 🔗 **Related APIs**

- [Auto Mode Detector API](auto-mode-detector.md)
- [Security Group Manager API](security-group-manager.md)
- [Configuration Manager API](configuration-manager.md)
