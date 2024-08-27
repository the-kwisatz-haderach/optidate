import { createContext } from 'preact'
import {
  PropsWithChildren,
  useCallback,
  useContext,
  useMemo,
  useState,
} from 'preact/compat'

type AppContextType = {
  state: {
    countryCode: string
    year: number
    threshold: number
  }
  updateState: (
    state: Partial<{
      countryCode: string
      year: number
      threshold: number
    }>
  ) => void
}

const AppContext = createContext<AppContextType>({
  state: {
    countryCode: '',
    year: 0,
    threshold: 3,
  },
  updateState: () => {},
})

export default function AppProvider({ children }: PropsWithChildren) {
  const [state, setState] = useState<AppContextType['state']>({
    countryCode: '',
    year: 0,
    threshold: 2,
  })

  const updateState = useCallback(
    (newState: Partial<AppContextType['state']>) => {
      setState((curr) => ({
        ...curr,
        ...newState,
      }))
    },
    []
  )

  const value = useMemo<AppContextType>(
    () => ({
      state,
      updateState,
    }),
    [state, updateState]
  )

  return <AppContext.Provider value={value}>{children}</AppContext.Provider>
}

export const useAppContext = () => useContext(AppContext)
