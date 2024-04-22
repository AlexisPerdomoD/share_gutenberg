import { FC } from "react"
import { models as m} from "../../wailsjs/go/models"
import Button from "./Button"
import { redirect } from "react-router-dom"


const BookCard:FC<{ book: m.Book }> = ({ book }) => {
  return (
    <div key={book.id}>
        <div>
            <picture>
                <img src={book.formats["image/jpeg"] || ""} alt={`${book.title} image`} />
            </picture>
            <h3>{book.title} {book.id}</h3>
        </div>
        <a href={`/${book.id}`}>aver</a>
        <Button
        value={`${book.id}`}
        onClick={()=> {redirect(`/${book.id}`)}}
        text={`Book Profile id: ${book.id}`}
        />
    </div>
  )
}

export default BookCard