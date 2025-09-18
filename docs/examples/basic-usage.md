# EKS Auto Mode SGPP Enhancement - Examples

## Basic Usage Examples

### Simple Pod with Security Group

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: web-server
  namespace: default
  annotations:
    vpc.amazonaws.com/security-groups: "sg-web-server"
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
    - containerPort: 80
```

### Deployment with Security Group

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-server
  namespace: production
spec:
  replicas: 3
  selector:
    matchLabels:
      app: api-server
  template:
    metadata:
      labels:
        app: api-server
      annotations:
        vpc.amazonaws.com/security-groups: "sg-api-server"
    spec:
      containers:
      - name: api-server
        image: my-api:latest
        ports:
        - containerPort: 8080
        env:
        - name: DATABASE_URL
          value: "postgresql://db:5432/mydb"
```

## Multi-Tier Application Example

### Web Tier

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-tier
  namespace: ecommerce
spec:
  replicas: 5
  selector:
    matchLabels:
      app: web-tier
  template:
    metadata:
      labels:
        app: web-tier
      annotations:
        vpc.amazonaws.com/security-groups: "sg-web-tier"
    spec:
      containers:
      - name: web-server
        image: nginx:latest
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: web-tier-service
  namespace: ecommerce
spec:
  selector:
    app: web-tier
  ports:
  - port: 80
    targetPort: 80
  type: LoadBalancer
```

### Application Tier

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-tier
  namespace: ecommerce
spec:
  replicas: 3
  selector:
    matchLabels:
      app: app-tier
  template:
    metadata:
      labels:
        app: app-tier
      annotations:
        vpc.amazonaws.com/security-groups: "sg-app-tier"
    spec:
      containers:
      - name: app-server
        image: my-app:latest
        ports:
        - containerPort: 8080
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: url
---
apiVersion: v1
kind: Service
metadata:
  name: app-tier-service
  namespace: ecommerce
spec:
  selector:
    app: app-tier
  ports:
  - port: 8080
    targetPort: 8080
  type: ClusterIP
```

### Database Tier

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: db-tier
  namespace: ecommerce
spec:
  serviceName: db-tier-service
  replicas: 1
  selector:
    matchLabels:
      app: db-tier
  template:
    metadata:
      labels:
        app: db-tier
      annotations:
        vpc.amazonaws.com/security-groups: "sg-db-tier"
    spec:
      containers:
      - name: postgres
        image: postgres:13
        ports:
        - containerPort: 5432
        env:
        - name: POSTGRES_DB
          value: "ecommerce"
        - name: POSTGRES_USER
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: username
        - name: POSTGRES_PASSWORD
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: password
        volumeMounts:
        - name: db-storage
          mountPath: /var/lib/postgresql/data
  volumeClaimTemplates:
  - metadata:
      name: db-storage
    spec:
      accessModes: ["ReadWriteOnce"]
      resources:
        requests:
          storage: 10Gi
---
apiVersion: v1
kind: Service
metadata:
  name: db-tier-service
  namespace: ecommerce
spec:
  selector:
    app: db-tier
  ports:
  - port: 5432
    targetPort: 5432
  type: ClusterIP
```

## Microservices Example

### API Gateway

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-gateway
  namespace: microservices
spec:
  replicas: 2
  selector:
    matchLabels:
      app: api-gateway
  template:
    metadata:
      labels:
        app: api-gateway
      annotations:
        vpc.amazonaws.com/security-groups: "sg-api-gateway"
    spec:
      containers:
      - name: kong
        image: kong:latest
        ports:
        - containerPort: 8000
        - containerPort: 8443
        env:
        - name: KONG_DATABASE
          value: "off"
        - name: KONG_DECLARATIVE_CONFIG
          value: "/kong/declarative/kong.yml"
```

### User Service

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service
  namespace: microservices
spec:
  replicas: 3
  selector:
    matchLabels:
      app: user-service
  template:
    metadata:
      labels:
        app: user-service
      annotations:
        vpc.amazonaws.com/security-groups: "sg-user-service"
    spec:
      containers:
      - name: user-service
        image: user-service:latest
        ports:
        - containerPort: 3000
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: user-db-secret
              key: url
