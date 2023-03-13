import * as fs from 'fs/promises';

let version = await fs.readFile('version.txt', { encoding: 'utf8'});
version = version.trim();

console.log(`Hello, world from ${version}!`);
