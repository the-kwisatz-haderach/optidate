import { Step } from './types'
import styles from './Stepper.module.css'
import { useCallback, useEffect, useMemo, useRef, useState } from 'preact/hooks'
import StepComponent from './Step'
import cn from 'clsx'
import { DefaultStepControl } from './DefaultControls'

export type Props = {
  steps: Step[]
  activeStepIndex: number
  onSetComplete: (completedStepIndex: number, completed?: boolean) => void
  onChangeActiveStep: (newActiveIndex: number) => void
}

export default function Stepper({
  steps,
  activeStepIndex,
  onChangeActiveStep,
  onSetComplete,
}: Props) {
  // const [deferredActiveStepIndex, setDeferredActiveStepIndex] =
  //   useState(activeStepIndex)
  const [transitionStart, setTransitionStart] = useState(false)
  const containerRef = useRef<HTMLElement>(null)

  // useEffect(() => {
  //   const timeout = setTimeout(() => {
  //     setDeferredActiveStepIndex(activeStepIndex)
  //   }, 2000)
  //   return () => {
  //     clearTimeout(timeout)
  //   }
  // }, [activeStepIndex])

  useEffect(() => {
    if (!containerRef.current) {
      return
    }
    if (transitionStart) {
      containerRef?.current?.classList.remove(styles.in)
      containerRef.current.classList.add(styles.out)
    } else {
      containerRef?.current?.classList.add(styles.in)
      containerRef.current.classList.remove(styles.out)
    }
  }, [transitionStart])

  const scrollStepIntoView = useCallback((id: string) => {
    const el = document.getElementById(id)
    el?.scrollIntoView({ behavior: 'smooth' })
  }, [])

  const Header = useMemo(
    () => steps[activeStepIndex].Header ?? DefaultStepControl,
    [steps, activeStepIndex]
  )

  const Footer = useMemo(
    () => steps[activeStepIndex].Footer ?? DefaultStepControl,
    [steps, activeStepIndex]
  )

  return (
    <main ref={containerRef} className={cn(styles.wrapper)}>
      <Header
        className={cn(styles.stepperControl, styles.top)}
        steps={steps}
        activeStepIndex={activeStepIndex}
        onStep={() => scrollStepIntoView(steps[activeStepIndex - 1].id)}
        ActionRenderer={<span style={{ position: 'relative', top: 3 }}>⌃</span>}
        title={
          steps[activeStepIndex].headerOverride ??
          steps[activeStepIndex - 1]?.title
        }
        hide={!steps[activeStepIndex - 1]}
        disabled={!steps[activeStepIndex - 1]}
      />
      <div>
        {steps
          .filter((_, index) => index === 0 || steps[index - 1]?.completed)
          .map((step, index) => (
            <StepComponent
              key={index}
              {...step}
              active={activeStepIndex === index}
              onTransitionIn={setTransitionStart}
              onChangeActiveStep={() => onChangeActiveStep(index)}
              onSetComplete={(completed = true) =>
                onSetComplete(index, completed)
              }
            />
          ))}
      </div>
      <Footer
        className={cn(styles.stepperControl, styles.bottom)}
        steps={steps}
        activeStepIndex={activeStepIndex}
        onStep={() => scrollStepIntoView(steps[activeStepIndex + 1].id)}
        ActionRenderer={
          <span
            style={{
              position: 'relative',
              bottom: 2,
              display: 'inline-block',
              transform: 'rotate(180deg)',
            }}
          >
            ⌃
          </span>
        }
        disabled={!steps[activeStepIndex].completed}
        title={steps[activeStepIndex + 1]?.title}
        hide={!steps[activeStepIndex + 1]}
      />
    </main>
  )
}