```

### Order Service

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order-service
  namespace: microservices
spec:
  replicas: 2
  selector:
    matchLabels:
      app: order-service
  template:
    metadata:
      labels:
        app: order-service
      annotations:
        vpc.amazonaws.com/security-groups: "sg-order-service"
    spec:
      containers:
      - name: order-service
        image: order-service:latest
        ports:
        - containerPort: 3001
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: order-db-secret
              key: url
        - name: USER_SERVICE_URL
          value: "http://user-service-service:3000"
```

## Security Group Configuration Examples

### Web Application Security Group

```bash
#!/bin/bash

# Create web application security group
WEB_SG_ID=$(aws ec2 create-security-group \
  --group-name web-app-sg \
  --description "Security group for web application" \
  --vpc-id vpc-12345678 \
  --query 'GroupId' \
  --output text)

# Allow HTTP traffic
aws ec2 authorize-security-group-ingress \
  --group-id $WEB_SG_ID \
  --protocol tcp \
  --port 80 \
  --cidr 0.0.0.0/0

# Allow HTTPS traffic
aws ec2 authorize-security-group-ingress \
  --group-id $WEB_SG_ID \
  --protocol tcp \
  --port 443 \
  --cidr 0.0.0.0/0

# Allow traffic from load balancer
aws ec2 authorize-security-group-ingress \
  --group-id $WEB_SG_ID \
  --protocol tcp \
  --port 80 \
  --source-group sg-load-balancer

echo "Web application security group created: $WEB_SG_ID"
```

### Database Security Group

```bash
#!/bin/bash

# Create database security group
DB_SG_ID=$(aws ec2 create-security-group \
  --group-name db-sg \
  --description "Security group for database" \
  --vpc-id vpc-12345678 \
  --query 'GroupId' \
  --output text)

# Allow PostgreSQL traffic from application tier
aws ec2 authorize-security-group-ingress \
  --group-id $DB_SG_ID \
  --protocol tcp \
  --port 5432 \
  --source-group sg-app-tier

# Allow MySQL traffic from application tier
aws ec2 authorize-security-group-ingress \
  --group-id $DB_SG_ID \
  --protocol tcp \
  --port 3306 \
  --source-group sg-app-tier

echo "Database security group created: $DB_SG_ID"
```

### API Gateway Security Group

```bash
#!/bin/bash

# Create API gateway security group
API_SG_ID=$(aws ec2 create-security-group \
  --group-name api-gateway-sg \
  --description "Security group for API gateway" \
  --vpc-id vpc-12345678 \
  --query 'GroupId' \
  --output text)

# Allow HTTP traffic
aws ec2 authorize-security-group-ingress \
  --group-id $API_SG_ID \
  --protocol tcp \
  --port 8000 \
  --cidr 0.0.0.0/0

# Allow HTTPS traffic
aws ec2 authorize-security-group-ingress \
  --group-id $API_SG_ID \
  --protocol tcp \
  --port 8443 \
  --cidr 0.0.0.0/0

# Allow traffic from microservices
aws ec2 authorize-security-group-ingress \
  --group-id $API_SG_ID \
  --protocol tcp \
  --port 8000 \
  --source-group sg-microservices

echo "API gateway security group created: $API_SG_ID"
```

## Configuration Examples

### VPC CNI Plugin ConfigMap

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: aws-node
  namespace: kube-system
data:
  # Standard VPC CNI configuration
  ENABLE_POD_ENI: "true"
  POD_SECURITY_GROUP_ENFORCING_MODE: "strict"
  ENABLE_PREFIX_DELEGATION: "true"
  WARM_IP_TARGET: "2"
  MINIMUM_IP_TARGET: "1"
  
  # Auto Mode SGPP configuration
  AUTO_MODE_SGPP_ENABLED: "true"
  AUTO_MODE_ENI_POOL_SIZE: "20"
  AUTO_MODE_SG_FALLBACK: "sg-cluster-default"
  AUTO_MODE_ENI_ALLOCATION_STRATEGY: "pool-based"
  AUTO_MODE_CACHE_TTL: "300"
  AUTO_MODE_LOG_LEVEL: "info"
  
  # Cluster information
  CLUSTER_NAME: "my-auto-mode-cluster"
  AWS_REGION: "us-west-2"
