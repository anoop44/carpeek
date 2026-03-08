import { fetch } from 'node-fetch';

async function testBackend() {
  try {
    const res = await fetch('http://localhost:8080/api/v1/makes', {
      method: 'GET',
      headers: { 'Accept': 'application/json' }
    });
    console.log('Status:', res.status);
    console.log('Headers:', Object.fromEntries(res.headers));
    const data = await res.json();
    console.log('Data:', JSON.stringify(data, null, 2));
  } catch (err) {
    console.error('Fetch failed:', err.message);
    if (err.cause) console.error('Cause:', err.cause);
  }
}

testBackend();