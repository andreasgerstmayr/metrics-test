## get keycloak admin password
```
kubectl get secret keycloak-initial-admin -o jsonpath='{.data.username}' | base64 --decode
kubectl get secret keycloak-initial-admin -o jsonpath='{.data.password}' | base64 --decode
```

## generate traces
```
telemetrygen traces --traces=1 --otlp-http --otlp-insecure --otlp-endpoint=otlp-http-otel-client-route-test.apps-crc.testing
```
