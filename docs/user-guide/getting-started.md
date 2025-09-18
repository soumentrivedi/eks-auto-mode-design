# EKS Auto Mode Security Groups per Pod - User Guide

## Overview

This guide explains how to use the EKS Auto Mode Security Groups per Pod (SGPP) enhancement, which enables fine-grained network security controls in EKS Auto Mode clusters.

## Prerequisites

- EKS cluster running in Auto Mode
- VPC CNI plugin version with Auto Mode SGPP support
- Appropriate AWS IAM permissions for ENI and security group management

## Configuration

### Environment Variables

Set the following environment variables in your VPC CNI plugin configuration:

#### Required Variables
```bash
export CLUSTER_NAME="your-auto-mode-cluster"
export AWS_REGION="us-west-2"
```

#### Optional Variables
```bash
export AUTO_MODE_SGPP_ENABLED="true"
export AUTO_MODE_ENI_POOL_SIZE="10"
export AUTO_MODE_SG_FALLBACK="sg-cluster-default"
export AUTO_MODE_ENI_ALLOCATION_STRATEGY="pool-based"
export AUTO_MODE_CACHE_TTL="300"
export AUTO_MODE_LOG_LEVEL="info"
```

### ConfigMap Configuration

Update your VPC CNI plugin ConfigMap:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: aws-node
  namespace: kube-system
data:
  # Existing VPC CNI configuration
  ENABLE_POD_ENI: "true"
  POD_SECURITY_GROUP_ENFORCING_MODE: "strict"
  
  # Auto Mode SGPP configuration
  AUTO_MODE_SGPP_ENABLED: "true"
  AUTO_MODE_ENI_POOL_SIZE: "10"
  AUTO_MODE_SG_FALLBACK: "sg-cluster-default"
  AUTO_MODE_ENI_ALLOCATION_STRATEGY: "pool-based"
  AUTO_MODE_CACHE_TTL: "300"
  AUTO_MODE_LOG_LEVEL: "info"
```

## Usage

### Pod Security Group Annotation

To assign a specific security group to a pod, use the following annotation:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
  namespace: default
  annotations:
    vpc.amazonaws.com/security-groups: "sg-12345678"
spec:
  containers:
  - name: app
    image: nginx:latest
```

### Multiple Security Groups

You can assign multiple security groups to a pod (up to 5):

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
  namespace: default
  annotations:
    vpc.amazonaws.com/security-groups: "sg-12345678,sg-87654321"
spec:
  containers:
  - name: app
    image: nginx:latest
```

### Deployment Example

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-app
  namespace: production
spec:
  replicas: 3
  selector:
    matchLabels:
      app: web-app
  template:
    metadata:
      labels:
        app: web-app
      annotations:
        vpc.amazonaws.com/security-groups: "sg-web-app"
    spec:
      containers:
      - name: web-app
        image: nginx:latest
        ports:
        - containerPort: 80
```

### Service Example

```yaml
apiVersion: v1
kind: Service
metadata:
  name: web-app-service
  namespace: production
spec:
  selector:
    app: web-app
  ports:
  - port: 80
    targetPort: 80
  type: ClusterIP
```

## Security Group Management

### Creating Security Groups

Create security groups for your applications:

```bash
# Web application security group
aws ec2 create-security-group \
  --group-name web-app-sg \
  --description "Security group for web application" \
  --vpc-id vpc-12345678

# Database security group
aws ec2 create-security-group \
  --group-name db-sg \
  --description "Security group for database" \
  --vpc-id vpc-12345678
```

### Security Group Rules

Configure security group rules:

```bash
# Allow HTTP traffic
aws ec2 authorize-security-group-ingress \
  --group-id sg-web-app \
  --protocol tcp \
  --port 80 \
  --cidr 0.0.0.0/0

# Allow HTTPS traffic
aws ec2 authorize-security-group-ingress \
  --group-id sg-web-app \
  --protocol tcp \
  --port 443 \
  --cidr 0.0.0.0/0

# Allow database access from web app
aws ec2 authorize-security-group-ingress \
  --group-id sg-db \
  --protocol tcp \
  --port 5432 \
  --source-group sg-web-app
```

## Monitoring and Troubleshooting

### Check Auto Mode Status

