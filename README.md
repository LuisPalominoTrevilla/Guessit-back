# GuessIt API

## Getting Started
Follow these instructions to set your environment succesfully.
Make sure you have go version 1.11 you can check it with `go version`
You will need a .env file for the project to work. If you don't have one, ask the owner of the repo to give you one

## Clone the project
First, clone the project to your desired location with `git@github.com:LuisPalominoTrevilla/Guessit-back.git`

## Setup docker containers
For the time being, GuessIt! runs via 2 docker containers.
Make sure you have docker-compose installed with `docker-compose -v`. If you don't have it, please [install docker-compose](https://docs.docker.com/compose/install/)
Build the images for the project, this should not take that long:
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
af6dea811c90        guessit-back_golang   "/app"                   8 hours ago         Up 4 seconds        0.0.0.0:5000->5000/tcp     guessit-backend
6f137255e098        mongo                 "docker-entrypoint.s…"   9 hours ago         Up 4 seconds        0.0.0.0:27017->27017/tcp   guessit-mongo
```
Docker is currently running 2 containers.
- `guessit-backend`: API container.
- `guessit-mongo`: Container for the main database.
You can stop your containers using `docker-compose stop` or, if you are not running your containers in the background, use `CTRL-C`

Whenever you make a change to a file, you need to rebuild the containers, it won't take long. We recommend you to use the command `docker-compose up --build` to run your containers and rebuild them.

## Accessing containers
If for any reason you need to access a container, use the following command:
`$ docker exec -it *container-name* bash`

## DB Management
For db management it is recommended that you use [Robo3T](https://robomongo.org/)
- Open Robo3T and create a new connection.
- Name the connection as you wish.
- Address: `localhost` Port `27017`
- Test the connection and save it

## Coding conventions
...

## Development work
...

## Other Guessit repositories
- [front end](https://github.com/LuisPalominoTrevilla/Guessit-front)

*Happy coding!*