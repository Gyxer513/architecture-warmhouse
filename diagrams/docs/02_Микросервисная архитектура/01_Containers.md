# Микросервисная архитектура - диаграмма контейнеров.

## Общая концепция
Монолитная система разбита на микросервисы, для увеличения гибкости и отказоустойчивости. 

Логика управления отоплением и устройствами разделена на отдельные сервисы с чёткими зонами ответственности и собственной БД.   

Обмен сообщениями между сервисами реализован через событийную шину (message bus).

---

## Участники системы


 **Пользователь** - управляет устройствами, настройками и просматривает телеметрию через веб-интерфейс.  
 **Администратор** - выполняет административные функции также через веб-интерфейс.  
 
 Техник исключён из архитектуры: пользователи и администраторы самостоятельно регистрируют и настраивают устройства. Процедура автоматизирована.

---

## Основные компоненты
### 1. **Web/Mobile Application**
Единый интерфейс для всех пользователей и админов для доступа к функциям системы.

### 2. **API Gateway**
Центральная точка обработки всех внешних запросов. Отвечает за:
- маршрутизацию
- контроль доступа
- балансировку нагрузки

### 3. **User Service**
Управление учётными записями, регистрация, аутентификация и права.

### 4. **Device Service**
Работа с IoT-устройствами: регистрация, настройка, опрос.

### 5. **Automation Service**
Хранение сценариев, расписаний, автоматизация действий.

### 6. **Telemetry Service**
Сбор, хранение и выдача информации с датчиков/устройств.

### 7. **Control Service**
Отправка команд на устройства (например, включение/выключение отопления).

### 8. **Partner Integration Service**
Интеграция с внешними облаками и экосистемами.

### 9. **Notification Service**
Оповещения пользователей (email, SMS, push и др.).

### 10. **Message Bus**
Асинхронная шина обмена событиями между сервисами.

---

## Работа с данными

- Для изоляции и изоляции хранимых данных микросервисов используется **отдельная база данных** (PostgreSQL, TSDB, Redis, очереди).

---

## Маршрутизация и взаимодействие

- Все внешние (клиентские) запросы проходят через **API Gateway**.
- Прямой обмен между сервисами — только через **message bus** (RabbitMQ, Kafka, NATS) в виде событий.

> В результате достигается
> слабая связность и
> упрощённое масштабирование

---

## Интеграция с устройствами и облаками

- Сервисы **Telemetry** и **Control** напрямую взаимодействуют с IoT-устройствами по защищённым протоколам (**MQTT/HTTP**).
- Для интеграций с внешними платформами предусмотрен **Partner Integration Service** с поддержкой стандартных API.

---

## Диаграмма контейнеров

```puml
@startuml
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

' Акторы
Person(user, "Пользователь")
Person(admin, "Администратор")

' Внешние устройства/системы
System_Ext(iot_device, "IoT-устройство", "Датчики, исполнительные устройства")
System_Ext(partner_cloud, "Облачная платформа партнёров", "Интеграция с внешними экосистемами")

' Frontend
Container(web, "Web/Mobile Application", "SPA (React/Vue)", "UI для пользователей и администраторов")

' API Gateway
Container(api_gw, "API Gateway", "NGINX/Kong/Custom", "Маршрутизация, ACL, throttling, входная точка")

' Микросервисы
Container(user_srv, "User Service", "Go/REST", "Аккаунты, аутентификация, права")
Container(device_srv, "Device Service", "Go/REST", "Управление устройствами, их настройка")
Container(automation_srv, "Automation Service", "Go/NodeJS", "Сценарии, расписания")
Container(telemetry_srv, "Telemetry Service", "Go", "Cбор и хранение телеметрии устройств, аналитика")
Container(control_srv, "Control Service", "Go", "Выполнение команд управления устройствами")
Container(partner_srv, "Partner Integration", "Go/Python", "Интеграция с внешними системами/устройствами")
Container(notification_srv, "Notification Service", "Node.js", "Оповещения пользователей (email/SMS/push)")

' Базы данных
ContainerDb(user_db, "Users DB", "PostgreSQL", "Профили и права пользователей")
ContainerDb(device_db, "Devices DB", "PostgreSQL", "Устройства и их статусы")
ContainerDb(telemetry_db, "Telemetry DB", "InfluxDB/Clickhouse", "История датчиков и событий")
ContainerDb(automation_db, "Automation DB", "MongoDB/Redis", "Хранение сценариев и очередей")
ContainerDb(notification_db, "Notifications DB", "PostgreSQL/Queue", "Статус уведомлений и логи")

' Message Broker
Container(message_bus, "Message Bus", "RabbitMQ/Kafka/NATS", "Обмен сообщениями/событиями между сервисами")

' Взаимодействия пользователей
Rel(user, web, "Взаимодействие через UI", "HTTPS")
Rel(admin, web, "Использование админ-функций", "HTTPS")
Rel(web, api_gw, "API-запросы", "REST/HTTPS")

' Взаимодействия API Gateway и сервисов
Rel(api_gw, user_srv, "Аутентификация/авторизация", "REST/gRPC")
Rel(api_gw, device_srv, "Управление устройствами", "REST/gRPC")
Rel(api_gw, automation_srv, "Работа с сценариями", "REST/gRPC")
Rel(api_gw, telemetry_srv, "Запрос исторических данных", "REST/gRPC")
Rel(api_gw, control_srv, "Команды управления устройствами", "REST/gRPC")
Rel(api_gw, notification_srv, "Статус и настройки уведомлений", "REST/gRPC")

' Микросервисы и БД
Rel(user_srv, user_db, "Чтение/запись", "SQL")
Rel(device_srv, device_db, "Чтение/запись", "SQL")
Rel(telemetry_srv, telemetry_db, "Чтение/запись", "TSDB")
Rel(automation_srv, automation_db, "Чтение/запись", "NoSQL/Redis")
Rel(notification_srv, notification_db, "Чтение/запись", "SQL/Queue")

' Message bus (events)
Rel(automation_srv, message_bus, "Публикация событий и сценариев", "event")
Rel(telemetry_srv, message_bus, "Публикация событий телеметрии", "event")
Rel(control_srv, message_bus, "Команды управления", "event")
Rel(notification_srv, message_bus, "Подписка на задачи для отправки уведомлений", "event")
Rel(partner_srv, message_bus, "Публикация/подписка внешних событий", "event")

' Сервис-устройство (IoT)
Rel(control_srv, iot_device, "Управление, команды", "MQTT/HTTP")
Rel(telemetry_srv, iot_device, "Сбор телеметрии", "MQTT/HTTP")

' Интеграция с внешними платформами
Rel(partner_srv, partner_cloud, "Двусторонний обмен", "REST/MQTT/стандарты")

@enduml
```