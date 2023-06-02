## Start the application
1.Copy the key for your mysql db source to __MYSQL_SOURCE__ variable <br>
2.Start mysql server. <br>
3.Run ```go run main.go```  <br>
Requests:
- (POST)/login, mock user is created{"username": "test", "password": 123456}
- (POST)/upload-picture (picture in form with key "image",in header must be the token, you received from login request)
- (GET)/images (in header must be the token, you received from login request)