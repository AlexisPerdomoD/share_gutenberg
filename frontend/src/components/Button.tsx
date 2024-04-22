import { FC } from 'react'
import { ButtonProps } from '../ts_models/ts.models'

const Button:FC<ButtonProps> = ({ text, className, onClick, ...props }) => {
    return (
        <button className={className} onClick={onClick} {...props}>
            {text}
        </button>
    )
}

export default Button;
