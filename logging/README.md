
## Setup

Execute the following command to install the EFK stack (`Elasticsearch`, `Kibana` and `Fluentd`):

```
kubectl apply -f . -n logging-system
```

You can watch the status of the created Pods being started:

```
kubectl get pods --selector='app in (elasticsearch,kibana,fluentd)' -n logging-system -w
```

## Kibana

You can get access to Kibana UI by executing a `kubectl port-forward` using this command:

```
kubectl port-forward $(kubectl get pod -o=name -lapp=kibana -n logging-system | awk -F/ '{print $2}' | head -1) 5601:5601 -n logging-system 
```

To be able to browse the logs, you need to perform these configuration steps on Kibana:

1. Create a new index named `logstash-*` on [this link](http://localhost:5601/app/kibana#/management/kibana/index_pattern?_g=()).
2. Click Next
3. Select `@timestamp` as Time Filter field.
4. Click on `Create index pattern` and wait for Elasticsearch to collect the index fields.
5. You can now view containers' logs following this [link](http://localhost:5601/app/kibana#/discover?_g=())


Here is a sample of a Kibana [Query](http://localhost:5601/app/kibana#/discover?_g=()&_a=(columns:!(kubernetes.container_name,log),filters:!((%27$state%27:(store:appState),meta:(alias:!n,disabled:!f,index:%2726471e70-df7a-11e9-8789-9958a9ac2aed%27,key:kubernetes.namespace_name,negate:!f,params:(query:strimzi),type:phrase,value:strimzi),query:(match:(kubernetes.namespace_name:(query:strimzi,type:phrase)))),(%27$state%27:(store:appState),meta:(alias:!n,disabled:!f,index:%2726471e70-df7a-11e9-8789-9958a9ac2aed%27,key:stream,negate:!f,params:(query:stderr),type:phrase,value:stderr),query:(match:(stream:(query:stderr,type:phrase)))),(%27$state%27:(store:appState),meta:(alias:!n,disabled:!f,index:%2726471e70-df7a-11e9-8789-9958a9ac2aed%27,key:kubernetes.container_name,negate:!t,params:(query:grafana-operator),type:phrase,value:grafana-operator),query:(match:(kubernetes.container_name:(query:grafana-operator,type:phrase))))),index:%2726471e70-df7a-11e9-8789-9958a9ac2aed%27,interval:auto,query:(language:kuery,query:%27%27),sort:!(%27@timestamp%27,desc)))

## Teardown

To cleanup the whole `logging-system` namespace and the created objects, execute the command: 

```
kubectl delete -f . -n logging-system
```
