import {FC} from "react"
import { type ReactElement } from "react"
import {InputProps, LabelProps} from "../ts_models/ts.models"
import Input from "./Input"

const  LabeledInput:FC<InputProps & LabelProps> = ({
  name, 
  type, 
  value,
  placeholder,
  onChange,
  title,
  ...props
}): ReactElement =>{
  return <>
    <label>
      <h4>
        {title}
      </h4>
      <Input
        name={name}
        type={type}
        value={value}
        placeholder={placeholder}
        onChange={onChange}
        {...props}
      />
    </label>
    </>
}
export default LabeledInput
