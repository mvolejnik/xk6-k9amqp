import http from 'k6/http';
import { sleep } from 'k6';
import k9amqp from 'k6/x/k9amqp';

export const options = {
  vus: 10,
  duration: '30s',
}

export function setup() {
  console.log("setup");
  k9amqp.initPool();
  k9amqp.connect("localhost", 5671, "/", "guest", "guest");
}

export default function() {
  console.log(k9amqp.test());
  sleep(1);
}
