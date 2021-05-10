# How to run project


1. Create .env file in project with following variables in it
       
       export DATABASE_DIALECT=xxxxxxx
       export PORT=xxxxxx
       export DATABASE_URL=xxxxxxx
       export GRPCPORT=xxxxxx

note := Make sure the database URL is correct and that database is exist in your system

2. Source .env file (change file env vars as per you configuration)

3. cd cmd/server and run server

    Currently server just pront the message and resture response with message string, havent im plimented email service yet.
4. cd cmd/clinet and run clinet
5. open postman and user following APIS

  i. to create user  POST
    ---- http://localhost:9003/v1/user/create/

    body --
         
         {
              "username": "ashish",
              "name": "ashish",
              "email": "ashis.nik16h@gmail.com"
         }

  ii.
      to get existing users GET
    --- http://localhost:9003/v1/user/


# You can also run the project using makefile

to run server 
  `make run-server`

to run client 
  `make run-client`

If you want to do changes in proto file, you can do that and to regernerate pb.go file use following command
  `make regenerate`
