# Auto Mode Detector API

The Auto Mode Detector provides APIs for detecting EKS clusters running in Auto Mode and retrieving cluster information.

## 📋 **API Overview**

The Auto Mode Detector is responsible for:
- Detecting if an EKS cluster is running in Auto Mode
- Caching cluster information for performance
- Providing cluster metadata and configuration

## 🔧 **Core APIs**

### **NewAutoModeDetector**

Creates a new Auto Mode detector instance.

```go
func NewAutoModeDetector(clusterName, region string) (*AutoModeDetector, error)
```

**Parameters:**
- `clusterName` (string): EKS cluster name
- `region` (string): AWS region where the cluster is located

**Returns:**
- `*AutoModeDetector`: Detector instance
- `error`: Error if initialization fails

**Example:**
```go
detector, err := automode.NewAutoModeDetector("my-cluster", "us-west-2")
if err != nil {
    log.Fatal(err)
}
```

### **IsAutoMode**

Checks if the cluster is running in Auto Mode.

```go
func (d *AutoModeDetector) IsAutoMode(ctx context.Context) (bool, error)
```

**Parameters:**
- `ctx` (context.Context): Context for the operation

**Returns:**
- `bool`: True if cluster is in Auto Mode, false otherwise
- `error`: Error if detection fails

**Example:**
```go
isAutoMode, err := detector.IsAutoMode(ctx)
if err != nil {
    return err
}

if isAutoMode {
    log.Info("Cluster is running in Auto Mode")
}
```

### **GetClusterInfo**

Retrieves detailed cluster information.

```go
func (d *AutoModeDetector) GetClusterInfo(ctx context.Context) (*ClusterInfo, error)
```

**Parameters:**
- `ctx` (context.Context): Context for the operation

**Returns:**
- `*ClusterInfo`: Cluster information structure
- `error`: Error if retrieval fails

**Example:**
```go
clusterInfo, err := detector.GetClusterInfo(ctx)
if err != nil {
    return err
}

log.Infof("Cluster %s is in %s mode", clusterInfo.Name, clusterInfo.Mode)
```

## 📊 **Data Structures**

### **ClusterInfo**

Contains detailed information about the EKS cluster.

```go
type ClusterInfo struct {
    Name      string            `json:"name"`
    Mode      ClusterMode       `json:"mode"`
    Region    string            `json:"region"`
    Tags      map[string]string `json:"tags"`
    LastCheck time.Time         `json:"lastCheck"`
}
```

**Fields:**
- `Name` (string): Cluster name
- `Mode` (ClusterMode): Cluster mode (auto, standard, unknown)
- `Region` (string): AWS region
- `Tags` (map[string]string): Cluster tags
- `LastCheck` (time.Time): Last check timestamp

### **ClusterMode**

Represents the mode of an EKS cluster.

```go
type ClusterMode string

const (
    ClusterModeAuto     ClusterMode = "auto"
    ClusterModeStandard ClusterMode = "standard"
    ClusterModeUnknown  ClusterMode = "unknown"
)
```

## ⚙️ **Configuration**

The Auto Mode Detector uses the following environment variables:

- `CLUSTER_NAME`: EKS cluster name (required)
- `AWS_REGION`: AWS region (required)

## 🔄 **Caching**

The detector implements caching for improved performance:
- Cluster information is cached for 5 minutes by default
- Cache TTL can be configured via environment variables
- Cache is automatically refreshed when expired

## 🐛 **Error Handling**

Common error scenarios:

- **Missing Configuration**: Returns error if required environment variables are not set
- **AWS API Errors**: Returns wrapped AWS SDK errors
- **Network Issues**: Returns connection timeout errors
- **Invalid Cluster**: Returns error if cluster is not found

## 📝 **Examples**

### **Basic Usage**

```go
package main

import (
    "context"
    "log"
    
    "github.com/aws/amazon-vpc-cni-k8s/pkg/automode"
)

func main() {
    // Create detector
    detector, err := automode.NewAutoModeDetector("my-cluster", "us-west-2")
    if err != nil {
        log.Fatal(err)
    }
    
    // Check Auto Mode
    ctx := context.Background()
    isAutoMode, err := detector.IsAutoMode(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    if isAutoMode {
        log.Info("Using Auto Mode SGPP features")
    } else {
        log.Info("Using standard mode")
    }
}
```

### **Advanced Usage with Caching**

```go
// Get cluster info (uses cache if available)
clusterInfo, err := detector.GetClusterInfo(ctx)
if err != nil {
    return err
}

// Check specific tags
if clusterInfo.Tags["eks.amazonaws.com/mode"] == "auto" {
    log.Info("Confirmed Auto Mode via tags")
}
```

## 🔗 **Related APIs**

- [ENI Manager API](eni-manager.md)
- [Security Group Manager API](security-group-manager.md)
- [Configuration Manager API](configuration-manager.md)
