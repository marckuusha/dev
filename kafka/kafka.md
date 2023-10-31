# Kafka

### Broker
Отвечает за прием сообщений, хранение сообщений, выдача сообщений. 

### Message
key - ключ для распредления сообщений по кластеру
value - массив байт
timestamp - время сообщения
headers - key-value пары

### Topic
Поток данных.

set replication-factor (>1) - для каждой партиции две реплики

два вида реплик: leader и follower

in-sync replicas (ISR) -> в партицию с такой настройкой так же попадают данные из leader реплики. синхронная запись. ISR кандидат, если лидер упадет.

min.insync.replicas=n-1

чтение и запись только с leader репликой

### Producer
Отправляет сообщение. Отправляет только в лидер реплику.

Гарантии доставки
* acks = 0 - не интерисует подтвреждение отправки (могут  теряться сообщения)
* acks = 1 - нужно подтвреждение только от лидер реплики
* acks = -1 (all) - подтверждение от лидера и во все ISR

Опредление отправки в партицию
* explicit partition - конкретная партиция
* round-robin - отдаем выбор партиции кафке или брокеру
* key-defined - партиция определяется по ключу (key_hash % count_node)

Accumulate batch
* batch.size - превышение батча инициирует отправку в брокер
* linger.ms - истечение таймаута инициирует отправку в брокер

### Consumer
Получаем сразу пачку сообшений. Если один консьюмер - подключается ко всем партициям.
Можно объединить консьюмеры в группу

__consumer_offsets - дефолтный топик с данными, кто что прочитал partition, group, offset

offsets.retention.minutes (default = 7 days) - если консьюмер не подключался к партиции в течении этого времени, то оффсет удалется из __consumer_offsets

auto.offset.reset = earliest or latest

Типы коммитов
* автокомит - коммитится как только консьюмер получил
* ручной коммит (возможны дубли)


### Статьи

* https://kafka.apache.org/quickstart - как быстро развернуть кафку
* https://www.cloudkarafka.com/blog/faq-how-to-view-kafka-messages.html - основные команды и как юзать для новичка
* https://www.youtube.com/watch?v=-AZOi3kP9Js&t=2s - основны кафки, круто помогло для старта
* https://kafka.apache.org/documentation.html#gettingStarted - дока по основам кафки
* https://stackoverflow.com/questions/39730126/difference-between-session-timeout-ms-and-max-poll-interval-ms-for-kafka-0-10 - объяснение разницы session.timeout.ms и max.poll.interval.ms
* https://www.confluent.io/en-gb/blog/error-handling-patterns-in-kafka/ - паттерны обработки сообщений с кафкой
* 

* https://nikgalushko.com/post/go_kafka_drivers_competition/ - сравнени либ кафки для go
* https://docs.confluent.io/kafka-clients/go/current/overview.html - дока по кафке
* https://docs.confluent.io/platform/current/clients/confluent-kafka-go/index.html - дока либы confluent для го
* https://docs.confluent.io/platform/current/installation/configuration/index.html - дока либы confluent для го
* https://github.com/confluentinc/confluent-kafka-go/tree/master - сама либа