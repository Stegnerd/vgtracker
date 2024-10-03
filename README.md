# VGTracker

[![Rust Version](https://img.shields.io/badge/rust-1.82.0-orange.svg?logo=rust)](https://www.rust-lang.org)
[![Vue Version](https://img.shields.io/badge/vue-3.3.4-green.svg?logo=vuedotjs)](https://vuejs.org)
[![Bun Version](https://img.shields.io/badge/bun-1.1.29-pink.svg?logo=bun)](https://bun.sh)

This is a project to learn rust and other tools that I do not use at the moment like vue and unocss.
Also creating a tool that I will use to keep track of my games

## Table of Contents

- [RoadMap](https://github.com/Stegnerd/vgtracker#roadmap)
- [Development](https://github.com/Stegnerd/vgtracker#technology)
- [Sources](https://github.com/Stegnerd/vgtracker#sources)

### RoadMap

These are goals for this project:

- [ ] search games from IGDB
- [ ] persist profile
  - [ ] inventory
  - [ ] wishlist
  - [ ] stats
- [ ] set up some kind of build
- [ ] make some kind of logo
  
#### v2

- [ ] include DLC
- [ ] auto sync from profiles (psn/steam/etc)
- [ ] notifications of sales from inventory

## Development

### Setup

- Install [Rust](https://www.rust-lang.org/)
- Install [Bun](https://bun.sh)
- run `bun install` at root of prject
- cd into src-tauri and run `cargo install`
- run `bun run tauri dev` at root and it should be running

### Testing

- Backend -
- UI - Vitest -

## Sources

- [Tauri](https://tauri.app/v1/guides/development/development-cycle) - application framework built with rust
