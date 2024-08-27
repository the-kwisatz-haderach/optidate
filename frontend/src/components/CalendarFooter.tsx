import { useAppContext } from '../context/AppProvider'
import { StepperControlRenderer } from './Stepper/types'
import styles from './CalendarFooter.module.css'

const CalendarFooter: StepperControlRenderer = ({ className }) => {
  const thresholds = [1, 2, 3]
  const { updateState, state } = useAppContext()
  return (
    <div className={className}>
      <span>Set sensitivity</span>
      <div className={styles.inputsContainer}>
        {thresholds.map((th) => (
          <div key={th} className={styles.inputWrapper}>
            <input
              id={th.toString()}
              type="radio"
              value={th}
              checked={state.threshold === th}
              onChange={() => {
                updateState({ threshold: th })
              }}
            />
            <label htmlFor={th.toString()}>{th}</label>
          </div>
        ))}
      </div>
    </div>
  )
}

export default CalendarFooter
