It's a todo app in go that we have here.
It's quite simple and straightforward.
Clone the project and go inside server folder.
Open the folder in Visual Studio and go over to the terminal.
Add the command go run main.go
Open Postman and punch in the url http://localhost:4000/api/todos/healthcheck. This url checks whether the app is up or not.
Next add the url http://localhost:4000/api/todos/1 and post the json - { "title":"whatever you want", "body":"whatever you wish"}
And that's all!
