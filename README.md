## Start the application
```git clone https://github.com/Folium1/testTask.git```<br>
1.Copy the key of your mysql db source to __MYSQL_SOURCE__ local variable <br>
2.Start mysql server. <br>
3.Run ```go run main.go```  <br>
Requests:
- (POST) /login: A mock user is created with the following JSON payload: {"username": "test", "password": 123456}.
- (POST) /upload-picture: Upload a picture in the form with the key "image". Include the token received from the login request in the header.
- (GET) /images: Include the token received from the login request in the header.
