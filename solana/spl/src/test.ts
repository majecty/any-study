// const web3 = require('@solana/web3.js');
import * as web3 from "@solana/web3.js"
console.log(web3);
const connection = new web3.Connection(web3.clusterApiUrl('devnet'), 'confirmed');
console.log(connection);