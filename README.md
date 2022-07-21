## registration

###  HANDLER: api.SignUp
### ADRESS : 85.143.174.57:80/auth/sing-up

### BODY :
{

    "name":"stepan",
    "email":"stepansheveljov@rambler.ru",
    "password":"12345",
    "password_confirm":"12345"

}

### RESPONSE :
{
"success": "user created"
}

{
"message": "registration success"
}

{
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQxNjM1MzQsImlhdCI6MTY1NDA3NzEzNCwiZW1haWwiOiIifQ.aRsGcHftu-eh8Xt-XYRbBWM8RBBxAKWnLZ75umxFX3s"
}



###  HANDLER: api.SignIn
### ADRESS : 85.143.174.57:80/auth/sing-in

### BODY :
{

    "email":"stepansheveljov@rambler.ru",
    "password":"12345",

}

### RESPONSE :
{
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQxNjQxOTQsImlhdCI6MTY1NDA3Nzc5NCwiZW1haWwiOiJzdGVwYW5zaGV2ZWxqb3ZAcmFtYmxlci5ydSJ9.g9x3H71SN-BNcev5EgCYp5_SxYJYmDbtbuNB6N3Jj_8"
}


###  HANDLER: api.CreatePost
### ADRESS : 85.143.174.57:80/API/create_post

### HEADER-Authorization :
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQxNjQxOTQsImlhdCI6MTY1NDA3Nzc5NCwiZW1haWwiOiJzdGVwYW5zaGV2ZWxqb3ZAcmFtYmxlci5ydSJ9.g9x3H71SN-BNcev5EgCYp5_SxYJYmDbtbuNB6N3Jj_8
### BODY :
{

    "title":"test post",
    "description":"test opis",
    "image":"image.png"

}

### RESPONSE :
{
"success": "Post Created"
}


###  HANDLER: api.UpdatePost
### ADRESS : 85.143.174.57:80/API/update_post

### HEADER-Authorization :
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQxNjQxOTQsImlhdCI6MTY1NDA3Nzc5NCwiZW1haWwiOiJzdGVwYW5zaGV2ZWxqb3ZAcmFtYmxlci5ydSJ9.g9x3H71SN-BNcev5EgCYp5_SxYJYmDbtbuNB6N3Jj_8
### BODY :
{
"id":1,
"title":"ппппппппппппппппппппл",
"description":"описание поста",
"image":"image.png"

}

### RESPONSE :
{
"success": "Post Updated"
}{
"desc": "описание поста",
"id": 1,
"image": "image.png",
"title": "ппппппппппппппппппппл"
}


###  HANDLER: api.DeletePost
### ADRESS : 85.143.174.57:80/API/delete_post

### HEADER-Authorization :
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQxNjQxOTQsImlhdCI6MTY1NDA3Nzc5NCwiZW1haWwiOiJzdGVwYW5zaGV2ZWxqb3ZAcmFtYmxlci5ydSJ9.g9x3H71SN-BNcev5EgCYp5_SxYJYmDbtbuNB6N3Jj_8
### BODY :
{
"id":1,

}

### RESPONSE :
{
"success": "Post Deleted"
}


###  HANDLER: api.GetUserProfile
### ADRESS : 85.143.174.57:80/API/profile

### HEADER-Authorization :
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQxNjQxOTQsImlhdCI6MTY1NDA3Nzc5NCwiZW1haWwiOiJzdGVwYW5zaGV2ZWxqb3ZAcmFtYmxlci5ydSJ9.g9x3H71SN-BNcev5EgCYp5_SxYJYmDbtbuNB6N3Jj_8
### BODY :
{

}

### RESPONSE :
{
"email": "stepansheveljov@mail.ru",
"id": 1,
"name": "stepan",
"posts": []
}



