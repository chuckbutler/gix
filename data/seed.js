use gix;
var bulk = db.gifs.initializeUnorderedBulkOp();
bulk.insert( { url: "https://i.imgur.com/yRe6x.gif", tags: "disgusted" });
bulk.insert( { url: "https://i.imgur.com/M0bWify.gif", tags: "nailedit" });
bulk.insert( { url: "https://i.imgur.com/2dPuX.gif", tags: "nailedit" });
bulk.execute();
