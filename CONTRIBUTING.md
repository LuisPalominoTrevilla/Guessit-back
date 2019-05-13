# Contributing to GuessIt!

🐙🐳🤯 Thanks for taking the time to contribute to this wonderful project! 🤯🐳🐙

So, you want to help GuessIt grow? 🙆🏻‍♂️

We have created this website to make your contribution process as smooth as possible.

## Before you start

Dont forget to do the following:
- Fork this repository.
- Email the developer so you can have the propper .env file required to run this project.
- Carefully read and understand the README of the project, as we state the naming conventions, endpoint documentation, app installation and any other prerequisite we might miss here.

## How to help?

Anyone can contribute to this project and it's very simple to start doing it right away!
**Don't forget to fork this repository before contributing!**

- Give a star on the repos. Don't forget [Guessit-front](https://github.com/LuisPalominoTrevilla/Guessit-front)
- If you see a bug or would like a feature implemented you can always add a new issue. If you are not sure how to do it you can [check the instructions](https://help.github.com/en/articles/creating-an-issue)
- Grab one of the issues (there might be some bugs or features) and work on it. Once you finish, create a pull request pointing to development. Make sure your branches follow the branch conventions stated below!
- If you see something and would like to implement it right away, you can! Just be sure to create your pull request.

## Branch conventions

- Every branch should start with *feat/* or *bug/* followed by a very brief description of your work (no more than 12 characters). Some examples of what is expected:
    - *feat/loginModal*
    - *feat/showImage*
    - *bug/loadingTimes*
- Every PR should be pointing to development.
- If your contribution does not meet the above criteria, we will be forced to decline your PR.

## Folder structure

The folder structure used for GuessIt-backend is pretty straightforward. 📏

At the root of the repository we have the basic files. `main.go` is the entry-point of the application.

There is also a Dockerfile 🐳 and a docker-compose. Docker-compose is the one responsible for declaring all necessary containers. Namely, the mongo database, redis and the golang application.

```
project
|   .env
|   .gitignore
|   docker-compose.yml
|   Dockerfile
|   go.mod
|   go.sum
|   CONTRIBUTING.md
|   README.md
|   main.go
|
|___static      <= Uploaded pictures from users will be saved here
|       5c6ba3df16625ad0ba67f2e0   <= Each user has its own folder
|           ...
|       ...
|
|___routers     <= All routers are stored here
|       apiRouter.go
|       routers.go
|       staticRouter.go
|
|___redis <= Module used for interaction with redis
|       operations.go
|       redi.go
|
|___modules <= This is where common modules go
|       authVerification.go
|       cookieRetriever.go
|       utils.go
|       ...
|
|___models <= All the models used within the code are stored here
|       image.go
|       parameters.go
|       rate.go
|       responses.go
|       user.go
|
|___middleware <= Middleware should go here
|       cors.go
|       ...
|
|___errors <= Not currently being used
|       error.go
|       ...
|
|___docs <= This is where swagger is auto-generated. DO NOT MODIFY ANYTHING INSIDE
|       ...
|
|___db <= Inside this folder are the files that instantiate the database and manipulate it directly
|       imageDb.go
|       mon.go
|       rateDb.go
|       userDb.go
|
|___controllers <= All logic related to the business-layer goes here
|       controller.go
|       imageController.go
|       userController.go
|
|___boot <= Code that runs whenever the app is initialized goes here
|       seeder
|           modelSeed.go
|           userSeeder.go
|
|___authentication <= Files used for athentication go here
|       auth.go
|       middleware.go
|
*
```

## How the rating system works

One of the core functionality of GuessIt is its rate system. We have a rather basic rate system. However, you are encouraged to improved if you fancy.

**Rating workflow**:

1. The user rates an image
2. The logic changes depending if the user has an account

| User has account    | Second Header |
| ------------------- | ------------- |
| Content Cell        | Content Cell  |
| Content Cell        | Content Cell  |