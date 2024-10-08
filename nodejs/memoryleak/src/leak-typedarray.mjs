import { setTimeout } from 'timers/promises';
import { memoryUsage } from 'node:process';

async function main() {
    const arr = [];
    for (let i = 0; i < 1_000_000_000; i++) {
        arr.push(new Uint8Array(1_000_000));
        await setTimeout(100);
        console.log(i);

        console.dir(memoryUsage());
    }
    console.log('Done');

    await setTimeout(1000 * 50);
}

main().catch(console.error);