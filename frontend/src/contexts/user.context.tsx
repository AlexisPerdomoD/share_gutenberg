import {ReactElement, ReactNode, createContext, useState} from "react";
import {Current, CurrentContext} from "../ts_models/ts.models";

export const userContext = createContext<CurrentContext>({
    current:null,
    setCurrent:()=>{}
})

export function CurrentContextProvider(
    {children}:{children:ReactNode}
):ReactElement{
    const [current, setCurrent] = useState<Current | null>(null)

    const handleCurrent = (current:Current | null) => setCurrent(current)
    
    return (
        <userContext.Provider 
          value={{current, setCurrent:handleCurrent}}
        >
            {children}
        </userContext.Provider>
    )

}
