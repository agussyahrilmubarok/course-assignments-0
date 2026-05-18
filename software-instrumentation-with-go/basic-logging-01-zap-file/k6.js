import http from "k6/http";
import { check, sleep } from "k6";
import { randomItem } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
  vus: 10,
  iterations: 1000,
  thresholds: {
    http_req_failed: ["rate<0.1"],
    http_req_duration: ["p(95)<500"],
  },
};

const BASE_URL = __ENV.BASE_URL || "http://localhost:8080";
const validNames = ['John', 'Jane', 'Bert'];
const shortNames = ["A", "Ed", "Jo"];
const emptyNames = ["", "", ""];

function randomName() {
  const type = randomItem(["valid", "short", "empty"]);
  if (type === "valid") {
    return randomItem(validNames);
  }

  if (type === "short") {
    return randomItem(shortNames);
  }

  return randomItem(emptyNames);
}

export default function () {
  const name = randomName();
  const url = `${BASE_URL}/?name=${name}`;
  const res = http.get(url);

  check(res, {
    "status ok": (r) => r.status === 200 || r.status === 400,
  });

  sleep(0.01);
}
