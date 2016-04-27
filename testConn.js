const net = require('net')

function testConn () {
  console.log('starting connection')
  const p = '\\\\.\\pipe\\kbservice\\Users\\IEUser\\AppData\\Roaming\\Keybase\\keybased.sock'

  const s = net.connect({path: p}, () => { console.log('connected') })

  s.on('end', () => {
    console.log('disconnected from server')
  }).on('error', (e) => {
    console.log('error from server', e)
  }).on('lookup', (l) => {
    console.log('lookup: ', l)
  })

  s.setTimeout(5e3, () => console.log('hit timeout'))
}

module.exports = testConn
