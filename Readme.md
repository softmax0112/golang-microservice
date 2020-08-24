# Golang-Microservice
A simple microservice with MySQL database connection. It contains  parts.

## Reqirements for running 
    1. Go 1.14
    2. MySQL 8.0+ for Database
    3. Make (optional but make larger commands in short)
    4. Docker 

## 1. Installing the Required Application
    I hope GO and MySQL is already installed.
### 1. Intalling Make in Windows the Easy Way
Install Chocolatey (Chocolatey is a software management solution)
1. Open PowerShell with Adminitrator
2. Run 
```sh
Get-ExecutionPolicy
```
If it return `Restricted` then run
```sh
Set-ExecutionPolicy AllSigned
```
OR
```sh
Set-ExecutionPolicy Bypass -Scope Process
```
3. Install by this command
```sh
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```
3. Install make 
```sh
choco install make
```
Unix like OS already comes with make preinstalled so no worries.
## 2. Installing Docker
It's a a bit long process but in short you need to install Docker from https://hub.docker.com/editions/community/docker-ce-desktop-windows/ and it also need specefic version of Windows and WSL2, so check online for for further help.

## 3. Running Web Application
After Cloning it, download all depencies once
```sh
go mod download
```
Run the command and make sure to set config.yaml file for configuration
```sh 
make run
```
OR 
```sh
go run test
```
## 4. Run the MySQL Server 
Create a DB with any name and make sure to give the set DB_NAME in Config file and create a table *users* with following fields.
1. id   AUTOINCREMENT PRIMARY KEY
2. username UNIQUE VARCHAR (20) 
3. password VARCHAR (20)
4. marks VARCHAR (20) NULL(by Defualt)

## 5. Running the Web Application in DOCKER
1. Buiding the image
Syntax :- docker build -t <tag name> <directory of Dockerfile>
Example :- 
```sh
docker build -t Test .
```
2. Running the image
Syntax :- docker run -p <internalPort>:<externalPort> --name <name of container> <image name>
Example :- 
```sh
docker run -d -p 8080:8080 --name TesingMyApp test
```
Extra Flags can also be added check them out by
```sh
docker run --help
```
### Flag that can be used when running the container
-d   for running conatainer in detach mode (running in background)
--rm for removing the conatiner after stopping it
3. For strating the container
```sh
docker start <container ID> or <container name>
```   
4. For stopping the container
```sh
docker stop <container ID> or <container name>
```
5. For listing the running and stopped container
```sh 
docker ps -a
```
6. For listing Docker images
```sh 
docker images
```
7. For inspecting the Conatainer
```sh
docker inspect <name of container>
```
### Note
-d  for running conatainer in detach mode
--rm for removing the conatiner after stopping it

## 6. Running MySQL in Another Container
1. Pull the MySQL image from docker hub
```sh
docker pull mysql/mysql-server:latest
```
2. To make sure the data persists even after stopping the container Create Volume in docker Syntax:
```sh
docker create -v <location> --name <nameOfTheImahe> <ImageName>
```
Example :-
```sh
docker create -v /var/lib/mysql --name mysqldata mysql/mysql-server:latest
```
3. Running the MySQL Container
Before running it, it required some ENV variable and config to set, atleast root password is required. You can check from here https://hub.docker.com/_/mysql.
Syntax :-
```sh
docker run -p <internalPort>:<externalPort> -e <ENV variable name>=<value> --volumes-from <volume location> --name=<name of Container> <image name> 
```
Example :-
```sh
docker run --rm -p 3306:3306 -d -e MYSQL_ROOT_PASSWORD=pwd --volumes-from mysqldata --name=mysqldb mysql/mysql-server:latest     
```
## 7. Getting inside the running the Container
1. A database could have been created when running the WebApllication but  here will create by another way.
Syntax :-
```sh 
docker exec -it <container ID> <bash or any shell or command>
```  
Example :-
```sh
docker exec -it mysqldb bash
```
2. Login into mysql as Root
Syntax :- 
```sh
mysql -u <username> -p
```
Example :- 
```sh
mysql -u root -p
```
3. You will get a error related to Access denied to solve that issue Create a new user with host with IP or % for any IP access to DB
Syntax :-
```sh 
CREATE USER '<username>'@'<IP or %>' IDENTIFIED BY '<password>';
```
Giving privilages (For Learing purpose give all)
Example :-
```sh
CREATE USER 'rajiv'@'%' IDENTIFIED BY 'pwd';
GRANT ALL PRIVILEGES ON *.* TO 'rajiv'@'%';
FLUSH PRIVILEGES;
``` 
4. Create the DB as stated at #4
        
## Note
1. A postman collection has been shared, import in the Postman to see the API endpoints.
2. For any help reguarding the docker type
```sh 
docker <commandName> --help
```
