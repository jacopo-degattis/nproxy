# Nproxy

## Description
Navidrome middleware to fetch songs and data from remote provider if not locally found 

## Structure

Here's a description of the most important folders and files in the project.

```
├── server/
├── middlewares/
├── redisdb/
```

The `lib` contains core library files.
The 

Inside the `middlewares` folder you can place your middlewares code, see [dabmusic](middlewares/dabmusic) implementation for reference

The `redisdb` folder contains the redis client necessary for albums' covers cache feature

## Disclaimer

This project, as of now, is  made with the main focus of serving my needs but any suggestion or contribution is welcome
