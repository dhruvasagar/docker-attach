const express = require('express')
const app = express()
const addr = '0.0.0.0'
const port = 1234

app.get("/", (req, res) => {
  debugger
  res.send('Dummy server, does nothing!')
})

app.listen(port, addr, () => {
  console.log(`Example app listening at http://0.0.0.0:${port}`)
})
