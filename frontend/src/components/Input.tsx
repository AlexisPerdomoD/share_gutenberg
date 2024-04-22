import { InputProps } from "../ts_models/ts.models";

const Input: React.FC<InputProps> = ({ type, value, placeholder, onChange, ...props }) => {
    return (
        <input
            type={type}
            value={value}
            placeholder={placeholder}
            onChange={onChange}
            {...props}
        />
    )
}

export default Input