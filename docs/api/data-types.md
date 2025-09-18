# Data Types Reference

This document provides a comprehensive reference for all data types used in the EKS Auto Mode SGPP enhancement.

## 📋 **Core Data Types**

### **AutoModeDetector**

Auto Mode detection service.

```go
type AutoModeDetector struct {
    clusterName string
    region      string
    eksClient   *eks.Client
    cache       map[string]*ClusterInfo
    cacheTTL    time.Duration
    mutex       sync.RWMutex
}
```

### **AutoModeENIManager**

Enhanced ENI manager for Auto Mode.

```go
type AutoModeENIManager struct {
    clusterName        string
    region             string
    ec2Client          *ec2.Client
    eniPool            map[string]*ENIInfo
    allocationTimeout  time.Duration
    cleanupInterval    time.Duration
    mutex              sync.RWMutex
}
```

### **SecurityGroupManager**

Security group management service.

```go
type SecurityGroupManager struct {
    clusterName           string
    region                string
    ec2Client             *ec2.Client
    securityGroupCache    map[string]*SecurityGroupInfo
    cacheTTL              time.Duration
    maxSecurityGroups     int
    mutex                 sync.RWMutex
}
```

### **ConfigurationManager**

Configuration management service.

```go
type ConfigurationManager struct {
    config *AutoModeConfig
    mutex  sync.RWMutex
}
```

## 📊 **Data Structures**

### **ClusterInfo**

Cluster information structure.

```go
type ClusterInfo struct {
    Name      string            `json:"name"`
    Mode      ClusterMode       `json:"mode"`
    Region    string            `json:"region"`
    Tags      map[string]string `json:"tags"`
    LastCheck time.Time         `json:"lastCheck"`
}
```

### **ClusterMode**

Cluster mode enumeration.

```go
type ClusterMode string

const (
    ClusterModeAuto     ClusterMode = "auto"
    ClusterModeStandard ClusterMode = "standard"
    ClusterModeUnknown  ClusterMode = "unknown"
)
```

### **ENIInfo**

ENI information structure.

```go
type ENIInfo struct {
    ENIID           string    `json:"eniId"`
    SecurityGroups  []string  `json:"securityGroups"`
    SubnetID        string    `json:"subnetId"`
    PrivateIP       string    `json:"privateIp"`
    AllocationTime  time.Time `json:"allocationTime"`
    Status          ENIStatus `json:"status"`
}
```

### **ENIStatus**

ENI status enumeration.

```go
type ENIStatus string

const (
    ENIStatusAvailable   ENIStatus = "available"
    ENIStatusAllocated   ENIStatus = "allocated"
    ENIStatusAttaching   ENIStatus = "attaching"
    ENIStatusAttached    ENIStatus = "attached"
    ENIStatusDetaching   ENIStatus = "detaching"
    ENIStatusDetached    ENIStatus = "detached"
    ENIStatusDeleting    ENIStatus = "deleting"
    ENIStatusDeleted     ENIStatus = "deleted"
)
```

### **SecurityGroupInfo**

Security group information structure.

```go
type SecurityGroupInfo struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Description string            `json:"description"`
    VPCID       string            `json:"vpcId"`
    Tags        map[string]string `json:"tags"`
    LastCheck   time.Time         `json:"lastCheck"`
}
```

### **PodSpec**

Pod specification structure.

```go
type PodSpec struct {
    Name            string            `json:"name"`
    Namespace       string            `json:"namespace"`
    SecurityGroups  []string          `json:"securityGroups"`
    Annotations     map[string]string `json:"annotations"`
    Labels          map[string]string `json:"labels"`
    UID             string            `json:"uid"`
    CreationTime    time.Time         `json:"creationTime"`
}
```

### **PodNetworkInfo**

Pod network information structure.

```go
type PodNetworkInfo struct {
    PodName         string    `json:"podName"`
    Namespace       string    `json:"namespace"`
    ENIID           string    `json:"eniId"`
    PrivateIP       string    `json:"privateIp"`
    SecurityGroups  []string  `json:"securityGroups"`
    SubnetID        string    `json:"subnetId"`
    AllocationTime  time.Time `json:"allocationTime"`
    Status          PodStatus `json:"status"`
}
```

### **PodStatus**

Pod status enumeration.

