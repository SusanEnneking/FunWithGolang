# PlayWithGo (and Next.js) #

An algorithm a day keeps the hubris away.  But, if you keep it to < 1/week you really aren't a geek.

## A place to practice with GoLang and Next.js ##

Building skills with go and Next.js by doing a few on-line code/algorithm challenges Next.js and/or Golang.  Maybe revisiting a few places where my pride has been hurt during a coding interview.

## Main components: ##

1. Client - Next.js page that allows playing with javascript algorithms and showing the result 

2. Server - Golang http server.  Just a way to invoke the code via Postman/Next.js.  Probably sould just add more tests, but then how would I dream up a need for a client.

## To Use ##
### Start Server:
1.  Navigate to server/main.go
2.  ```go run main.go```\
    --OR--\
    Use vsCode and choose **Run** , **Start Debugging** from the menu.  You have to be editing **main.go** for this to work.

### Start Client:
1.  Navigate to client/nextjs-trystuff
2.  ```npm install```
2.  ```npm run dev```
3.  navigate to http://localhost:3000

## Testing ##
### Server ###
Navigate to server and run the following commands\
```make test``` runs all server tests\
```make test_coverage``` shows test coverage
### Client ###
[Used this setup advice](https://nextjs.org/docs/app/building-your-application/testing/jest)
```npm run test``` runs all client tests\
