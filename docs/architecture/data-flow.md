# Data Flow Documentation

This document describes the data flow and interactions in the EKS Auto Mode SGPP enhancement.

## 📋 **Data Flow Overview**

The EKS Auto Mode SGPP enhancement implements a comprehensive data flow that handles pod lifecycle, ENI allocation, security group assignment, and resource management.

## 🔄 **Main Data Flow**

### **Pod Creation Flow**

```mermaid
graph TD
    A[Pod Creation Request] --> B[Auto Mode Detection]
    B --> C{Is Auto Mode?}
    C -->|Yes| D[Parse Security Group Annotations]
    C -->|No| E[Standard Mode Processing]
    D --> F[Validate Security Groups]
    F --> G{Validation Success?}
    G -->|Yes| H[Allocate ENI from Pool]
    G -->|No| I[Return Validation Error]
    H --> J{ENI Available?}
    J -->|Yes| K[Attach Security Groups to ENI]
    J -->|No| L[Create New ENI]
    L --> M[Add to Pool]
    M --> K
    K --> N[Update Pod Network Config]
    N --> O[Pod Ready]
    E --> O
```

### **Pod Deletion Flow**

```mermaid
graph TD
    A[Pod Deletion Request] --> B[Identify Pod Resources]
    B --> C[Detach Security Groups]
    C --> D[Release ENI]
    D --> E{Return to Pool?}
    E -->|Yes| F[Add ENI to Available Pool]
    E -->|No| G[Delete ENI]
    F --> H[Update Resource Counters]
    G --> H
    H --> I[Cleanup Complete]
```

## 📊 **Component Data Flow**

### **Auto Mode Detection Flow**

```mermaid
sequenceDiagram
    participant Client
    participant Detector
    participant EKS API
    participant Cache
    
    Client->>Detector: IsAutoMode()
    Detector->>Cache: Check Cache
    Cache-->>Detector: Cache Miss
    Detector->>EKS API: DescribeCluster()
    EKS API-->>Detector: Cluster Info
    Detector->>Detector: Parse Mode
    Detector->>Cache: Store Result
    Detector-->>Client: Auto Mode Status
```

### **ENI Allocation Flow**

```mermaid
sequenceDiagram
    participant Manager
    participant Pool
    participant EC2 API
    participant SG Manager
    
    Manager->>Pool: Check Available ENI
    Pool-->>Manager: ENI Available
    Manager->>SG Manager: Get Security Groups
    SG Manager-->>Manager: Security Group IDs
    Manager->>EC2 API: Attach Security Groups
    EC2 API-->>Manager: Success
    Manager->>Pool: Mark ENI as Allocated
    Manager-->>Client: ENI Info
```

### **Security Group Validation Flow**

```mermaid
sequenceDiagram
    participant Manager
    participant Cache
    participant EC2 API
    
    Manager->>Cache: Check Security Group Cache
    Cache-->>Manager: Cache Miss
    Manager->>EC2 API: DescribeSecurityGroups()
    EC2 API-->>Manager: Security Group Info
    Manager->>Manager: Validate Security Group
    Manager->>Cache: Store Validated SG
    Manager-->>Client: Validation Result
```

## 🔧 **Data Structures Flow**

### **Pod Specification Flow**

```go
// Pod creation data flow
type PodCreationFlow struct {
    // Input
    PodSpec *PodSpec
    
    // Processing
    SecurityGroups []string
    ENIInfo       *ENIInfo
    NetworkInfo   *PodNetworkInfo
    
    // Output
    Result *PodNetworkInfo
    Error  error
}
```

### **ENI Allocation Flow**

```go
// ENI allocation data flow
type ENIAllocationFlow struct {
    // Input
    PodSpec        *PodSpec
    SecurityGroups []string
    
    // Processing
    PoolCheck      bool
    ENICreation    bool
    SGAttachment   bool
    
    // Output
    ENIInfo *ENIInfo
    Error   error
}
```

## 🔄 **Resource Management Flow**

### **ENI Pool Management**

