import { TextAreaProps } from "../ts_models/ts.models";

const TextArea: React.FC<TextAreaProps> = ({ value, placeholder, onChange, ...props }) => {
    return (
        <textarea
            value={value}
            placeholder={placeholder}
            onChange={onChange}
            {...props}
        />
    )
}

export default TextArea;