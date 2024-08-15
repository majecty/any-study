import { setTimeout } from 'timers/promises';
import { createRequire } from 'module';
import { memoryUsage } from 'node:process';
import os from 'node:os';

const require = createRequire(import.meta.url);
const addon = require('../build/Release/addon');

async function main() {
    for (let i = 0; i < 1_000_000_000; i++) {
        addon.helloLeakMalloc();
        await setTimeout(100);
        console.log(i);

        console.dir(memoryUsage());
    }
    console.log('Done');

    await setTimeout(1000 * 50);
}

main().catch(console.error);