import {sleep} from 'k6'
import tcp from 'k6/x/tcp'

const opts = {
    host: '127.0.0.1',
    port: 8000
};
const client = new tcp.Client();
client.on('error', (err) => {
    console.log('-------this is error-----------');
    console.log(err);
})
client.on('data', (data) => {
    console.log('-------this is data-----------');
    console.log(data);
})
client.connect(opts);

export default function () {
    client.write('/time\n');
    sleep(1);
}