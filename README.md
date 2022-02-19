# radiant_cloud_assessment

## i) Installation

### 1. Install golang on your machine
   
   go to this [page](https://go.dev/dl/) and follow the instructions to install golang on your machine

### 2. Install MongoDB on your machine
   
   go to this [page](https://docs.mongodb.com/manual/installation/) and follow the instructions to install MongoDB on your machine
  
## ii) Database setup

### For linux
   
   step 1: Open terminal. Using your favorite editor, Open the `.bashrc` file located in the /home/<username> directory. I am using vim
   
   ```console
   $ vim ~/.bashrc
   ```
   
   step 2: Declare the environment variable for mongodb
   
   ```bash
   export RADIANT_DB_USERNAME=<your_username>
   export RADIANT_DB_PASSWORD=<your_password>   
   export RADIANT_DB='radiant_cloud_db'
   ```
   
   step 3: save the changes and restart the terminal
   
   step 4: login to your mongodb cli using your username and password
   ```console
   $ sudo mongo -u <username> -p <password>
   ```
   once you are logged in, the terminal prompt looks like
   ```console
   >
   ```
   
   step 5: Create a database called radiant_cloud_db
   ```console
   > use radiant_cloud_db;
   ```
   
 ## iii) Server setup
   step 1: Declare the environment variable for go server
   ```bash
   export PORT=8080
   ```
   
   step 2: Clone the project
   ```console
   $ git clone git@github.com:Harsha-S2604/radiant_cloud_assesment.git
   ```
   
   step 3: change directory to the cloned project
   ```console
   $ cd radiant_cloud_assesment
   ```
   
   step 4(optional): download the dependencies
   this step is optional since `go run/go build` automatically runs the `go mod download`
   ```console
   $ go mod download
   ```
   
   step 5: run the server
   ```console
   $ go run main.go
   ```
   
   now the server runs on localhost:8080 or the port you have configured
   
 ## iv) Working
   There are totally 8 APIs
   user => adduser, updateuser, deleteuser, getuser
   group => addgroup, getgroupusers, updategroup, deletegroup
   
   step 1: open the terminal
   ### Add User
   ```console
   $ curl --header "Content-Type: application/json" \
		--request POST \
		--data '{"userid":"tuser","first_name":"test","last_name":"user","email":"test@123.com"}' \
		'localhost:8080/api/v1/users'
   ```
   
   ### Add Group
   ```console
   $ curl --header "Content-Type: application/json" \
		--request POST \
		--data '{"group_name":"admin", "users":["tuser", "jreese"]}' \
		'localhost:8080/api/v1/groups
   ```
   
   ### Get User
   ```console
   $ curl --request GET 'localhost:8080/api/v1/users/tuser'
   ```
   
   ### Get group users
   ```console
   $ curl --request GET 'localhost:8080/api/v1/groups/admin'
   ```
   ### Update user
   ```console
   $ curl --request PUT 'localhost:8080/api/v1/users/tuser' 
   --data '{"email": "test456@gmail.com"}' \
   ```
   
   ### Update group
   ```console
   $ curl --request PUT 'localhost:8080/api/v1/groups/admin' \
   --data '{"users":["treese","kjackson"]}' \
   ```
   
   ### Delete User
   ```console
   $ curl --request DELETE 'localhost:8080/api/v1/users/tuser'
   ```
   
   ### Delete Group
   ```console
   $ curl --request DELETE 'localhost:8080/api/v1/groups/admin' \
   ```
   
