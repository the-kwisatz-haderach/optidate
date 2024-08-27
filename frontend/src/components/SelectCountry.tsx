import { useEffect, useState } from 'preact/hooks'
import { StepperContentRenderer } from './Stepper/types'
import { Country } from '../types'
import { useAppContext } from '../context/AppProvider'
import styles from './SelectCountry.module.css'

export const SelectCountryStep: StepperContentRenderer = ({
  onSetComplete,
  active,
}) => {
  const { state, updateState } = useAppContext()
  const [query, setQuery] = useState('')
  const [countries, setCountries] = useState<Country[]>([])

  useEffect(() => {
    if (active) {
      fetch('/countries')
        .then((v) => v.json())
        .then((c: Country[]) =>
          setCountries(c.sort((a, b) => a.name.localeCompare(b.name)))
        )
    }
  }, [active])

  return (
    <div className={styles.wrapper}>
      <div>
        <h2>Select country</h2>
        <div className={styles.inputWrapper}>
          <input
            className={styles.input}
            value={query}
            onChange={(e) => {
              setQuery(e.currentTarget.value)
              if (state.countryCode) {
                onSetComplete(false)
                updateState({ countryCode: '' })
              }
            }}
          />
        </div>
      </div>
      <ul className={styles.countryList}>
        {countries
          .filter((country) =>
            country.name.toLowerCase().includes(query.toLowerCase())
          )
          .map((country) => (
            <li key={country.countryCode}>
              <div className={styles.countryInputContainer}>
                <input
                  id={country.countryCode}
                  className={styles.countryInput}
                  type="radio"
                  checked={state.countryCode === country.countryCode}
                  onChange={() => {
                    setQuery(country.name)
                    updateState({ countryCode: country.countryCode })
                    onSetComplete(true)
                  }}
                />
                <label
                  htmlFor={country.countryCode}
                  className={styles.countryLabel}
                >
                  <img
                    className={styles.flag}
                    src={`https://flagsapi.com/${country.countryCode}/flat/24.png`}
                  />
                  {country.name}
                </label>
              </div>
            </li>
          ))}
      </ul>
    </div>
  )
}
