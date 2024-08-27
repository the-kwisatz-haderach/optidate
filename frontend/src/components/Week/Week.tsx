import { PropsWithChildren } from 'preact/compat'
import styles from './Week.module.css'

type Props = {
  weekNumber: number
}

export default function Week({
  weekNumber,
  children,
}: PropsWithChildren<Props>) {
  return (
    <div className={styles.wrapper}>
      <span>Week: {weekNumber}</span>
      <div className={styles.container}>{children}</div>
    </div>
  )
}
