import * as web3 from "@solana/web3.js";

const SOLANA_CLUSTER = "devnet";
// program id from https://docs.solana.com/getstarted/hello-world
const PROGRAM_ID = "6PPuTKJHypp5vuJMD2owSAomzy5oi39vgST45s1HLJY8";

async function main() {
    const connection = new web3.Connection(web3.clusterApiUrl(SOLANA_CLUSTER));

    let payer = web3.Keypair.generate();
    console.log("Generated payer address:", payer.publicKey.toBase58());

    console.log("Requesting airdrop...");
    let airdropSignature = await connection.requestAirdrop(
        payer.publicKey,
        web3.LAMPORTS_PER_SOL,
    );

    await connection.confirmTransaction(airdropSignature);

    console.log(
        "Airdrop submitted:",
        `https://explorer.solana.com/tx/${airdropSignature}?cluster=${SOLANA_CLUSTER}`,
    );

    const transaction = new web3.Transaction();

    // add a single instruction to the transaction
    transaction.add(
        new web3.TransactionInstruction({
            keys: [
                {
                    pubkey: payer.publicKey,
                    isSigner: true,
                    isWritable: false,
                },
                {
                    pubkey: web3.SystemProgram.programId,
                    isSigner: false,
                    isWritable: false,
                },
            ],
            programId: new web3.PublicKey(PROGRAM_ID),
        }),
    );

    console.log("Sending transaction...");
    let txid = await web3.sendAndConfirmTransaction(connection, transaction, [
        payer,
    ]);
    console.log(
        "Transaction submitted:",
        `https://explorer.solana.com/tx/${txid}?cluster=${SOLANA_CLUSTER}`,
    );
}

main().catch(console.error);
