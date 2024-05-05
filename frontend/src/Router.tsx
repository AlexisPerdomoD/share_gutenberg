import { createBrowserRouter } from "react-router-dom";
import Layout from "./pages/Layout";
import Home from "./pages/home/Home";
import BookProfile,{loader as bookProfileLoader} from "./pages/book-profile/BookProfile";
import Login from "./pages/login/Login";
import {IsLogged} from "./utils/midlewares";
import Register from "./pages/register/Register";

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
                path:"/:id",
                element:<BookProfile />,
                loader:bookProfileLoader
                //view and crud operations with book s to collections
            },
            {
                path:"/collections",
                //children maybe for collections id and crud operations with it 
            } 
        ]
    },
    {
      path:"/profile",
      element:<IsLogged children={<p>profile</p>} />
    },
    {
        path:"/login",
      element: <Login/>
    },
    {
        path:"/register",
        element:<Register />
    },
    {
        path:"/about"
    },
    {
        path:"/error"
    }
])

export default router
