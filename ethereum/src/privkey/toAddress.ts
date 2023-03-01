import { ethers } from "ethers";

const PRIVATE_KEY =
  "0x9828ff62af78703913729710c69af932393e96feb8e388aaf1c81feccfdbebfa";

async function main() {
  const wallet = new ethers.Wallet(PRIVATE_KEY);
  const address = await wallet.getAddress();
  console.log("address", address);
}

main().catch(console.error);
