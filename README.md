# TBOX Backend Homework

## Requirements
An ecommerce application needs to provide authentication feature for users.
The Product Manager wants the app supports login via phone number
Here is the flow of login

As a user:
1. I will be asked to enter their phone number. 
2. Then the app will send to user an OTP token. 
3. Then I will have 60 seconds to enter my OTP token
4. If the token is correct, the app will return a token to me.
5. If the token is incorrect, the app will show an error message to me, then I could enter the OTP again.
6. If I did not receive the OTP token, I could ask the system to send the token again to me after 30 seconds     
7. After login successfully, I will never need to login again
8. The system need to have a mechanism to prevent attacker to call our send SMS API

Your task is implement APIs for this feature by any programming languages. But if you could write it by Golang, you will have bonus point :)

You will have 7 days to finish implement this tasks. 
But we think that it will take you about 8 hours only.

After finish the task, please create a public Github Repository, and share with us.

The good source code should provide

+ Document for each API
+ Unit test to test all API
+ Document explain about how to run the code, and how to run the unit test

## Project Architect
Using Domain-driven design 
[Domain-driven design ](https://github.com/vektra/mockery) with 5 layers:
+ **Transport**: place appy input and out return can use for GRPC or Restful (gin)
+ **Handler**: Place call 1 or many repository for solve require from transport layer.
+ **Transport**: All business is centralize in this layer. All unit test should apply here.
+ **Storage**: only solve problem CRUD. Don't have any logic here. Help solve problem replace another DB with mininum change source code
+ **Model**: Place to define structure, validation, input and out put

I also design package common and middleware for reuse source code in project

## Setup & Run
```
go run main.
```

## UnitTEST and Mock

Use **Mockery** [package](https://github.com/vektra/mockery).<br/>
Generate single mock file (example) :
```bash
mockery -name=StoreStorage -case=underscore -dir=./storage -output=./mocks/storage
```
- `name`: The name of `interface` to generate.
- `case`: The name of mock file will be generate.
- `dir`: The input `inferface` file path.
- `output`: The output mock file will be generated.

Generate all interfaces with recursive:
```bash
mockery -all -recursive -case=underscore
```
or use `make` command
```
make mock
```

## Run test coverage
Can use unit test feature of golang cmd for run all test file
```
    go test -run ''
```

or use `make` command
```
make coverage
```

