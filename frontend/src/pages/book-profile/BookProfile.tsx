
import {LoaderFunction, LoaderFunctionArgs, useLoaderData } from 'react-router-dom';
import { models as m } from '../../../wailsjs/go/models';
import { GetBook } from '../../../wailsjs/go/main/App';
import validate, { bookSchema } from '../../ts_models/validations';
import { ReactNode } from 'react';


 export const loader:LoaderFunction<LoaderFunctionArgs> = async({params})=>{
     const response = await GetBook(params["id"] || "2")
     if ("status" in response && response.status === 500 )throw new Error(response.message)
     return response
 }


const BookProfile = () => {
  const book = validate<m.Book>(bookSchema, useLoaderData())

  return (
    <div>
      {book && <p>aqui falta renderizar el book pero hasta ahora aqui el id:{book.id} mas nombre: {book.title} y bueno de una a ver los formatos</p>}
      {book && <ul>{renderFormats(book)}</ul>}
      {!book && <p>holaaaaaa</p>}
    </div>
  )
}

export default BookProfile

//provitional

function renderFormats(book:m.Book):ReactNode{
  const collected = Object.entries(book.formats)
  return collected.map(e =>(
    <li>
      <h3>
        {e[0]}
      </h3>
      <p>{e[1]}</p>
    </li>
  ))
}
/*
application/epub+zip

https://www.gutenberg.org/ebooks/84.epub3.images

application/octet-stream

https://www.gutenberg.org/cache/epub/84/pg84-h.zip

application/rdf+xml

https://www.gutenberg.org/ebooks/84.rdf

application/x-mobipocket-ebook

https://www.gutenberg.org/ebooks/84.kf8.images

image/jpeg

https://www.gutenberg.org/cache/epub/84/pg84.cover.medium.jpg

text/html

https://www.gutenberg.org/ebooks/84.html.images

text/plain; charset=us-ascii

https://www.gutenberg.org/ebooks/84.txt.utf-8 */