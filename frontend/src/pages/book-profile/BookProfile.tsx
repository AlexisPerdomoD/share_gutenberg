
import {LoaderFunction, LoaderFunctionArgs, useLoaderData, useParams } from 'react-router-dom';
import { models as m } from '../../../wailsjs/go/models';
import { GetBook } from '../../../wailsjs/go/main/App';
import { useEffect, useState } from 'react';
import validate, { bookSchema } from '../../ts_models/validations';


 export const loader:LoaderFunction<LoaderFunctionArgs> = async({params})=>{
     const response = await GetBook(params["id"] || "2")
     if (response.Error && response.status === 500)throw new Error("se rompio")
     return response
 }


const BookProfile = () => {
  const book = validate<m.Book>(bookSchema, useLoaderData())
  // const [book, setBook] = useState<m.Book>()
  // const params = useParams()
  // useEffect(()=>{
  //   params["id"] && GetBook(params["id"]).then(book=> setBook(book))
  // },[])
  return (
    <div>
      {book && <p>aqui falta renderizar el book pero hasta ahora aqui el id:{book.id} mas nombre: {book.title} y bueno de una a ver los formatos</p>}
      {book && renderFormats(book)}
      
    </div>
  )
}

export default BookProfile

//provitional

function renderFormats(book:m.Book) {
  let response = "<div>";
  for (let key in book.formats) {
    const value = book.formats[key];
    response += `<p>Key: ${key}, Value: ${value}</p>`;
  }
  response += "</div>";
  return response
}