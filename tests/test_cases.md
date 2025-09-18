# EKS Auto Mode SGPP Enhancement - Test Cases

## Test Categories

### 1. Unit Tests

#### Auto Mode Detection Tests
```go
func TestAutoModeDetector_IsAutoMode(t *testing.T) {
    tests := []struct {
        name        string
        clusterName string
        region      string
        expected    bool
        expectError bool
    }{
        {
            name:        "Auto Mode cluster",
            clusterName: "auto-mode-cluster",
            region:      "us-west-2",
            expected:    true,
            expectError: false,
        },
        {
            name:        "Standard Mode cluster",
            clusterName: "standard-mode-cluster",
            region:      "us-west-2",
            expected:    false,
            expectError: false,
        },
        {
            name:        "Invalid cluster",
            clusterName: "invalid-cluster",
            region:      "us-west-2",
            expected:    false,
            expectError: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            detector := &AutoModeDetector{
                clusterName: tt.clusterName,
                region:      tt.region,
            }
            
            result, err := detector.IsAutoMode()
            
            if tt.expectError && err == nil {
                t.Error("Expected error but got none")
            }
            
            if !tt.expectError && err != nil {
                t.Errorf("Unexpected error: %v", err)
            }
            
            if result != tt.expected {
                t.Errorf("Expected %v, got %v", tt.expected, result)
            }
        })
    }
}
```

#### ENI Management Tests
```go
func TestAutoModeENIManager_AllocateENIForPod(t *testing.T) {
    tests := []struct {
        name        string
        podSpec     *PodSpec
        autoMode    bool
        expectError bool
    }{
        {
            name: "Auto Mode with security group annotation",
            podSpec: &PodSpec{
                Name:      "test-pod",
                Namespace: "default",
                Annotations: map[string]string{
                    "vpc.amazonaws.com/security-groups": "sg-12345678",
                },
            },
            autoMode:    true,
            expectError: false,
        },
        {
            name: "Auto Mode without security group annotation",
            podSpec: &PodSpec{
                Name:      "test-pod",
                Namespace: "default",
                Annotations: map[string]string{},
            },
            autoMode:    true,
            expectError: false,
        },
        {
            name: "Standard Mode",
            podSpec: &PodSpec{
                Name:      "test-pod",
                Namespace: "default",
                Annotations: map[string]string{
                    "vpc.amazonaws.com/security-groups": "sg-12345678",
                },
            },
            autoMode:    false,
            expectError: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            manager := &AutoModeENIManager{
                detector: &AutoModeDetector{},
                eniPool:  &ENIPool{},
                sgManager: &SecurityGroupManager{},
            }
            
            eni, err := manager.AllocateENIForPod(tt.podSpec)
            
            if tt.expectError && err == nil {
                t.Error("Expected error but got none")
            }
            
            if !tt.expectError && err != nil {
                t.Errorf("Unexpected error: %v", err)
            }
            
            if !tt.expectError && eni == nil {
                t.Error("Expected ENI to be allocated")
            }
        })
    }
}
```

#### Security Group Management Tests
```go
func TestSecurityGroupManager_GetSecurityGroupForPod(t *testing.T) {
    tests := []struct {
        name        string
        podSpec     *PodSpec
        expectedSG  string
        expectError bool
    }{
        {
            name: "Pod with security group annotation",
            podSpec: &PodSpec{
                Name:      "test-pod",
                Namespace: "default",
                Annotations: map[string]string{
                    "vpc.amazonaws.com/security-groups": "sg-12345678",
                },
            },
            expectedSG:  "sg-12345678",
            expectError: false,
        },
        {
            name: "Pod with multiple security groups",
            podSpec: &PodSpec{
                Name:      "test-pod",
                Namespace: "default",
                Annotations: map[string]string{
                    "vpc.amazonaws.com/security-groups": "sg-12345678,sg-87654321",
                },
            },
            expectedSG:  "sg-12345678",
            expectError: false,
        },
        {
            name: "Pod without security group annotation",
            podSpec: &PodSpec{
                Name:      "test-pod",
                Namespace: "default",
                Annotations: map[string]string{},
            },
            expectedSG:  "sg-default-auto-mode",
            expectError: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            sgManager := &SecurityGroupManager{
                autoMode: true,
            }
            
            sg, err := sgManager.GetSecurityGroupForPod(tt.podSpec)
            
            if tt.expectError && err == nil {
                t.Error("Expected error but got none")
            }
            
            if !tt.expectError && err != nil {
                t.Errorf("Unexpected error: %v", err)
            }
            
            if !tt.expectError && sg.ID != tt.expectedSG {
                t.Errorf("Expected security group %s, got %s", tt.expectedSG, sg.ID)
            }
        })
    }
}
```

### 2. Integration Tests

#### VPC CNI Plugin Integration Tests
```go
func TestVPCCNIPluginIntegration(t *testing.T) {
    // Test integration with existing VPC CNI plugin
    // This would test the complete workflow from pod creation to ENI allocation
    
    t.Run("Pod creation with SGPP in Auto Mode", func(t *testing.T) {
        // Create test pod with security group annotation
        pod := &v1.Pod{
            ObjectMeta: metav1.ObjectMeta{
                Name:      "test-pod",
                Namespace: "default",
                Annotations: map[string]string{
                    "vpc.amazonaws.com/security-groups": "sg-12345678",
                },
            },
            Spec: v1.PodSpec{
                Containers: []v1.Container{
                    {
                        Name:  "test-container",
                        Image: "nginx:latest",
                    },
                },
            },
        }
        
        // Test pod creation and ENI allocation
        // This would involve creating the pod in a test cluster
        // and verifying that the ENI is allocated with the correct security group
    })
}
```

