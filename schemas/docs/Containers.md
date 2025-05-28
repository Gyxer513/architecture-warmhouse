# Новая микросервисная архитектура

## 
```puml
@startuml
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
LAYOUT_WITH_LEGEND()
' Actors
Person(user, "User", "Обычный пользователь системы")
Person(admin, "Admin", "Администратор")

' Frontend
Container(web, "Web Application", "SPA (React/Vue/etc)", "UI для управления домом")

' Микросервисы
Container(api_gw, "API Gateway", "NGINX/Kong/Custom", "Единая входная точка для внешних клиентов")
Container(user_srv, "User Service", "Go/REST", "Регистрация, авторизация, управление аккаунтом")
Container(device_srv, "Device Service", "Go/REST", "Добавление/удаление/управление устройствами")
Container(automation_srv, "Automation Service", "Go/NodeJS", "Сценарии, расписания, обработки событий")
Container(telemetry_srv, "Telemetry Service", "Go", "Сбор телеметрии (температура и др.)")
Container(control_srv, "Control Service", "Go", "Управление командами устройств")
Container(partner_srv, "Partner Integration", "Go/Python", "Интеграция сторонних устройств по стандартам")
Container(notification_srv, "Notification Service", "Node.js", "Уведомления пользователя")

' Data
ContainerDb(user_db, "Users DB", "PostgreSQL", "Пользователи, профили, права")
ContainerDb(device_db, "Devices DB", "PostgreSQL", "Устройства, параметры, привязки")
ContainerDb(telemetry_db, "Telemetry DB", "InfluxDB/Clickhouse", "Лог телеметрии")
ContainerDb(automation_db, "Automation DB", "MongoDB/Redis", "Сценарии, расписания")
ContainerDb(notification_db, "Notifications DB", "PostgreSQL/Queue", "Статус уведомлений")

' Message broker
Container(message_bus, "Message Bus", "RabbitMQ/Kafka/NATS", "Асинхронная событийная шина")

' Внешний мир (Partner device/cloud)
System_Ext(partner_cloud, "Partner Cloud", "Сторонние экосистемы устройств")
System_Ext(iot_device, "IoT Device", "Физические устройства, установленные в доме")

' Взаимодействия
Rel(user, web, "Использует", "HTTPS")
Rel(web, api_gw, "REST API", "HTTPS")
Rel(admin, api_gw, "REST API", "HTTPS")

Rel(api_gw, user_srv, "Реализует методы авторизации", "REST/gRPC")
Rel(api_gw, device_srv, "Опрашивает инфо об устройствах", "REST/gRPC")
Rel(api_gw, automation_srv, "Вызовы сценариев и управлений", "REST/gRPC")
Rel(api_gw, telemetry_srv, "Просмотр истории сенсоров", "REST/gRPC")
Rel(api_gw, control_srv, "Команды управления устройствами", "REST/gRPC")

Rel(device_srv, device_db, "Чтение/запись", "SQL")
Rel(user_srv, user_db, "Чтение/запись", "SQL")
Rel(automation_srv, automation_db, "Чтение/запись", "NoSQL/Redis")
Rel(telemetry_srv, telemetry_db, "Чтение/запись", "TSDB")
Rel(notification_srv, notification_db, "Чтение/запись", "SQL")

Rel(automation_srv, message_bus, "Посылает события", "event")
Rel(telemetry_srv, message_bus, "Посылает события", "event")
Rel(control_srv, message_bus, "Отправляет команды", "event")
Rel(notification_srv, message_bus, "Получает задачи для уведомления", "event")

Rel(control_srv, iot_device, "Команды управляющих сигналов", "MQTT/HTTP")
Rel(telemetry_srv, iot_device, "Получает телеметрию", "MQTT/HTTP")
Rel(partner_srv, partner_cloud, "Интеграция", "REST/MQTT/стандарты")

@enduml
```