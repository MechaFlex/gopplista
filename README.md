# Gopplista

Gopplista is a replacement for the Topplista I made in PHP during gymnasiet. 

Gopplista uses Go, and the web server framework Fiber. (I like Fiber, even though Echo and the standard package seems more popular. Fiber has many features I like such as nested layouts.)

On top of Go and the built in HTML templating in Go, HTMX is used for the frontend as well as some web components using Lit. 

The frontend uses UnoCSS for styling.

The database used is SQLite and sqlc is used to generate queries.

## Development

First run `npm i` to install dev dependencies for UnoCSS and Prettier.

To develop with live reload, use Air. Install Air, then run `air`.

Create a .env file in the root directory and add `ENV=dev` to it to use UnoCSS CDN. You will see a development box in the bottom left corner when in development.

## Building

Once everything is ready the project can be built locally using the Makefile, as well as the Dockerfile. 

## Environment variables

```
ENV=dev # if you are in development
SEED_DATABASE=true # if you want to populate the database with some predefined data
ADMIN_PASSWORD=<password>
```

## What I've learned from this project

SQLite is good for data storage, but the triggers and other tools are not as powerful as when using Postgres. This means the business logic has to live in the application code, which isn't so bad after all.

I've learned to set up and work with Go, and I like the language in general. I don't like the module structure, or I have just not figured it out yet, but I want to be able to nest my files more nicely.

I've learned to create a Dockerfile with separate build steps. It's very convenient that you can simply run a binary with Go, very efficient.

