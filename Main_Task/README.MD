# Appointy Task (Instagram Clone)

This respository contains the server side code required to start a server, the server has api end points for creating and finding users, and their posts

## Running the Code

simply clone the repo and navigate to the "Main_Task" directory and use the "go run ." command
(alternatively you can copy the code below and paste it in your terminal)
```bash
git clone https://github.com/Gaurav4604/GolangCrash.git && cd GolangCrash/Main_Task && go run .
```

## Usage

Better to open up the URL end points in a API platform such as postman

There are 5 major end points

Link for the example demonstration in postman
(https://github.com/Gaurav4604/GolangCrash/blob/master/Main_Task/postman_examples.pdf)

## Note, 


```python
127.0.0.1:8080/api/v1/posts/user/ # this is a POST request, with user data in the request body

"""
if anything except a post request is passed to this route, it will be considered a bad request (response status code 400)
"""

example_request_body = 
   {
	"name" : "Gaurav Singh",
    "age" : 20
	"username" : "Gaurav4604",
	"email" : "abc@gmail.com",
	"password" : "abcde12345"
   }

example_response_body = 
    {
    "_id": "ObjectID(\"6161971b7fd64937d6c2f164\")",
	"Name": "Gaurav Singh",
	"Username": "Gaurav4604",
	"Age": 20,
	"Email": "abc@gmail.com",
	"Password": ""
    }
  
# the encrypted password is stored in mongodb server, and is encrypted using the MD5 algorithm
# the code for the same is in the file named "encryption.go"

//////////////////////////////////////////////////////////////////////////////////////////////////////////

127.0.0.1:8080/api/v1/posts/users/<user_mongo_object_id> # this is a GET request

# since this is a simple GET request, nothing is to be passed in the request body
""" if the the userID specified in the URL is valid, the response will return a user object with user details"""
example_response_body = 
    {
	"Name": "Gaurav Singh",
	"Username": "Gaurav4604",
	"Age": 20,
	"Email": "abc@gmail.com",
	"Password": ""
    }

""" if an invalid userID is specified in the URL, the response will return an error message and a status code of 404 """

//////////////////////////////////////////////////////////////////////////////////////////////////////////

127.0.0.1:8080/api/v1/posts/posts/ # this is a POST request, with post data in the request body

""" the given url endpoint accepts POST request, the post can 
only be created by valid users, thus the 
request body must contain userID as well, while creating a post request """

example_request_body = {
	"userid": "6161971b7fd64937d6c2f164",
	"caption": "what a beautiful day",
	"body": "this is my description of the beautiful day",
	"imageurl": "abc/123/ok.png"
}

""" after the request reaches the server, user is validated, and if user is valid, the post body object is assigned a
UNIX format timestamp, then the data is stored in the db  """

""" if a valid user id is passed, the post is saved and the response contains the data 
same as what is saved in the db"""

example_response_body = {
		"_id": "6161af4862ad9d3c3ac73391",
		"userID": "6161971b7fd64937d6c2f164",
		"Caption": "what a beautiful day",
		"Body": "this is my description of the beautiful day",
		"ImageURL": "abc/123/ok.png",
		"Timestamp": "1633791816"
}

""" if invalid userID is entered, the response will return a 404 error with a error message """

//////////////////////////////////////////////////////////////////////////////////////////////////////////

127.0.0.1:8080/api/v1/posts/posts/<post_mongo_object_id> # this is a GET request

""" if a valid postID is passed in the URL, the post response contains the data 
for that particular post """

example_response_body = {
		"_id": "6161af4862ad9d3c3ac73391",
		"userID": "6161971b7fd64937d6c2f164",
		"Caption": "what a beautiful day",
		"Body": "this is my description of the beautiful day",
		"ImageURL": "abc/123/ok.png",
		"Timestamp": "1633791816"
}

""" if invalid postID is entered, the response will return a 404 error with a error message """

//////////////////////////////////////////////////////////////////////////////////////////////////////////

127.0.0.1:8080/api/v1/posts/posts/users/<user_mongo_object_id> # this is a GET request

""" if a valid userID is passed in the URL, the post response contains the data 
for all the posts of that particular user """

# additionally you can pass a param in the query called "pages" with a value greater than 1
# this will enable pagination on the api response, showing you only those many posts as specified by the parameter

#example
127.0.0.1:8080/api/v1/posts/user/6161971b7fd64937d6c2f164?pages=2

""" here the url will show you only 2 pages at a time, every time 
you hit the endpoint, the next two pages are sent to you in the response """
""" cookies need to be enabled for pagination to work, since your current page state is stored in cookies """


example_response_body = [
	{
		"_id": "6161b2fdec815141e29298de",
		"userID": "6161971b7fd64937d6c2f164",
		"Caption": "what a beautiful day",
		"Body": "this is my description of the beautiful day",
		"ImageURL": "abc/123/ok.png",
		"Timestamp": "1633792765"
	},
	{
		"_id": "6161b49aec815141e29298df",
		"userID": "6161971b7fd64937d6c2f164",
		"Caption": "what a beautiful day",
		"Body": "this is my description of the beautiful day",
		"ImageURL": "abc/123/ok.png",
		"Timestamp": "1633793178"
	}
]


""" if all posts in the db are already displayed, the state counter resets and response sends a message to tell the same """

example_response_body = {
	"Message": "all pages done, reseting counter"
}

```

//////////////////////////////////////////////////////////////////////////////////////////////////////////

## Note
All the endpoints are declared as functions handled via routes
If more functions for those specific routes are to be added,
that can be done in the respective files, if any endpoint, 
except the ones mentioned above, ends up being used,
error Handler (declared in "utils.go" file) will invalidate
it and send appropriate response

for User Routes -> "userRoutes.go"
for Post Routes -> "postRoutes.go"

A simple password validation function has also been implemented
on the user struct, thereby easing the process of validation during
login ("visible" in the file named Schema)


## License
[MIT](https://choosealicense.com/licenses/mit/)