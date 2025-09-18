# Error Handling Reference

This document provides a comprehensive reference for error handling patterns, error codes, and best practices in the EKS Auto Mode SGPP enhancement.

## 📋 **Error Handling Overview**

The EKS Auto Mode SGPP enhancement implements comprehensive error handling with:
- Structured error types with error codes
- Detailed error messages with context
- Error wrapping for root cause analysis
- Graceful degradation and fallback mechanisms

## 🔧 **Error Types**

### **AutoModeError**

Custom error type for Auto Mode operations.

```go
type AutoModeError struct {
    Code    ErrorCode `json:"code"`
    Message string    `json:"message"`
    Cause   error     `json:"cause,omitempty"`
    Context map[string]interface{} `json:"context,omitempty"`
}

func (e *AutoModeError) Error() string {
    if e.Cause != nil {
        return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
    }
    return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AutoModeError) Unwrap() error {
    return e.Cause
}
```

### **ErrorCode**

Comprehensive error code enumeration.

```go
type ErrorCode string

const (
    // Configuration Errors
    ErrorCodeInvalidConfiguration         ErrorCode = "INVALID_CONFIGURATION"
    ErrorCodeMissingRequiredConfig        ErrorCode = "MISSING_REQUIRED_CONFIG"
    ErrorCodeInvalidEnvironmentVariable  ErrorCode = "INVALID_ENVIRONMENT_VARIABLE"
    
    // Auto Mode Detection Errors
    ErrorCodeAutoModeDetectionFailed      ErrorCode = "AUTO_MODE_DETECTION_FAILED"
    ErrorCodeClusterNotFound             ErrorCode = "CLUSTER_NOT_FOUND"
    ErrorCodeInvalidClusterMode          ErrorCode = "INVALID_CLUSTER_MODE"
    
    // ENI Management Errors
    ErrorCodeENIAllocationFailed          ErrorCode = "ENI_ALLOCATION_FAILED"
    ErrorCodeENIReleaseFailed             ErrorCode = "ENI_RELEASE_FAILED"
    ErrorCodeENIPoolExhausted             ErrorCode = "ENI_POOL_EXHAUSTED"
    ErrorCodeENINotFound                  ErrorCode = "ENI_NOT_FOUND"
    ErrorCodeENIAttachmentFailed          ErrorCode = "ENI_ATTACHMENT_FAILED"
    
    // Security Group Errors
    ErrorCodeSecurityGroupValidationFailed ErrorCode = "SECURITY_GROUP_VALIDATION_FAILED"
    ErrorCodeSecurityGroupNotFound        ErrorCode = "SECURITY_GROUP_NOT_FOUND"
    ErrorCodeTooManySecurityGroups        ErrorCode = "TOO_MANY_SECURITY_GROUPS"
    ErrorCodeSecurityGroupAttachmentFailed ErrorCode = "SECURITY_GROUP_ATTACHMENT_FAILED"
    
    // Resource Management Errors
    ErrorCodeResourceExhausted            ErrorCode = "RESOURCE_EXHAUSTED"
    ErrorCodeResourceNotFound              ErrorCode = "RESOURCE_NOT_FOUND"
    ErrorCodeResourceConflict              ErrorCode = "RESOURCE_CONFLICT"
    
    // Timeout Errors
    ErrorCodeTimeout                       ErrorCode = "TIMEOUT"
    ErrorCodeAllocationTimeout             ErrorCode = "ALLOCATION_TIMEOUT"
    ErrorCodeValidationTimeout             ErrorCode = "VALIDATION_TIMEOUT"
    ErrorCodeDetectionTimeout              ErrorCode = "DETECTION_TIMEOUT"
    
    // Permission Errors
    ErrorCodePermissionDenied              ErrorCode = "PERMISSION_DENIED"
    ErrorCodeInsufficientPermissions       ErrorCode = "INSUFFICIENT_PERMISSIONS"
    ErrorCodeAccessDenied                  ErrorCode = "ACCESS_DENIED"
    
    // Input Validation Errors
    ErrorCodeInvalidInput                  ErrorCode = "INVALID_INPUT"
    ErrorCodeInvalidPodSpec                ErrorCode = "INVALID_POD_SPEC"
    ErrorCodeInvalidSecurityGroupID        ErrorCode = "INVALID_SECURITY_GROUP_ID"
    ErrorCodeInvalidENIID                  ErrorCode = "INVALID_ENI_ID"
    
    // AWS API Errors
    ErrorCodeAWSServiceError               ErrorCode = "AWS_SERVICE_ERROR"
    ErrorCodeEC2ServiceError               ErrorCode = "EC2_SERVICE_ERROR"
    ErrorCodeEKSServiceError                ErrorCode = "EKS_SERVICE_ERROR"
    
    // Integration Errors
    ErrorCodeIPAMIntegrationFailed         ErrorCode = "IPAM_INTEGRATION_FAILED"
    ErrorCodeKubernetesAPIError            ErrorCode = "KUBERNETES_API_ERROR"
    ErrorCodeCNIPluginError                ErrorCode = "CNI_PLUGIN_ERROR"
)
```

