import {
    Keypair, Transaction, SystemProgram,
    LAMPORTS_PER_SOL,
    sendAndConfirmTransaction,
    clusterApiUrl,
    Connection
} from "@solana/web3.js";
import { loadKey } from "./keys/loadKey";

const SOLANA_CLUSTER = "devnet";

async function main() {
    const fromKeypair = loadKey();
    const toKeypair = Keypair.generate();
    const transaction = new Transaction();

    transaction.add(
        SystemProgram.transfer({
            fromPubkey: fromKeypair.publicKey,
            toPubkey: toKeypair.publicKey,
            lamports: 100
        })
    );

    const connection = new Connection(clusterApiUrl(SOLANA_CLUSTER));

    console.log("Sending transaction...");
    const txid = await sendAndConfirmTransaction(connection, transaction, [fromKeypair]);
    console.log(
        "Tansaction submitted:",
        `https://explorer.solana.com/tx/${txid}?cluster=${SOLANA_CLUSTER}`,
    )

}

main().catch(console.error);