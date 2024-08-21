import { defineConfig } from 'vite'
import preact from '@preact/preset-vite'

// https://vitejs.dev/config/
export default defineConfig(({ command, mode, isSsrBuild, isPreview }) => ({
  server: {
    port: 3000,
    ...(mode === 'production'
      ? {
          proxy: {
            '/': 'http://localhost:8000',
          },
        }
      : {}),
  },
  plugins: [preact()],
}))
