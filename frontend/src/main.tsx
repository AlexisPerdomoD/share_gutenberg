import * as React from 'react'
import {createRoot} from 'react-dom/client'
import { RouterProvider } from 'react-router-dom'
import router from './Router'
import "./styles/styles.css"
import {CurrentContextProvider} from "./contexts/user.context"


const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <React.StrictMode>
      <CurrentContextProvider>
        <RouterProvider router={router}/>
      </CurrentContextProvider>
    </React.StrictMode>
)
