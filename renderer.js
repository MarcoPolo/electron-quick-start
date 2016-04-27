const testConn = require('./testConn')
const net = require('net')

function windowsHack () {
  var fake = net.connect({})
  fake.on('error', function () {})
}

window.windowsHack = windowsHack

testConn()
