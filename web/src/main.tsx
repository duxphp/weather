import { createRoot } from 'react-dom/client'
import '@unocss/reset/tailwind.css'
import 'virtual:uno.css'
import App from './App.tsx'


createRoot(document.getElementById('root')!).render(<App />)
