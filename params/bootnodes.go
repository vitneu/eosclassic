// Copyright 2015 The eosclassic & go-ethereum Authors
// This file is part of the eosclassic library.
//
// The eosclassic library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The eosclassic library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the eosclassic library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303", // IE
	"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:30303",  // US-WEST
	"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:30303", // BR
	"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:30303", // AU
	"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:30303",  // SG

	// Ethereum Foundation C++ Bootnodes
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303", // DE
}

// EOSClassicBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the EOSClassic network.
var EOSClassicBootnodes = []string{
	"enode://68141aad001d5984bdd6c935e5f7b89fb55fd6a59585658dcfa6f6ed1b020977107e2baaf0e157bc65afd63e9a6d45f636f6afc3c6208e9c9da185a0ec06f1dd@13.125.38.131:25252",
	"enode://6d5b4fe7e9fc27616b3bdc272ea89c7cbe368432bf299b4597002ad9c3aff65a1bc570c032a2accb03ca3d2955074956758411d89de6ab5feef2d10a8f963a6a@13.124.177.177:25252",
	"enode://b9f636ae1b9bd29b09438de87fc2bca9631bee0f35d145afd0f9fa522366a07d8ca5dd2b1ee89a0f86744b7a28e1cba1566c720238690e596d57f406bf01f0e7@13.209.67.52:25252",
	"enode://1b3f420caacf4fc868befb32199560d150777054bf06e2e557444d1e9bce4b704ca012ef10d89f124f04c2152dd1ccb0ad219ff4f535bd36e18f915692801775@13.125.1.40:25252",
	"enode://4ac82fd3d0ea65e339e4aac17ff7a05ffcb1a3e218523003e07d7217cb6eee4b76a7f7a091c7f6be1dbd3cf3f0b616b7b8feeab0f740dd30b14921db4a84ca88@13.209.41.126:25252",
	"enode://b34facef380fde111908348207320b56a75fece06c6c1aef31b35903ec8b2195861c3833c45b6c91019c9ce35423f8e37b04db6e4419d0eea48fff42ee4ee812@13.209.75.168:25252",
	"enode://14e28c54a0e37dcc53db3f69d1cee2451203f9463ff3e02441f8a3078e249601c21485f71d06422fe99574670fe09d6379a63d0bafa796bcdedb1d4822c064eb@13.209.77.72:25252",
	"enode://747958fcb02c6ed0d2269d96fb02bb43e7e08b2c47545d7396b71b265845a658a274a68409e1c674d9feded17641a4788c61c7feed109e62602ea13d3121c355@54.180.8.50:30303",
	"enode://a3bc90ad80c66fcdc9ee048040444261d740bdafb1aad71020ac8c3b3918819ac75628493527831feab2cddbebbfd782a538f1f889386c9524a9053ff3385ba8@13.125.38.131:30303",
}

// EOSCTestBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the EOSCTest network.
var EOSCTestBootnodes = []string{
	"enode://37a0a55096b3750bd19d5033f1b58d58f48fa5e954b12f8d4c1ca52a41243be532a0ca68881e230b202e3a220761fc43deaae95be26099c6eeedbc0749ea0837@13.209.35.2:25252",
	"enode://61db43578388d0c65b50ffb2eca6c9dc9de5677374aeb254457dde35da3666bcfd0c4118536de83703e9ce5c710ff22e7db8997a5b3df5fcfa348083e47fcfc9@13.209.35.2:25262",
	"enode://35fbcccca93e4a9ca71d2e5f9eada2d4295cf09dbaac89054df184dd8f56af86c04ca03c9a48b371514d5685eb9cf800828495e17b2216819b6861e0c99debd9@13.209.35.2:25272",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://6d5b4fe7e9fc27616b3bdc272ea89c7cbe368432bf299b4597002ad9c3aff65a1bc570c032a2accb03ca3d2955074956758411d89de6ab5feef2d10a8f963a6a@13.124.177.177:25252",
	"enode://b9f636ae1b9bd29b09438de87fc2bca9631bee0f35d145afd0f9fa522366a07d8ca5dd2b1ee89a0f86744b7a28e1cba1566c720238690e596d57f406bf01f0e7@13.209.67.52:25252",
	"enode://1b3f420caacf4fc868befb32199560d150777054bf06e2e557444d1e9bce4b704ca012ef10d89f124f04c2152dd1ccb0ad219ff4f535bd36e18f915692801775@13.125.1.40:25252",
	"enode://4ac82fd3d0ea65e339e4aac17ff7a05ffcb1a3e218523003e07d7217cb6eee4b76a7f7a091c7f6be1dbd3cf3f0b616b7b8feeab0f740dd30b14921db4a84ca88@13.209.41.126:25252",
}