#### EKS Control Plane Integration Tests
```go
func TestEKSControlPlaneIntegration(t *testing.T) {
    // Test integration with EKS Control Plane
    // This would test the complete workflow from cluster creation to pod deployment
    
    t.Run("Auto Mode cluster with SGPP", func(t *testing.T) {
        // Create test cluster in Auto Mode
        // Deploy test pod with security group annotation
        // Verify ENI allocation and security group assignment
    })
}
```

### 3. End-to-End Tests

#### Complete Workflow Tests
```go
func TestCompleteSGPPWorkflow(t *testing.T) {
    // Test complete workflow from cluster creation to pod deployment
    // This would involve:
    // 1. Creating EKS cluster in Auto Mode
    // 2. Deploying VPC CNI plugin with SGPP support
    // 3. Creating pod with security group annotation
    // 4. Verifying ENI allocation and security group assignment
    // 5. Testing pod connectivity with assigned security group
}
```

#### Error Scenario Tests
```go
func TestErrorScenarios(t *testing.T) {
    tests := []struct {
        name        string
        scenario    string
        expectError bool
    }{
        {
            name:        "Invalid security group ID",
            scenario:    "pod-with-invalid-sg",
            expectError: true,
        },
        {
            name:        "Non-existent security group",
            scenario:    "pod-with-nonexistent-sg",
            expectError: true,
        },
        {
            name:        "Security group in different VPC",
            scenario:    "pod-with-different-vpc-sg",
            expectError: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test error scenarios
            // This would test various error conditions and ensure proper error handling
        })
    }
}
```

### 4. Performance Tests

#### ENI Allocation Performance Tests
```go
func BenchmarkAutoModeENIAllocation(b *testing.B) {
    manager, err := NewAutoModeENIManager("test-cluster", "us-west-2")
    if err != nil {
        b.Fatalf("Failed to create ENI manager: %v", err)
    }
    
    podSpec := &PodSpec{
        Name:      "test-pod",
        Namespace: "default",
        Annotations: map[string]string{
            "vpc.amazonaws.com/security-groups": "sg-12345678",
        },
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := manager.AllocateENIForPod(podSpec)
        if err != nil {
            b.Fatalf("Failed to allocate ENI: %v", err)
        }
    }
}
```

#### Security Group Assignment Performance Tests
```go
func BenchmarkSecurityGroupAssignment(b *testing.B) {
    sgManager := &SecurityGroupManager{
        autoMode: true,
    }
    
    podSpec := &PodSpec{
        Name:      "test-pod",
        Namespace: "default",
        Annotations: map[string]string{
            "vpc.amazonaws.com/security-groups": "sg-12345678",
        },
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := sgManager.GetSecurityGroupForPod(podSpec)
        if err != nil {
            b.Fatalf("Failed to get security group: %v", err)
        }
    }
}
```

## Test Data

### Test Pod Specifications
```yaml
# test-pod-with-sg.yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-pod-with-sg
  namespace: default
  annotations:
    vpc.amazonaws.com/security-groups: "sg-12345678"
spec:
  containers:
  - name: test-container
    image: nginx:latest
    ports:
    - containerPort: 80
```

### Test Security Groups
```yaml
# test-security-groups.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-security-groups
  namespace: kube-system
data:
  security-groups: |
    - id: sg-12345678
      name: test-sg-1
      description: Test security group 1
    - id: sg-87654321
      name: test-sg-2
      description: Test security group 2
```

## Test Execution

### Running Unit Tests
```bash
# Run all unit tests
go test ./pkg/eniconfig/... -v

# Run specific test
go test ./pkg/eniconfig/... -run TestAutoModeDetector_IsAutoMode -v

# Run tests with coverage
go test ./pkg/eniconfig/... -cover -v
```

### Running Integration Tests
```bash
# Run integration tests
go test ./tests/integration/... -v

# Run integration tests with cluster
go test ./tests/integration/... -v -cluster=test-cluster
```

### Running Performance Tests
```bash
# Run performance tests
go test ./pkg/eniconfig/... -bench=. -v

# Run specific benchmark
go test ./pkg/eniconfig/... -bench=BenchmarkAutoModeENIAllocation -v
```

## Test Environment Setup

### Prerequisites
- Go 1.19+ installed
- Docker installed
- AWS CLI configured
- kubectl installed
- eksctl installed

### Test Cluster Setup
```bash
# Create test cluster in Auto Mode
eksctl create cluster \
  --name test-auto-mode-cluster \
  --region us-west-2 \
  --managed

# Deploy test VPC CNI plugin
kubectl apply -f tests/manifests/vpc-cni-plugin.yaml

# Deploy test security groups
kubectl apply -f tests/manifests/test-security-groups.yaml
```

### Test Cleanup
```bash
# Clean up test cluster
eksctl delete cluster \
  --name test-auto-mode-cluster \
  --region us-west-2
```
