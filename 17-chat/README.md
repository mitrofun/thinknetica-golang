# Задание
-Групповой чат по WebSocket.

## Описание
- Требуется разработать приложение, которое бы позволяло группе пользователей обмениваться текстовыми сообщениями.
- Приложение должно состоять из сервера с WS API и текстового клиента.

### Задача №1
- Разработать приложение-сервер, соответствующее требованиям задачи.
- Сервер имеет две конечных точки: “/send”, позволяющей клиенту отправить сообщение по описанному в задаче протоколу и “/messages”, которая отправляет всем подключенным клиентам все сообщения.


### Задача №2
- Разработать приложение-клиент, которое позволяет в текстовом интерактивном режиме (CLI) отправлять и получать сообщения одновременно.
- При вводе сообщения пользователем клиент вызывает метод сервера, отправляет сообщение и завершает соединение.
- Сообщение транслируется всем подключенным клиентам.