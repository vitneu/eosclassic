// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/eosclassic/go-eosclassic/log"
)

// dashboardContent is the actual dashboard HTML content to serve up when users
// load the dashboard website.
var dashboardContent = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<!-- Meta, title, CSS, favicons, etc. -->
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">

		<title>{{.NetworkTitle}}: Ethereum Testnet</title>

		<link href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
		<link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">
		<link href="https://cdnjs.cloudflare.com/ajax/libs/gentelella/1.3.0/css/custom.min.css" rel="stylesheet">
		<style>
			.vertical-center {
				min-height: 100%;
				min-height: 95vh;
				display: flex;
				align-items: center;
			}
			.nav.side-menu li a {
				font-size: 18px;
			}
			.nav-sm .nav.side-menu li a {
				font-size: 10px;
			}
			pre{
				white-space: pre-wrap;
			}
		</style>
	</head>

	<body class="nav-sm" style="overflow-x: hidden">
		<div class="container body">
			<div class="main_container">
				<div class="col-md-3 left_col">
					<div class="left_col scroll-view">
						<div class="navbar nav_title" style="border: 0; margin-top: 8px;">
							<a class="site_title"><i class="fa fa-globe" style="margin-left: 6px"></i> <span>{{.NetworkTitle}} Testnet</span></a>
						</div>
						<div class="clearfix"></div>
						<br />
						<div id="sidebar-menu" class="main_menu_side hidden-print main_menu">
							<div class="menu_section">
								<ul class="nav side-menu">
									{{if .EthstatsPage}}<li id="stats_menu"><a onclick="load('#stats')"><i class="fa fa-tachometer"></i> Network Stats</a></li>{{end}}
									{{if .ExplorerPage}}<li id="explorer_menu"><a onclick="load('#explorer')"><i class="fa fa-database"></i> Block Explorer</a></li>{{end}}
									{{if .WalletPage}}<li id="wallet_menu"><a onclick="load('#wallet')"><i class="fa fa-address-book-o"></i> Browser Wallet</a></li>{{end}}
									{{if .FaucetPage}}<li id="faucet_menu"><a onclick="load('#faucet')"><i class="fa fa-bath"></i> Crypto Faucet</a></li>{{end}}
									<li id="connect_menu"><a><i class="fa fa-plug"></i> Connect Yourself</a>
										<ul id="connect_list" class="nav child_menu">
											<li><a onclick="$('#connect_menu').removeClass('active'); $('#connect_list').toggle(); load('#geth')">Go Ethereum: Geth</a></li>
											<li><a onclick="$('#connect_menu').removeClass('active'); $('#connect_list').toggle(); load('#mist')">Go Ethereum: Wallet & Mist</a></li>
											<li><a onclick="$('#connect_menu').removeClass('active'); $('#connect_list').toggle(); load('#mobile')">Go Ethereum: Android & iOS</a></li>{{if .Ethash}}
											<li><a onclick="$('#connect_menu').removeClass('active'); $('#connect_list').toggle(); load('#other')">Other Ethereum Clients</a></li>{{end}}
										</ul>
									</li>
									<li id="about_menu"><a onclick="load('#about')"><i class="fa fa-heartbeat"></i> About Puppeth</a></li>
								</ul>
							</div>
						</div>
					</div>
				</div>
				<div class="right_col" role="main" style="padding: 0 !important">
					<div id="geth" hidden style="padding: 16px;">
						<div class="page-title">
							<div class="title_left">
								<h3>Connect Yourself &ndash; Go Ethereum: Geth</h3>
							</div>
						</div>
						<div class="clearfix"></div>
						<div class="row">
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2><i class="fa fa-archive" aria-hidden="true"></i> Archive node <small>Retains all historical data</small></h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>An archive node synchronizes the blockchain by downloading the full chain from the genesis block to the current head block, executing all the transactions contained within. As the node crunches through the transactions, all past historical state is stored on disk, and can be queried for each and every block.</p>
										<p>Initial processing required to execute all transactions may require non-negligible time and disk capacity required to store all past state may be non-insignificant. High end machines with SSD storage, modern CPUs and 8GB+ RAM are recommended.</p>
										<br/>
										<p>To run an archive node, download <a href="/{{.GethGenesis}}"><code>{{.GethGenesis}}</code></a> and start Geth with:
											<pre>geth --datadir=$HOME/.{{.Network}} init {{.GethGenesis}}</pre>
											<pre>geth --networkid={{.NetworkID}} --datadir=$HOME/.{{.Network}} --cache=1024 --syncmode=full{{if .Ethstats}} --ethstats='{{.Ethstats}}'{{end}} --bootnodes={{.BootnodesFlat}}</pre>
										</p>
										<br/>
										<p>You can download Geth from <a href="https://geth.ethereum.org/downloads/" target="about:blank">https://geth.ethereum.org/downloads/</a>.</p>
									</div>
								</div>
							</div>
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2><i class="fa fa-laptop" aria-hidden="true"></i> Full node <small>Retains recent data only</small></h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>A full node synchronizes the blockchain by downloading the full chain from the genesis block to the current head block, but does not execute the transactions. Instead, it downloads all the transactions receipts along with the entire recent state. As the node downloads the recent state directly, historical data can only be queried from that block onward.</p>
										<p>Initial processing required to synchronize is more bandwidth intensive, but is light on the CPU and has significantly reduced disk requirements. Mid range machines with HDD storage, decent CPUs and 4GB+ RAM should be enough.</p>
										<br/>
										<p>To run a full node, download <a href="/{{.GethGenesis}}"><code>{{.GethGenesis}}</code></a> and start Geth with:
											<pre>geth --datadir=$HOME/.{{.Network}} init {{.GethGenesis}}</pre>
											<pre>geth --networkid={{.NetworkID}} --datadir=$HOME/.{{.Network}} --cache=512{{if .Ethstats}} --ethstats='{{.Ethstats}}'{{end}} --bootnodes={{.BootnodesFlat}}</pre>
										</p>
										<br/>
										<p>You can download Geth from <a href="https://geth.ethereum.org/downloads/" target="about:blank">https://geth.ethereum.org/downloads/</a>.</p>
									</div>
								</div>
							</div>
						</div>
						<div class="clearfix"></div>
						<div class="row">
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2><i class="fa fa-mobile" aria-hidden="true"></i> Light node <small>Retrieves data on demand</small></h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>A light node synchronizes the blockchain by downloading and verifying only the chain of headers from the genesis block to the current head, without executing any transactions or retrieving any associated state. As no state is available locally, any interaction with the blockchain relies on on-demand data retrievals from remote nodes.</p>
										<p>Initial processing required to synchronize is light, as it only verifies the validity of the headers; similarly required disk capacity is small, tallying around 500 bytes per header. Low end machines with arbitrary storage, weak CPUs and 512MB+ RAM should cope well.</p>
										<br/>
										<p>To run a light node, download <a href="/{{.GethGenesis}}"><code>{{.GethGenesis}}</code></a> and start Geth with:
											<pre>geth --datadir=$HOME/.{{.Network}} init {{.GethGenesis}}</pre>
											<pre>geth --networkid={{.NetworkID}} --datadir=$HOME/.{{.Network}} --syncmode=light{{if .Ethstats}} --ethstats='{{.Ethstats}}'{{end}} --bootnodes={{.BootnodesFlat}}</pre>
										</p>
										<br/>
										<p>You can download Geth from <a href="https://geth.ethereum.org/downloads/" target="about:blank">https://geth.ethereum.org/downloads/</a>.</p>
									</div>
								</div>
							</div>
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2><i class="fa fa-microchip" aria-hidden="true"></i> Embedded node <small>Conserves memory vs. speed</small></h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>An embedded node is a variation of the light node with configuration parameters tuned towards low memory footprint. As such, it may sacrifice processing and disk IO performance to conserve memory. It should be considered an <strong>experimental</strong> direction for now without hard guarantees or bounds on the resources used.</p>
										<p>Initial processing required to synchronize is light, as it only verifies the validity of the headers; similarly required disk capacity is small, tallying around 500 bytes per header. Embedded machines with arbitrary storage, low power CPUs and 128MB+ RAM may work.</p>
										<br/>
										<p>To run an embedded node, download <a href="/{{.GethGenesis}}"><code>{{.GethGenesis}}</code></a> and start Geth with:
											<pre>geth --datadir=$HOME/.{{.Network}} init {{.GethGenesis}}</pre>
											<pre>geth --networkid={{.NetworkID}} --datadir=$HOME/.{{.Network}} --cache=16 --ethash.cachesinmem=1 --syncmode=light{{if .Ethstats}} --ethstats='{{.Ethstats}}'{{end}} --bootnodes={{.BootnodesFlat}}</pre>
										</p>
										<br/>
										<p>You can download Geth from <a href="https://geth.ethereum.org/downloads/" target="about:blank">https://geth.ethereum.org/downloads/</a>.</p>
									</div>
								</div>
							</div>
						</div>
					</div>
					<div id="mist" hidden style="padding: 16px;">
						<div class="page-title">
							<div class="title_left">
								<h3>Connect Yourself &ndash; Go Ethereum: Wallet &amp; Mist</h3>
							</div>
						</div>
						<div class="clearfix"></div>
						<div class="row">
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2><i class="fa fa-credit-card" aria-hidden="true"></i> Desktop wallet <small>Interacts with accounts and contracts</small></h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>The Ethereum Wallet is an <a href="https://electron.atom.io/" target="about:blank">Electron</a> based desktop application to manage your Ethereum accounts and funds. Beside the usual account life-cycle operations you would expect to perform, the wallet also provides a means to send transactions from your accounts and to interact with smart contracts deployed on the network.</p>
										<p>Under the hood the wallet is backed by a go-ethereum full node, meaning that a mid range machine is assumed. Similarly, synchronization is based on <strong>fast-sync</strong>, which will download all blockchain data from the network and make it available to the wallet. Light nodes cannot currently fully back the wallet, but it's a target actively pursued.</p>
										<br/>
										<p>To connect with the Ethereum Wallet, you'll need to initialize your private network first via Geth as the wallet does not currently support calling Geth directly. To initialize your local chain, download <a href="/{{.GethGenesis}}"><code>{{.GethGenesis}}</code></a> and run:
											<pre>geth --datadir=$HOME/.{{.Network}} init {{.GethGenesis}}</pre>
										</p>
										<p>With your local chain initialized, you can start the Ethereum Wallet:
											<pre>ethereumwallet --rpc $HOME/.{{.Network}}/geth.ipc --node-networkid={{.NetworkID}} --node-datadir=$HOME/.{{.Network}}{{if .Ethstats}} --node-ethstats='{{.Ethstats}}'{{end}} --node-bootnodes={{.BootnodesFlat}}</pre>
										<p>
										<br/>
										<p>You can download the Ethereum Wallet from <a href="https://github.com/ethereum/mist/releases" target="about:blank">https://github.com/ethereum/mist/releases</a>.</p>
									</div>
								</div>
							</div>
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2><i class="fa fa-picture-o" aria-hidden="true"></i> Mist browser <small>Interacts with third party DApps</small></h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>The Mist browser is an <a href="https://electron.atom.io/" target="about:blank">Electron</a> based desktop application to load and interact with Ethereum enabled third party web DApps. Beside all the functionality provided by the Ethereum Wallet, Mist is an extended web-browser where loaded pages have access to the Ethereum network via a web3.js provider, and may also interact with users' own accounts (given proper authorization and confirmation of course).</p>
										<p>Under the hood the browser is backed by a go-ethereum full node, meaning that a mid range machine is assumed. Similarly, synchronization is based on <strong>fast-sync</strong>, which will download all blockchain data from the network and make it available to the wallet. Light nodes cannot currently fully back the wallet, but it's a target actively pursued.</p>
										<br/>
										<p>To connect with the Mist browser, you'll need to initialize your private network first via Geth as Mist does not currently support calling Geth directly. To initialize your local chain, download <a href="/{{.GethGenesis}}"><code>{{.GethGenesis}}</code></a> and run:
											<pre>geth --datadir=$HOME/.{{.Network}} init {{.GethGenesis}}</pre>
										</p>
										<p>With your local chain initialized, you can start Mist:
											<pre>mist --rpc $HOME/.{{.Network}}/geth.ipc --node-networkid={{.NetworkID}} --node-datadir=$HOME/.{{.Network}}{{if .Ethstats}} --node-ethstats='{{.Ethstats}}'{{end}} --node-bootnodes={{.BootnodesFlat}}</pre>
										<p>
										<br/>
										<p>You can download the Mist browser from <a href="https://github.com/ethereum/mist/releases" target="about:blank">https://github.com/ethereum/mist/releases</a>.</p>
									</div>
								</div>
							</div>
						</div>
					</div>
					<div id="mobile" hidden style="padding: 16px;">
						<div class="page-title">
							<div class="title_left">
								<h3>Connect Yourself &ndash; Go Ethereum: Android &amp; iOS</h3>
							</div>
						</div>
						<div class="clearfix"></div>
						<div class="row">
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2><i class="fa fa-android" aria-hidden="true"></i> Android devices <small>Accesses Ethereum via Java</small></h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>Starting with the 1.5 release of go-ethereum, we've transitioned away from shipping only full blown Ethereum clients and started focusing on releasing the code as reusable packages initially for Go projects, then later for Java based Android projects too. Mobile support is still evolving, hence is bound to change often and hard, but the Ethereum network can nonetheless be accessed from Android too.</p>
										<p>Under the hood the Android library is backed by a go-ethereum light node, meaning that given a not-too-old Android device, you should be able to join the network without significant issues. Certain functionality is not yet available and rough edges are bound to appear here and there, please report issues if you find any.</p>
										<br/>
										<p>The stable Android archives are distributed via Maven Central, and the develop snapshots via the Sonatype repositories. Before proceeding, please ensure you have a recent version configured in your Android project. You can find details in <a href="https://github.com/ethereum/go-ethereum/wiki/Mobile:-Introduction#android-archive" target="about:blank">Mobile: Introduction &ndash; Android archive</a>.
										<p>Before connecting to the Ethereum network, download the <a href="/{{.GethGenesis}}"><code>{{.GethGenesis}}</code></a> genesis json file and either store it in your Android project as a resource file you can access, or save it as a string in a variable. You're going to need to to initialize your client.</p>
										<p>Inside your Java code you can now import the geth archive and connect to Ethereum:
											<pre>import org.ethereum.geth.*;</pre>
