import { PTTServer } from './src'

async function run() {
  const s = new PTTServer({
    socket: "ptt.sock"
  })

  await s.start()

  const r = await s.parse({ 
    torrent_titles:["One Piece S01E02"],
  })

  console.log('r', r)
}

run()
