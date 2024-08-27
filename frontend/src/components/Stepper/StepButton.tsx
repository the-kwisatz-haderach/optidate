import { PropsWithChildren } from 'preact/compat'
import styles from './StepButton.module.css'
import cn from 'clsx'

type Props = PropsWithChildren<{
  onClick: () => void
  className?: string
  disabled?: boolean
}>

export default function StepButton({ children, className, ...props }: Props) {
  return (
    <button className={cn(styles.root, className)} {...props}>
      {children}
    </button>
  )
}
