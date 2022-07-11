# GO Package Header Server

[![Build Status](https://ci.mrmelon54.xyz/api/badges/alfred/GOPackageHeaderServer/status.svg)](https://ci.mrmelon54.xyz/alfred/GOPackageHeaderServer)

This allows for the required meta headers to be outputted in order for the GO package system to find the source files of the package.

The middleware can be configured in runtime, the server has a YAML configuration.

Maintainer: 
[Captain ALM](https://code.mrmelon54.xyz/alfred)

License: 
[BSD 3-Clause](https://code.mrmelon54.xyz/alfred/GOPackageHeaderServer/src/branch/master/LICENSE.md)

Example configuration: 
[config.example.yml](https://code.mrmelon54.xyz/alfred/GOPackageHeaderServer/src/branch/master/config.example.yml) 
The configuration must by placed in a .data sub-directory from the executable. A .env file must also be generated (Can be empty).