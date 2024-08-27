import { useEffect, useMemo, useState } from 'preact/hooks'
import { CalendarEntry } from '../types'
import { useAppContext } from '../context/AppProvider'
import { chunk } from '../utils/chunk'
import DateEntry from './DateEntry/DateEntry'
import styles from './Calendar.module.css'
import { StepperContentRenderer } from './Stepper/types'
import cn from 'clsx'
import { Fragment } from 'preact/jsx-runtime'

export const Calendar: StepperContentRenderer = ({ active }) => {
  const {
    state: { countryCode, year, threshold },
  } = useAppContext()
  const [weeks, setWeeks] = useState<CalendarEntry[][]>([])

  useEffect(() => {
    if (active && year && countryCode) {
      fetch(`/calendar/${countryCode}?threshold=4&year=${year}`)
        .then((v) => v.json())
        .then((dates: CalendarEntry[]) => {
          setWeeks(chunk(dates, 7))
        })
    }
  }, [year, countryCode, active])

  const filteredWeeks = useMemo(
    () =>
      weeks.filter((w) =>
        w.some((d) => !d.isHoliday && !d.isWeekend && d.gapSize <= threshold)
      ),
    [weeks, threshold]
  )

  const bridgeDays = useMemo(
    () =>
      filteredWeeks.reduce(
        (count, week) =>
          count +
          week.filter(
            (date) =>
              !date.isHoliday && !date.isWeekend && date.gapSize <= threshold
          ).length,
        0
      ),
    [filteredWeeks, threshold]
  )

  return (
    <>
      {filteredWeeks.length > 0 && (
        <div className={styles.container}>
          <div className={cn(styles.column, styles.mainColumn)}>
            <span className={cn(styles.columnHeader, styles.weekDay)}>
              Week
            </span>
            {filteredWeeks[0]?.map((d) => (
              <span key={d.timestamp} className={styles.weekDay}>
                {new Date(d.timestamp).toLocaleString(navigator.language, {
                  weekday: 'short',
                })}
              </span>
            ))}
          </div>
          <div className={styles.weeks}>
            {filteredWeeks.map((w) => (
              <div key={w[0].week} className={cn(styles.column, styles.snap)}>
                {w.map((date, i) => (
                  <Fragment key={date.timestamp}>
                    {i === 0 && (
                      <span
                        className={cn(styles.columnHeader, styles.entryWrapper)}
                      >
                        {date.week}
                      </span>
                    )}
                    <span
                      className={cn(styles.entryWrapper, {
                        [styles.holiday]: date.isHoliday,
                        [styles.squeeze]:
                          !date.isHoliday &&
                          !date.isWeekend &&
                          date.gapSize <= threshold,
                      })}
                    >
                      <DateEntry {...date} />
                    </span>
                  </Fragment>
                ))}
              </div>
            ))}
          </div>
        </div>
      )}
      <div className={styles.summary}>
        <div>
          <span>{year}</span> contains{' '}
          <span className={cn(styles.bridge, styles.squeeze)}>
            {bridgeDays}
          </span>{' '}
          bridge days with a gap of <span>{threshold}</span> or less spread over{' '}
          <span>{filteredWeeks.length}</span> weeks.
        </div>
      </div>
    </>
  )
}
