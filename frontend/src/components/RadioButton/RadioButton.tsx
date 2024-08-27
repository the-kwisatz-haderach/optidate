import { HTMLAttributes, useId } from 'preact/compat'
import styles from './RadioButton.module.css'

type Props<T> = {
  value: T
  onChange: (value: T) => void
  selected: boolean
  label?: string
}

export default function RadioButton<
  T extends HTMLAttributes<HTMLInputElement>['value']
>({ value, onChange, selected, label = '' }: Props<T>) {
  const id = useId()
  return (
    <div>
      <input
        id={id}
        name={label}
        className={styles.input}
        type="radio"
        value={value}
        checked={selected}
        onChange={() => onChange(value)}
      />
      <label htmlFor={id} className={styles.label}>
        {label || value}
      </label>
    </div>
  )
}
