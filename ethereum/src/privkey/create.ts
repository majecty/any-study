import { ethers } from "ethers";

async function main() {
  const wallet = ethers.Wallet.createRandom();
  console.log("private key", wallet.signingKey.privateKey);
}

main().catch(console.error);
