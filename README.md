# VGTracker

[![Go](https://img.shields.io/badge/go-1.24-blue.svg?logo=go)](https://go.dev/)
[![Vue](https://img.shields.io/badge/vue-3.5-green.svg?logo=Vue.js)](https://vuejs.org/)
[![Bun](https://img.shields.io/badge/bun-1.2.4-orange.svg?logo=bun)](https://github.com/oven-sh/bun)

## Table of Contents

- [About](https://github.com/Stegnerd/vgtracker#about)
- [Setup](https://github.com/Stegnerd/vgtracker#setup)
- [Architecture](https://github.com/Stegnerd/vgtracker#architecture)
- [Technology](https://github.com/Stegnerd/vgtracker#technologies)
- [Road Map](https://github.com/Stegnerd/vgtracker#roadmap)
- [Sources](https://github.com/Stegnerd/vgtracker#sources)

## About

This is a project of mine to keep track of my video games and get some practice with Vue! I will implement features as I think of them

## Setup

### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on `http://localhost:34115`. Connect
to this in your browser, and you can call your Go code from devtools.

### Building

To build a redistributable, production mode package, use `wails build`.

### Migrations

- install [Mage](https://magefile.org/) (do not need to do mage init here)
- install [Golang-Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain) -
  - make sure that you install via toolchain with the tag "sqlite3" and NOT homebrew
  - you will now if you did it wrong if you see the error: "Unknown Driver"
- To make a new migration run `mage MigrateCreate <name of migration>` and you should see it in internal/db/migration
- To run migrations up `mage MigrateUp`
- To run migrations down `mage MigrateDown`

## Architecture

## Technologies

## Roadmap

## Sources
