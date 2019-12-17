# Review 360

# Challenge Requirement
  - Admin Views
    - Add/remove/update/view employees
      - [X] add employee
      - [X] remove employee
      - [X] update employee
      - [X] view employees
    - Add/update/view performance reviews
      - [X] add review
      - [X] update review
      - [X] view review
    - [X] Assign employees to participate in another employee's performance review
  - Employee Views
    - [X] List of performance reviews requiring feedback
    - [X] Submit feedback

# Development Requirement
  - dotenv
  - golang
  - mysql
  - goose (migration tools)
  - npm
  - yarn
  - react-native-cli

# Development

we suggest using dotenv for environment variables management.
you can choose `dotenv` or `godotenv`, then create .env.local first.

`.env.local` recommended content is below.
```
MYSQL_HOST=127.0.0.1
MYSQL_PASSWORD=
MYSQL_URL=root:$MYSQL_PASSWORD@tcp($MYSQL_HOST:3306)/review360
SECRET=
PORT=8080
```

## Backend Setup

Use golang as the backend language. so you will need to install golang in your environment. use go module for dependencies and setup mysql first.
Please create an account and database, `db.sql` will be helpful. and update your `.env.local` by those values.

making test first!

`$ dotenv -f .env.local -- make test`

if anything goes well, you can try the command below, it will start a api server on 8080.

`$ dotenv -f .env.local -- go run main.go`

## Frontend Setup

Use react as the frontend framework, we could develop three platform at the same time by react native and react native web. But android is lack of testing. so the first recommendation is trying ios or web.

install npm packages

`$ yarn`

build ios environment, start simulator, and run metro bundler.

`$ yarn ios`

build web by webpack

`$ yarn web`

## Docker environment

if you don't have golang and nodejs environment, you can simply build a runnable service by docker.

build service, init container, and setup database at the same time.

`$ make docker`

only build service, it will use multi-stage for building backend and frontend.

`$ make docker-review360`

build init container, it will build a image for database migrations at the first.

`$ make docker-init`

build database

`$ make docker-mysql`

Run services by `docker-compose`

`$ make run`

Shutdown all services

`$ make down`

Build frontend code by local yarn. static files is served at `build` directory, and then you can test frontend production app in the local.

`$ make build-app`

## Testing

Server side use simply go test, but frontend is more complicated. we separate test files to native and web for different environment testing.

Web platform is using js-dom, and native part is using node.
Ideally, we could hold one test file for both environment, but actually
we will encounter many problems from testing lib, test runner, 3rd lib and our own code. so just separating them at the beginning will save lots of time.

more details can see `app/config`

# Full Stack Developer Challenge
This is an interview challengs. Please feel free to fork. Pull Requests will be ignored.

## Requirements
Design a web application that allows employees to submit feedback toward each other's performance review.

*Partial solutions are acceptable.*  It is not necessary to submit a complete solution that implements every requirement.

### Admin view
* Add/remove/update/view employees
* Add/update/view performance reviews
* Assign employees to participate in another employee's performance review

### Employee view
* List of performance reviews requiring feedback
* Submit feedback

## Challenge Scope
* High level description of design and technologies used
* Server side API (using a programming language and/or framework of your choice)
  * Implementation of at least 3 API calls
  * Most full stack web developers at PayPay currently use Java, Ruby on Rails, or Node.js on the server(with MySQL for the database), but feel free to use other tech if you prefer
* Web app
  * Implementation of 2-5 web pages using a modern web framework (e.g. React or Angular) that talks to server side
    * This should integrate with your API, but it's fine to use static responses for some of it
* Document all assumptions made
* Complete solutions aren't required, but what you do submit needs to run.

## How to complete this challenge
* Fork this repo in github
* Complete the design and code as defined to the best of your abilities
* Place notes in your code to help with clarity where appropriate. Make it readable enough to present to the PayPay interview team
* Complete your work in your own github repo and send the results to us and/or present them during your interview

## What are we looking for? What does this prove?
* Assumptions you make given limited requirements
* Technology and design choices
* Identify areas of your strengths
* This is not a pass or fail test, this will serve as a common ground that we can deep dive together into specific issues
