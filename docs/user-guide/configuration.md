# Configuration Guide

This document provides comprehensive configuration options and settings for the EKS Auto Mode SGPP enhancement.

## 📋 **Configuration Overview**

The EKS Auto Mode SGPP enhancement supports extensive configuration through environment variables, with sensible defaults for all settings.

## ⚙️ **Core Configuration**

### **Auto Mode SGPP Enablement**

```bash
# Enable Auto Mode SGPP functionality
export AUTO_MODE_SGPP_ENABLED="true"
```

**Default**: `false`  
**Description**: Enables Security Groups per Pod functionality in Auto Mode clusters.

### **Cluster Information**

```bash
# EKS cluster name (required)
export CLUSTER_NAME="my-auto-mode-cluster"

# AWS region (required)
export AWS_REGION="us-west-2"
```

**Required**: Yes  
**Description**: Basic cluster information required for Auto Mode detection and AWS API calls.

## 🔧 **ENI Management Configuration**

### **ENI Pool Settings**

```bash
# Maximum number of ENIs in the pool
export ENI_POOL_SIZE="20"

# Timeout for ENI allocation operations
export ENI_ALLOCATION_TIMEOUT="45s"

# Interval for ENI cleanup operations
export ENI_CLEANUP_INTERVAL="3m"
```

**Defaults**:
- `ENI_POOL_SIZE`: `10`
- `ENI_ALLOCATION_TIMEOUT`: `30s`
- `ENI_CLEANUP_INTERVAL`: `5m`

**Description**: Controls ENI pool management and allocation behavior.

### **ENI Pool Scaling**

```bash
# Minimum ENIs to maintain in pool
export ENI_POOL_MIN_SIZE="5"

# Maximum ENIs allowed in pool
export ENI_POOL_MAX_SIZE="50"

# Scale-up threshold (percentage of pool used)
export ENI_POOL_SCALE_UP_THRESHOLD="80"

# Scale-down threshold (percentage of pool used)
export ENI_POOL_SCALE_DOWN_THRESHOLD="20"
```

**Defaults**:
- `ENI_POOL_MIN_SIZE`: `5`
- `ENI_POOL_MAX_SIZE`: `50`
- `ENI_POOL_SCALE_UP_THRESHOLD`: `80`
- `ENI_POOL_SCALE_DOWN_THRESHOLD`: `20`

## 🛡️ **Security Group Configuration**

### **Security Group Limits**

```bash
# Maximum security groups per pod
export MAX_SECURITY_GROUPS_PER_POD="5"

# Cache TTL for security group information
export SECURITY_GROUP_CACHE_TTL="10m"

# Timeout for security group validation
export SECURITY_GROUP_VALIDATION_TIMEOUT="15s"
```

**Defaults**:
- `MAX_SECURITY_GROUPS_PER_POD`: `5`
- `SECURITY_GROUP_CACHE_TTL`: `5m`
- `SECURITY_GROUP_VALIDATION_TIMEOUT`: `10s`

**Description**: Controls security group management and validation behavior.

### **Security Group Validation**

```bash
# Enable strict security group validation
export SECURITY_GROUP_STRICT_VALIDATION="true"

# Validate security groups on pod creation
export SECURITY_GROUP_VALIDATE_ON_CREATE="true"

# Cache security group validation results
export SECURITY_GROUP_CACHE_VALIDATION="true"
```

**Defaults**:
- `SECURITY_GROUP_STRICT_VALIDATION`: `true`
- `SECURITY_GROUP_VALIDATE_ON_CREATE`: `true`
- `SECURITY_GROUP_CACHE_VALIDATION`: `true`

## 🔍 **Auto Mode Detection Configuration**

### **Detection Settings**

```bash
# Cache TTL for Auto Mode detection
export AUTO_MODE_CACHE_TTL="5m"

# Timeout for Auto Mode detection
export AUTO_MODE_DETECTION_TIMEOUT="10s"

# Enable Auto Mode detection retry
export AUTO_MODE_DETECTION_RETRY="true"

# Maximum retry attempts for detection
export AUTO_MODE_DETECTION_MAX_RETRIES="3"
```

