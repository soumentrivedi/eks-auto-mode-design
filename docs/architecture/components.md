# Architecture Components

This document provides detailed descriptions of all components in the EKS Auto Mode SGPP enhancement architecture.

## 📋 **Component Overview**

The EKS Auto Mode SGPP enhancement consists of several key components that work together to provide Security Groups per Pod functionality in Auto Mode clusters.

## 🏗️ **Core Components**

### **Auto Mode Detector**

**Purpose**: Detects if an EKS cluster is running in Auto Mode and provides cluster information.

**Key Responsibilities**:
- Cluster mode detection via AWS EKS API
- Cluster metadata retrieval and caching
- Configuration validation for Auto Mode

**Key Features**:
- Automatic detection of Auto Mode clusters
- Caching for improved performance
- Fallback mechanisms for detection failures

**Integration Points**:
- AWS EKS API for cluster information
- Configuration Manager for settings
- Other components for mode-specific behavior

### **Enhanced ENI Manager**

**Purpose**: Manages Elastic Network Interfaces (ENIs) with pool-based allocation for Auto Mode clusters.

**Key Responsibilities**:
- Pre-allocation of ENIs in pools
- ENI lifecycle management (create, attach, detach, delete)
- Security group assignment to ENIs
- Resource optimization and cleanup

**Key Features**:
- Pool-based ENI allocation for faster pod startup
- Automatic scaling based on demand
- Periodic cleanup of unused ENIs
- Support for multiple security groups per ENI

**Integration Points**:
- AWS EC2 API for ENI operations
- Security Group Manager for security group assignment
- IPAM context for integration with existing CNI

### **Security Group Manager**

**Purpose**: Manages security groups for pods with annotation-based specification.

**Key Responsibilities**:
- Parsing pod annotations and labels for security groups
- Validation of security group IDs
- Caching of security group information
- Support for multiple security groups per pod (up to 5)

**Key Features**:
- Annotation-based security group specification
- Multiple security group support
- Security group validation and caching
- Integration with AWS EC2 security groups

**Integration Points**:
- AWS EC2 API for security group operations
- Pod lifecycle management
- ENI Manager for security group attachment

### **Configuration Manager**

**Purpose**: Manages Auto Mode specific configuration and environment variables.

**Key Responsibilities**:
- Environment variable parsing and validation
- Default value management
- Configuration caching and hot-reloading
- Integration with existing VPC CNI configuration

**Key Features**:
- Environment variable configuration
- Default value fallbacks
- Configuration validation
- Hot-reloading support

**Integration Points**:
- All components for configuration access
- Environment variables for settings
- Existing VPC CNI configuration

## 🔧 **Integration Components**

### **Auto Mode Manager**

**Purpose**: Main integration component that coordinates all Auto Mode functionality.

**Key Responsibilities**:
- Orchestrating Auto Mode operations
- Managing component lifecycle
- Error handling and fallback mechanisms
- Integration with IPAM context

**Key Features**:
- Component coordination
- Error handling and recovery
- Fallback to standard mode
- Integration with existing CNI

**Integration Points**:
- All core components
- IPAM context for CNI integration
- Kubernetes API for pod events

### **IPAM Integration**

**Purpose**: Integrates Auto Mode functionality with the existing VPC CNI IPAM context.

**Key Responsibilities**:
- Extending IPAM context with Auto Mode support
- Maintaining backward compatibility
- Pod lifecycle integration
- Resource management

**Key Features**:
- Seamless integration with existing IPAM
- Backward compatibility
- Pod lifecycle management
- Resource cleanup

**Integration Points**:
- Existing VPC CNI IPAM context
- Pod lifecycle events
- Kubernetes API

## 📊 **Data Flow Components**

### **Pod Lifecycle Manager**

**Purpose**: Manages the complete pod lifecycle with Auto Mode SGPP support.

**Key Responsibilities**:
- Pod creation and deletion workflows
- Security group assignment and validation
- ENI lifecycle management
- Integration with Kubernetes pod events

**Key Features**:
- Complete pod lifecycle management
- Security group assignment
- ENI management
- Event handling

**Integration Points**:
- Kubernetes API for pod events
- All core components
- IPAM context

### **Resource Pool Manager**

**Purpose**: Manages resource pools for efficient resource utilization.

**Key Responsibilities**:
- ENI pool management
- Resource allocation and deallocation
- Pool scaling and optimization
- Resource monitoring

**Key Features**:
- Pool-based resource management
- Automatic scaling
- Resource optimization
- Monitoring and metrics

**Integration Points**:
- ENI Manager for ENI operations
- Configuration Manager for settings
- Monitoring systems

## 🔄 **Supporting Components**

### **Cache Manager**

**Purpose**: Provides caching functionality for improved performance.

**Key Responsibilities**:
- Caching cluster information
- Caching security group information
- Cache invalidation and refresh
- Performance optimization

**Key Features**:
- Multi-level caching
- TTL-based expiration
- Cache invalidation
- Performance monitoring

**Integration Points**:
- All components for caching
- Configuration Manager for cache settings

### **Error Handler**

**Purpose**: Provides comprehensive error handling and recovery mechanisms.

**Key Responsibilities**:
- Error detection and classification
- Error recovery and fallback
- Error logging and monitoring
- Graceful degradation

**Key Features**:
- Structured error handling
- Error recovery mechanisms
- Fallback strategies
- Comprehensive logging

**Integration Points**:
- All components for error handling
- Logging systems
- Monitoring systems

## 🔗 **Component Interactions**

### **Component Dependencies**

```
Configuration Manager
    ↓
Auto Mode Detector
    ↓
Security Group Manager ← → ENI Manager
    ↓
Auto Mode Manager
    ↓
IPAM Integration
```

### **Data Flow**

1. **Configuration**: Configuration Manager provides settings to all components
2. **Detection**: Auto Mode Detector identifies cluster mode
3. **Validation**: Security Group Manager validates security groups
4. **Allocation**: ENI Manager allocates ENIs with security groups
5. **Integration**: Auto Mode Manager coordinates operations
6. **IPAM**: IPAM Integration provides CNI integration

## 📝 **Component Configuration**

Each component can be configured through environment variables:

### **Auto Mode Detector**
- `AUTO_MODE_CACHE_TTL`: Cache TTL for cluster information
- `AUTO_MODE_DETECTION_TIMEOUT`: Timeout for detection operations

### **ENI Manager**
- `ENI_POOL_SIZE`: Maximum ENIs in pool
- `ENI_ALLOCATION_TIMEOUT`: Timeout for ENI allocation
- `ENI_CLEANUP_INTERVAL`: Interval for ENI cleanup

### **Security Group Manager**
- `MAX_SECURITY_GROUPS_PER_POD`: Maximum security groups per pod
- `SECURITY_GROUP_CACHE_TTL`: Cache TTL for security group info
- `SECURITY_GROUP_VALIDATION_TIMEOUT`: Timeout for validation

### **Configuration Manager**
- `AUTO_MODE_SGPP_ENABLED`: Enable Auto Mode SGPP
- `CLUSTER_NAME`: EKS cluster name
- `AWS_REGION`: AWS region

## 🔗 **Related Documentation**

- [Architecture Overview](overview.md)
- [Data Flow](data-flow.md)
- [Auto Mode Detector API](../api/auto-mode-detector.md)
- [ENI Manager API](../api/eni-manager.md)
- [Security Group Manager API](../api/security-group-manager.md)
