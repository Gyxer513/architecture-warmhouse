# Уровень кода Device Service — Command Handler

**Device Service** Обрабатывает команды добавления, обновления, удаления устройств, 
валидирует входные данные, транслирует задачи в manager/repository, инициирует публикацию событий.

## Device Service: Command Handler - UML (Class Diagram)

```puml
@startuml

' --- Классы команд ---
class AddDeviceCommand {
    +type: string
    +name: string
    +ownerId: uuid
    +params: Map
}
class UpdateDeviceCommand {
    +deviceId: uuid
    +params: Map
}
class RemoveDeviceCommand {
    +deviceId: uuid
}

' --- Основные бизнес-обработчики ---
class DeviceService {
    +addDevice(cmd: AddDeviceCommand): Device
    +updateDevice(cmd: UpdateDeviceCommand): Device
    +removeDevice(cmd: RemoveDeviceCommand): void
    +getDevice(id: uuid): Device
    +listDevices(ownerId: uuid): List<Device>
}

DeviceService --> AddDeviceCommand
DeviceService --> UpdateDeviceCommand
DeviceService --> RemoveDeviceCommand

class DeviceStateManager {
    +save(device: Device): void
    +update(device: Device): void
    +delete(device: Device): void
    +findById(deviceId: uuid): Device
    +listByOwner(ownerId: uuid): List<Device>
}

DeviceService --> DeviceStateManager

class EventPublisher {
    +publish(event: DeviceEvent): void
}

DeviceService --> EventPublisher

class DeviceRepository {
    +save(device: Device): void
    +update(device: Device): void
    +delete(device: Device): void
    +findById(deviceId: uuid): Device
    +listByOwner(ownerId: uuid): List<Device>
}

DeviceStateManager --> DeviceRepository

' --- Модель устройства ---
class Device {
    +id: uuid
    +name: string
    +type: string
    +ownerId: uuid
    +status: DeviceStatus
    +params: Map
}

enum DeviceStatus {
    INACTIVE
    ACTIVE
    ERROR
    REMOVED
}

' --- События ---
class DeviceEvent {
    +deviceId: uuid
    +eventType: string
    +timestamp: datetime
    +payload: Map
}

@enduml
```
