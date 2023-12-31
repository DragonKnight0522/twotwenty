// import "@nomiclabs/hardhat-waffle";
// import "@nomiclabs/hardhat-etherscan";
import "@nomicfoundation/hardhat-ethers";
import "@openzeppelin/hardhat-upgrades";
import "@nomicfoundation/hardhat-toolbox";
import "hardhat-gas-reporter";
import "solidity-coverage";


import * as dotenv from "dotenv";
dotenv.config();

import "./scripts/deploy";

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */

module.exports = {
  allowUnlimitedContractSize: true,
  solidity: {
    compilers: [
      {
        version: "0.8.4",
        settings: {
          optimizer: {
            enabled: true,
            runs: 1000,
          },          
        },        
      },
    ],
  },  
  gasReporter: {
    currency: "USD", // can be set to ETH and other currencies (see coinmarketcap api documentation)
    coinmarketcap: process.env.coinMarketCap_API,
  },  
  networks: {
    hardhat: {
      // allowUnlimitedContractSize: true,
    },
    goerli: {
      url: `https://goerli.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
    ropsten: {
      url: `https://ropsten.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
    rinkeby: {
      url: `https://rinkeby.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
    bsctest: {
      url: "https://data-seed-prebsc-1-s1.binance.org:8545",
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
    kovan: {
      url: `https://kovan.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
    fantom: {
      url: "https://rpc.ftm.tools/",
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
    fantom_test: {
      url: "https://rpc.testnet.fantom.network/",
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
    mumbai: {
      url: process.env.MUMBAI_URL || "",
      accounts: [`0x${process.env.PRIVATE_KEY}`],      
    },    
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
};