```go
type PodStatus string

const (
    PodStatusPending    PodStatus = "pending"
    PodStatusRunning    PodStatus = "running"
    PodStatusSucceeded  PodStatus = "succeeded"
    PodStatusFailed     PodStatus = "failed"
    PodStatusUnknown    PodStatus = "unknown"
)
```

### **AutoModeConfig**

Auto Mode configuration structure.

```go
type AutoModeConfig struct {
    Enabled                     bool          `json:"enabled"`
    ENIPoolSize                 int           `json:"eniPoolSize"`
    MaxSecurityGroupsPerPod     int           `json:"maxSecurityGroupsPerPod"`
    AllocationTimeout           time.Duration `json:"allocationTimeout"`
    CleanupInterval             time.Duration `json:"cleanupInterval"`
    CacheTTL                    time.Duration `json:"cacheTtl"`
    ValidationTimeout           time.Duration `json:"validationTimeout"`
    ClusterName                 string        `json:"clusterName"`
    Region                      string        `json:"region"`
}
```

## 🔧 **Error Types**

### **AutoModeError**

Custom error type for Auto Mode operations.

```go
type AutoModeError struct {
    Code    ErrorCode `json:"code"`
    Message string    `json:"message"`
    Cause   error     `json:"cause,omitempty"`
}

func (e *AutoModeError) Error() string {
    return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
```

### **ErrorCode**

Error code enumeration.

```go
type ErrorCode string

const (
    ErrorCodeInvalidConfiguration ErrorCode = "INVALID_CONFIGURATION"
    ErrorCodeAutoModeDetectionFailed ErrorCode = "AUTO_MODE_DETECTION_FAILED"
    ErrorCodeENIAllocationFailed ErrorCode = "ENI_ALLOCATION_FAILED"
    ErrorCodeSecurityGroupValidationFailed ErrorCode = "SECURITY_GROUP_VALIDATION_FAILED"
    ErrorCodeResourceExhausted ErrorCode = "RESOURCE_EXHAUSTED"
    ErrorCodeTimeout ErrorCode = "TIMEOUT"
    ErrorCodePermissionDenied ErrorCode = "PERMISSION_DENIED"
    ErrorCodeInvalidInput ErrorCode = "INVALID_INPUT"
)
```

## 📝 **Constants**

### **Environment Variables**

```go
const (
    EnvAutoModeSGPPEnabled                = "AUTO_MODE_SGPP_ENABLED"
    EnvClusterName                        = "CLUSTER_NAME"
    EnvAWSRegion                          = "AWS_REGION"
    EnvENIPoolSize                        = "ENI_POOL_SIZE"
    EnvENIAllocationTimeout                = "ENI_ALLOCATION_TIMEOUT"
    EnvENICleanupInterval                 = "ENI_CLEANUP_INTERVAL"
    EnvMaxSecurityGroupsPerPod            = "MAX_SECURITY_GROUPS_PER_POD"
    EnvSecurityGroupCacheTTL              = "SECURITY_GROUP_CACHE_TTL"
    EnvSecurityGroupValidationTimeout     = "SECURITY_GROUP_VALIDATION_TIMEOUT"
    EnvAutoModeCacheTTL                   = "AUTO_MODE_CACHE_TTL"
    EnvAutoModeDetectionTimeout           = "AUTO_MODE_DETECTION_TIMEOUT"
)
```

### **Default Values**

```go
const (
    DefaultENIPoolSize                     = 10
    DefaultMaxSecurityGroupsPerPod         = 5
    DefaultAllocationTimeout               = 30 * time.Second
    DefaultCleanupInterval                 = 5 * time.Minute
    DefaultCacheTTL                        = 5 * time.Minute
    DefaultValidationTimeout               = 10 * time.Second
    DefaultDetectionTimeout                = 10 * time.Second
)
```

### **Pod Annotations**

```go
const (
    PodSecurityGroupAnnotation            = "vpc.amazonaws.com/security-groups"
    PodSecurityGroupLabel                 = "vpc.amazonaws.com/security-groups"
    PodSecurityGroupIDsAnnotation         = "vpc.amazonaws.com/security-group-ids"
)
```

## 🔗 **Related Documentation**

- [Auto Mode Detector API](auto-mode-detector.md)
- [ENI Manager API](eni-manager.md)
- [Security Group Manager API](security-group-manager.md)
- [Configuration Manager API](configuration-manager.md)
- [Error Handling](error-handling.md)
