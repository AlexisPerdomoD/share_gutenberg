import { FC } from 'react';
import { models as m } from '../../wailsjs/go/models';
import BookCard from './BookCard';
const Books:FC<{books:m.Book[]}> = ({books}) => {
  return (
    <div>{books.map(book => <BookCard book={book}/>)}</div>
  )
}

export default Books