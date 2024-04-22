import { createBrowserRouter } from "react-router-dom";
import Layout from "./pages/Layout";
import Home from "./pages/home/Home";


const router = createBrowserRouter([
    {
        path:"/",
        element:<Layout/>,
        //search bar 
        //advance search modal 
        //nav bar to navigate betwend 
        children:[
            {
                index:true,
                element: <Home />
                //small message to complLa pagineete the layout
                //book container

            },
            {
                path:"/:id"
                //view and crud operations with books to collections
            },
            {
                path:"/collections",
                //children maybe for collections id and crud operations with it 
            }
        ]
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