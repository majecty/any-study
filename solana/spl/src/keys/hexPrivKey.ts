import * as crypto from "crypto";
import * as web3 from "@solana/web3.js";

function main() {
  // generate random 32 byte array
  const randomBytes = crypto.randomBytes(32);
  const hexBytes = randomBytes.toString("hex");
  console.log(hexBytes, hexBytes.length);

  // convert hexBytes to solana privae key
  const buffer = Buffer.from(hexBytes, "hex");
  const uint8Array = new Uint8Array(buffer);
  console.log(uint8Array.length);
  const privateKey = web3.Keypair.fromSeed(uint8Array);
  console.log(privateKey.secretKey.toString(), privateKey.secretKey.length);
}

main();
