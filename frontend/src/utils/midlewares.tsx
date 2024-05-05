import {ReactElement, useContext} from "react";
//import {LoaderFunction, LoaderFunctionArgs, redirectDocument} from "react-router-dom";
import {userContext} from "../contexts/user.context";
import {Navigate} from "react-router-dom";
// export const isLoggedLoader:LoaderFunction<LoaderFunctionArgs> 
//     = async ({request, params}) =>{
   
// }


export function IsLogged({children}:{children:ReactElement})  {
    const {current} = useContext(userContext)
    if (current === null){
      return <Navigate to="/login" />
    }
    return children
}