## 🔄 **Error Handling Patterns**

### **Error Creation**

```go
// Create error with code and message
func NewAutoModeError(code ErrorCode, message string) *AutoModeError {
    return &AutoModeError{
        Code:    code,
        Message: message,
    }
}

// Create error with cause
func NewAutoModeErrorWithCause(code ErrorCode, message string, cause error) *AutoModeError {
    return &AutoModeError{
        Code:    code,
        Message: message,
        Cause:   cause,
    }
}

// Create error with context
func NewAutoModeErrorWithContext(code ErrorCode, message string, context map[string]interface{}) *AutoModeError {
    return &AutoModeError{
        Code:    code,
        Message: message,
        Context: context,
    }
}
```

### **Error Wrapping**

```go
// Wrap AWS SDK errors
func wrapAWSError(operation string, err error) *AutoModeError {
    if err == nil {
        return nil
    }
    
    var awsErr *AutoModeError
    if errors.As(err, &awsErr) {
        return awsErr
    }
    
    return NewAutoModeErrorWithCause(
        ErrorCodeAWSServiceError,
        fmt.Sprintf("AWS operation %s failed", operation),
        err,
    )
}

// Wrap Kubernetes API errors
func wrapKubernetesError(operation string, err error) *AutoModeError {
    if err == nil {
        return nil
    }
    
    return NewAutoModeErrorWithCause(
        ErrorCodeKubernetesAPIError,
        fmt.Sprintf("Kubernetes operation %s failed", operation),
        err,
    )
}
```

### **Error Context**

```go
// Add context to errors
func addErrorContext(err *AutoModeError, key string, value interface{}) *AutoModeError {
    if err.Context == nil {
        err.Context = make(map[string]interface{})
    }
    err.Context[key] = value
    return err
}

// Example usage
err := NewAutoModeError(ErrorCodeENIAllocationFailed, "Failed to allocate ENI")
err = addErrorContext(err, "podName", podSpec.Name)
err = addErrorContext(err, "namespace", podSpec.Namespace)
err = addErrorContext(err, "securityGroups", podSpec.SecurityGroups)
```

## 🐛 **Error Handling by Component**

### **Auto Mode Detector**

```go
// Auto Mode detection errors
func (d *AutoModeDetector) IsAutoMode(ctx context.Context) (bool, error) {
    // Check configuration
    if d.clusterName == "" {
        return false, NewAutoModeError(
            ErrorCodeMissingRequiredConfig,
            "Cluster name is required for Auto Mode detection",
        )
    }
    
    // Detect Auto Mode with timeout
    ctx, cancel := context.WithTimeout(ctx, DefaultDetectionTimeout)
    defer cancel()
    
    // Implementation with error handling
    clusterInfo, err := d.getClusterInfo(ctx)
    if err != nil {
        return false, NewAutoModeErrorWithCause(
            ErrorCodeAutoModeDetectionFailed,
            "Failed to detect Auto Mode",
            err,
        )
    }
    
    return clusterInfo.Mode == ClusterModeAuto, nil
}
```

### **ENI Manager**

