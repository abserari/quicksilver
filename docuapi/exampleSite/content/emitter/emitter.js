var client = require('emitter-io').connect({host:"127.0.0.1", port:"8080"}); 
// use  on NodeJS 
const emitterKey = 'DdYjd8w_zH4UTj3OLWOqM8kSmbk9c68H';
const channel = 'personinfo';

// once we're connected, subscribe to the 'chat' channel
client.subscribe({
	key: emitterKey,
	channel: channel
});
    
// on every message, print it out
client.on('message', function(msg){
	console.log( msg.asString());
});

// publish a message to the chat channel
client.publish({
    key: emitterKey,
    channel: channel,
    message: "hello world"
})