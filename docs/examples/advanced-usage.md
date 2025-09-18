# Advanced Usage Examples

This document provides advanced usage examples and complex scenarios for the EKS Auto Mode SGPP enhancement.

## 📋 **Advanced Usage Overview**

This guide covers complex usage patterns, advanced configurations, and real-world scenarios for the EKS Auto Mode SGPP enhancement.

## 🔧 **Advanced Pod Configurations**

### **Multiple Security Groups per Pod**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: multi-sg-pod
  namespace: production
  annotations:
    vpc.amazonaws.com/security-groups: "sg-web-tier,sg-app-tier,sg-db-tier"
spec:
  containers:
  - name: web-server
    image: nginx:latest
    ports:
    - containerPort: 80
  - name: app-server
    image: node:18-alpine
    ports:
    - containerPort: 3000
```

### **Dynamic Security Group Assignment**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dynamic-sg-deployment
  namespace: staging
spec:
  replicas: 3
  selector:
    matchLabels:
      app: dynamic-sg-app
  template:
    metadata:
      labels:
        app: dynamic-sg-app
        environment: staging
      annotations:
        vpc.amazonaws.com/security-groups: "sg-staging-{{ .Values.environment }}"
    spec:
      containers:
      - name: app
        image: myapp:latest
        env:
        - name: ENVIRONMENT
          value: "staging"
```

### **Namespace-based Security Groups**

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: secure-namespace
  annotations:
    vpc.amazonaws.com/security-groups: "sg-namespace-default"
---
apiVersion: v1
kind: Pod
metadata:
  name: namespace-sg-pod
  namespace: secure-namespace
  annotations:
    vpc.amazonaws.com/security-groups: "sg-pod-specific"
spec:
  containers:
  - name: app
    image: myapp:latest
```

## 🏗️ **Advanced Architecture Patterns**

### **Microservices with Security Groups**

```yaml
# Frontend Service
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: frontend
  template:
    metadata:
      labels:
        app: frontend
      annotations:
        vpc.amazonaws.com/security-groups: "sg-frontend,sg-public"
    spec:
      containers:
      - name: frontend
        image: frontend:latest
        ports:
        - containerPort: 80
---
# Backend Service
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
      annotations:
        vpc.amazonaws.com/security-groups: "sg-backend,sg-internal"
    spec:
      containers:
      - name: backend
        image: backend:latest
        ports:
        - containerPort: 8080
---
# Database Service
apiVersion: apps/v1
kind: Deployment
metadata:
  name: database-service
spec:
  replicas: 1
  selector:
    matchLabels:
      app: database
  template:
    metadata:
      labels:
        app: database
      annotations:
        vpc.amazonaws.com/security-groups: "sg-database,sg-internal"
    spec:
      containers:
      - name: database
        image: postgres:13
        ports:
        - containerPort: 5432
```

### **Multi-Tier Application**

```yaml
# Web Tier
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-tier
spec:
  replicas: 5
  selector:
    matchLabels:
      tier: web
  template:
    metadata:
      labels:
        tier: web
      annotations:
        vpc.amazonaws.com/security-groups: "sg-web-tier"
    spec:
      containers:
      - name: web
        image: nginx:latest
---
# Application Tier
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-tier
spec:
  replicas: 3
  selector:
    matchLabels:
      tier: app
  template:
    metadata:
      labels:
        tier: app
      annotations:
        vpc.amazonaws.com/security-groups: "sg-app-tier"
    spec:
      containers:
      - name: app
        image: myapp:latest
---
# Database Tier
apiVersion: apps/v1
kind: Deployment
metadata:
  name: db-tier
spec:
  replicas: 1
  selector:
    matchLabels:
      tier: db
  template:
    metadata:
      labels:
        tier: db
      annotations:
        vpc.amazonaws.com/security-groups: "sg-db-tier"
    spec:
      containers:
      - name: db
        image: postgres:13
```

## 🔄 **Advanced Configuration Patterns**

### **Environment-specific Configuration**

```bash
#!/bin/bash
# Environment-specific configuration script

ENVIRONMENT=${1:-development}

