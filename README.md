# Angular Go CRUD
A demo CRUD application using Angular and Go


# How to rebuild locally
Before proceeding, you must have Angular CLI and Go installed in your system.

**Install npm packages**
Start from the root directory
```bash
cd client
npm install
```
**Install go dependencies**
Start from the root directory
```bash
cd server
go mod tidy
```
**Run the server**
Under the /server directory
```bash
go run main/main.go
```
**Run the client**
Under the /client directory
```bash
ng serve --o
```


