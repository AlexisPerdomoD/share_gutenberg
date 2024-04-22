
import { GetBooks } from '../../../wailsjs/go/main/App';
import { FC, useContext, useEffect, useState } from "react";
import Books from "../../components/Books";
import { models as m } from "../../../wailsjs/go/models";
import { redirect } from 'react-router-dom';
import { QueryContext } from '../../contexts/Queries.context';


// export const loader:LoaderFunction<{request:Request}> = async({request}) =>{
//        try {
//         const url = new URL(request.url)
//         console.log(url)
//         const data = await GetBooks({})
//         return data
//        } catch (error) {
//         console.log(error)
//        }
// }


const Home:FC = () => {
    const [books, setBooks] = useState<m.Book[]>()
    const [loading, setLoading] = useState<boolean>(false)
    const {queries} = useContext(QueryContext)
     useEffect(()=>{
        setLoading(true)
         GetBooks(queries)
            .then(gutenbex => setBooks(gutenbex.results))
            .catch(e => redirect("/error"))
            .finally(() => setLoading(false))
    }, [])
  return <>
    {loading && <p>loading</p>}
    {books && <Books books={books}/>}
  </>
}

export default Home