```go
// ENI allocation errors
func (m *AutoModeENIManager) AllocateENIForPod(ctx context.Context, podSpec *PodSpec) (*ENIInfo, error) {
    // Validate input
    if podSpec == nil {
        return nil, NewAutoModeError(
            ErrorCodeInvalidInput,
            "Pod specification cannot be nil",
        )
    }
    
    // Check pool availability
    if len(m.eniPool) >= m.maxPoolSize {
        return nil, NewAutoModeErrorWithContext(
            ErrorCodeENIPoolExhausted,
            "ENI pool is exhausted",
            map[string]interface{}{
                "currentPoolSize": len(m.eniPool),
                "maxPoolSize":     m.maxPoolSize,
            },
        )
    }
    
    // Allocate ENI with timeout
    ctx, cancel := context.WithTimeout(ctx, m.allocationTimeout)
    defer cancel()
    
    // Implementation with error handling
    eniInfo, err := m.allocateENI(ctx, podSpec)
    if err != nil {
        return nil, NewAutoModeErrorWithCause(
            ErrorCodeENIAllocationFailed,
            "Failed to allocate ENI for pod",
            err,
        )
    }
    
    return eniInfo, nil
}
```

### **Security Group Manager**

```go
// Security group validation errors
func (m *SecurityGroupManager) ValidateSecurityGroup(ctx context.Context, sgID string) error {
    // Validate input
    if sgID == "" {
        return NewAutoModeError(
            ErrorCodeInvalidInput,
            "Security group ID cannot be empty",
        )
    }
    
    // Check cache first
    if cached, exists := m.securityGroupCache[sgID]; exists {
        if time.Since(cached.LastCheck) < m.cacheTTL {
            return nil // Valid cached entry
        }
    }
    
    // Validate with AWS
    ctx, cancel := context.WithTimeout(ctx, m.validationTimeout)
    defer cancel()
    
    _, err := m.ec2Client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
        GroupIds: []string{sgID},
    })
    
    if err != nil {
        return NewAutoModeErrorWithCause(
            ErrorCodeSecurityGroupValidationFailed,
            fmt.Sprintf("Failed to validate security group %s", sgID),
            err,
        )
    }
    
    return nil
}
```

## 🔄 **Fallback Mechanisms**

### **Graceful Degradation**

```go
// Fallback to standard mode
func (m *AutoModeManager) ProcessPodWithAutoMode(ctx context.Context, podSpec *PodSpec) (*PodNetworkInfo, error) {
    // Try Auto Mode first
    networkInfo, err := m.processPodAutoMode(ctx, podSpec)
    if err != nil {
        // Log error and fallback
        log.Printf("Auto Mode processing failed: %v", err)
        
        // Fallback to standard mode
        return m.processPodStandard(ctx, podSpec)
    }
    
    return networkInfo, nil
}
```

### **Retry Mechanisms**

```go
// Retry with exponential backoff
func (m *AutoModeENIManager) allocateENIWithRetry(ctx context.Context, podSpec *PodSpec) (*ENIInfo, error) {
    maxRetries := 3
    baseDelay := 1 * time.Second
    
    for attempt := 0; attempt < maxRetries; attempt++ {
        eniInfo, err := m.allocateENI(ctx, podSpec)
        if err == nil {
            return eniInfo, nil
        }
        
        // Check if error is retryable
        if !isRetryableError(err) {
            return nil, err
        }
        
        // Wait before retry
        if attempt < maxRetries-1 {
            delay := baseDelay * time.Duration(1<<attempt)
            select {
            case <-ctx.Done():
                return nil, ctx.Err()
            case <-time.After(delay):
                // Continue to next attempt
            }
        }
    }
    
    return nil, NewAutoModeError(
        ErrorCodeENIAllocationFailed,
        fmt.Sprintf("Failed to allocate ENI after %d attempts", maxRetries),
    )
}
```

## 📝 **Error Logging**

### **Structured Logging**

```go
// Log errors with context
func logError(err error, context map[string]interface{}) {
    if autoModeErr, ok := err.(*AutoModeError); ok {
        log.WithFields(logrus.Fields{
            "errorCode":    autoModeErr.Code,
            "errorMessage": autoModeErr.Message,
            "context":      context,
        }).Error("Auto Mode operation failed")
        
        if autoModeErr.Cause != nil {
            log.WithError(autoModeErr.Cause).Debug("Root cause")
        }
    } else {
        log.WithError(err).WithFields(logrus.Fields{
            "context": context,
        }).Error("Operation failed")
    }
}
```

## 🔗 **Related Documentation**

- [Data Types](data-types.md)
- [Auto Mode Detector API](auto-mode-detector.md)
- [ENI Manager API](eni-manager.md)
- [Security Group Manager API](security-group-manager.md)
