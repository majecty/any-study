import * as solanaWeb3 from "@solana/web3.js";

function main() {
  const publicKey =
    "0x5a2d93d35f692d7ca74996c11265efd5493201049da15611066bbcca8f1d753a";
  const buffer = Buffer.from(publicKey.slice(2), "hex");
  const pub = solanaWeb3.PublicKey.decode(buffer);
  console.log(pub);
}

main();