case $ENVIRONMENT in
  development)
    export AUTO_MODE_SGPP_ENABLED="true"
    export CLUSTER_NAME="dev-cluster"
    export AWS_REGION="us-west-2"
    export ENI_POOL_SIZE="5"
    export MAX_SECURITY_GROUPS_PER_POD="3"
    export LOG_LEVEL="debug"
    ;;
  staging)
    export AUTO_MODE_SGPP_ENABLED="true"
    export CLUSTER_NAME="staging-cluster"
    export AWS_REGION="us-west-2"
    export ENI_POOL_SIZE="20"
    export MAX_SECURITY_GROUPS_PER_POD="5"
    export LOG_LEVEL="info"
    ;;
  production)
    export AUTO_MODE_SGPP_ENABLED="true"
    export CLUSTER_NAME="prod-cluster"
    export AWS_REGION="us-west-2"
    export ENI_POOL_SIZE="100"
    export MAX_SECURITY_GROUPS_PER_POD="5"
    export LOG_LEVEL="warn"
    export ENABLE_METRICS="true"
    export ENABLE_HEALTH_CHECKS="true"
    ;;
esac
```

### **Dynamic Configuration Management**

```yaml
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
      allocationTimeout: 45s
      cleanupInterval: 3m
    securityGroups:
      strictValidation: true
      cacheTTL: 10m
      validationTimeout: 15s
    performance:
      enableCaching: true
      maxConcurrentAllocations: 20
      workerPoolSize: 10
```

### **Configuration Validation**

```go
package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

func validateConfiguration() error {
    // Required configuration
    if os.Getenv("CLUSTER_NAME") == "" {
        return fmt.Errorf("CLUSTER_NAME is required")
    }
    
    if os.Getenv("AWS_REGION") == "" {
        return fmt.Errorf("AWS_REGION is required")
    }
    
    // Validate ENI pool size
    if poolSize := os.Getenv("ENI_POOL_SIZE"); poolSize != "" {
        if size, err := strconv.Atoi(poolSize); err != nil || size <= 0 {
            return fmt.Errorf("ENI_POOL_SIZE must be a positive integer")
        }
    }
    
    // Validate timeout values
    if timeout := os.Getenv("ENI_ALLOCATION_TIMEOUT"); timeout != "" {
        if _, err := time.ParseDuration(timeout); err != nil {
            return fmt.Errorf("ENI_ALLOCATION_TIMEOUT must be a valid duration")
        }
    }
    
    return nil
}
```

## 📊 **Advanced Monitoring and Observability**

### **Custom Metrics Collection**

```go
package main

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    eniAllocationDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "auto_mode_eni_allocation_duration_seconds",
            Help: "Duration of ENI allocation operations",
        },
        []string{"status", "error_type"},
    )
    
    securityGroupValidationDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "auto_mode_sg_validation_duration_seconds",
            Help: "Duration of security group validation operations",
        },
        []string{"status", "error_type"},
    )
    
    eniPoolUtilization = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "auto_mode_eni_pool_utilization",
            Help: "ENI pool utilization percentage",
        },
        []string{"pool_name"},
    )
)
```

### **Health Check Implementation**

```go
package main

import (
    "context"
    "net/http"
    "time"
)

type HealthChecker struct {
    detector           *AutoModeDetector
    eniManager         *AutoModeENIManager
    securityGroupManager *SecurityGroupManager
}

func (h *HealthChecker) CheckHealth(ctx context.Context) error {
    // Check Auto Mode detection
    if _, err := h.detector.IsAutoMode(ctx); err != nil {
        return fmt.Errorf("Auto Mode detection failed: %v", err)
    }
    
    // Check ENI manager
    if err := h.eniManager.CheckHealth(ctx); err != nil {
        return fmt.Errorf("ENI manager health check failed: %v", err)
    }
    
    // Check security group manager
    if err := h.securityGroupManager.CheckHealth(ctx); err != nil {
        return fmt.Errorf("Security group manager health check failed: %v", err)
    }
    
    return nil
}

