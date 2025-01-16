import { useEffect, useState } from 'preact/hooks'
import { StepperContentRenderer } from './Stepper/types'
import { Country } from '../types'
import { useAppContext } from '../context/AppProvider'
import styles from './SelectCountry.module.css'
import cx from 'clsx'

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

  const filteredCountries = countries.filter((country) =>
    country.name.toLowerCase().includes(query.toLowerCase())
  )

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
              if (
                filteredCountries.length === 1 &&
                filteredCountries[0].name.toLowerCase() ===
                  e.currentTarget.value.toLowerCase()
              ) {
                updateState({ countryCode: filteredCountries[0].countryCode })
                onSetComplete(true)
              } else if (state.countryCode) {
                onSetComplete(false)
                updateState({ countryCode: '' })
              }
            }}
          />
        </div>
      </div>
      <ul className={styles.countryList}>
        {filteredCountries.map((country) => (
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
                className={cx(styles.countryLabel, {
                  [styles.selectedCountry]:
                    country.countryCode === state.countryCode,
                })}
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
