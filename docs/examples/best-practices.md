# Best Practices Guide

This document provides best practices and recommendations for using the EKS Auto Mode SGPP enhancement effectively.

## 📋 **Best Practices Overview**

This guide covers best practices for configuration, security, performance, monitoring, and operational excellence with the EKS Auto Mode SGPP enhancement.

## 🔧 **Configuration Best Practices**

### **Environment-specific Configuration**

```bash
# Use environment-specific configuration files
# development.env
AUTO_MODE_SGPP_ENABLED=true
CLUSTER_NAME=dev-cluster
AWS_REGION=us-west-2
ENI_POOL_SIZE=5
MAX_SECURITY_GROUPS_PER_POD=3
LOG_LEVEL=debug

# production.env
AUTO_MODE_SGPP_ENABLED=true
CLUSTER_NAME=prod-cluster
AWS_REGION=us-west-2
ENI_POOL_SIZE=100
MAX_SECURITY_GROUPS_PER_POD=5
LOG_LEVEL=info
ENABLE_METRICS=true
ENABLE_HEALTH_CHECKS=true
```

### **Configuration Validation**

```go
// Always validate configuration on startup
func validateConfiguration() error {
    requiredVars := []string{"CLUSTER_NAME", "AWS_REGION"}
    for _, varName := range requiredVars {
        if os.Getenv(varName) == "" {
            return fmt.Errorf("%s is required", varName)
        }
    }
    
    // Validate numeric values
    if poolSize := os.Getenv("ENI_POOL_SIZE"); poolSize != "" {
        if size, err := strconv.Atoi(poolSize); err != nil || size <= 0 {
            return fmt.Errorf("ENI_POOL_SIZE must be a positive integer")
        }
    }
    
    return nil
}
```

### **Configuration Management**

```yaml
# Use ConfigMaps for configuration
apiVersion: v1
kind: ConfigMap
metadata:
  name: auto-mode-config
  namespace: kube-system
data:
  config.yaml: |
    autoMode:
      enabled: true
      eniPoolSize: 50
      maxSecurityGroupsPerPod: 5
    securityGroups:
      strictValidation: true
      cacheTTL: 10m
    performance:
      enableCaching: true
      maxConcurrentAllocations: 20
```

## 🛡️ **Security Best Practices**

### **Security Group Design**

```yaml
# Use descriptive security group names
apiVersion: v1
kind: Pod
metadata:
  name: web-server
  annotations:
    vpc.amazonaws.com/security-groups: "sg-web-tier,sg-public-access"
spec:
  containers:
  - name: web
    image: nginx:latest
```

### **Least Privilege Principle**

```yaml
# Create specific security groups for each service
# Web tier security group
apiVersion: v1
kind: Pod
metadata:
  name: web-pod
  annotations:
    vpc.amazonaws.com/security-groups: "sg-web-tier"
spec:
  containers:
  - name: web
    image: nginx:latest
    ports:
    - containerPort: 80
---
# API tier security group
apiVersion: v1
kind: Pod
metadata:
  name: api-pod
  annotations:
    vpc.amazonaws.com/security-groups: "sg-api-tier"
spec:
  containers:
  - name: api
    image: api:latest
    ports:
    - containerPort: 8080
```

### **Security Group Validation**

```bash
# Always validate security groups before deployment
aws ec2 describe-security-groups --group-ids sg-xxx

# Check security group rules
aws ec2 describe-security-groups --group-ids sg-xxx --query 'SecurityGroups[0].IpPermissions'
```

## 📊 **Performance Best Practices**

### **ENI Pool Sizing**

```bash
# Size ENI pool based on expected workload
# For high-traffic applications
export ENI_POOL_SIZE="100"
export ENI_POOL_MIN_SIZE="20"
export ENI_POOL_MAX_SIZE="200"

# For low-traffic applications
export ENI_POOL_SIZE="10"
export ENI_POOL_MIN_SIZE="5"
export ENI_POOL_MAX_SIZE="20"
```

