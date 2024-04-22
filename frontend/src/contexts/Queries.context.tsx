import React, {ReactNode, createContext, useState } from "react";
import {Queries} from '../ts_models/ts.models';

export const QueryContext = createContext<{queries:Queries,handleQueries?:Function}>({
    queries:{
        search:"",
        category:"",
        page:"",
        topic:"",
        copyright:"",
        ids:"",
        languages:""
    }
})
export function QueryContextProvider({children}:{children:ReactNode}){
    const [queries, setQueries] = useState<Queries>({
        search:"",
        category:"",
        page:"",
        topic:"",
        copyright:"",
        ids:"",
        languages:""
    })
    const handleQueries = (params:Queries) => setQueries(params)
     return <>
    <QueryContext.Provider value={{queries, handleQueries}}>
        {children}
    </QueryContext.Provider>
    </>
}
