import {setTimeout} from "timers/promises";
import { createRequire } from 'module';

const require = createRequire(import.meta.url);
const addon = require('../build/Release/addon');

async function main() {
    for (let i = 0; i < 1_000_000_000; i++) {
        addon.helloLeakJS();
        await setTimeout(100);
    }
    console.log('Done');

    await setTimeout(1000 * 50);
}

main().catch(console.error);