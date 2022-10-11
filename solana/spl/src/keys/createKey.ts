import * as web3 from "@solana/web3.js";

async function main() {
    const keypair = web3.Keypair.generate();
    console.log(keypair);
    console.log("publickey", keypair.publicKey.toBase58());
    console.log("secretkey", keypair.secretKey.toString());
}

main().catch(console.error);