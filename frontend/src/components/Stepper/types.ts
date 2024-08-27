import { ComponentType } from 'preact'
import { ReactNode } from 'preact/compat'

export type StepperContentRenderer = ComponentType<
  {
    onSetComplete: (completed?: boolean) => void
    completed: boolean
    active: boolean
  } & Omit<Step, 'Content'>
>

export type StepperControlRenderer = ComponentType<{
  className: string
  steps: Step[]
  activeStepIndex: number
  onStep: () => void
  ActionRenderer?: ReactNode
  disabled?: boolean
  title: string
  hide?: boolean
}>

export type Step = {
  id: string
  title: string
  completed: boolean
  headerOverride?: string
  Footer?: StepperControlRenderer
  Header?: StepperControlRenderer
  headerIconOverride?: string
  disableScroll?: boolean
  Content: StepperContentRenderer
}
