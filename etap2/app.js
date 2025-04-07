const http = require('http');
const os = require('os');

const hostname = os.hostname();
const ip = os.networkInterfaces()['eth0']?.[0]?.address || 'IP not found';
const version = os.release();

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/html');
  res.end(`
    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>System Information</title>
      </head>
      <body>
        <p><strong>Hostname:</strong> ${hostname}</p>
        <p><strong>IP Address:</strong> ${ip}</p>
        <p><strong>System Version:</strong> ${version}</p>
      </body>
    </html>
  `);
});

const port = 3000;
server.listen(port, () => {
  console.log(`Server running at http://localhost:${port}/`);
});
