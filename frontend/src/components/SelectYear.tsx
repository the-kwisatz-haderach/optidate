import { StepperContentRenderer } from './Stepper/types'
import { useAppContext } from '../context/AppProvider'
import RadioButton from './RadioButton/RadioButton'
import styles from './SelectYear.module.css'

export const SelectYearStep: StepperContentRenderer = ({
  onSetComplete,
  title,
}) => {
  const { state, updateState } = useAppContext()
  const currentYear = new Date().getUTCFullYear()

  const setYear = (year: number) => {
    updateState({ year })
    onSetComplete()
  }

  return (
    <div>
      <h2>{title}</h2>
      <div className={styles.inputContainer}>
        <RadioButton
          value={currentYear}
          selected={state.year === currentYear}
          onChange={setYear}
        />
        <RadioButton
          value={currentYear + 1}
          selected={state.year === currentYear + 1}
          onChange={setYear}
        />
      </div>
    </div>
  )
}
