# EKS Auto Mode SGPP Enhancement - Prototype Implementation

## Overview

This document outlines the prototype implementation for adding Security Groups per Pod (SGPP) support to EKS Auto Mode.

## Prototype Components

### 1. Auto Mode Detection Service

```go
// pkg/eniconfig/auto_mode_detector.go
package eniconfig

import (
    "context"
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go-v2/service/eks"
    "github.com/aws/aws-sdk-go-v2/config"
)

type AutoModeDetector struct {
    clusterName string
    region      string
    eksClient   *eks.Client
}

type ClusterInfo struct {
    Name string
    Mode string
    Tags map[string]string
}

func NewAutoModeDetector(clusterName, region string) (*AutoModeDetector, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        return nil, fmt.Errorf("failed to load AWS config: %v", err)
    }
    
    return &AutoModeDetector{
        clusterName: clusterName,
        region:      region,
        eksClient:   eks.NewFromConfig(cfg),
    }, nil
}

func (d *AutoModeDetector) IsAutoMode() (bool, error) {
    clusterInfo, err := d.getClusterInfo()
    if err != nil {
        return false, fmt.Errorf("failed to get cluster info: %v", err)
    }
    
    // Check if cluster is in Auto Mode
    // This could be determined by checking cluster tags or API responses
    return clusterInfo.Mode == "AUTO", nil
}

func (d *AutoModeDetector) getClusterInfo() (*ClusterInfo, error) {
    // Implementation to get cluster information
    // This would use AWS EKS API to check cluster configuration
    return &ClusterInfo{
        Name: d.clusterName,
        Mode: "AUTO", // This would be determined by API call
        Tags: make(map[string]string),
    }, nil
}
```

### 2. Enhanced ENI Manager for Auto Mode

```go
// pkg/eniconfig/auto_mode_eni.go
package eniconfig

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type ENI struct {
    ID           string
    SubnetID     string
    SecurityGroups []string
    Status       string
    AllocatedAt  time.Time
}

type ENIPool struct {
    availableENIs []*ENI
    allocatedENIs map[string]*ENI
    mutex        sync.RWMutex
}

type PodSpec struct {
    Name        string
    Namespace   string
    Annotations map[string]string
    Labels      map[string]string
}

type AutoModeENIManager struct {
    detector     *AutoModeDetector
    eniPool      *ENIPool
    sgManager    *SecurityGroupManager
    mutex        sync.RWMutex
}

func NewAutoModeENIManager(clusterName, region string) (*AutoModeENIManager, error) {
    detector, err := NewAutoModeDetector(clusterName, region)
    if err != nil {
        return nil, fmt.Errorf("failed to create Auto Mode detector: %v", err)
    }
    
    eniPool := &ENIPool{
        availableENIs: make([]*ENI, 0),
        allocatedENIs: make(map[string]*ENI),
    }
    
    sgManager, err := NewSecurityGroupManager(region)
    if err != nil {
        return nil, fmt.Errorf("failed to create security group manager: %v", err)
    }
    
    return &AutoModeENIManager{
        detector:  detector,
        eniPool:   eniPool,
        sgManager: sgManager,
    }, nil
}

func (m *AutoModeENIManager) AllocateENIForPod(podSpec *PodSpec) (*ENI, error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    // Check if Auto Mode is enabled
    isAutoMode, err := m.detector.IsAutoMode()
    if err != nil {
        return nil, fmt.Errorf("failed to detect Auto Mode: %v", err)
    }
    
    if !isAutoMode {
        // Use standard ENI allocation
        return m.standardENIAllocation(podSpec)
    }
    
    // Use Auto Mode compatible ENI allocation
    return m.autoModeENIAllocation(podSpec)
}

func (m *AutoModeENIManager) autoModeENIAllocation(podSpec *PodSpec) (*ENI, error) {
    // Get available ENI from pool
    eni, err := m.eniPool.GetAvailableENI()
    if err != nil {
        return nil, fmt.Errorf("failed to get ENI from pool: %v", err)
    }
    
    // Assign security group based on pod annotations
    securityGroup, err := m.sgManager.GetSecurityGroupForPod(podSpec)
    if err != nil {
        return nil, fmt.Errorf("failed to get security group: %v", err)
    }
    
    // Attach security group to ENI
    err = m.sgManager.AttachSecurityGroupToENI(eni.ID, securityGroup.ID)
    if err != nil {
        return nil, fmt.Errorf("failed to attach security group: %v", err)
    }
    
    return eni, nil
}

func (m *AutoModeENIManager) standardENIAllocation(podSpec *PodSpec) (*ENI, error) {
    // Standard ENI allocation logic
    return &ENI{
        ID: "eni-standard",
        Status: "available",
    }, nil
}
```

### 3. Security Group Manager for Auto Mode

