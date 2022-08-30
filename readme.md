# Objective 

We will create a simple product api and webpage using the golang for backend language. 

We will create the CRUD for product and users. 

User can post product, 
User can update product detail,
User can delete product, 

---

## HOW ?

For the first step we can try to replicate the guide from this link :
<a href="https://tutorialedge.net/golang/creating-restful-api-with-golang/">Resftul Api With Golang</a>
to import the package 

---

## Steps

first you'll need to create the go mod init then after that you can imprt the package

1. go mod init

```bash
go mod init <nameofyourproject/youraccountname> 
```
2. After that we can import with the go mod tidy to import the gorrila/mux package

```
go mod tidy
```
list of endpoint

```
/api/article
/api/article/{id}
/api/create-article
/api/delete-article/{id}
/api/update-article/{id}
```
---

## Implement the 

1. We can create the database first.


```sql
CREATE TABLE Users ( 
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL, 
  fullname VARCHAR(50) NOT NULL, 
  username VARCHAR(10) NOT NULL,
  password VARCHAR(10) NOT NULL,
  created_at TIMESTAMP, updated_at TIMESTAMP
);

CREATE TABLE Posts (
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL, 
  user_id INT, 
  post_name VARCHAR(50) NOT NULL, 
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY user_id REFERENCES Users(id)
);
```

2. Create the function

type User struct 

// User must have these method
- createUser
- updateUser
- deleteUser


// Function for Posts
- createPost(user User)
- updatePost(user User)
- deletePost(user User)
___
