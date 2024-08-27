import { StepperControlRenderer } from './types'
import styles from './Stepper.module.css'
import cn from 'clsx'

export const DefaultStepControl: StepperControlRenderer = ({
  onStep,
  ActionRenderer = 'Next',
  disabled,
  title,
  hide = false,
  className,
}) => {
  return (
    <button className={className} onClick={onStep} disabled={disabled}>
      <span className={styles.controlTitle}>{title}</span>
      <span className={cn(hide && styles.hide)}>{ActionRenderer}</span>
    </button>
  )
}