```mermaid
graph TD
    A[ENI Pool Manager] --> B[Monitor Pool Size]
    B --> C{Pool Size < Min?}
    C -->|Yes| D[Create New ENI]
    C -->|No| E{Pool Size > Max?}
    E -->|Yes| F[Remove Excess ENI]
    E -->|No| G[Maintain Current Size]
    D --> H[Add to Pool]
    F --> I[Remove from Pool]
    H --> J[Update Metrics]
    I --> J
    G --> J
    J --> K[Continue Monitoring]
```

### **Security Group Cache Management**

```mermaid
graph TD
    A[Security Group Request] --> B[Check Cache]
    B --> C{Cache Hit?}
    C -->|Yes| D{Cache Valid?}
    C -->|No| E[Fetch from AWS]
    D -->|Yes| F[Return Cached Data]
    D -->|No| G[Refresh Cache]
    E --> H[Validate Security Group]
    G --> H
    H --> I[Update Cache]
    I --> J[Return Data]
    F --> J
```

## 📝 **Error Flow Handling**

### **Error Propagation Flow**

```mermaid
graph TD
    A[Operation Error] --> B[Error Classification]
    B --> C{Error Type}
    C -->|Retryable| D[Retry with Backoff]
    C -->|Non-Retryable| E[Return Error]
    C -->|Fallback| F[Fallback to Standard Mode]
    D --> G{Retry Success?}
    G -->|Yes| H[Continue Processing]
    G -->|No| I[Max Retries Reached]
    I --> E
    F --> J[Standard Mode Processing]
    J --> H
    E --> K[Log Error]
    H --> L[Operation Complete]
    K --> L
```

### **Fallback Mechanism Flow**

```mermaid
sequenceDiagram
    participant Client
    participant AutoMode
    participant Standard
    participant Logger
    
    Client->>AutoMode: Process Pod
    AutoMode->>AutoMode: Try Auto Mode
    AutoMode-->>AutoMode: Error
    AutoMode->>Logger: Log Error
    AutoMode->>Standard: Fallback to Standard
    Standard-->>AutoMode: Standard Result
    AutoMode-->>Client: Result
```

## 🔍 **Monitoring and Metrics Flow**

### **Metrics Collection Flow**

```mermaid
graph TD
    A[Operation Start] --> B[Record Start Time]
    B --> C[Execute Operation]
    C --> D[Record End Time]
    D --> E[Calculate Duration]
    E --> F[Update Metrics]
    F --> G[Check Thresholds]
    G --> H{Threshold Exceeded?}
    H -->|Yes| I[Generate Alert]
    H -->|No| J[Continue]
    I --> K[Send Alert]
    J --> L[Operation Complete]
    K --> L
```

## 📊 **Configuration Flow**

### **Configuration Loading Flow**

```mermaid
graph TD
    A[Application Start] --> B[Load Environment Variables]
    B --> C[Parse Configuration]
    C --> D[Validate Configuration]
    D --> E{Validation Success?}
    E -->|Yes| F[Apply Defaults]
    E -->|No| G[Return Error]
    F --> H[Cache Configuration]
    H --> I[Initialize Components]
    I --> J[Application Ready]
    G --> K[Exit Application]
```

## 🔗 **Integration Flow**

### **IPAM Integration Flow**

```mermaid
sequenceDiagram
    participant IPAM
    participant AutoMode
    participant Components
    participant AWS
    
    IPAM->>AutoMode: Process Pod Request
    AutoMode->>Components: Initialize Components
    Components-->>AutoMode: Components Ready
    AutoMode->>Components: Detect Auto Mode
    Components->>AWS: Check Cluster Mode
    AWS-->>Components: Mode Info
    Components-->>AutoMode: Auto Mode Detected
    AutoMode->>Components: Allocate ENI
    Components->>AWS: Create/Attach ENI
    AWS-->>Components: ENI Info
    Components-->>AutoMode: ENI Allocated
    AutoMode-->>IPAM: Pod Network Info
```

## 🔗 **Related Documentation**

- [Architecture Overview](overview.md)
- [Components](components.md)
- [Auto Mode Detector API](../api/auto-mode-detector.md)
- [ENI Manager API](../api/eni-manager.md)
- [Security Group Manager API](../api/security-group-manager.md)
