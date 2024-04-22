import { FC } from "react"
import { models as m} from "../../wailsjs/go/models"

const BookCard:FC<{ book: m.Book }> = ({ book }) => {
  return (
    <div>
        <div>
            <picture>
                <img src={book.formats["image/jpeg"] || ""} alt={`${book.title} image`} />
            </picture>
            <h3>{book.title}</h3>
        </div>
    </div>
  )
}

export default BookCard