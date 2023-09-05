## Helm Chart for QLVENTURES APPS


#### How to template the chart?

```
helm template ./qlventures-chart -f ./qlventures-chart/values.yaml --output-dir out --dry-run --debug
```

#### How to install/upgrade the chart ?

Install the chart into specific namespace by modifying the argument -n ${namespace}

```
helm upgrade --install -n staging ./qlventures-chart -f ./qlventures-chart/values.yaml
```

```
helm ls -n staging
```

### Helm Lint Validation

```
helm lint ./qlventures-chart
```

### Configuration Options

TBA