### **Caching Configuration**

```bash
# Optimize caching for performance
export ENABLE_COMPONENT_CACHING="true"
export GLOBAL_CACHE_TTL="10m"
export SECURITY_GROUP_CACHE_TTL="15m"
export AUTO_MODE_CACHE_TTL="5m"
```

### **Concurrency Settings**

```bash
# Configure concurrency based on cluster size
export MAX_CONCURRENT_ENI_ALLOCATIONS="20"
export MAX_CONCURRENT_SG_VALIDATIONS="50"
export WORKER_POOL_SIZE="10"
```

## 🔍 **Monitoring Best Practices**

### **Key Metrics to Monitor**

```go
// Monitor critical metrics
var (
    eniAllocationSuccessRate = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "auto_mode_eni_allocation_success_rate",
            Help: "ENI allocation success rate",
        },
        []string{"cluster", "namespace"},
    )
    
    securityGroupValidationLatency = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "auto_mode_sg_validation_duration_seconds",
            Help: "Security group validation latency",
        },
        []string{"status"},
    )
    
    eniPoolUtilization = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "auto_mode_eni_pool_utilization_percent",
            Help: "ENI pool utilization percentage",
        },
        []string{"pool_name"},
    )
)
```

### **Alerting Rules**

```yaml
# Prometheus alerting rules
groups:
- name: auto-mode-sgpp
  rules:
  - alert: ENIPoolExhausted
    expr: auto_mode_eni_pool_utilization_percent > 90
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "ENI pool utilization is high"
      description: "ENI pool utilization is {{ $value }}%"
  
  - alert: SecurityGroupValidationFailed
    expr: rate(auto_mode_sg_validation_failures_total[5m]) > 0.1
    for: 2m
    labels:
      severity: critical
    annotations:
      summary: "Security group validation failures"
      description: "Security group validation failure rate is {{ $value }}"
```

### **Health Checks**

```go
// Implement comprehensive health checks
func (h *HealthChecker) CheckHealth(ctx context.Context) error {
    checks := []func(context.Context) error{
        h.checkAutoModeDetection,
        h.checkENIManager,
        h.checkSecurityGroupManager,
        h.checkConfiguration,
    }
    
    for _, check := range checks {
        if err := check(ctx); err != nil {
            return err
        }
    }
    
    return nil
}
```

## 🔄 **Operational Best Practices**

### **Deployment Strategy**

```yaml
# Use rolling updates for deployments
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auto-mode-sgpp
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 1
  selector:
    matchLabels:
      app: auto-mode-sgpp
  template:
    metadata:
      labels:
        app: auto-mode-sgpp
    spec:
      containers:
      - name: auto-mode-sgpp
        image: auto-mode-sgpp:latest
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
```

### **Resource Management**

```yaml
# Set appropriate resource limits
apiVersion: v1
kind: Pod
metadata:
  name: resource-managed-pod
spec:
  containers:
  - name: app
    image: myapp:latest
    resources:
      requests:
        memory: "128Mi"
        cpu: "100m"
      limits:
        memory: "256Mi"
        cpu: "200m"
```

### **Graceful Shutdown**

```go
// Implement graceful shutdown
func (s *Server) Shutdown(ctx context.Context) error {
    // Stop accepting new requests
    s.server.SetKeepAlivesEnabled(false)
    
    // Wait for existing requests to complete
    timeout := 30 * time.Second
    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()
    
    return s.server.Shutdown(ctx)
}
```

## 🧪 **Testing Best Practices**

### **Unit Testing**

