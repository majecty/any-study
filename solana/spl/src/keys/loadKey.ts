import * as web3 from "@solana/web3.js";

export function loadKey(): web3.Keypair {
    const secretKey = Uint8Array.from([
        122, 37, 226, 139, 21, 247, 99, 74, 182, 150, 203, 128, 89, 205, 56, 60, 63, 51, 184, 206, 19, 203, 134, 27, 210, 187, 140, 27, 220, 98, 61, 105, 62, 62, 49, 39, 79, 233, 155, 188, 12, 5, 175, 28, 166, 236, 82, 144, 162, 21, 115, 94, 80, 104, 140, 220, 102, 44, 36, 184, 14, 185, 193, 26
    ]);
    const keypair = web3.Keypair.fromSecretKey(secretKey);
    return keypair;
}