```bash
# Check if Auto Mode is enabled
kubectl get configmap aws-node -n kube-system -o yaml | grep AUTO_MODE_SGPP_ENABLED

# Check cluster mode
aws eks describe-cluster --name your-cluster-name --query 'cluster.tags'
```

### View ENI Pool Status

```bash
# Check ENI pool status
kubectl logs -n kube-system -l app=aws-node | grep "ENI pool status"
```

### Pod Security Group Assignment

```bash
# Check pod security group assignment
kubectl describe pod my-pod -n default | grep -A 5 "Annotations"

# Check ENI allocation
kubectl logs -n kube-system -l app=aws-node | grep "ENI allocated"
```

### Common Issues

#### Issue: Pod not getting security group
**Symptoms**: Pod starts but doesn't have the expected security group
**Solution**: 
1. Verify the security group annotation is correct
2. Check if the security group exists and is accessible
3. Ensure Auto Mode SGPP is enabled

#### Issue: ENI allocation failures
**Symptoms**: Pods stuck in Pending state
**Solution**:
1. Check ENI limits for your instance type
2. Verify subnet has available IP addresses
3. Check security group permissions

#### Issue: Performance degradation
**Symptoms**: Slower pod startup times
**Solution**:
1. Adjust ENI pool size based on workload
2. Consider using different allocation strategy
3. Monitor ENI pool utilization

## Best Practices

### Security Group Design

1. **Principle of Least Privilege**: Only allow necessary traffic
2. **Layered Security**: Use multiple security groups for different layers
3. **Regular Review**: Periodically review and update security group rules

### ENI Pool Management

1. **Right-sizing**: Set appropriate ENI pool size based on workload
2. **Monitoring**: Monitor ENI pool utilization and adjust as needed
3. **Cleanup**: Ensure proper cleanup of unused ENIs

### Pod Annotations

1. **Consistency**: Use consistent annotation patterns across deployments
2. **Documentation**: Document security group assignments
3. **Validation**: Validate security group IDs before deployment

## Advanced Configuration

### Custom Allocation Strategy

```bash
# Use on-demand allocation for dynamic workloads
export AUTO_MODE_ENI_ALLOCATION_STRATEGY="on-demand"

# Use hybrid strategy for mixed workloads
export AUTO_MODE_ENI_ALLOCATION_STRATEGY="hybrid"
```

### Cache Configuration

```bash
# Adjust cache TTL for different environments
export AUTO_MODE_CACHE_TTL="600"  # 10 minutes for production
export AUTO_MODE_CACHE_TTL="60"   # 1 minute for development
```

### Logging Configuration

```bash
# Enable debug logging for troubleshooting
export AUTO_MODE_LOG_LEVEL="debug"

# Reduce logging for production
export AUTO_MODE_LOG_LEVEL="warn"
```

## Migration Guide

### From Standard Mode to Auto Mode

1. **Prepare Security Groups**: Create necessary security groups
2. **Update Pod Annotations**: Add security group annotations to pods
3. **Test Configuration**: Test in a development environment first
4. **Deploy Gradually**: Migrate workloads gradually

### From Manual ENI Management

1. **Audit Current Setup**: Review existing ENI configurations
2. **Map Security Groups**: Map current security groups to pods
3. **Update Annotations**: Add security group annotations
4. **Monitor Performance**: Monitor performance during migration

## Support and Resources

### Documentation
- [AWS EKS User Guide](https://docs.aws.amazon.com/eks/)
- [VPC CNI Plugin Documentation](https://github.com/aws/amazon-vpc-cni-k8s)
- [Enhanced VPC CNI Plugin (Auto Mode SGPP)](https://github.com/soumentrivedi/amazon-vpc-cni-k8s)
- [Security Groups Documentation](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html)

### Community
- [AWS EKS GitHub](https://github.com/aws/eks-cluster)
- [VPC CNI GitHub](https://github.com/aws/amazon-vpc-cni-k8s)
- [Enhanced VPC CNI GitHub](https://github.com/soumentrivedi/amazon-vpc-cni-k8s)
- [AWS Forums](https://forums.aws.amazon.com/forum.jspa?forumID=30)

### Support
- AWS Support for production issues
- GitHub Issues for bug reports and feature requests
- AWS re:Post for community questions
