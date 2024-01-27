# Patient API

Patient API I used Gin for routing, ent as ORM and migrations
I have used repository pattern for this API
I have integrated Air for auto-reload as well

## Start up

Add local.env for postgres database and start the application
```
make apply (in makefile it will apply migration into your database) 
air (run golang with autoreload, just in case you want to test something)
```

## API Docs
[Swagger](http://localhost:8080)
