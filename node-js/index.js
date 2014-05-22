#!/usr/bin/env node
/*jslint node: true */
// Run like: `node node-benchmark.js ~/Desktop/twitter.json`
var fs = require('fs');
// prerequisite: `npm install streaming`
var streaming = require('streaming');

// process.argv: ["node", "node-benchmark.js", ...]
var input = fs.createReadStream(process.argv[2]);
var splitter = new streaming.Splitter();
input.pipe(splitter);

var nlines = 0;
var ndeletes = 0;
var nentities = 0;
var colors = {};

var started = Date.now();

var obj;

splitter.on('data', function(chunk) {
  if (chunk) {
    // eval('obj = ' + chunk); // eval is actually a lot slower (221s!)
    obj = JSON.parse(chunk);
    if ('delete' in obj) {
      ndeletes++;
    }
    else {
      var profile_background_color = obj.user.profile_background_color;
      if (profile_background_color !== 'C0DEED') {
        colors[profile_background_color] = (colors[profile_background_color] || 0) + 1;
      }

      // count entities
      nentities += (
        obj.entities.hashtags.length +
        obj.entities.symbols.length +
        obj.entities.urls.length +
        obj.entities.user_mentions.length
      );
    }
  }
  nlines++;
});

splitter.on('end', function(chunk) {
  // for a 1M file of tweets,
  // takes about 13 seconds just to iterate through all the lines and count them
  // wc -l can do it in 3s, wc takes 10s
  // iwc -l can do it in 5s, iwc takes 13s
  var ended = Date.now();

  // for same 2.9G file of 1M tweets, the complete benchmark (to the line below) takes about 44s
  console.log('elapsed: %d seconds', (ended - started) / 1000);

  console.log('nlines = %d', nlines);
  console.log('ndeletes = %d', ndeletes);
  console.log('nentities = %d', nentities);
  // console.log('colors:', JSON.stringify(colors, null, 2));
});
