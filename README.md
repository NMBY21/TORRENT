# Bit_Torrent Client Implementation

This is a BitTorrent client implementation built with [Golang](https://www.golang.org/) using the [BitTorrent Protocol Specification](https://www.bittorrent.org/beps/bep_0003.html).

This project is a BitTorrent client implementation for downloading and sharing files over the BitTorrent protocol. It's easy to use, customizable, and supports downloading from multiple sources.

## Overview

This implementation supports the following features:

- Downloading and seeding of torrent files
- Magnet link support
- n-leechers and n-seeders
- n-peers and etc

## Running the BitTorrent Client Implementation

To run the BitTorrent client implementation, please follow these steps:

1.  Initialize a new Go module for the BitTorrent project using the command

             go mod init bittorrent

    This will create a new file named `go.mod` in the root directory of the project.

2.  Tidy up the module dependencies using the command

             go mod tidy

    This will ensure that all the required dependencies are downloaded and available in the project.

3.  Build the BitTorrent client using the command

             go build

    This will compile the project and generate the executable file.

4.  Run the BitTorrent client using the command

           ./bittorrent -torrentfile <magnet link or .torrentfile>

    Replace `<magnet link or .torrentfile>`
    with the actual magnet link or .torrent file you want to download. This will start the BitTorrent client and initiate the download process.

Once you have completed these steps, the BitTorrent client should be up and running, and you should be able to download files using the BitTorrent protocol.
