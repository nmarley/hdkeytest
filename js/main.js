const bitcore = require('bitcore-lib-cash');
const { HDPrivateKey } = bitcore;

const serialized = 'xprv9s21ZrQH143K25nXfd4XuZ6d955Y7XQ7GzSrew7bEYCx1eC6JixjXoaULDjgpB4BJaLL4BmCTka4L9Jq5rxstLcXLnVkGK4tkMk8D9iv1y8';

const rootHDKey = new HDPrivateKey(serialized);

console.log("serialized xprv: " + serialized);
console.log("xpub key: " + rootHDKey.xpubkey);
