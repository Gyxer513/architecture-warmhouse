# Текущая архитектура

## Краткое описание

**Монолитное приложение** реализовано на Go с использованием базы данных PostgreSQL.

Главные функции:
- Управление отоплением (включить/выключить)
- Мониторинг температуры

### Проблемы архитектуры

- Монолитная архитектура: всё в одном приложении
- Нет разделения по доменам (Domain-Driven Design не реализован)
- Подключение устройств — только через специалиста
- Синхронное взаимодействие и невозможность горизонтального масштабирования
- Невозможность гибкой интеграции новых устройств

---

## Контекстная диаграмма: Текущий монолит

```puml
@startuml Current_Context_Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

Person(admin, "Администратор", "Настраивает систему")
Person(user, "Пользователь", "Включает/выключает отопление, смотрит температуру")
Person(technician, "Техник", "Подключает датчики вручную")

System(monolith, "Монолитное приложение", "Go + PostgreSQL", "Управление отоплением, температура, устройства")
SystemDb(database, "PostgreSQL", "Реляционная база данных", "Хранит пользователей, устройства, историю температур")

System_Ext(sensor, "Датчик температуры", "Отправляет данные по HTTP")

Rel(admin, monolith, "Настройка параметров", "HTTP")
Rel(user, monolith, "Запросы температуры/отопления", "HTTP")
Rel(technician, monolith, "Регистрация датчиков", "HTTP")
Rel(monolith, database, "Чтение/запись данных", "SQL")
Rel(monolith, sensor, "Опрос температуры", "HTTP")
Rel_Back(sensor, monolith, "Ответ с температурой", "HTTP")

legend right
**Легенда:**
→ — Исходящий запрос (HTTP, SQL)
← — Ответ (HTTP)
endlegend

@enduml
```

## Компонентная диаграмма: Текущий монолит

```puml
@startuml Current_Component_Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

' Контур всей системы
Container(monolith, "Монолитное приложение", "Go", "Выполняет управление отоплением и мониторинг температуры")

' Внутренние компоненты
Component(api, "Web API", "REST API", "Обрабатывает запросы пользователей, админов, техников")
Component(user, "User Management", "Управление пользователями", "Аутентификация, авторизация, работа с ролями")
Component(device, "Device Management", "Управление устройствами", "Регистрация устройств, привязка к пользователям")
Component(temp, "Temperature Monitoring", "Мониторинг температуры", "Опрос датчиков, хранение истории температур")
Component(heating, "Heating Control", "Управление отоплением", "Отправка команд на реле отопления")
Component(db, "Database Access Layer", "Работа с PostgreSQL", "ORM, репозитории, SQL")
Component(ext, "External Interfaces", "Протоколы взаимодействия", "Связь с датчиками и реле")

SystemDb(postgresql, "PostgreSQL", "Реляционная БД", "Хранит пользователей, устройства, историю температур")
System_Ext(sensor, "Датчик температуры", "Физический датчик")
System_Ext(relay, "Реле отопления", "Физическое реле управления отоплением")

' Связи
Rel(api, user, "Вызовы", "внутренние Go вызовы")
Rel(api, device, "Вызовы", "внутренние Go вызовы")
Rel(api, temp, "Вызовы", "внутренние Go вызовы")
Rel(api, heating, "Вызовы", "внутренние Go вызовы")
Rel(user, db, "Чтение/запись пользователей", "SQL")
Rel(device, db, "Чтение/запись устройств", "SQL")
Rel(temp, db, "Чтение/запись температуры", "SQL")
Rel(heating, db, "Чтение настроек", "SQL")
Rel(heating, ext, "Посылка команд на реле", "HTTP/MQTT/и др.")
Rel(temp, ext, "Опрос датчиков", "HTTP/MQTT/и др.")
Rel(ext, sensor, "Получение температуры", "HTTP/MQTT")
Rel(ext, relay, "Передача команд", "HTTP/MQTT")
Rel_Back(sensor, ext, "Ответ с температурой", "HTTP/MQTT")
Rel_Back(relay, ext, "Подтверждение команды", "HTTP/MQTT")

' API — самая верхняя точка для пользователей
Person(user_person, "Пользователь", "Работает с Web API")
Rel(user_person, api, "Запросы управления отоплением, просмотра температуры", "HTTP")

Person(admin_person, "Администратор", "Настраивает систему через Web API")
Rel(admin_person, api, "Администрирование", "HTTP")

Person(technician, "Техник", "Добавляет устройства через Web API")
Rel(technician, api, "Добавление устройств", "HTTP")

@enduml
```

## Контейнерная диаграмма: Текущий монолит

```puml
@startuml Current_Container_Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

Person(user, "Пользователь", "Взаимодействует через Web UI")
Person(admin, "Администратор", "Взаимодействует через Web UI")
Person(technician, "Техник", "Взаимодействует через Web UI")
Container(webui, "Web Browser", "HTML/JS", "Веб-интерфейс для пользователей, админа, техника")

Container(monolith, "Монолитное приложение", "Go", "Обрабатывает все бизнес-процессы: управление отоплением, мониторинг температуры")
ContainerDb(db, "PostgreSQL", "Реляционная база данных", "Хранение данных пользователей, устройств, температур")

Container_Ext(sensor, "Датчик температуры", "Внешнее устройство", "Передает/отдает данные о температуре")
Container_Ext(relay, "Реле отопления", "Внешнее устройство", "Исполняет команды на включение/выключение отопления")

Rel(user, webui, "HTTPs", "Работает с интерфейсом")
Rel(admin, webui, "HTTPs", "Администрирует через интерфейс")
Rel(technician, webui, "HTTPs", "Регистрирует устройства через интерфейс")
Rel(webui, monolith, "REST/HTTP", "API-запросы")

Rel(monolith, db, "SQL", "Чтение/запись данных")
Rel(monolith, sensor, "HTTP", "Опрос данных температуры")
Rel(sensor, monolith, "HTTP", "Ответ/передача температуры")

Rel(monolith, relay, "HTTP", "Посылка команд на реле")
Rel(relay, monolith, "HTTP", "Ответ о статусе/ошибках")

@enduml
```

## Диаграмма развертывания: Текущий монолит

```puml
@startuml
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Deployment.puml

Node(web_client, "Устройство пользователя", "ПК, планшет, смартфон")
Node(house, "Дом пользователя", "Физическое здание") {
Node(sensor, "Датчик температуры", "IoT устройство / сенсор")
Node(relay, "Реле отопления", "IoT устройство / исполнительное устройство")
}

Node(datacenter, "Основной сервер компании", "Дата-центр, облако или on-premise сервер") {
Node(app_server, "Сервер приложений", "Linux, Go") {
Container(monolith, "Монолитное приложение", "Go + встроенный web сервер")
}
Node(db_server, "Сервер БД", "Linux, PostgreSQL") {
ContainerDb(db, "PostgreSQL", "База данных")
}
}

Rel(web_client, monolith, "HTTP/HTTPS", "Веб-доступ через браузер")
Rel(monolith, db, "SQL", "читает/пишет данные")
Rel(monolith, sensor, "HTTP", "Опрос данных температуры (LAN/VPN/интернет)")
Rel(sensor, monolith, "HTTP", "Ответ с температурой")
Rel(monolith, relay, "HTTP", "Посылка команд на реле (LAN/VPN/интернет)")
Rel(relay, monolith, "HTTP", "Ответ (статус)")

@enduml
```