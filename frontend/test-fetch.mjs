// test-fetch.mjs
async function test() {
  try {
    const res = await fetch('http://localhost:8080/api/v1/makes');
    console.log('Status:', res.status);
    console.log('OK?', res.ok);
    const data = await res.json();
    console.log('Data length:', Array.isArray(data) ? data.length : 'not array');
    console.log('Sample:', JSON.stringify(data.slice?.(0, 2) || data, null, 2));
  } catch (err) {
    console.error('Fetch error:', err.message);
    if (err.cause) console.error('Cause:', err.cause);
  }
}

test();