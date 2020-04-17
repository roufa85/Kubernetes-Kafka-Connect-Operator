#!/bin/bash

oc login ${OPENSHIFT_URL} -u ${OPENSHIFT_USR} -p ${OPENSHIFT_PSW} ${OPENSHIFT_URL} --insecure-skip-tls-verify=true

oc new-project ${OPENSHIFT_NS}

#####################################################################################
#####################################################################################

oc apply -f strimzi/cluster-operator/00-service-account/

oc apply -f strimzi/cluster-operator/01-rbacs/

oc apply -f strimzi/cluster-operator/02-crds/

oc apply -f strimzi/cluster-operator/03-deployments/

oc rollout status deployment/strimzi-cluster-operator

oc apply -f strimzi/cluster-operator/04-crs/

sleep 10s;

oc get sts/ephemeral-zookeeper --watch --request-timeout=30s

oc get sts/ephemeral-kafka --watch --request-timeout=30s

#####################################################################################
#####################################################################################

oc apply -f strimzi/topic-operator/00-service-account/

oc apply -f strimzi/topic-operator/01-rbacs/

oc apply -f strimzi/topic-operator/02-crds/

oc apply -f strimzi/topic-operator/03-deployments/00-zoo-entrance.yaml

oc rollout status deployment/zoo-entrance

oc apply -f strimzi/topic-operator/03-deployments/01-ephemeral-topic-operator-deployment.yaml

oc rollout status deployment/ephemeral-topic-operator

#####################################################################################
#####################################################################################

oc apply -f strimzi/topic-operator/04-crs/00-create-test-topic-kt.yaml

oc get kt

sleep 5s;
oc exec -it ephemeral-kafka-0 -c kafka -- bin/kafka-topics.sh --zookeeper localhost:2181 --describe --topic test-topic

#####################################################################################
#####################################################################################

oc apply -f elastic/

#####################################################################################
#####################################################################################

oc apply -f deploy/service_account.yaml
oc apply -f deploy/crds/
oc apply -f deploy/role.yaml
oc apply -f deploy/role_binding.yaml
oc apply -f deploy/cluster_role.yaml
oc apply -f deploy/cluster_role_binding.yaml
oc apply -f deploy/operator.yaml
oc rollout status deployment/kubernetes-kafka-connect-operator

#####################################################################################
#####################################################################################

oc apply -f examples/v1alpha1/experiment-ssp.yaml
oc rollout status deployment/experiment-ssp

curl -k -XPUT http://localhost:8083/connectors/elastic-sink-test/config/ -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{
    "connector.class": "io.confluent.connect.elasticsearch.ElasticsearchSinkConnector",
    "tasks.max": "2",
    "topics": "test-topic",
    "key.ignore": "true",
    "connection.url": "http://elastic-server:9200",
    "type.name": "index",
    "name": "elastic-sink-test",
    "schema.ignore": "true",
    "value.converter": "org.apache.kafka.connect.json.JsonConverter",
    "value.converter.schemas.enable": "false"
}'

echo '{"tasks.max": "xx", "key.ignore": "YO"}' | oc exec -i -c kafka ephemeral-kafka-0 -- /opt/kafka/bin/kafka-console-producer.sh \
                    --broker-list ephemeral-kafka-bootstrap:9092 \
                    --topic test-topic

oc exec -i -c kafka ephemeral-kafka-0 -- /opt/kafka/bin/kafka-console-consumer.sh \
    --bootstrap-server ephemeral-kafka-bootstrap:9092 \
    --topic test-topic --from-beginning

curl -k -XPUT http://localhost:8083/connectors/connector-elastic/config/ -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{
    "connector.class": "io.confluent.connect.elasticsearch.ElasticsearchSinkConnector",
    "tasks.max": "2",
    "topics": "test-topic",
    "key.ignore": "true",
    "connection.url": "http://elastic-svc:9200",
    "type.name": "index",
    "name": "connector-elastic",
    "schema.ignore": "true",
    "value.converter": "org.apache.kafka.connect.json.JsonConverter",
    "value.converter.schemas.enable": "false"
}'

curl -k -XDELETE http://localhost:8083/connectors/connector-elastic/



curl -k -XPUT http://localhost:8083/connectors/connector-elastic/config/ -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{
    "connector.class": "io.confluent.connect.elasticsearch.ElasticsearchSinkConnector",
    "tasks.max": "2",
    "topics": "test-topic",
    "key.ignore": "true",
    "connection.url": "http://elastic-svc:9200",
    "type.name": "index",
    "name": "connector-elastic",
    "schema.ignore": "true",
    "value.converter": "org.apache.kafka.connect.json.JsonConverter",
    "value.converter.schemas.enable": "false"
}'