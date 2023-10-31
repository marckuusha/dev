
bin/kafka-topics.sh --create --topic hrtech_hr_employee_ib --bootstrap-server localhost:9092


### TOPIC

# create topic
bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic my-topic-3 --partitions 3 --replication-factor 1

# info topic
bin/kafka-topics.sh \
--topic my-topic_2 --describe \
--bootstrap-server localhost:9092

# topic list
bin/kafka-topics.sh --bootstrap-server localhost:9092 --list 


### PRODUCER

# send message
bin/kafka-console-producer.sh --topic my-topic-3 --bootstrap-server localhost:9092


### CONSUMER

# listener message
bin/kafka-console-consumer.sh --topic my-topic-3 --group first-group --from-beginning --bootstrap-server localhost:9092

bin/kafka-console-consumer.sh --bootstrap-server wbc-kafka-1.alljobswb-stage.vm.prod-4.cloud.datapro:9092,wbc-kafka-2.alljobswb-stage.vm.prod-5.cloud.dataline:9092,wbc-kafka-3.alljobswb-stage.vm.prod-6.cloud.el:9092 --consumer.config consumer.properties --topic hrtech_org_structure_sync --group first-group

# Read topic
bin/kafka-console-consumer.sh --topic __consumer_offsets --from-beginning --bootstrap-server localhost:9092

bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic my-topic_2 --from-beginning --group my-created-consumer-group
