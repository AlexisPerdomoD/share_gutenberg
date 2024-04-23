import { FC } from "react"
import { models as m} from "../../wailsjs/go/models"
import Button from "./Button"
import { NavLink } from "react-router-dom"

const BookCard:FC<{ book: m.Book }> = ({ book }) => {
  return (
    <div key={book.id}>
        <div>
            <picture>
                <img src={book.formats["image/jpeg"] || ""} alt={`${book.title} image`} />
            </picture>
            <h3>{book.title} {book.id}</h3>
        </div>
        <NavLink 
          to={`/${book.id}`} children={
          <Button
            value={`${book.id}`}
            text={`Book Profile id: ${book.id}`}
          />}
        />
    </div>
  )
}

export default BookCard