<h1 align="center">VGTracker</h1>

<div align="center">
  <strong>Keep track of your backlog</strong>
</div>
<div align="center">
  An application for tracking your video game backlog across platforms
</div>

<br />

<div align="center">
  <!-- Tauri -->
  <a href="https://tauri.app/">
    <img src="https://img.shields.io/badge/tauri-1.6.1-orange.svg?style=flat-square&logo=tauri"
      alt="API stability" />
  </a>
  <!-- Rust -->
  <a href="https://www.rust-lang.org/">
    <img src="https://img.shields.io/badge/Rust Version-1.80 nightly-orange.svg?style=flat-square&logo=rust"
      alt="Rust Version" />
  </a>
  <!-- Bun -->
  <a href="https://github.com/oven-sh/bun">
    <img src="https://img.shields.io/badge/bun-1.12.0-FBF0DF.svg?style=flat-square&logo=bun"
      alt="NPM version" />
  </a>
  <!-- Vue -->
  <a href="https://vuejs.org/">
    <img src="https://img.shields.io/badge/Vue-3.3.4-43B883.svg?style=flat-square&logo=vue.js"
      alt="NPM version" />
  </a>
</div>

## Table of Contents
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Features](#features)
    - [Wishlist Features](#wishlist-features)
- [Getting Started](#getting-started)
  - [Running Application (Non Development)](#running-application-non-development)
  - [Recommended Dev Setup](#recommended-dev-setup)
- [Resources](#resources)

## Introduction
I was having trouble keeping track of all of my video games across all platforms or media(physical/digital) so I am building this app to 
help me keep track of my backlog. This app will be built using Tauri, Vue, and Rust; all tools that I don't know as a learning experience.

## Features
- [ ] Search games to add to your backlog via IGDB
- [ ] Light/Dark mode
- [ ] Categories for games (Playing, Completed, Wishlist, etc)
- [ ] Trophy/Achievement tracking
#### Wishlist Features
- [ ] Sync with Steam
- [ ] Sync with PSN
- [ ] Notifications for sales
- [ ] Custom theme colors
- [ ] Auto Updater

## Getting Started
### Running Application (Non Development)
Running the application for yourself, you can download the assets for your 
appropriate platform (mac/linux/windows) from the releases page. It is currently not signed 
so run at your own risk. You will also need to follow the setup for IGDB and acquiring a token for the application to work properly. It can be found [here](https://api-docs.igdb.com/#getting-started)

### Recommended Dev Setup
- [VS Code](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) + [Tauri](https://marketplace.visualstudio.com/items?itemName=tauri-apps.tauri-vscode) + [rust-analyzer](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust-analyzer)
- Requirements:
  - [Bun](https://bun.sh/docs/installation)
  - [Rust nightly](https://www.rust-lang.org/tools/install)
  - [Tauri cli](https://crates.io/crates/tauri-cli)
- Running application:
  - `bun install`
  - `bun run tauri dev`
    - configuration files and db/migrations should be created in os specific area
      - Linux: Resolves to $HOME.
      - macOS: Resolves to $HOME.
      - Windows: Resolves to {FOLDERID_Profile}.

## Resources
This is just a list of things that I used that helped me in some way.
- [PrimeVue + TailwindCSS](https://tailwind.primevue.org/overview/)
