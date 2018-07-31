## Go EOS Classic

Official golang implementation of the EOS Classic protocol.

[![Travis](https://travis-ci.org/eosclassic/go-eosclassic.svg?branch=master)](https://travis-ci.org/eosclassic/go-eosclassic)
[![Appveyor](https://ci.appveyor.com/api/projects/status/kj14asyrfkgg8vmn/branch/master?svg=true)](https://ci.appveyor.com/project/eosclassicteam/go-eosclassic/branch/master)
[![Docker](https://img.shields.io/docker/build/eosclassic/go-eosclassic.svg)](https://hub.docker.com/r/eosclassic/go-eosclassic)
[![Github All Releases](https://img.shields.io/github/downloads/eosclassic/go-eosclassic/total.svg)](https://github.com/eosclassic/go-eosclassic/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/eosclassic/go-eosclassic)](https://goreportcard.com/report/github.com/eosclassic/go-eosclassic)

Binary archives are published at https://github.com/eosclassic/go-eosclassic/releases/.

## Building the source

Building eosc requires both a Go (version 1.7 or later) and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

    make eosc

or, to build the full suite of utilities:

    make all

## Executables

The go-eosclassic project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`eosc`** | Our main EOS Classic CLI client. It is the entry point into the EOS Classic network (main-, test- or private net), capable of running as a full node (default) archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the EOS Classic network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `eosc --help` and the [CLI Wiki page](https://github.com/ethereum/go-ethereum/wiki/Command-Line-Options) for command line options. |

## Running eosc

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://github.com/ethereum/go-ethereum/wiki/Command-Line-Options)), but we've
enumerated a few common parameter combos to get you up to speed quickly on how you can run your
own EOSC instance.

### Full node on the main EOS Classic network

By far the most common scenario is people wanting to simply interact with the EOS Classic network:
create accounts; transfer funds; deploy and interact with contracts. For this particular use-case
the user doesn't care about years-old historical data, so we can fast-sync quickly to the current
state of the network. To do so:

```
$ eosc console
```

This command will:

 * Start eosc in fast sync mode (default, can be changed with the `--syncmode` flag), causing it to
   download more data in exchange for avoiding processing the entire history of the EOS Classic network,
   which is very CPU intensive.
 * Start up EOSC's built-in interactive [JavaScript console](https://github.com/ethereum/go-ethereum/wiki/JavaScript-Console),
   (via the trailing `console` subcommand) through which you can invoke all official [`web3` methods](https://github.com/ethereum/wiki/wiki/JavaScript-API)
   as well as EOSC's own [management APIs](https://github.com/ethereum/go-ethereum/wiki/Management-APIs).
   This too is optional and if you leave it out you can always attach to an already running EOSC instance
   with `eosc attach`.

### Configuration

As an alternative to passing the numerous flags to the `eosc` binary, you can also pass a configuration file via:

```
$ eosc --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to export your existing configuration:

```
$ eosc --your-favourite-flags dumpconfig
```

*Note: This works only with eosc v1.6.0 and above.*

#### Docker quick start

One of the quickest ways to get EOS Classic up and running on your machine is by using Docker:

```
docker run -d --name ethereum-node -v /Users/alice/eosclassic:/root \
           -p 8282:8282 -p 25252:25252 \
           eosclassic/go-eosclassic
```

This will start eosc in fast-sync mode with a DB memory allowance of 1GB just as the above command does.  It will also create a persistent volume in your home directory for saving your blockchain as well as map the default ports. There is also an `alpine` tag available for a slim version of the image.

Do not forget `--rpcaddr 0.0.0.0`, if you want to access RPC from other containers and/or hosts. By default, `eosc` binds to the local interface and RPC endpoints is not accessible from the outside.

### Programatically interfacing EOSC nodes

As a developer, sooner rather than later you'll want to start interacting with EOSC and the EOS Classic
network via your own programs and not manually through the console. To aid this, EOSC has built-in
support for a JSON-RPC based APIs ([standard APIs](https://github.com/ethereum/wiki/wiki/JSON-RPC) and
[EOSC specific APIs](https://github.com/ethereum/go-ethereum/wiki/Management-APIs)). These can be
exposed via HTTP, WebSockets and IPC (unix sockets on unix based platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by EOSC, whereas the HTTP
and WS interfaces need to manually be enabled and only expose a subset of APIs due to security reasons.
These can be turned on/off and configured as you'd expect.

HTTP based JSON-RPC API options:

  * `--rpc` Enable the HTTP-RPC server
  * `--rpcaddr` HTTP-RPC server listening interface (default: "localhost")
  * `--rpcport` HTTP-RPC server listening port (default: 8282)
  * `--rpcapi` API's offered over the HTTP-RPC interface (default: "eth,net,web3")
  * `--rpccorsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--wsaddr` WS-RPC server listening interface (default: "localhost")
  * `--wsport` WS-RPC server listening port (default: 8546)
  * `--wsapi` API's offered over the WS-RPC interface (default: "eth,net,web3")
  * `--wsorigins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: "admin,debug,eth,miner,net,personal,shh,txpool,web3")
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to connect
via HTTP, WS or IPC to a EOSC node configured with the above flags and you'll need to speak [JSON-RPC](http://www.jsonrpc.org/specification)
on all transports. You can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based transport before
doing so! Hackers on the internet are actively trying to subvert EOS Classic nodes with exposed APIs!
Further, all browser tabs can access locally running webservers, so malicious webpages could try to
subvert locally available APIs!**

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.

## Credits

Originally made by [go-ethereum authors](https://github.com/ethereum/go-ethereum)

Modified by eosclassicteam for EOS Classic usage.
