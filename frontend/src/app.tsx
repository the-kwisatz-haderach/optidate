import { useCallback, useState } from 'preact/hooks'
import Stepper, { Props as StepperProps } from './components/Stepper/Stepper'
import { SelectCountryStep } from './components/SelectCountry'
import { SelectYearStep } from './components/SelectYear'
import { AboutStep } from './components/About'
import { Calendar } from './components/Calendar'
import CalendarFooter from './components/CalendarFooter'

export function App() {
  const [activeStep, setActiveStep] = useState(0)
  const [steps, setSteps] = useState<StepperProps['steps']>([
    {
      id: 'about',
      completed: true,
      title: 'Welcome',
      Content: AboutStep,
    },
    {
      id: 'year',
      completed: false,
      headerOverride: 'Start',
      title: 'Select year',
      Content: SelectYearStep,
    },
    {
      id: 'country',
      completed: false,
      title: 'Select country',
      Content: SelectCountryStep,
    },
    {
      id: 'calendar',
      completed: false,
      headerOverride: 'Calendar',
      title: 'Show calendar',
      disableScroll: true,
      Footer: CalendarFooter,
      Content: Calendar,
    },
  ])

  const onSetComplete = useCallback(
    (index: number, completed = true) =>
      setSteps((currSteps) =>
        currSteps.map((step, i) =>
          i === index
            ? {
                ...step,
                completed,
              }
            : step
        )
      ),
    []
  )

  return (
    <Stepper
      steps={steps}
      activeStepIndex={activeStep}
      onChangeActiveStep={setActiveStep}
      onSetComplete={onSetComplete}
    />
  )
}
