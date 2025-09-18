# Troubleshooting Guide

This document provides comprehensive troubleshooting information for common issues with the EKS Auto Mode SGPP enhancement.

## 📋 **Troubleshooting Overview**

This guide covers common issues, error messages, diagnostic steps, and solutions for the EKS Auto Mode SGPP enhancement.

## 🔍 **Common Issues**

### **Auto Mode Detection Issues**

#### **Issue**: Auto Mode detection fails

**Symptoms**:
- Error: `AUTO_MODE_DETECTION_FAILED`
- Pods fall back to standard mode
- Security groups not applied

**Diagnostic Steps**:
1. Check cluster configuration:
   ```bash
   aws eks describe-cluster --name your-cluster --region your-region
   ```

2. Verify cluster tags:
   ```bash
   aws eks describe-cluster --name your-cluster --region your-region --query 'cluster.tags'
   ```

3. Check IAM permissions:
   ```bash
   aws sts get-caller-identity
   ```

**Solutions**:
- Ensure cluster is actually in Auto Mode
- Verify IAM permissions for EKS API access
- Check network connectivity to AWS APIs
- Review cluster configuration

#### **Issue**: Auto Mode detection timeout

**Symptoms**:
- Error: `DETECTION_TIMEOUT`
- Slow pod startup
- Intermittent failures

**Diagnostic Steps**:
1. Check network connectivity:
   ```bash
   curl -I https://eks.us-west-2.amazonaws.com
   ```

2. Review timeout settings:
   ```bash
   echo $AUTO_MODE_DETECTION_TIMEOUT
   ```

**Solutions**:
- Increase timeout value: `export AUTO_MODE_DETECTION_TIMEOUT="30s"`
- Check network connectivity
- Verify AWS region configuration
- Review DNS resolution

### **ENI Allocation Issues**

#### **Issue**: ENI allocation fails

**Symptoms**:
- Error: `ENI_ALLOCATION_FAILED`
- Pods stuck in Pending state
- Resource exhaustion errors

**Diagnostic Steps**:
1. Check ENI pool status:
   ```bash
   kubectl logs -n kube-system -l app=aws-node
   ```

2. Verify subnet capacity:
   ```bash
   aws ec2 describe-subnets --subnet-ids subnet-xxx
   ```

3. Check ENI limits:
   ```bash
   aws ec2 describe-network-interfaces --filters "Name=subnet-id,Values=subnet-xxx"
   ```

**Solutions**:
- Increase ENI pool size: `export ENI_POOL_SIZE="20"`
- Check subnet IP availability
- Verify ENI limits per instance type
- Review VPC configuration

#### **Issue**: ENI pool exhausted

**Symptoms**:
- Error: `ENI_POOL_EXHAUSTED`
- Pods cannot be scheduled
- High resource utilization

**Diagnostic Steps**:
1. Check pool utilization:
   ```bash
   kubectl get pods -n kube-system -l app=aws-node
   ```

2. Review pod distribution:
   ```bash
   kubectl get pods --all-namespaces -o wide
   ```

**Solutions**:
- Increase pool size: `export ENI_POOL_SIZE="50"`
- Enable pool scaling: `export ENABLE_POOL_SCALING="true"`
- Review pod scheduling policies
- Check for resource leaks

### **Security Group Issues**

#### **Issue**: Security group validation fails

**Symptoms**:
- Error: `SECURITY_GROUP_VALIDATION_FAILED`
- Pods cannot start
- Network connectivity issues

**Diagnostic Steps**:
1. Verify security group exists:
   ```bash
   aws ec2 describe-security-groups --group-ids sg-xxx
   ```

2. Check security group permissions:
   ```bash
   aws ec2 describe-security-groups --group-ids sg-xxx --query 'SecurityGroups[0].IpPermissions'
   ```

3. Validate pod annotations:
   ```bash
   kubectl get pod pod-name -o yaml | grep security-groups
   ```

**Solutions**:
- Verify security group IDs are correct
- Check security group exists in correct VPC
- Validate IAM permissions for EC2 API
- Review security group rules

#### **Issue**: Too many security groups

**Symptoms**:
- Error: `TOO_MANY_SECURITY_GROUPS`
- Pod creation fails
- Validation errors

**Diagnostic Steps**:
1. Count security groups in pod annotation:
   ```bash
   kubectl get pod pod-name -o jsonpath='{.metadata.annotations.vpc\.amazonaws\.com/security-groups}' | tr ',' '\n' | wc -l
   ```

2. Check maximum limit:
   ```bash
   echo $MAX_SECURITY_GROUPS_PER_POD
   ```

**Solutions**:
- Reduce number of security groups per pod
- Increase limit: `export MAX_SECURITY_GROUPS_PER_POD="10"`
- Consolidate security group rules
- Review security group design

### **Configuration Issues**

#### **Issue**: Invalid configuration

**Symptoms**:
- Error: `INVALID_CONFIGURATION`
- Application fails to start
- Missing required settings

**Diagnostic Steps**:
1. Check required environment variables:
   ```bash
   env | grep -E "(CLUSTER_NAME|AWS_REGION|AUTO_MODE_SGPP_ENABLED)"
   ```

2. Validate configuration values:
   ```bash
   echo "Cluster: $CLUSTER_NAME"
   echo "Region: $AWS_REGION"
   echo "Enabled: $AUTO_MODE_SGPP_ENABLED"
   ```

**Solutions**:
- Set required environment variables
- Validate configuration values
- Check configuration format
- Review configuration documentation

#### **Issue**: Configuration not loaded

**Symptoms**:
- Default values being used
- Configuration changes not applied
- Unexpected behavior