```go
// pkg/eniconfig/auto_mode_sg.go
package eniconfig

import (
    "context"
    "fmt"
    "strings"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
    "github.com/aws/aws-sdk-go-v2/config"
)

type SecurityGroup struct {
    ID          string
    Name        string
    Description string
    VpcId       string
}

type SecurityGroupManager struct {
    ec2Client    *ec2.Client
    autoMode     bool
    region       string
}

func NewSecurityGroupManager(region string) (*SecurityGroupManager, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        return nil, fmt.Errorf("failed to load AWS config: %v", err)
    }
    
    return &SecurityGroupManager{
        ec2Client: ec2.NewFromConfig(cfg),
        autoMode:  true,
        region:    region,
    }, nil
}

func (sgm *SecurityGroupManager) GetSecurityGroupForPod(podSpec *PodSpec) (*SecurityGroup, error) {
    // Check pod annotations for security group specification
    sgAnnotation := podSpec.Annotations["vpc.amazonaws.com/security-groups"]
    if sgAnnotation == "" {
        // Use default security group for Auto Mode
        return sgm.getDefaultSecurityGroup()
    }
    
    // Parse security group IDs from annotation
    sgIDs := strings.Split(sgAnnotation, ",")
    
    // Validate security groups exist and are accessible
    for _, sgID := range sgIDs {
        sgID = strings.TrimSpace(sgID)
        if err := sgm.validateSecurityGroup(sgID); err != nil {
            return nil, fmt.Errorf("invalid security group %s: %v", sgID, err)
        }
    }
    
    // Return the first valid security group
    return &SecurityGroup{ID: strings.TrimSpace(sgIDs[0])}, nil
}

func (sgm *SecurityGroupManager) validateSecurityGroup(sgID string) error {
    // Validate that the security group exists and is accessible
    // This would use EC2 API to check security group existence
    return nil
}

func (sgm *SecurityGroupManager) getDefaultSecurityGroup() (*SecurityGroup, error) {
    // Get default security group for Auto Mode
    return &SecurityGroup{
        ID: "sg-default-auto-mode",
        Name: "default-auto-mode",
    }, nil
}

func (sgm *SecurityGroupManager) AttachSecurityGroupToENI(eniID, sgID string) error {
    // Attach security group to ENI
    // This would use EC2 API to modify ENI security groups
    return nil
}
```

### 4. Configuration Updates

```go
// pkg/eniconfig/auto_mode_config.go
package eniconfig

import (
    "os"
    "strconv"
)

type AutoModeConfig struct {
    SGPPEnabled           bool
    ENIPoolSize          int
    SGFallback           string
    ENIAllocationStrategy string
}

func LoadAutoModeConfig() *AutoModeConfig {
    return &AutoModeConfig{
        SGPPEnabled:           getBoolEnv("AUTO_MODE_SGPP_ENABLED", true),
        ENIPoolSize:          getIntEnv("AUTO_MODE_ENI_POOL_SIZE", 10),
        SGFallback:           getStringEnv("AUTO_MODE_SG_FALLBACK", "cluster-security-group"),
        ENIAllocationStrategy: getStringEnv("AUTO_MODE_ENI_ALLOCATION_STRATEGY", "pool-based"),
    }
}

func getBoolEnv(key string, defaultValue bool) bool {
    if value := os.Getenv(key); value != "" {
        if parsed, err := strconv.ParseBool(value); err == nil {
            return parsed
        }
    }
    return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if parsed, err := strconv.Atoi(value); err == nil {
            return parsed
        }
    }
    return defaultValue
}

func getStringEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

## Testing Framework

### Unit Tests

```go
// pkg/eniconfig/auto_mode_detector_test.go
package eniconfig

import (
    "testing"
)

func TestAutoModeDetector_IsAutoMode(t *testing.T) {
    detector := &AutoModeDetector{
        clusterName: "test-cluster",
        region:      "us-west-2",
    }
    
    isAutoMode, err := detector.IsAutoMode()
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    
    if !isAutoMode {
        t.Error("Expected Auto Mode to be detected")
    }
}

func TestAutoModeENIManager_AllocateENIForPod(t *testing.T) {
    manager, err := NewAutoModeENIManager("test-cluster", "us-west-2")
    if err != nil {
        t.Fatalf("Failed to create ENI manager: %v", err)
    }
    
    podSpec := &PodSpec{
        Name:      "test-pod",
        Namespace: "default",
        Annotations: map[string]string{
            "vpc.amazonaws.com/security-groups": "sg-12345678",
        },
    }
    
    eni, err := manager.AllocateENIForPod(podSpec)
    if err != nil {
        t.Fatalf("Failed to allocate ENI: %v", err)
    }
    
    if eni == nil {
        t.Error("Expected ENI to be allocated")
    }
}
```

## Usage Example

```go
// Example usage of Auto Mode SGPP
func main() {
    // Create Auto Mode ENI manager
    manager, err := NewAutoModeENIManager("my-auto-mode-cluster", "us-west-2")
    if err != nil {
        log.Fatalf("Failed to create ENI manager: %v", err)
    }
    
    // Create pod specification with security group annotation
    podSpec := &PodSpec{
        Name:      "my-pod",
        Namespace: "default",
        Annotations: map[string]string{
            "vpc.amazonaws.com/security-groups": "sg-12345678,sg-87654321",
        },
    }
    
    // Allocate ENI with security group
    eni, err := manager.AllocateENIForPod(podSpec)
    if err != nil {
        log.Fatalf("Failed to allocate ENI: %v", err)
    }
    
    log.Printf("Allocated ENI: %s with security groups: %v", eni.ID, eni.SecurityGroups)
}
```

## Next Steps

1. **Implement Core Components**: Complete the implementation of all components
2. **Add Tests**: Implement comprehensive unit and integration tests
3. **Integration**: Integrate with existing VPC CNI plugin
4. **Validation**: Test with real EKS Auto Mode clusters
5. **Documentation**: Update documentation and user guides
