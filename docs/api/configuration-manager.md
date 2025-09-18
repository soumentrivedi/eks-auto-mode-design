# Configuration Manager API

The Configuration Manager provides APIs for managing Auto Mode specific configuration and environment variables.

## 📋 **API Overview**

The Configuration Manager is responsible for:
- Environment variable parsing and validation
- Default value management
- Configuration caching and hot-reloading
- Integration with existing VPC CNI configuration

## 🔧 **Core APIs**

### **NewConfigurationManager**

Creates a new Configuration Manager instance.

```go
func NewConfigurationManager() *ConfigurationManager
```

**Returns:**
- `*ConfigurationManager`: Manager instance

### **GetAutoModeConfig**

Retrieves the complete Auto Mode configuration.

```go
func (m *ConfigurationManager) GetAutoModeConfig() *AutoModeConfig
```

**Returns:**
- `*AutoModeConfig`: Complete configuration object

### **IsAutoModeEnabled**

Checks if Auto Mode SGPP is enabled.

```go
func (m *ConfigurationManager) IsAutoModeEnabled() bool
```

**Returns:**
- `bool`: True if Auto Mode SGPP is enabled

### **GetENIPoolSize**

Gets the ENI pool size configuration.

```go
func (m *ConfigurationManager) GetENIPoolSize() int
```

**Returns:**
- `int`: ENI pool size

### **GetMaxSecurityGroupsPerPod**

Gets the maximum security groups per pod.

```go
func (m *ConfigurationManager) GetMaxSecurityGroupsPerPod() int
```

**Returns:**
- `int`: Maximum security groups per pod

## 📊 **Data Structures**

### **AutoModeConfig**

Contains all Auto Mode configuration settings.

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

## ⚙️ **Environment Variables**

The Configuration Manager supports the following environment variables:

### **Core Configuration**
- `AUTO_MODE_SGPP_ENABLED`: Enable Auto Mode SGPP (default: "false")
- `CLUSTER_NAME`: EKS cluster name (required)
- `AWS_REGION`: AWS region (required)

### **ENI Management**
- `ENI_POOL_SIZE`: Maximum ENIs in pool (default: "10")
- `ENI_ALLOCATION_TIMEOUT`: ENI allocation timeout (default: "30s")
- `ENI_CLEANUP_INTERVAL`: ENI cleanup interval (default: "5m")

### **Security Group Management**
- `MAX_SECURITY_GROUPS_PER_POD`: Max security groups per pod (default: "5")
- `SECURITY_GROUP_CACHE_TTL`: Security group cache TTL (default: "5m")
- `SECURITY_GROUP_VALIDATION_TIMEOUT`: Validation timeout (default: "10s")

### **Performance Tuning**
- `AUTO_MODE_CACHE_TTL`: Auto Mode detection cache TTL (default: "5m")
- `AUTO_MODE_DETECTION_TIMEOUT`: Detection timeout (default: "10s")

## 🔄 **Configuration Loading**

The Configuration Manager loads configuration in the following order:

1. **Environment Variables**: Primary source of configuration
2. **Default Values**: Fallback values for missing configuration
3. **Validation**: Ensures all required values are present
4. **Caching**: Caches configuration for performance

## 🐛 **Error Handling**

Common error scenarios:
- **Missing Required Config**: Returns error if required environment variables are missing
- **Invalid Values**: Returns error for invalid configuration values
- **Parse Errors**: Returns error if values cannot be parsed

## 📝 **Examples**

### **Basic Configuration Usage**

```go
package main

import (
    "log"
    
    "github.com/aws/amazon-vpc-cni-k8s/pkg/automode"
)

func main() {
    // Create configuration manager
    configManager := automode.NewConfigurationManager()
    
    // Check if Auto Mode is enabled
    if configManager.IsAutoModeEnabled() {
        log.Info("Auto Mode SGPP is enabled")
        
        // Get configuration
        config := configManager.GetAutoModeConfig()
        log.Printf("ENI Pool Size: %d", config.ENIPoolSize)
        log.Printf("Max Security Groups per Pod: %d", config.MaxSecurityGroupsPerPod)
    }
}
```

### **Environment Variable Setup**

```bash
# Core configuration
export AUTO_MODE_SGPP_ENABLED="true"
export CLUSTER_NAME="my-auto-mode-cluster"
export AWS_REGION="us-west-2"

# ENI management
export ENI_POOL_SIZE="20"
export ENI_ALLOCATION_TIMEOUT="45s"
export ENI_CLEANUP_INTERVAL="3m"

# Security group management
export MAX_SECURITY_GROUPS_PER_POD="3"
export SECURITY_GROUP_CACHE_TTL="10m"
export SECURITY_GROUP_VALIDATION_TIMEOUT="15s"
```

### **Configuration Validation**

```go
// Validate configuration before use
config := configManager.GetAutoModeConfig()

if config.ClusterName == "" {
    return fmt.Errorf("CLUSTER_NAME is required")
}

if config.Region == "" {
    return fmt.Errorf("AWS_REGION is required")
}

if config.ENIPoolSize <= 0 {
    return fmt.Errorf("ENI_POOL_SIZE must be positive")
}
```

## 🔗 **Related APIs**

- [Auto Mode Detector API](auto-mode-detector.md)
- [ENI Manager API](eni-manager.md)
- [Security Group Manager API](security-group-manager.md)
