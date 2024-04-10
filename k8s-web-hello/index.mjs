import express from 'express'
import os from 'os'

const app = express()
const PORT = 3000

app.get("/", (req, res) => {
  const helloMessage = `<h1>VERSION 3.5: Hello from the ${os.hostname()}</h1>`
  console.log(helloMessage)
  res.send(helloMessage)
})

app.listen(PORT, () => {
  console.log(`Web server is listening at port ${PORT}`)
})

import pkg from 'pg';
const { Pool } = pkg;

const connectionString = process.env.DATABASE_URL;

const pool = new Pool({
  connectionString,
});

app.get("/pg", (req, res) => {
  pool.query(
    "SELECT * FROM users u WHERE u.username = 'nir'", (err, res) => {
    console.log(err, res);
  });
})
