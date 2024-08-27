import { render } from 'preact'
import { App } from './app.tsx'
import './index.css'
import AppProvider from './context/AppProvider.tsx'

async function enableMocking() {
  if (process.env.NODE_ENV !== 'development') {
    return
  }

  const { worker } = await import('./mocks/browser')

  // `worker.start()` returns a Promise that resolves
  // once the Service Worker is up and ready to intercept requests.
  return worker.start()
}

enableMocking().then(() => {
  render(
    <AppProvider>
      <App />
    </AppProvider>,
    document.getElementById('app')!
  )
})