func (h *HealthChecker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()
    
    if err := h.CheckHealth(ctx); err != nil {
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte(fmt.Sprintf("Health check failed: %v", err)))
        return
    }
    
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}
```

## 🔄 **Advanced Error Handling**

### **Retry with Exponential Backoff**

```go
package main

import (
    "context"
    "math"
    "time"
)

func retryWithBackoff(ctx context.Context, operation func() error, maxRetries int) error {
    baseDelay := 1 * time.Second
    maxDelay := 30 * time.Second
    
    for attempt := 0; attempt < maxRetries; attempt++ {
        if err := operation(); err == nil {
            return nil
        }
        
        if attempt == maxRetries-1 {
            return err
        }
        
        // Calculate delay with exponential backoff
        delay := time.Duration(float64(baseDelay) * math.Pow(2, float64(attempt)))
        if delay > maxDelay {
            delay = maxDelay
        }
        
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(delay):
            // Continue to next attempt
        }
    }
    
    return fmt.Errorf("operation failed after %d attempts", maxRetries)
}
```

### **Circuit Breaker Pattern**

```go
package main

import (
    "context"
    "sync"
    "time"
)

type CircuitBreaker struct {
    maxFailures int
    timeout     time.Duration
    failures    int
    lastFailure time.Time
    mutex       sync.RWMutex
}

func (cb *CircuitBreaker) Execute(ctx context.Context, operation func() error) error {
    cb.mutex.RLock()
    if cb.failures >= cb.maxFailures {
        if time.Since(cb.lastFailure) < cb.timeout {
            cb.mutex.RUnlock()
            return fmt.Errorf("circuit breaker is open")
        }
    }
    cb.mutex.RUnlock()
    
    err := operation()
    
    cb.mutex.Lock()
    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
    } else {
        cb.failures = 0
    }
    cb.mutex.Unlock()
    
    return err
}
```

## 🔧 **Advanced Testing Patterns**

### **Integration Testing**

```go
package main

import (
    "context"
    "testing"
    "time"
)

func TestAutoModeIntegration(t *testing.T) {
    // Setup test environment
    ctx := context.Background()
    
    // Create test cluster
    cluster := createTestCluster(t)
    defer cleanupTestCluster(t, cluster)
    
    // Test Auto Mode detection
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
    
    // Verify ENI allocation
    if eniInfo.ENIID == "" {
        t.Error("ENI ID is empty")
    }
    
    // Cleanup
    if err := eniManager.ReleaseENIForPod(ctx, podSpec); err != nil {
        t.Errorf("ENI release failed: %v", err)
    }
}
```

### **Load Testing**

```go
package main

import (
    "context"
    "sync"
    "testing"
    "time"
)

func TestLoadENIAllocation(t *testing.T) {
    const numGoroutines = 100
    const numAllocations = 10
    
    ctx := context.Background()
    eniManager := createTestENIManager(t)
    
    var wg sync.WaitGroup
    errors := make(chan error, numGoroutines*numAllocations)
    
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func(goroutineID int) {
            defer wg.Done()
            
            for j := 0; j < numAllocations; j++ {
                podSpec := &PodSpec{
                    Name:      fmt.Sprintf("load-test-pod-%d-%d", goroutineID, j),
                    Namespace: "default",
                    SecurityGroups: []string{"sg-load-test"},
                }
                
                eniInfo, err := eniManager.AllocateENIForPod(ctx, podSpec)
                if err != nil {
                    errors <- err
                    return
                }
                
                // Simulate pod lifecycle
                time.Sleep(100 * time.Millisecond)
                
                if err := eniManager.ReleaseENIForPod(ctx, podSpec); err != nil {
                    errors <- err
                    return
                }
            }
        }(i)
    }
    
    wg.Wait()
    close(errors)
    
    // Check for errors
    var errorCount int
    for err := range errors {
        t.Errorf("Load test error: %v", err)
        errorCount++
    }
    
    if errorCount > 0 {
        t.Fatalf("Load test failed with %d errors", errorCount)
    }
}
```

## 🔗 **Related Documentation**

- [Basic Usage](basic-usage.md)
- [Best Practices](best-practices.md)
- [Architecture Overview](../architecture/overview.md)
- [Configuration Guide](../user-guide/configuration.md)
