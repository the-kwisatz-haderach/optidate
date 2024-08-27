import { useEffect, useRef } from 'preact/hooks'
import { Step } from './types'
import styles from './Stepper.module.css'

interface Props extends Step {
  onTransitionIn: (state: boolean) => void
  onSetComplete: (completed?: boolean) => void
  onChangeActiveStep: () => void
  active: boolean
}

export default function StepComponent({
  onChangeActiveStep,
  Content,
  onTransitionIn,
  ...step
}: Props) {
  const ref = useRef<HTMLDivElement>(null)

  useEffect(() => {
    const intersectionObserver = new IntersectionObserver(
      ([entry]) => {
        if (entry.intersectionRatio > 0.8) {
          onChangeActiveStep()
          onTransitionIn(false)
        } else if (entry.intersectionRatio > 0) {
          onTransitionIn(true)
        }
      },
      { threshold: [0, 0.8, 1.0] }
    )

    if (ref.current) {
      intersectionObserver.observe(ref.current)
    }
    return () => {
      intersectionObserver.disconnect()
    }
  }, [])

  return (
    <section id={step.id} ref={ref} className={styles.step}>
      <Content {...step} />
    </section>
  )
}