**Diagnostic Steps**:
1. Check environment variable loading:
   ```bash
   printenv | grep AUTO_MODE
   ```

2. Verify configuration source:
   ```bash
   kubectl get configmap aws-node -n kube-system -o yaml
   ```

**Solutions**:
- Restart application after configuration changes
- Verify environment variable names
- Check configuration file format
- Review configuration loading order

## 🔧 **Diagnostic Commands**

### **Cluster Information**

```bash
# Check cluster status
aws eks describe-cluster --name your-cluster --region your-region

# Check cluster tags
aws eks describe-cluster --name your-cluster --region your-region --query 'cluster.tags'

# Check cluster version
aws eks describe-cluster --name your-cluster --region your-region --query 'cluster.version'
```

### **ENI Information**

```bash
# List ENIs in subnet
aws ec2 describe-network-interfaces --filters "Name=subnet-id,Values=subnet-xxx"

# Check ENI limits
aws ec2 describe-account-attributes --attribute-names max-network-interfaces

# Check ENI status
aws ec2 describe-network-interfaces --network-interface-ids eni-xxx
```

### **Security Group Information**

```bash
# List security groups
aws ec2 describe-security-groups --group-ids sg-xxx

# Check security group rules
aws ec2 describe-security-groups --group-ids sg-xxx --query 'SecurityGroups[0].IpPermissions'

# Check security group tags
aws ec2 describe-security-groups --group-ids sg-xxx --query 'SecurityGroups[0].Tags'
```

### **Pod Information**

```bash
# Check pod status
kubectl get pods -o wide

# Check pod annotations
kubectl get pod pod-name -o yaml | grep -A 10 annotations

# Check pod events
kubectl describe pod pod-name

# Check pod logs
kubectl logs pod-name -c aws-node
```

## 📊 **Monitoring and Metrics**

### **Key Metrics to Monitor**

1. **Auto Mode Detection**:
   - Detection success rate
   - Detection latency
   - Cache hit rate

2. **ENI Management**:
   - ENI allocation success rate
   - ENI pool utilization
   - Allocation latency

3. **Security Group Management**:
   - Validation success rate
   - Cache hit rate
   - Validation latency

### **Monitoring Commands**

```bash
# Check component health
kubectl get pods -n kube-system -l app=aws-node

# Check metrics
kubectl top pods -n kube-system

# Check logs
kubectl logs -n kube-system -l app=aws-node --tail=100

# Check events
kubectl get events --sort-by=.metadata.creationTimestamp
```

## 🚨 **Error Codes Reference**

### **Configuration Errors**
- `INVALID_CONFIGURATION`: Invalid configuration values
- `MISSING_REQUIRED_CONFIG`: Required configuration missing
- `INVALID_ENVIRONMENT_VARIABLE`: Invalid environment variable format

### **Auto Mode Detection Errors**
- `AUTO_MODE_DETECTION_FAILED`: Auto Mode detection failed
- `CLUSTER_NOT_FOUND`: Cluster not found
- `INVALID_CLUSTER_MODE`: Invalid cluster mode

### **ENI Management Errors**
- `ENI_ALLOCATION_FAILED`: ENI allocation failed
- `ENI_RELEASE_FAILED`: ENI release failed
- `ENI_POOL_EXHAUSTED`: ENI pool exhausted
- `ENI_NOT_FOUND`: ENI not found

### **Security Group Errors**
- `SECURITY_GROUP_VALIDATION_FAILED`: Security group validation failed
- `SECURITY_GROUP_NOT_FOUND`: Security group not found
- `TOO_MANY_SECURITY_GROUPS`: Too many security groups per pod

### **Resource Errors**
- `RESOURCE_EXHAUSTED`: Resource exhausted
- `RESOURCE_NOT_FOUND`: Resource not found
- `RESOURCE_CONFLICT`: Resource conflict

### **Timeout Errors**
- `TIMEOUT`: Operation timeout
- `ALLOCATION_TIMEOUT`: ENI allocation timeout
- `VALIDATION_TIMEOUT`: Security group validation timeout

## 🔄 **Recovery Procedures**

### **ENI Pool Recovery**

```bash
# Restart ENI manager
kubectl delete pod -n kube-system -l app=aws-node

# Clear ENI pool
kubectl exec -n kube-system deployment/aws-node -- /bin/sh -c "rm -rf /tmp/eni-pool"

# Restart with clean state
kubectl rollout restart deployment/aws-node -n kube-system
```

### **Security Group Cache Recovery**

```bash
# Clear security group cache
kubectl exec -n kube-system deployment/aws-node -- /bin/sh -c "rm -rf /tmp/sg-cache"

# Restart security group manager
kubectl delete pod -n kube-system -l app=aws-node
```

### **Configuration Recovery**

```bash
# Reset to default configuration
kubectl delete configmap aws-node -n kube-system

# Apply default configuration
kubectl apply -f aws-node-configmap.yaml

# Restart pods
kubectl rollout restart deployment/aws-node -n kube-system
```

## 📝 **Best Practices**

### **Prevention**

1. **Regular Monitoring**: Monitor key metrics and logs regularly
2. **Resource Planning**: Plan ENI and security group usage
3. **Configuration Management**: Use configuration management tools
4. **Testing**: Test changes in non-production environments

### **Response**

1. **Quick Diagnosis**: Use diagnostic commands to identify issues
2. **Gradual Recovery**: Apply fixes gradually to avoid cascading failures
3. **Documentation**: Document issues and solutions for future reference
4. **Communication**: Communicate issues and resolutions to team

## 🔗 **Related Documentation**

- [Getting Started](getting-started.md)
- [Configuration](configuration.md)
- [Error Handling API](../api/error-handling.md)
- [Architecture Overview](../architecture/overview.md)
