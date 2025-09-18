# IPAM Integration API

The IPAM Integration provides APIs for integrating Auto Mode SGPP functionality with the existing VPC CNI plugin's IPAM context.

## 📋 **API Overview**

The IPAM Integration is responsible for:
- Seamless integration with existing IPAM context
- Pod lifecycle management with Auto Mode support
- ENI allocation and security group assignment
- Error handling and fallback mechanisms

## 🔧 **Core APIs**

### **NewAutoModeManager**

Creates a new Auto Mode Manager for IPAM integration.

```go
func NewAutoModeManager(ipamContext *IPAMContext) *AutoModeManager
```

**Parameters:**
- `ipamContext` (*IPAMContext): Existing IPAM context

**Returns:**
- `*AutoModeManager`: Manager instance

### **InitializeAutoMode**

Initializes Auto Mode functionality in the IPAM context.

```go
func (m *AutoModeManager) InitializeAutoMode() error
```

**Returns:**
- `error`: Error if initialization fails

### **ProcessPodWithAutoMode**

Processes a pod with Auto Mode SGPP support.

```go
func (m *AutoModeManager) ProcessPodWithAutoMode(ctx context.Context, podSpec *PodSpec) (*PodNetworkInfo, error)
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `podSpec` (*PodSpec): Pod specification

**Returns:**
- `*PodNetworkInfo`: Pod network information
- `error`: Error if processing fails

### **CleanupPodWithAutoMode**

Cleans up resources when a pod is terminated.

```go
func (m *AutoModeManager) CleanupPodWithAutoMode(ctx context.Context, podSpec *PodSpec) error
```

**Parameters:**
- `ctx` (context.Context): Context for the operation
- `podSpec` (*PodSpec): Pod specification

**Returns:**
- `error`: Error if cleanup fails

## 📊 **Data Structures**

### **PodNetworkInfo**

Contains network information for a pod.

```go
type PodNetworkInfo struct {
    PodName         string   `json:"podName"`
    Namespace       string   `json:"namespace"`
    ENIID           string   `json:"eniId"`
    PrivateIP       string   `json:"privateIp"`
    SecurityGroups  []string `json:"securityGroups"`
    SubnetID        string   `json:"subnetId"`
    AllocationTime  time.Time `json:"allocationTime"`
}
```

### **AutoModeManager**

Main manager for Auto Mode integration.

```go
type AutoModeManager struct {
    detector           *AutoModeDetector
    eniManager         *AutoModeENIManager
    securityGroupManager *SecurityGroupManager
    configManager      *ConfigurationManager
    ipamContext        *IPAMContext
}
```

## ⚙️ **Integration Points**

The IPAM Integration integrates with the following VPC CNI components:

### **IPAM Context**
- Extends existing `IPAMContext` with Auto Mode support
- Maintains backward compatibility with standard mode
- Provides seamless fallback mechanisms

### **Pod Lifecycle**
- Integrates with pod creation and deletion workflows
- Handles security group assignment during pod startup
- Manages ENI allocation and cleanup

### **Configuration**
- Uses existing VPC CNI configuration as base
- Extends with Auto Mode specific settings
- Maintains configuration consistency

## 🔄 **Processing Flow**

The IPAM Integration follows this processing flow:

1. **Pod Creation**:
   - Detect if cluster is in Auto Mode
   - Parse pod security group annotations
   - Allocate ENI with security groups
   - Update pod network configuration

2. **Pod Deletion**:
   - Release allocated ENI
   - Clean up security group associations
   - Update resource pools

3. **Error Handling**:
   - Fallback to standard mode if Auto Mode fails
   - Log errors for debugging
   - Maintain system stability

## 🐛 **Error Handling**

Common error scenarios:
- **Auto Mode Detection Failure**: Falls back to standard mode
- **ENI Allocation Failure**: Returns error with details
- **Security Group Validation Failure**: Returns validation error
- **Integration Errors**: Returns wrapped integration errors

## 📝 **Examples**

### **Basic IPAM Integration**

```go
package main

import (
    "context"
    "log"
    
    "github.com/aws/amazon-vpc-cni-k8s/pkg/ipamd"
    "github.com/aws/amazon-vpc-cni-k8s/pkg/automode"
)

func main() {
    // Create IPAM context
    ipamContext := ipamd.NewIPAMContext()
    
    // Create Auto Mode manager
    autoModeManager := automode.NewAutoModeManager(ipamContext)
    
    // Initialize Auto Mode
    if err := autoModeManager.InitializeAutoMode(); err != nil {
        log.Fatal(err)
    }
    
    // Process pod with Auto Mode
    podSpec := &automode.PodSpec{
        Name:      "my-pod",
        Namespace: "default",
        Annotations: map[string]string{
            "vpc.amazonaws.com/security-groups": "sg-12345678",
        },
    }
    
    ctx := context.Background()
    networkInfo, err := autoModeManager.ProcessPodWithAutoMode(ctx, podSpec)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Pod %s allocated ENI %s with IP %s", 
        networkInfo.PodName, networkInfo.ENIID, networkInfo.PrivateIP)
}
```

### **Error Handling Example**

```go
// Process pod with error handling
networkInfo, err := autoModeManager.ProcessPodWithAutoMode(ctx, podSpec)
if err != nil {
    // Log error and fallback to standard mode
    log.Printf("Auto Mode processing failed: %v", err)
    
    // Fallback to standard IPAM processing
    return ipamContext.ProcessPodStandard(ctx, podSpec)
}

return networkInfo, nil
```

## 🔗 **Related APIs**

- [Auto Mode Detector API](auto-mode-detector.md)
- [ENI Manager API](eni-manager.md)
- [Security Group Manager API](security-group-manager.md)
- [Configuration Manager API](configuration-manager.md)
