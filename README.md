# API-News

Simple API Project & Invoice

# Config DB Mysql
```console
Setting config db with 'db_emi' as name of db
```

# Build
```console
go build main.go
```

# Build From Docker
```console
docker build . -t api-project-and-invoice:latest
docker run --name api-project-and-invoice -p 8585:8585 -d api-project-and-invoice:latest
```

# API
```console
Import Insomnia.json
```

## Core library

Library | Usage
-- | --
gin | Base framework
gorm | ORM library
mysql | Database
logrus | Logger library
viper | Config library

And others library are listed on `go.mod` file
