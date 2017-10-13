use gix;
var bulk = db.gifs.initializeUnorderedBulkOp();
bulk.insert( { Url: "https://i.imgur.com/yRe6x.gif", Tags: "disgusted", ID: 1 });
bulk.insert( { Url: "https://i.imgur.com/M0bWify.gif", Tags: "nailedit", ID: 2 });
bulk.insert( { Url: "https://i.imgur.com/2dPuX.gif", Tags: "nailedit", ID: 3 });
bulk.execute();