**Defaults**:
- `AUTO_MODE_CACHE_TTL`: `5m`
- `AUTO_MODE_DETECTION_TIMEOUT`: `10s`
- `AUTO_MODE_DETECTION_RETRY`: `true`
- `AUTO_MODE_DETECTION_MAX_RETRIES`: `3`

## 📊 **Performance Configuration**

### **Caching Settings**

```bash
# Enable component-level caching
export ENABLE_COMPONENT_CACHING="true"

# Global cache TTL
export GLOBAL_CACHE_TTL="5m"

# Cache cleanup interval
export CACHE_CLEANUP_INTERVAL="1m"

# Maximum cache size
export MAX_CACHE_SIZE="1000"
```

**Defaults**:
- `ENABLE_COMPONENT_CACHING`: `true`
- `GLOBAL_CACHE_TTL`: `5m`
- `CACHE_CLEANUP_INTERVAL`: `1m`
- `MAX_CACHE_SIZE`: `1000`

### **Concurrency Settings**

```bash
# Maximum concurrent ENI allocations
export MAX_CONCURRENT_ENI_ALLOCATIONS="10"

# Maximum concurrent security group validations
export MAX_CONCURRENT_SG_VALIDATIONS="20"

# Worker pool size for background operations
export WORKER_POOL_SIZE="5"
```

**Defaults**:
- `MAX_CONCURRENT_ENI_ALLOCATIONS`: `10`
- `MAX_CONCURRENT_SG_VALIDATIONS`: `20`
- `WORKER_POOL_SIZE`: `5`

## 🔄 **Retry and Timeout Configuration**

### **Retry Settings**

```bash
# Enable retry for failed operations
export ENABLE_OPERATION_RETRY="true"

# Maximum retry attempts
export MAX_RETRY_ATTEMPTS="3"

# Base delay for retry backoff
export RETRY_BASE_DELAY="1s"

# Maximum delay for retry backoff
export RETRY_MAX_DELAY="30s"
```

**Defaults**:
- `ENABLE_OPERATION_RETRY`: `true`
- `MAX_RETRY_ATTEMPTS`: `3`
- `RETRY_BASE_DELAY`: `1s`
- `RETRY_MAX_DELAY`: `30s`

### **Timeout Settings**

```bash
# Global operation timeout
export GLOBAL_OPERATION_TIMEOUT="60s"

# AWS API call timeout
export AWS_API_TIMEOUT="30s"

# Kubernetes API call timeout
export K8S_API_TIMEOUT="15s"
```

**Defaults**:
- `GLOBAL_OPERATION_TIMEOUT`: `60s`
- `AWS_API_TIMEOUT`: `30s`
- `K8S_API_TIMEOUT`: `15s`

## 📝 **Logging Configuration**

### **Log Level Settings**

```bash
# Log level (debug, info, warn, error)
export LOG_LEVEL="info"

# Enable structured logging
export ENABLE_STRUCTURED_LOGGING="true"

# Log format (json, text)
export LOG_FORMAT="json"
```

**Defaults**:
- `LOG_LEVEL`: `info`
- `ENABLE_STRUCTURED_LOGGING`: `true`
- `LOG_FORMAT`: `json`

### **Log Output Settings**

```bash
# Log output destination (stdout, stderr, file)
export LOG_OUTPUT="stdout"

# Log file path (if using file output)
export LOG_FILE_PATH="/var/log/auto-mode-sgpp.log"

# Enable log rotation
export ENABLE_LOG_ROTATION="true"

# Maximum log file size
export MAX_LOG_FILE_SIZE="100MB"
```

**Defaults**:
- `LOG_OUTPUT`: `stdout`
- `LOG_FILE_PATH`: `/var/log/auto-mode-sgpp.log`
- `ENABLE_LOG_ROTATION`: `true`
- `MAX_LOG_FILE_SIZE`: `100MB`

## 🔧 **Advanced Configuration**

### **Feature Flags**

