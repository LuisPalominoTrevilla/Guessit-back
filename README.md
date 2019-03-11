# GuessIt API

## Getting Started
Follow these instructions to set your environment succesfully.
Make sure you have go version 1.11 you can check it with `go version`
You will need a .env file for the project to work. If you don't have one, ask the owner of the repo to give you one

## Clone the project
First, clone the project to your desired location with `git@github.com:LuisPalominoTrevilla/Guessit-back.git`

## Setup docker containers
For the time being, GuessIt! runs via 3 docker containers.
Make sure you have docker-compose installed with `docker-compose -v`. If you don't have it, please [install docker-compose](https://docs.docker.com/compose/install/)
Build the images for the project, it might take a while:
```
$ docker-compose build
```
Run the containers with:
```
$ docker-compose up
```
The web server should be running on port 5000 and you should have mongoDB running too.
Run `docker ps` and you should see an output like the following:
```
CONTAINER ID        IMAGE                 COMMAND                  CREATED             STATUS              PORTS                      NAMES
e0b82cc45f3c        guessit-back_golang   "/app"                   3 minutes ago       Up 3 minutes        0.0.0.0:5000->5000/tcp     guessit-backend
bcf419b266f9        redis                 "docker-entrypoint.s…"   3 days ago          Up 3 minutes        0.0.0.0:6379->6379/tcp     guessit-redis
6f137255e098        mongo                 "docker-entrypoint.s…"   3 weeks ago         Up 3 minutes        0.0.0.0:27017->27017/tcp   guessit-mongo
```
Docker is currently running 3 containers.
- `guessit-backend`: API container.
- `guessit-redis`: Redis Database for the project.
- `guessit-mongo`: Container for the main database.
You can stop your containers using `docker-compose stop` or, if you are not running your containers in the background, use `CTRL-C`

Whenever you make a change to a file, you need to rebuild the containers, it won't take long. We recommend you to use the command `docker-compose up --build` to run your containers and rebuild them.

## Accessing containers
If for any reason you need to access a container, use the following command:
`$ docker exec -it *container-name* bash`

If you want to access Redis CLI just use
`$ docker exec -it guessit-redis bash`
and then
`$ redis-cli`

## DB Management
For db management it is recommended that you use [Robo3T](https://robomongo.org/)
- Open Robo3T and create a new connection.
- Name the connection as you wish.
- Address: `localhost` Port `27017`
- Test the connection and save it

## Coding conventions
...

## Endpoint documentation
Yes pal, you need to document every endpoint you write so the development team knows how each endpoint works.

We are using swaggo/swag, a tool that allows easy documentation of your code using Swagger UI

#### Usage
To use it follow these instructions:
- You need to download swag `go get -u github.com/swaggo/swag/cmd/swag`
- Whenever you change or add new documentation, run the following command inside the root folder of your app `swag init -g routers/apiRouter.go`
- If there were no errors, just build the container again and run it
- To view the documentation UI just enter *localhost:5000/swagger/index.html*
- If you have questions about the Swagger UI, ask the repo owner

#### Troubleshooting
**What happens if the command swag is not found?**
This means that you don't have GOPATH exported. If you are using Linux, or macOS, follow the next instructions:
1. Run the command `go env` and look for the variables GOPATH and GOROOT, you will need them soon
2. Edit your bash config file (~/.bash_profile for bash and ~/.zshrc for Zsh)
3. Add the following lines somewhere in the file, replacing GOPATH and GOROOT with the previous output
```
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
4. Run the command `source ~/.bash_profile` or `~/.zshrc` depending on your bash


We encourage you to visit the [Swag docs](https://github.com/swaggo/swag#api-operation) so you have an idea about how to document your code.
You may also check current endpoints so you can have an idea about how to document new endpoints.

## Development work
Since you will be working with a team, it's important that you follow coding conventions and get used to the work flow.

There are bassically to branches, the main and the development branch. Every other branch has to derive from development.

Whenever you are going to start working on a ticket for the backend, there should be specified a branch name in the header of the ticket. It looks something like this: [GI-01]

You need to create the branch GI-01 from development with the command `git checkout -b GI-01`

There will be times when a set of tickets are to be developed as a story. When this happens, you need to create the branch from a specified branch (not development! the description of the ticket will specify the branch).

Once you are done with your work, push the changes to the remote branch with `git push origin *branchName*` and then create the pull request for that branch.
The pull request must contain the following elements:
- The name should be the name of the branch followed by the description
- You should say what you changes
- And how to test your branch
After that, you can create the pull request and wait for your teammates to review it.

Some useful git commands:
- `git commit -m "message"`: Makes a commit with a message specified
- `git pull origin branchNamr`: Pulls changes from the remote branch
- `git add .`: Adds modified files to the staging area
- `git reset --hard`: <span style="color:red">DANGEROUS!</span> The staging area and working directory are reset to the last commit (You'll lose all changes)
- `git reset --soft HEAD~1`: Remove the last commit but keep the changed files

## Other Guessit repositories
- [front end](https://github.com/LuisPalominoTrevilla/Guessit-front)

**Happy coding!**
