import { type ReactElement } from "react"
import {Link, useFetcher} from "react-router-dom"
import LabeledInput from "../../components/LabeledInput"
import Button from "../../components/Button"

export default function Register(): ReactElement {
  const f = useFetcher()

  return (
    <div>
      <div>
        {/* header */} 
        <h2>Sign up</h2>
        <p>Create an account in order to have saved any time all your books and others' collections</p>
      </div>
      <div>
        <f.Form>
          <LabeledInput
            name="first_name"
            title="First Name"
            type="text"
            placeholder="Brenda"
            onBlur={()=>{}}
          />
          <LabeledInput
            name="last_name"
            title="Last Name"
            type="text"
            placeholder="Gutenberg"
            onBlur={()=>{}}
          />
          <LabeledInput
            name="email"
            title="Email"
            type="email"
            placeholder="correo@service.com"
            onBlur={()=>{}}
          />
          <LabeledInput
            name="username"
            title="Username"
            type="text"
            placeholder="gutenberg"
            onBlur={()=>{}}
          />
          <LabeledInput
            name="password"
            title="Password"
            type="password"
            placeholder="********"
            onBlur={()=>{}}
          />
          <Button
            text="Sign up"
          />
        </f.Form>
        <Link to={"/login"}>
          <h4>
            Already have an account? Sign in
          </h4>
        </Link>
      </div>
    </div>
  )
}
