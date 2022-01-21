# 安装elk

```bash
helm install --dry-run --debug elasticsearch --generate-name 

helm install elasticsearch elasticsearch --set volumes.type=pvc --set replicas=3

```