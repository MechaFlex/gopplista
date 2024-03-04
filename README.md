## Development

First run `npm i` to install dev dependencies. This project uses UnoCSS and Prettier.

Unocss runtime used for development, but be sure to run build when building to ship css directly.

To develop with live reload, use Air. Install Air, then run `air`.

## Building
When building the go application needs to be built I guess
We also want to compile the unocss to not constantly rely on the CDN.



## Tech stack

Server: Fiber
Fiber is a bit controversial since it is not based on the regular Go net/http package, meaning it is not as compatible with middleware.
For me this has not yet been an issue as I have not used any additional middleware.
What I like about fiber is that nested layouts are supported out of the box.

Database: SQLite and sqlc

Go html templates with htmx and unocss.

## Environment variables

```
ENV=dev # if you are in development
SEED_DATABASE=true # if you want to populate the database with some predefined data
ADMIN_PASSWORD=<password>
```

## Learnings

SQLite is good for data storage, but the triggers and other tools are not as powerful as when using Postgres. This means the business logic has to live in the application code which isn't so bad.

TODO: Create way to update an order for section. But needs to be controlled so overlap is avoided. Don't expose native?

Fixed alot but takes time. 

Need to not use delete but use update instead. Current setup seems to be glitchy for some reason...