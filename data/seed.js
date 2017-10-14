use gix;
var bulk = db.gifs.initializeUnorderedBulkOp();
 bulk.insert({"url":"https://i.imgur.com/yRe6x.gif","tags":"disgusted"});
 bulk.insert({"url":"https://i.imgur.com/M0bWify.gif","tags":"nailedit"});
 bulk.insert({"url":"https://i.imgur.com/2dPuX.gif","tags":"nailedit"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/fngrgnz.gif","tags":"yisss"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/kpo.gif","tags":"peace"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/asm.gif","tags":"awesome"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/sbbn.gif","tags":"sad"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/jack-tripper.gif","tags":"dumb"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/omgshoop.gif","tags":"amazing"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/wwhet.gif","tags":"what"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/mkms.gif","tags":"cha-ching"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/whid1.gif","tags":"ohcrap"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/aawt.gif","tags":"waiting"});
 bulk.insert({"url":"http://www.reactiongifs.com/r/htp.gif","tags":"hack"});
bulk.execute();



