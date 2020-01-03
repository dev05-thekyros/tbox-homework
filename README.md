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

## Assumptions
We assume that you could send SMS to any phone number by calling this API

```
curl -XPOST https://5db83e44177b350014ac77c6.mockapi.io/v1/sms \
    -d '{"phone_number": <phone_number>, "content": <content>}'
```

E.g: you could send a SMS to phone 09812345678 a message "Welcome to TBOX" by calling

```
curl -XPOST https://5db83e44177b350014ac77c6.mockapi.io/v1/sms \
    -d '{"phone_number": "09812345678", "content": "Welcome to TBOX"}'
```

## Scope
Items 1, 2, 3, 4, 8 are required for this task.

Items 5, 6, 7 are optional. 