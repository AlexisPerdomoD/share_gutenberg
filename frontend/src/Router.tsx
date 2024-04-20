import { createBrowserRouter } from "react-router-dom";
import Layout from "./pages/Layout";


const router = createBrowserRouter([
    {
        path:"/",
        element:<Layout/>
    },
    {
        path:"/login"
    },
    {
        path:"/register"
    },
    {
        path:"/about"
    },
    {
        path:"/error"
    }
])

export default router