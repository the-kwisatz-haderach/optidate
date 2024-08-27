import cx from 'clsx'
import { CalendarEntry } from '../../types'
import styles from './DateEntry.module.css'

export default function DateEntry({ timestamp, holidayName }: CalendarEntry) {
  const parsed = new Date(timestamp)
  return (
    <>
      <div className={cx(styles.date)}>{parsed.getUTCDate()}</div>
      <div className={cx(styles.dateName)}>
        {parsed.toLocaleString(navigator.language, { month: 'short' })}
      </div>
      {holidayName && (
        <div className={cx(styles.holidayName)}>{holidayName}</div>
      )}
    </>
  )
}
