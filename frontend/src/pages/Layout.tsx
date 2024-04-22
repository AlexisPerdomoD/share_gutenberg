import { FC, Suspense} from "react"
import {  Outlet } from "react-router-dom"
import NavBar from "../components/NavBar"
import { QueryContextProvider } from "../contexts/Queries.context"

const Layout:FC = () => {
return (
<QueryContextProvider>
    <div>
        <h1>La pagine</h1>
            <Outlet/>
        <NavBar/>
    </div>
</QueryContextProvider>
)
}

export default Layout