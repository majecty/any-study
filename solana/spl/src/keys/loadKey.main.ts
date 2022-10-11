import { loadKey } from "./loadKey";

async function main() {
    const keypair = loadKey();
    console.log(keypair);
    console.log("publickey", keypair.publicKey.toBase58());
    console.log("secretkey", keypair.secretKey.toString());
}

main().catch(console.error);