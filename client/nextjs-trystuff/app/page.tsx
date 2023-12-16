import Image from 'next/image'
import Page from './play-code/palendrome'
import Trie from './play-code/trie'
import styles from './page.module.css'

export default function Home() {
  return (
    
    <main className={styles.main}>
      <div className={styles.top}>
          <Page />
          <Trie />
      </div>
      <div className={styles.left}>
        <Image
          src="/EnnekingSS_Logo_Full.png"
          alt="ESS Logo"
          width={300}
          height={136}
        />
      </div>


    </main>
  )
}
