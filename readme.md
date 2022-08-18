# Objective 

We will create a simple product api and webpage using the golang. 

We will create the CRUD for product and users. 

For the first step we can try to replicate the guide from this link :
<a href="https://tutorialedge.net/golang/creating-restful-api-with-golang/">Resftul Api With Golang</a>

to import the package 

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
