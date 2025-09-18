# EKS Auto Mode SGPP Enhancement - Architecture Overview

## System Architecture

The EKS Auto Mode Security Groups per Pod (SGPP) enhancement introduces several new components to support fine-grained network security in Auto Mode clusters.

## High-Level Architecture

```mermaid
graph TB
    subgraph "EKS Auto Mode Enhanced Architecture"
        EKS_CONTROL[EKS Control Plane<br/>Auto Managed]
        AUTO_NETWORKING[Auto Networking<br/>Enhanced Model]
        VPC_CNI[VPC CNI Plugin<br/>SGPP Enabled]
        AUTO_NODEGROUP[Auto Node Groups<br/>SGPP Aware]
        
        subgraph "New Components"
            SGPP_MANAGER[SGPP Manager<br/>Auto Mode Compatible]
            ENI_POOL[ENI Pool Manager<br/>Auto Scaling]
            SG_POLICY[Security Group Policy<br/>Auto Assignment]
        end
        
        EKS_CONTROL --> AUTO_NETWORKING
        AUTO_NETWORKING --> VPC_CNI
        AUTO_NODEGROUP --> VPC_CNI
        VPC_CNI --> SGPP_MANAGER
        SGPP_MANAGER --> ENI_POOL
        SGPP_MANAGER --> SG_POLICY
    end
```

## Component Architecture

### 1. Auto Mode Detection Service

**Purpose**: Detect when a cluster is running in Auto Mode

**Key Features**:
- Cluster mode detection via API calls
- Configuration validation
- Fallback mechanisms

**Location**: `pkg/eniconfig/auto_mode_detector.go`

### 2. Enhanced ENI Manager

**Purpose**: Manage ENIs for SGPP in Auto Mode

**Key Features**:
- Pool-based ENI allocation
- Security group assignment
- Auto-scaling ENI pools

**Location**: `pkg/eniconfig/auto_mode_eni.go`

### 3. Security Group Manager

**Purpose**: Assign security groups to pods in Auto Mode

**Key Features**:
- Annotation-based security group specification
- Security group validation
- Fallback to default security groups

**Location**: `pkg/eniconfig/auto_mode_sg.go`

## Data Flow Architecture

### Pod Creation Flow

```mermaid
sequenceDiagram
    participant User
    participant K8s
    participant VPC_CNI
    participant SGPP_MANAGER
    participant ENI_POOL
    participant SG_POLICY
    participant AWS_EC2
    
    User->>K8s: Create Pod with SG annotation
    K8s->>VPC_CNI: Pod creation event
    VPC_CNI->>SGPP_MANAGER: Check Auto Mode
    SGPP_MANAGER->>ENI_POOL: Request ENI
    ENI_POOL->>AWS_EC2: Allocate ENI
    AWS_EC2-->>ENI_POOL: ENI allocated
    ENI_POOL-->>SGPP_MANAGER: ENI available
    SGPP_MANAGER->>SG_POLICY: Get security group
    SG_POLICY-->>SGPP_MANAGER: Security group
    SGPP_MANAGER->>AWS_EC2: Attach security group
    AWS_EC2-->>SGPP_MANAGER: Security group attached
    SGPP_MANAGER-->>VPC_CNI: ENI with SG ready
    VPC_CNI-->>K8s: Pod ready
    K8s-->>User: Pod running
```

### ENI Allocation Flow

```mermaid
flowchart TD
    A[Pod Creation] --> B{Auto Mode?}
    B -->|Yes| C[Use Auto Mode ENI Manager]
    B -->|No| D[Use Standard ENI Manager]
    C --> E[Check ENI Pool]
    E --> F{ENI Available?}
    F -->|Yes| G[Allocate from Pool]
    F -->|No| H[Create New ENI]
    G --> I[Assign Security Group]
    H --> I
    I --> J[Attach to Pod]
    D --> K[Standard Allocation]
    K --> J
```

## Security Architecture

### Security Group Assignment

```mermaid
graph LR
    A[Pod Annotation] --> B[Parse Security Groups]
    B --> C[Validate Security Groups]
    C --> D{Security Group Valid?}
    D -->|Yes| E[Assign Security Group]
    D -->|No| F[Use Default Security Group]
    E --> G[Attach to ENI]
    F --> G
```

### Network Isolation

- **Pod-level isolation**: Each pod can have its own security group
- **Namespace isolation**: Security groups can be scoped to namespaces
- **Tenant isolation**: Multi-tenant security group management

## Performance Architecture

### ENI Pool Management

```mermaid
graph TB
    A[ENI Pool] --> B[Available ENIs]
    A --> C[Allocated ENIs]
    A --> D[Reserved ENIs]
    
    B --> E[Pod Request]
    E --> F[Allocate ENI]
    F --> C
    
    C --> G[Pod Termination]
    G --> H[Return ENI]
    H --> B
    
    D --> I[High Availability]
    I --> B
```

### Caching Strategy

- **Security Group Cache**: Cache security group information
- **ENI Pool Cache**: Cache available ENIs
- **Configuration Cache**: Cache Auto Mode configuration

## Scalability Architecture

### Horizontal Scaling

- **ENI Pool Scaling**: Automatically scale ENI pools based on demand
- **Security Group Scaling**: Support multiple security groups per pod
- **Node Scaling**: Scale nodes to accommodate ENI requirements

### Vertical Scaling

- **Memory Optimization**: Optimize memory usage for ENI management
- **CPU Optimization**: Optimize CPU usage for security group operations
- **Network Optimization**: Optimize network performance

## Monitoring Architecture

### Metrics Collection

- **ENI Metrics**: ENI allocation and deallocation metrics
- **Security Group Metrics**: Security group assignment metrics
- **Performance Metrics**: Performance and latency metrics

### Logging

- **Structured Logging**: JSON-formatted logs
- **Log Levels**: Debug, Info, Warn, Error
- **Log Aggregation**: Centralized log collection

## Deployment Architecture

### Component Deployment

```mermaid
graph TB
    A[VPC CNI Plugin] --> B[Auto Mode Detector]
    A --> C[ENI Manager]
    A --> D[Security Group Manager]
    
    B --> E[Configuration]
    C --> E
    D --> E
    
    E --> F[ConfigMap]
    E --> G[Secrets]
    E --> H[Environment Variables]
```

### Configuration Management

- **ConfigMap**: Configuration for Auto Mode components
- **Secrets**: Sensitive configuration data
- **Environment Variables**: Runtime configuration

## Error Handling Architecture

### Error Types

- **Configuration Errors**: Invalid configuration
- **API Errors**: AWS API errors
- **Network Errors**: Network connectivity issues
- **Resource Errors**: Resource allocation failures

### Error Recovery

- **Retry Logic**: Automatic retry for transient errors
- **Fallback Mechanisms**: Fallback to default configurations
- **Circuit Breakers**: Prevent cascading failures

## Future Architecture Considerations

### Planned Enhancements

- **Multi-Region Support**: Support for multi-region deployments
- **Advanced Security**: Enhanced security features
- **Performance Optimization**: Further performance improvements

### Scalability Improvements

- **Distributed ENI Management**: Distributed ENI pool management
- **Advanced Caching**: More sophisticated caching strategies
- **Load Balancing**: Load balancing for ENI allocation