```go
func TestAutoModeDetector(t *testing.T) {
    tests := []struct {
        name     string
        cluster  string
        region   string
        expected bool
        wantErr  bool
    }{
        {
            name:     "valid auto mode cluster",
            cluster:  "test-cluster",
            region:   "us-west-2",
            expected: true,
            wantErr:  false,
        },
        {
            name:     "invalid cluster",
            cluster:  "",
            region:   "us-west-2",
            expected: false,
            wantErr:  true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            detector, err := NewAutoModeDetector(tt.cluster, tt.region)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewAutoModeDetector() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            
            if detector != nil {
                result, err := detector.IsAutoMode(context.Background())
                if (err != nil) != tt.wantErr {
                    t.Errorf("IsAutoMode() error = %v, wantErr %v", err, tt.wantErr)
                    return
                }
                
                if result != tt.expected {
                    t.Errorf("IsAutoMode() = %v, want %v", result, tt.expected)
                }
            }
        })
    }
}
```

### **Integration Testing**

```go
func TestIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }
    
    // Setup test environment
    ctx := context.Background()
    cluster := createTestCluster(t)
    defer cleanupTestCluster(t, cluster)
    
    // Test complete workflow
    detector, err := NewAutoModeDetector(cluster.Name, cluster.Region)
    if err != nil {
        t.Fatalf("Failed to create detector: %v", err)
    }
    
    isAutoMode, err := detector.IsAutoMode(ctx)
    if err != nil {
        t.Fatalf("Auto Mode detection failed: %v", err)
    }
    
    if !isAutoMode {
        t.Skip("Cluster is not in Auto Mode")
    }
    
    // Test ENI allocation
    eniManager, err := NewAutoModeENIManager(cluster.Name, cluster.Region)
    if err != nil {
        t.Fatalf("Failed to create ENI manager: %v", err)
    }
    
    podSpec := &PodSpec{
        Name:      "test-pod",
        Namespace: "default",
        SecurityGroups: []string{"sg-test"},
    }
    
    eniInfo, err := eniManager.AllocateENIForPod(ctx, podSpec)
    if err != nil {
        t.Fatalf("ENI allocation failed: %v", err)
    }
    
    // Verify allocation
    if eniInfo.ENIID == "" {
        t.Error("ENI ID is empty")
    }
    
    // Cleanup
    if err := eniManager.ReleaseENIForPod(ctx, podSpec); err != nil {
        t.Errorf("ENI release failed: %v", err)
    }
}
```

## 📝 **Documentation Best Practices**

### **Code Documentation**

```go
// AutoModeDetector detects if an EKS cluster is running in Auto Mode
// and provides cluster information with caching for improved performance.
type AutoModeDetector struct {
    clusterName string
    region      string
    eksClient   *eks.Client
    cache       map[string]*ClusterInfo
    cacheTTL    time.Duration
    mutex       sync.RWMutex
}

// IsAutoMode checks if the cluster is running in Auto Mode.
// It uses caching to improve performance and falls back to AWS API calls
// when cache is empty or expired.
func (d *AutoModeDetector) IsAutoMode(ctx context.Context) (bool, error) {
    // Implementation details...
}
```

### **API Documentation**

```go
// AllocateENIForPod allocates an ENI for a pod with specified security groups.
// It uses a pool-based allocation strategy for improved performance and
// supports multiple security groups per pod (up to 5).
//
// Parameters:
//   - ctx: Context for the operation
//   - podSpec: Pod specification with security group requirements
//
// Returns:
//   - *ENIInfo: ENI information including ID, IP, and security groups
//   - error: Error if allocation fails
//
// Example:
//   podSpec := &PodSpec{
//       Name: "my-pod",
//       Namespace: "default",
//       SecurityGroups: []string{"sg-12345678"},
//   }
//   eniInfo, err := manager.AllocateENIForPod(ctx, podSpec)
func (m *AutoModeENIManager) AllocateENIForPod(ctx context.Context, podSpec *PodSpec) (*ENIInfo, error) {
    // Implementation details...
}
```

## 🔗 **Related Documentation**

- [Advanced Usage](advanced-usage.md)
- [Configuration Guide](../user-guide/configuration.md)
- [Troubleshooting Guide](../user-guide/troubleshooting.md)
- [Architecture Overview](../architecture/overview.md)
