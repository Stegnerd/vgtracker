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

- Install [Bun](https://github.com/oven-sh/bun) - homebrew is fine
  - when you need to add a new depedency, use `bun add`
- Install [Go](https://go.dev/) - homebrew is fine
- install [Golang-Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain)
  - - make sure that you install via toolchain with the tag "sqlite3" and NOT homebrew
- Install [Wails](https://wails.io/docs/gettingstarted/installation)
- run `wails dev` - this will run go mod tidy, install frontend dependencies, and db migrations
  - This will run a Vite development
    server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
    and have access to your Go methods, there is also a dev server that runs on `http://localhost:34115`. Connect
    to this in your browser, and you can call your Go code from devtools.
- Set up your own [twitch account for igdb access](https://api-docs.igdb.com/#getting-started)
  - this is needed for igdb api access and you need to supply your own api creds, not going to share mine!
- To enable syncing with steam you will need to sign up for a steam api key [here](https://steamcommunity.com/dev/apikey)
  <!--
  potentially other links?
  - [PlayStation Network](https://www.playstation.com/en-us/develop/)
  - [Xbox Live](https://developer.microsoft.com/en-us/games/xbox/)
  - [Steam](https://partner.steamgames.com/)
  - [Nintendo](https://developer.nintendo.com/) -->

### Building

To build a redistributable, production mode package, use `wails build`. You can also grab a build from the releases page <br>
[here](https://github.com/Stegnerd/vgtracker/releases). Currently only Linux and macos are supported. I have no windows devices to test
and I don't support windows os. <br> I also am not signing builds, not going to pay for that but thats a talk for a different time!
You can get around that with [here](https://support.apple.com/guide/mac-help/open-a-mac-app-from-an-unknown-developer-mh40616/mac).

- [Bumping versions](https://github.com/mathieudutour/github-tag-action?tab=readme-ov-file#bumping)

### Migrations

- install [Mage](https://magefile.org/) (do not need to do mage init here)
- install [Golang-Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain) -
  - make sure that you install via toolchain with the tag "sqlite3" and NOT homebrew
  - you will now if you did it wrong if you see the error: "Unknown Driver"
- To make a new migration run `mage MigrateCreate <name of migration>` and you should see it in internal/db/migration
- To run migrations up `mage DB:MigrateUp`
- To run migrations down `mage DB:MigrateDown`

## Architecture

## Technologies

## Roadmap

### todos

- [ ] double check codegen works between time.Time and frontend
- [ ] look into making the theme composable more stream lined/pinia?
- [ ] more unit test coverage
- [ ] keyboard short cuts (vueUse?)

### v1

- [x] able to save backlog stats
- [ ] analytics dashboard for stats
- [ ] wishlist/ list making
- [ ] e2e with playwright
- [ ] steam syncing

### cool ideas

- [ ] image recognition of game covers to add to backlog
- [ ] ability to sync psn/xbox/steam to pull in your digital libraries
  - not sure how to check nintendo library
- [ ] auto updating?

## Sources

- [IGDB API](https://api-docs.igdb.com/#getting-started)
- [Go Embed for Migrations](https://oscarforner.com/blog/2023-10-10-go-embed-for-migrations/)
- [Generating enums for wails](https://wails.io/docs/guides/application-development#dealing-with-context-when-binding-multiple-structs)
- [Icon look up](https://icones.js.org/)
- [Steam Web API](https://partner.steamgames.com/doc/webapi)
