import http from 'k6/http';
import { sleep } from 'k6';
import { k9amqp } from 'k6/x/k9amqp';

export const options = {
  vus: 10,
  duration: '30s',
}

export default function() {
  console.log(k9amqp);
  //console.log(k9amqp.test());
  sleep(1);
}
