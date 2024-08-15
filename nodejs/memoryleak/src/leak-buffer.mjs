import {setTimeout} from "timers/promises";

async function main() {
    const arr = [];
    for (let i = 0; i < 1_000_000_000; i++) {
        arr.push(Buffer.alloc(1_000_000));
        await setTimeout(0);
    }
    console.log('Done');

    await setTimeout(1000 * 50);
}

main().catch(console.error);