<pre>
Enodes bootnodes = new Enodes();{{range .Bootnodes}}
bootnodes.append(new Enode("{{.}}"));{{end}}

NodeConfig config = new NodeConfig();
config.setBootstrapNodes(bootnodes);
config.setEthereumNetworkID({{.NetworkID}});
config.setEthereumGenesis(genesis);{{if .Ethstats}}
config.setEthereumNetStats("{{.Ethstats}}");{{end}}

Node node = new Node(getFilesDir() + "/.{{.Network}}", config);
node.start();
</pre>
										<p>
									</div>
								</div>
							</div>
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2><i class="fa fa-apple" aria-hidden="true"></i> iOS devices <small>Accesses Ethereum via ObjC/Swift</small></h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>Starting with the 1.5 release of go-ethereum, we've transitioned away from shipping only full blown Ethereum clients and started focusing on releasing the code as reusable packages initially for Go projects, then later for ObjC/Swift based iOS projects too. Mobile support is still evolving, hence is bound to change often and hard, but the Ethereum network can nonetheless be accessed from iOS too.</p>
										<p>Under the hood the iOS library is backed by a go-ethereum light node, meaning that given a not-too-old Apple device, you should be able to join the network without significant issues. Certain functionality is not yet available and rough edges are bound to appear here and there, please report issues if you find any.</p>
										<br/>
										<p>Both stable and develop builds of the iOS framework are available via CocoaPods. Before proceeding, please ensure you have a recent version configured in your iOS project. You can find details in <a href="https://github.com/ethereum/go-ethereum/wiki/Mobile:-Introduction#ios-framework" target="about:blank">Mobile: Introduction &ndash; iOS framework</a>.
										<p>Before connecting to the Ethereum network, download the <a href="/{{.GethGenesis}}"><code>{{.GethGenesis}}</code></a> genesis json file and either store it in your iOS project as a resource file you can access, or save it as a string in a variable. You're going to need to to initialize your client.</p>
										<p>Inside your Swift code you can now import the geth framework and connect to Ethereum (ObjC should be analogous):
											<pre>import Geth</pre>