```

### DaemonSet Configuration

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: aws-node
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: aws-node
  template:
    metadata:
      labels:
        app: aws-node
    spec:
      containers:
      - name: aws-node
        image: amazon/aws-k8s-cni:v1.12.0
        env:
        - name: CLUSTER_NAME
          value: "my-auto-mode-cluster"
        - name: AWS_REGION
          value: "us-west-2"
        - name: AUTO_MODE_SGPP_ENABLED
          value: "true"
        - name: AUTO_MODE_ENI_POOL_SIZE
          value: "20"
        - name: AUTO_MODE_SG_FALLBACK
          value: "sg-cluster-default"
        - name: AUTO_MODE_ENI_ALLOCATION_STRATEGY
          value: "pool-based"
        - name: AUTO_MODE_CACHE_TTL
          value: "300"
        - name: AUTO_MODE_LOG_LEVEL
          value: "info"
        volumeMounts:
        - name: cni-bin-dir
          mountPath: /opt/cni/bin
        - name: cni-cache-dir
          mountPath: /var/lib/cni
        - name: log-dir
          mountPath: /var/log
      volumes:
      - name: cni-bin-dir
        hostPath:
          path: /opt/cni/bin
      - name: cni-cache-dir
        hostPath:
          path: /var/lib/cni
      - name: log-dir
        hostPath:
          path: /var/log
```

## Monitoring and Observability Examples

### Prometheus Monitoring

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: prometheus-config
  namespace: monitoring
data:
  prometheus.yml: |
    global:
      scrape_interval: 15s
    scrape_configs:
    - job_name: 'aws-node'
      static_configs:
      - targets: ['aws-node:61678']
      metrics_path: /metrics
      scrape_interval: 30s
```

### Grafana Dashboard

```json
{
  "dashboard": {
    "title": "EKS Auto Mode SGPP",
    "panels": [
      {
        "title": "ENI Pool Status",
        "type": "stat",
        "targets": [
          {
            "expr": "aws_node_eni_pool_total",
            "legendFormat": "Total ENIs"
          }
        ]
      },
      {
        "title": "Security Group Assignments",
        "type": "graph",
        "targets": [
          {
            "expr": "aws_node_security_group_assignments_total",
            "legendFormat": "Security Group Assignments"
          }
        ]
      }
    ]
  }
}
```

## Troubleshooting Examples

### Debug Pod Security Group Assignment

```bash
#!/bin/bash

# Check pod annotations
kubectl get pod my-pod -o jsonpath='{.metadata.annotations.vpc\.amazonaws\.com/security-groups}'

# Check ENI allocation logs
kubectl logs -n kube-system -l app=aws-node | grep "ENI allocated"

# Check security group attachment
kubectl logs -n kube-system -l app=aws-node | grep "Security group attached"

# Check Auto Mode status
kubectl logs -n kube-system -l app=aws-node | grep "Auto Mode"
```

### Validate Security Group Configuration

```bash
#!/bin/bash

# Check if security group exists
aws ec2 describe-security-groups --group-ids sg-12345678

# Check security group rules
aws ec2 describe-security-groups --group-ids sg-12345678 --query 'SecurityGroups[0].IpPermissions'

# Check ENI security groups
aws ec2 describe-network-interfaces --network-interface-ids eni-12345678 --query 'NetworkInterfaces[0].Groups'
```

### Performance Monitoring

```bash
#!/bin/bash

# Check ENI pool utilization
kubectl logs -n kube-system -l app=aws-node | grep "ENI pool status"

# Check allocation performance
kubectl logs -n kube-system -l app=aws-node | grep "ENI allocation time"

# Check cache performance
kubectl logs -n kube-system -l app=aws-node | grep "Cache hit rate"
```
