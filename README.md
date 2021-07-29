# tradebot

Decentralized cryptocurrency trading platform with portfolio, accounting / tax reports, and automated trading bot. See [README-OVERVIEW](README-OVERVIEW.md) for details.

## Current Features
* Run locally (your computer, your data), centralized (hosted), or decentralized (Ethereum).
* User friendly web interface
* Plugin architecture supports financial indicators, trading strategies, exchanges and wallets
* Portfolio shows hosted exchange and offline wallet balances
* Exchange order / trade history import via API and CSV
* Accounting / tax reporting (form 8949 statement)
* Trading bot to automatically execute trades based on configured trading strategies / indicators
* Json Web Token (JWT) protected APIs
* REST and WebSocket APIs


## Roadmap
* Support LILO, dollar value, average cost, and specific identification tax reporting strategies
* Decentralized trading protocol
* Feature voting / bounties
* Marketplace for custom trading strategies and indicators
* Live Charts
* Integrate with open source, decentralized atomic swap exchange(s)
* Financial lending
* Social and communication integration (community chat, share portfolio to social media, etc)


## Build

Due to limitations of [Golang plugins](https://golang.org/pkg/plugin/), this software requires a Linux or Mac operating system to run natively. [Debian](https://www.debian.org/) is a great choice. [Docker](https://www.docker.com/) support is included for Windows users.


    git clone git@github.com:jeremyhahn/tradebot.git $GOPATH/src/https://github.com/wattson-coder/tradebot
    cd $GOPATH/src/https://github.com/wattson-coder/tradebot
    make deps
    make
    make test


#### Dependencies

1. [Golang](https://golang.org/)
2. [Make](https://www.gnu.org/software/make/)
3. [Yarn](https://yarnpkg.com/lang/en/docs/install/)
4. [OpenSSL](https://www.openssl.org/)
5. [Docker](https://www.docker.com/) (Optional)


#### Linux / Mac OS - Native

After installing the dependencies listed above, simply run `make` to build the `tradebot` binary. The database will need to be initialized the first time and then it can be started.

    # Install dependencies
    make deps

    # Build backend and webapp
    make

    # Initialize the database
    ./tradebot --initdb

    # Start in debug mode / logging
    ./tradebot --debug


#### Linux / Mac OS - Docker

Simply install docker and run the `build-docker.sh` bash script in the project root. Run `docker-run.sh` to start a container with the resulting image.

    ./docker-build.sh
    ./docker-run.sh


#### Windows

1. [Docker Toolkit](https://docs.docker.com/toolbox/toolbox_install_windows/) (Windows Home Edition or < Windows 10)
2. [Docker Community Edition](https://store.docker.com/editions/community/docker-ce-desktop-windows) (Windows 10 Professional or Enterprise)

Windows Home Edition and versions prior to Windows 10 do not support native virtualization. As such, they require `docker-machine` included in the Docker Toolkit. Newer versions of Windows that support native virtualization can take advantage of the latest Docker CE with enhanced performance.


## Development :: Tech Stack


#### User Interface
* [React.js](https://reactjs.org/)
* [Material-UI](https://material-ui-next.com/)


#### Backend
* [Golang](https://golang.org/)
* [Ethereum](https://www.ethereum.org/)
* [GORM](http://gorm.io/)
* [SQLite](https://www.sqlite.org/)


## Requirements


#### Firewall

Geth requires both TCP and UDP port 30303, otherwise it will not be able to synchronize the chain data with peers on the public network. Be sure these ports are forwarded to the system hosting Geth if it's behind a router/firewall.


## Broken

1. Fix broken indicators due to float64 -> decimal refactoring
2. Add fiat conversion to Bittrex order history (currently only supported by CSV import)


## Support

Join the [Telegram Channel](https://t.me/joinchat/AAAAAE3ha9a8OpK4bJFomQ) for assistance.
