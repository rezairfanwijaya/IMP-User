# Simple User Auth

This repository provides services about user authentication and authorization

## Run Locally


#### Clone the project

```bash
  git clone https://github.com/rezairfanwijaya/IMP-User.git
```

#### Go to the project directory

```bash
  cd IMP-User
```

#### Get Dependency
```bash
  go mod tidy
```

#### Open the project with your favorite text editor

#### Setup ENV
##### Edit the .env.example file to .env and adjust it to the env that you will use
##### example :
```bash
DATABASE_USERNAME = "root"
DATABASE_PASSWORD = "12345"
DATABASE_HOST = "127.0.0.1"
DATABASE_PORT = "3306"
DATABASE_NAME = "pos"
DOMAIN = ":8080"
SECRET_KEY = "1213v-dhgfvh2342fved"
```

#### Run application
```bash
  go run main.go
```

## API Documentation
[API DOCS](https://documenter.getpostman.com/view/11940636/2s93K1pzzk)