```bash
# Enable experimental features
export ENABLE_EXPERIMENTAL_FEATURES="false"

# Enable debug mode
export ENABLE_DEBUG_MODE="false"

# Enable metrics collection
export ENABLE_METRICS="true"

# Enable health checks
export ENABLE_HEALTH_CHECKS="true"
```

**Defaults**:
- `ENABLE_EXPERIMENTAL_FEATURES`: `false`
- `ENABLE_DEBUG_MODE`: `false`
- `ENABLE_METRICS`: `true`
- `ENABLE_HEALTH_CHECKS`: `true`

### **Integration Settings**

```bash
# Enable IPAM integration
export ENABLE_IPAM_INTEGRATION="true"

# Enable Kubernetes integration
export ENABLE_K8S_INTEGRATION="true"

# Enable AWS integration
export ENABLE_AWS_INTEGRATION="true"
```

**Defaults**:
- `ENABLE_IPAM_INTEGRATION`: `true`
- `ENABLE_K8S_INTEGRATION`: `true`
- `ENABLE_AWS_INTEGRATION`: `true`

## 📋 **Configuration Validation**

### **Required Configuration**

The following configuration is required for the enhancement to function:

```bash
# Required settings
export AUTO_MODE_SGPP_ENABLED="true"
export CLUSTER_NAME="your-cluster-name"
export AWS_REGION="your-aws-region"
```

### **Configuration Validation**

The system validates configuration on startup:

```go
// Example validation
func validateConfiguration() error {
    if os.Getenv("CLUSTER_NAME") == "" {
        return fmt.Errorf("CLUSTER_NAME is required")
    }
    
    if os.Getenv("AWS_REGION") == "" {
        return fmt.Errorf("AWS_REGION is required")
    }
    
    return nil
}
```

## 🔄 **Configuration Hot Reloading**

### **Hot Reload Support**

```bash
# Enable configuration hot reloading
export ENABLE_CONFIG_HOT_RELOAD="true"

# Configuration reload interval
export CONFIG_RELOAD_INTERVAL="30s"

# Reload on configuration change
export RELOAD_ON_CONFIG_CHANGE="true"
```

**Defaults**:
- `ENABLE_CONFIG_HOT_RELOAD`: `true`
- `CONFIG_RELOAD_INTERVAL`: `30s`
- `RELOAD_ON_CONFIG_CHANGE`: `true`

## 📝 **Configuration Examples**

### **Development Configuration**

```bash
# Development settings
export AUTO_MODE_SGPP_ENABLED="true"
export CLUSTER_NAME="dev-cluster"
export AWS_REGION="us-west-2"
export LOG_LEVEL="debug"
export ENABLE_DEBUG_MODE="true"
export ENI_POOL_SIZE="5"
export MAX_SECURITY_GROUPS_PER_POD="3"
```

### **Production Configuration**

```bash
# Production settings
export AUTO_MODE_SGPP_ENABLED="true"
export CLUSTER_NAME="prod-cluster"
export AWS_REGION="us-west-2"
export LOG_LEVEL="info"
export ENABLE_DEBUG_MODE="false"
export ENI_POOL_SIZE="50"
export MAX_SECURITY_GROUPS_PER_POD="5"
export ENABLE_METRICS="true"
export ENABLE_HEALTH_CHECKS="true"
```

### **High Performance Configuration**

```bash
# High performance settings
export AUTO_MODE_SGPP_ENABLED="true"
export CLUSTER_NAME="perf-cluster"
export AWS_REGION="us-west-2"
export ENI_POOL_SIZE="100"
export MAX_CONCURRENT_ENI_ALLOCATIONS="20"
export MAX_CONCURRENT_SG_VALIDATIONS="50"
export WORKER_POOL_SIZE="10"
export CACHE_TTL="10m"
```

## 🔗 **Related Documentation**

- [Getting Started](getting-started.md)
- [Troubleshooting](troubleshooting.md)
- [Configuration Manager API](../api/configuration-manager.md)
- [Architecture Overview](../architecture/overview.md)