<pre>
var error: NSError?

let bootnodes = GethNewEnodesEmpty(){{range .Bootnodes}}
bootnodes?.append(GethNewEnode("{{.}}", &error)){{end}}

let config = GethNewNodeConfig()
config?.setBootstrapNodes(bootnodes)
config?.setEthereumNetworkID({{.NetworkID}})
config?.setEthereumGenesis(genesis){{if .Ethstats}}
config?.setEthereumNetStats("{{.Ethstats}}"){{end}}

let datadir = NSSearchPathForDirectoriesInDomains(.documentDirectory, .userDomainMask, true)[0]
let node = GethNewNode(datadir + "/.{{.Network}}", config, &error);
try! node?.start();
</pre>
										<p>
									</div>
								</div>
							</div>
						</div>
					</div>{{if .Ethash}}
					<div id="other" hidden style="padding: 16px;">
						<div class="page-title">
							<div class="title_left">
								<h3>Connect Yourself &ndash; Other Ethereum Clients</h3>
							</div>
						</div>
						<div class="clearfix"></div>
						<div class="row">
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2>
											<svg height="14px" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 115 115"><path fill="#5C8DBC" d="M9.7 83.3V35.5s0-3.4 3.3-5.2c3.3-1.8 39.6-23.5 39.6-23.5s4.6-3.1 9.4 0c0 0 43.1 23.9 42.4 25.3L85.3 43.3s-3.6-8.4-13.1-13c-11.3-5.5-29.7-6.2-42.9 13.3 0 0-8.6 13.5.3 31.6l-19 10.7s-.9-.6-.9-2.6z"/><path fill="#5C8DBC" d="M71 51.3c-2.8-4.7-7.9-7.9-13.8-7.9-8.8 0-16 7.2-16 16 0 2.8.7 5.4 2 7.7L71 51.3z"/><path fill="#194674" d="M43.1 67c2.8 4.7 7.9 7.9 13.8 7.9 8.8 0 16-7.2 16-16 0-2.8-.7-5.4-2-7.7L43.1 67z"/><path fill="#1B598E" d="M104.4 32.1s1.3 52.6-.3 53.6L58 58.6l46.4-26.5z"/><path fill="#FFF" d="M90 57h-3.9v-4.1h-4.2V57h-4v4.1h4V65h4.2v-3.9H90zm13.6 0h-3.9v-4.1h-4.2V57h-4v4.1h4V65h4.2v-3.9h3.9z"/><path fill="#194674" d="M29.5 75.1s9.2 17 28.5 16.1 27.3-16.6 27.3-16.6L104 85.4s4.1.8-41.6 25.7c0 0-4.9 3.3-10.2 0 0 0-41.3-23.1-41.6-25.3l18.9-10.7z"/></svg>
											C++ Ethereum <small>Official C++ client from the Ethereum Foundation</small>
										</h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>C++ Ethereum is the third most popular of the Ethereum clients, focusing on code portability to a broad range of operating systems and hardware. The client is currently a full node with transaction processing based synchronization.</p>
										<br/>
										<p>To run a cpp-ethereum node, download <a href="/{{.CppGenesis}}"><code>{{.CppGenesis}}</code></a> and start the node with:
											<pre>eth --config {{.CppGenesis}} --datadir $HOME/.{{.Network}} --peerset "{{.CppBootnodes}}"</pre>
										</p>
										<br/>
										<p>You can find cpp-ethereum at <a href="https://github.com/ethereum/cpp-ethereum/" target="about:blank">https://github.com/ethereum/cpp-ethereum/</a>.</p>
									</div>
								</div>
							</div>
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2>
											<svg height="14px" version="1.1" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64"><path d="M46.42,13.07S24.51,18.54,35,30.6c3.09,3.55-.81,6.75-0.81,6.75s7.84-4,4.24-9.11C35,23.51,32.46,21.17,46.42,13.07ZM32.1,16.88C45.05,6.65,38.4,0,38.4,0c2.68,10.57-9.46,13.76-13.84,20.34-3,4.48,1.46,9.3,7.53,14.77C29.73,29.77,21.71,25.09,32.1,16.88Z" transform="translate(-8.4)" fill="#e57125"/><path d="M23.6,49.49c-9.84,2.75,6,8.43,18.51,3.06a23.06,23.06,0,0,1-3.52-1.72,36.62,36.62,0,0,1-13.25.56C21.16,50.92,23.6,49.49,23.6,49.49Zm17-5.36a51.7,51.7,0,0,1-17.1.82c-4.19-.43-1.45-2.46-1.45-2.46-10.84,3.6,6,7.68,21.18,3.25A7.59,7.59,0,0,1,40.62,44.13ZM51.55,54.68s1.81,1.49-2,2.64c-7.23,2.19-30.1,2.85-36.45.09-2.28-1,2-2.37,3.35-2.66a8.69,8.69,0,0,1,2.21-.25c-2.54-1.79-16.41,3.51-7,5C37.15,63.67,58.17,57.67,51.55,54.68ZM42.77,39.12a20.42,20.42,0,0,1,2.93-1.57s-4.83.86-9.65,1.27A87.37,87.37,0,0,1,20.66,39c-7.51-1,4.12-3.77,4.12-3.77A22,22,0,0,0,14.7,37.61C8.14,40.79,31,42.23,42.77,39.12Zm2.88,7.77a1,1,0,0,1-.24.31C61.44,43,55.54,32.35,47.88,35a2.19,2.19,0,0,0-1,.79,9,9,0,0,1,1.37-.37C52.1,34.66,57.65,40.65,45.64,46.89Zm0.43,14.75a94.76,94.76,0,0,1-29.17.45s1.47,1.22,9,1.7c11.53,0.74,29.22-.41,29.64-5.86C55.6,57.94,54.79,60,46.08,61.65Z" transform="translate(-8.4)" fill="#5482a2"/></svg>
											Ethereum Harmony<small>Third party Java client from EtherCamp</small>
										</h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>Ethereum Harmony is a web user-interface based graphical Ethereum client built on top of the EthereumJ Java implementation of the Ethereum protocol. The client currently is a full node with state download based synchronization.</p>
										<br/>
										<p>To run an Ethereum Harmony node, download <a href="/{{.HarmonyGenesis}}"><code>{{.HarmonyGenesis}}</code></a> and start the node with:
											<pre>./gradlew runCustom -DgenesisFile={{.HarmonyGenesis}} -Dpeer.networkId={{.NetworkID}} -Ddatabase.dir=$HOME/.harmony/{{.Network}} {{.HarmonyBootnodes}} </pre>
										</p>
										<br/>
										<p>You can find Ethereum Harmony at <a href="https://github.com/ether-camp/ethereum-harmony/" target="about:blank">https://github.com/ether-camp/ethereum-harmony/</a>.</p>
									</div>
								</div>
							</div>
						</div>
						<div class="clearfix"></div>
						<div class="row">
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2>
											<svg height="14px" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:cc="http://creativecommons.org/ns#" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:svg="http://www.w3.org/2000/svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 104.56749 104.56675" version="1.1" viewbox="0 0 144 144" y="0px" x="0px"><metadata id="metadata10"><rdf:RDF><cc:Work rdf:about=""><dc:format>image/svg+xml</dc:format><dc:type rdf:resource="http://purl.org/dc/dcmitype/StillImage" /></cc:Work></rdf:RDF></metadata><defs id="defs8" /><path style="fill:#676767;" id="path2" d="m 49.0125,12.3195 a 3.108,3.108 0 0 1 6.216,0 3.108,3.108 0 0 1 -6.216,0 m -37.077,28.14 a 3.108,3.108 0 0 1 6.216,0 3.108,3.108 0 0 1 -6.216,0 m 74.153,0.145 a 3.108,3.108 0 0 1 6.216,0 3.108,3.108 0 0 1 -6.216,0 m -65.156,4.258 c 1.43,-0.635 2.076,-2.311 1.441,-3.744 l -1.379,-3.118 h 5.423 v 24.444 h -10.941 a 38.265,38.265 0 0 1 -1.239,-14.607 z m 22.685,0.601 v -7.205 h 12.914 c 0.667,0 4.71,0.771 4.71,3.794 0,2.51 -3.101,3.41 -5.651,3.41 z m -17.631,38.793 a 3.108,3.108 0 0 1 6.216,0 3.108,3.108 0 0 1 -6.216,0 m 46.051,0.145 a 3.108,3.108 0 0 1 6.216,0 3.108,3.108 0 0 1 -6.216,0 m 0.961,-7.048 c -1.531,-0.328 -3.037,0.646 -3.365,2.18 l -1.56,7.28 a 38.265,38.265 0 0 1 -31.911,-0.153 l -1.559,-7.28 c -0.328,-1.532 -1.834,-2.508 -3.364,-2.179 l -6.427,1.38 a 38.265,38.265 0 0 1 -3.323,-3.917 h 31.272 c 0.354,0 0.59,-0.064 0.59,-0.386 v -11.062 c 0,-0.322 -0.236,-0.386 -0.59,-0.386 h -9.146 v -7.012 h 9.892 c 0.903,0 4.828,0.258 6.083,5.275 0.393,1.543 1.256,6.562 1.846,8.169 0.588,1.802 2.982,5.402 5.533,5.402 h 16.146 a 38.265,38.265 0 0 1 -3.544,4.102 z m 17.365,-29.207 a 38.265,38.265 0 0 1 0.081,6.643 h -3.926 c -0.393,0 -0.551,0.258 -0.551,0.643 v 1.803 c 0,4.244 -2.393,5.167 -4.49,5.402 -1.997,0.225 -4.211,-0.836 -4.484,-2.058 -1.178,-6.626 -3.141,-8.041 -6.241,-10.486 3.847,-2.443 7.85,-6.047 7.85,-10.871 0,-5.209 -3.571,-8.49 -6.005,-10.099 -3.415,-2.251 -7.196,-2.702 -8.216,-2.702 h -40.603 a 38.265,38.265 0 0 1 21.408,-12.082 l 4.786,5.021 c 1.082,1.133 2.874,1.175 4.006,0.092 l 5.355,-5.122 a 38.265,38.265 0 0 1 26.196,18.657 l -3.666,8.28 c -0.633,1.433 0.013,3.109 1.442,3.744 z m 9.143,0.134 -0.125,-1.28 3.776,-3.522 c 0.768,-0.716 0.481,-2.157 -0.501,-2.523 l -4.827,-1.805 -0.378,-1.246 3.011,-4.182 c 0.614,-0.85 0.05,-2.207 -0.984,-2.377 l -5.09,-0.828 -0.612,-1.143 2.139,-4.695 c 0.438,-0.956 -0.376,-2.179 -1.428,-2.139 l -5.166,0.18 -0.816,-0.99 1.187,-5.032 c 0.24,-1.022 -0.797,-2.06 -1.819,-1.82 l -5.031,1.186 -0.992,-0.816 0.181,-5.166 c 0.04,-1.046 -1.184,-1.863 -2.138,-1.429 l -4.694,2.14 -1.143,-0.613 -0.83,-5.091 c -0.168,-1.032 -1.526,-1.596 -2.376,-0.984 l -4.185,3.011 -1.244,-0.377 -1.805,-4.828 c -0.366,-0.984 -1.808,-1.267 -2.522,-0.503 l -3.522,3.779 -1.28,-0.125 -2.72,-4.395 c -0.55,-0.89 -2.023,-0.89 -2.571,0 l -2.72,4.395 -1.281,0.125 -3.523,-3.779 c -0.714,-0.764 -2.156,-0.481 -2.522,0.503 l -1.805,4.828 -1.245,0.377 -4.184,-3.011 c -0.85,-0.614 -2.209,-0.048 -2.377,0.984 l -0.83,5.091 -1.143,0.613 -4.694,-2.14 c -0.954,-0.436 -2.178,0.383 -2.138,1.429 l 0.18,5.166 -0.992,0.816 -5.031,-1.186 c -1.022,-0.238 -2.06,0.798 -1.82,1.82 l 1.185,5.032 -0.814,0.99 -5.166,-0.18 c -1.042,-0.03 -1.863,1.183 -1.429,2.139 l 2.14,4.695 -0.613,1.143 -5.09,0.828 c -1.034,0.168 -1.594,1.527 -0.984,2.377 l 3.011,4.182 -0.378,1.246 -4.828,1.805 c -0.98,0.366 -1.267,1.807 -0.501,2.523 l 3.777,3.522 -0.125,1.28 -4.394,2.72 c -0.89,0.55 -0.89,2.023 0,2.571 l 4.394,2.72 0.125,1.28 -3.777,3.523 c -0.766,0.714 -0.479,2.154 0.501,2.522 l 4.828,1.805 0.378,1.246 -3.011,4.183 c -0.612,0.852 -0.049,2.21 0.985,2.376 l 5.089,0.828 0.613,1.145 -2.14,4.693 c -0.436,0.954 0.387,2.181 1.429,2.139 l 5.164,-0.181 0.816,0.992 -1.185,5.033 c -0.24,1.02 0.798,2.056 1.82,1.816 l 5.031,-1.185 0.992,0.814 -0.18,5.167 c -0.04,1.046 1.184,1.864 2.138,1.428 l 4.694,-2.139 1.143,0.613 0.83,5.088 c 0.168,1.036 1.527,1.596 2.377,0.986 l 4.182,-3.013 1.246,0.379 1.805,4.826 c 0.366,0.98 1.808,1.269 2.522,0.501 l 3.523,-3.777 1.281,0.128 2.72,4.394 c 0.548,0.886 2.021,0.888 2.571,0 l 2.72,-4.394 1.28,-0.128 3.522,3.777 c 0.714,0.768 2.156,0.479 2.522,-0.501 l 1.805,-4.826 1.246,-0.379 4.183,3.013 c 0.85,0.61 2.208,0.048 2.376,-0.986 l 0.83,-5.088 1.143,-0.613 4.694,2.139 c 0.954,0.436 2.176,-0.38 2.138,-1.428 l -0.18,-5.167 0.991,-0.814 5.031,1.185 c 1.022,0.24 2.059,-0.796 1.819,-1.816 l -1.185,-5.033 0.814,-0.992 5.166,0.181 c 1.042,0.042 1.866,-1.185 1.428,-2.139 l -2.139,-4.693 0.612,-1.145 5.09,-0.828 c 1.036,-0.166 1.598,-1.524 0.984,-2.376 l -3.011,-4.183 0.378,-1.246 4.827,-1.805 c 0.982,-0.368 1.269,-1.808 0.501,-2.522 l -3.776,-3.523 0.125,-1.28 4.394,-2.72 c 0.89,-0.548 0.891,-2.021 10e-4,-2.571 z" /></svg>
											Parity<small>Third party Rust client from Parity Technologies</small>
										</h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>Parity is a fast, light and secure Ethereum client, supporting both headless mode of operation as well as a web user interface for direct manual interaction. The client is currently a full node with transaction processing based synchronization and state pruning enabled.</p>
										<br/>
										<p>To run a Parity node, download <a href="/{{.ParityGenesis}}"><code>{{.ParityGenesis}}</code></a> and start the node with:
											<pre>parity --chain={{.ParityGenesis}}</pre>
										</p>
										<br/>
										<p>You can find Parity at <a href="https://parity.io/" target="about:blank">https://parity.io/</a>.</p>
									</div>
								</div>
							</div>
							<div class="col-md-6">
								<div class="x_panel">
									<div class="x_title">
										<h2>
											<svg height="14px" version="1.1" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64"><defs><linearGradient id="a" x1="13.79" y1="38.21" x2="75.87" y2="-15.2" gradientTransform="matrix(0.56, 0, 0, -0.57, -8.96, 23.53)" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#5c9fd3"/><stop offset="1" stop-color="#316a99"/></linearGradient><linearGradient id="b" x1="99.87" y1="-47.53" x2="77.7" y2="-16.16" gradientTransform="matrix(0.56, 0, 0, -0.57, -8.96, 23.53)" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#ffd43d"/><stop offset="1" stop-color="#fee875"/></linearGradient></defs><g><path d="M31.62,0a43.6,43.6,0,0,0-7.3.62c-6.46,1.14-7.63,3.53-7.63,7.94v5.82H32v1.94H11a9.53,9.53,0,0,0-9.54,7.74,28.54,28.54,0,0,0,0,15.52c1.09,4.52,3.68,7.74,8.11,7.74h5.25v-7a9.7,9.7,0,0,1,9.54-9.48H39.58a7.69,7.69,0,0,0,7.63-7.76V8.56c0-4.14-3.49-7.25-7.63-7.94A47.62,47.62,0,0,0,31.62,0ZM23.37,4.68A2.91,2.91,0,1,1,20.5,7.6,2.9,2.9,0,0,1,23.37,4.68Z" transform="translate(-0.35)" fill="url(#a)"/><path d="M49.12,16.32V23.1a9.79,9.79,0,0,1-9.54,9.68H24.33a7.79,7.79,0,0,0-7.63,7.76V55.08c0,4.14,3.6,6.57,7.63,7.76a25.55,25.55,0,0,0,15.25,0c3.84-1.11,7.63-3.35,7.63-7.76V49.26H32V47.32H54.85c4.44,0,6.09-3.1,7.63-7.74s1.53-9.38,0-15.52c-1.1-4.42-3.19-7.74-7.63-7.74H49.12ZM40.54,53.14A2.91,2.91,0,1,1,37.67,56,2.88,2.88,0,0,1,40.54,53.14Z" transform="translate(-0.35)" fill="url(#b)"/></g></svg>
											PyEthApp<small>Official Python client from the Ethereum Foundation</small>
										</h2>
										<div class="clearfix"></div>
									</div>
									<div class="x_content">
										<p>Pyethapp is the Ethereum Foundation's research client, aiming to provide an easily hackable and extendable codebase. The client is currently a full node with transaction processing based synchronization and state pruning enabled.</p>
										<br/>
										<p>To run a pyethapp node, download <a href="/{{.PythonGenesis}}"><code>{{.PythonGenesis}}</code></a> and start the node with:
											<pre>mkdir -p $HOME/.config/pyethapp/{{.Network}}</pre>
											<pre>pyethapp -c eth.genesis="$(cat {{.PythonGenesis}})" -c eth.network_id={{.NetworkID}} -c data_dir=$HOME/.config/pyethapp/{{.Network}} -c discovery.bootstrap_nodes="[{{.PythonBootnodes}}]" -c eth.block.HOMESTEAD_FORK_BLKNUM={{.Homestead}} -c eth.block.ANTI_DOS_FORK_BLKNUM={{.Tangerine}} -c eth.block.SPURIOUS_DRAGON_FORK_BLKNUM={{.Spurious}} -c eth.block.METROPOLIS_FORK_BLKNUM={{.Byzantium}} -c eth.block.DAO_FORK_BLKNUM=18446744073709551615 run --console</pre>
										</p>
										<br/>
										<p>You can find pyethapp at <a href="https://github.com/ethereum/pyethapp/" target="about:blank">https://github.com/ethereum/pyethapp/</a>.</p>
									</div>
								</div>
							</div>
						</div>
					</div>{{end}}
					<div id="about" hidden>
						<div class="row vertical-center">
							<div style="margin: 0 auto;">
								<div class="x_panel">
									<div class="x_title">
										<h3>Puppeth &ndash; Your Ethereum private network manager</h3>
										<div class="clearfix"></div>
									</div>
									<div style="display: inline-block; vertical-align: bottom; width: 623px; margin-top: 16px;">
										<p>Puppeth is a tool to aid you in creating a new Ethereum network down to the genesis block, bootnodes, signers, ethstats server, crypto faucet, wallet browsers, block explorer, dashboard and more; without the hassle that it would normally entail to manually configure all these services one by one.</p>
										<p>Puppeth uses ssh to dial in to remote servers, and builds its network components out of docker containers using docker-compose. The user is guided through the process via a command line wizard that does the heavy lifting and topology configuration automatically behind the scenes.</p>
										<br/>
										<p>Puppeth is distributed as part of the <a href="https://geth.ethereum.org/downloads/" target="about:blank">Geth &amp; Tools</a> bundles, but can also be installed separately via:<pre>go get github.com/ethereum/go-ethereum/cmd/puppeth</pre></p>
										<br/>
										<p><em>Copyright 2017. The go-ethereum Authors.</em></p>
									</div>
									<div style="display: inline-block; vertical-align: bottom; width: 217px;">
										<img src="puppeth.png" style="height: 256px; margin: 16px 16px 16px 16px"></img>
									</div>
								</div>
							</div>
						</div>
					</div>
					<div id="frame-wrapper" hidden style="position: absolute; height: 100%;">
						<iframe id="frame" style="position: absolute; width: 1920px; height: 100%; border: none;" onload="if ($(this).attr('src') != '') { resize(); $('#frame-wrapper').fadeIn(300); }"></iframe>
					</div>
				</div>
			</div>
		</div>

		<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.0/jquery.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/gentelella/1.3.0/js/custom.min.js"></script>
		<script>
			var load = function(hash) {
				window.location.hash = hash;

				// Fade out all possible pages (yes, ugly, no, don't care)
				$("#geth").fadeOut(300)
				$("#mist").fadeOut(300)
				$("#mobile").fadeOut(300)
				$("#other").fadeOut(300)
				$("#about").fadeOut(300)
				$("#frame-wrapper").fadeOut(300);

				// Depending on the hash, resolve it into a local or remote URL
				var url = hash;
				switch (hash) {
					case "#stats":
						url = "//{{.EthstatsPage}}";
						break;
					case "#explorer":
						url = "//{{.ExplorerPage}}";
						break;
					case "#wallet":
						url = "//{{.WalletPage}}";
						break;
					case "#faucet":
						url = "//{{.FaucetPage}}";
						break;
				}
				setTimeout(function() {
					if (url.substring(0, 1) == "#") {
						$('.body').css({overflowY: 'auto'});
						$("#frame").attr("src", "");
						$(url).fadeIn(300);
					} else {
						$('.body').css({overflowY: 'hidden'});
						$("#frame").attr("src", url);
					}
				}, 300);
			}
			var resize = function() {
				var sidebar = $($(".navbar")[0]).width();
				var limit   = document.body.clientWidth - sidebar;
				var scale   = limit / 1920;

				$("#frame-wrapper").width(limit);
				$("#frame-wrapper").height(document.body.clientHeight / scale);
				$("#frame-wrapper").css({
					transform: 'scale(' + (scale) + ')',
					transformOrigin: "0 0"
				});
			};
			$(window).resize(resize);

			if (window.location.hash == "") {
				var item = $(".side-menu").children()[0];
				$(item).children()[0].click();
				$(item).addClass("active");
			} else {
				load(window.location.hash);
				var menu = $(window.location.hash + "_menu");
				if (menu !== undefined) {
					$(menu).addClass("active");
				}
			}
		</script>
	</body>
</html>
`

// dashboardMascot is the png dump of the mascot to display on the dashboard about page.
// nolint: misspell
