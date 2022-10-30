#Рассылка сообщений

## Архитектура программы
- cmd
  - main.go
- internal
	- handlers
		- handler.go
	- server
		- server.go
- pkg
	- emails
		- emails.go
- template
	- index.html
	- email.html
	- css
	- favicon

---
*Для запуска программы используйте*
```
go run cmd/main/main.go
```
На локальном сервере имеются две кнопки "Sending" (для моментальной отправки писем) и "delayed send" (для отложенной отправки) с вводом времени, в которое нужно отправить письма.
---
*Рассылка писем реализована с помощью пакета smtp*
---
*Рассылка отложенных писем реализована, с помощью планировщика заданий Cron*
---
Список подписчиков хранится в структуре Receiver
```
type Receiver struct {
	Name     string
	Surname  string
	Mail     []string
	Birthday string
}
```
Каждому подписчику отправляется письмо с сообщением, а также его имя, фамилия и дата рождения
```
<div class="container">
        <h3>Name:</h3><span>{{.Name}}</span><br/><br/>
        <h3>Surname:</h3><span>{{.Surname}}</span><br/><br/>
        <h3>Birthday:</h3><span>{{.Message}}</span><br/>
        <p>Привет это мое тестовое сообщение!</p>
    </div>
```
