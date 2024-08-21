import { useCallback, useEffect, useState } from 'preact/hooks'
import styles from './app.module.css'
import { CalendarEntry, Country } from './types'
import cx from 'clsx'

export function App() {
  const [dates, setDates] = useState<CalendarEntry[]>([])
  const [countries, setCountries] = useState<Country[]>([])
  const [selectedYear, setSelectedYear] = useState(new Date().getUTCFullYear())
  const [selectedCountry, setSelectedCountry] = useState('')
  const [selectedThreshold, setSelectedThreshold] = useState(3)

  const onUpdate = useCallback(() => {
    fetch(
      `/calendar/${selectedCountry}?threshold=${selectedThreshold}&year=${selectedYear}`
    )
      .then((v) => v.json())
      .then((entries: CalendarEntry[]) => {
        setDates(entries)
      })
  }, [selectedCountry, selectedThreshold, selectedYear])

  useEffect(() => {
    fetch('/countries')
      .then((v) => v.json())
      .then((countries) => {
        setCountries(countries)
        setSelectedCountry(countries[0].countryCode)
      })
  }, [])

  return (
    <main className={styles.main}>
      <header className={styles.header}>
        <h1>Squeeze Days</h1>
      </header>
      <button onClick={onUpdate}>Update</button>
      <select
        onChange={(e) => {
          setSelectedCountry(e.currentTarget.value)
        }}
      >
        {countries.map((c) => (
          <option value={c.countryCode}>{c.name}</option>
        ))}
      </select>
      <input
        list="markers"
        onChange={(e) => {
          setSelectedThreshold(Number.parseInt(e.currentTarget.value))
        }}
        type="range"
        min="1"
        max="4"
        value={selectedThreshold}
      />
      <datalist id="markers">
        <option value="1"></option>
        <option value="2"></option>
        <option value="3"></option>
        <option value="4"></option>
      </datalist>
      <input
        type="number"
        value={selectedYear}
        onChange={(e) => {
          const parsedYear = Number.parseInt(e.currentTarget.value, 10)
          if (!Number.isNaN(parsedYear)) {
            setSelectedYear(parsedYear)
          }
        }}
      />
      <ul className={styles.grid}>
        {dates.map((d) => (
          <li
            className={cx(styles.item, {
              [styles.weekend]: d.isWeekend,
              [styles.holiday]: d.isHoliday,
              [styles.squeeze]: d.isSqueeze,
            })}
          >
            <p>{d.timestamp}</p>
            <p>{d.holidayName}</p>
          </li>
        ))}
      </ul>
    </main>
  )
}
