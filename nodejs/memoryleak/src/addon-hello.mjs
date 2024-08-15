/*
node-gyp configure
node-gyp build
*/

import { createRequire } from 'module';
const require = createRequire(import.meta.url);
const addon = require('../build/Release/addon');
console.log(addon.hello());