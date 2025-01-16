import { StepperContentRenderer } from './Stepper/types'
import styles from './About.module.css'

export const AboutStep: StepperContentRenderer = ({ title }) => {
  return (
    <div>
      <h2>{title}</h2>
      <p className={styles.description}>
        This app will help you find bridge days for a particular country and
        year. A bridge day is a working day in between non-working days.
      </p>
    </div>
  )
}
