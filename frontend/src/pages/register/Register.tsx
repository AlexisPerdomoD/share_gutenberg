import { useEffect, type ReactElement } from "react"
import {Link, useFetcher} from "react-router-dom"
import LabeledInput from "../../components/LabeledInput"
import Button from "../../components/Button"





export default function Register(): ReactElement {
  const f = useFetcher()
  useEffect(()=>{
    f.data && console.log(f.data)
  },[f])

  return (
    <div>
      <div>
        {/* header */} 
        <h2>Sign up</h2>
        <p>Create an account in order to have saved any time all your books and others' collections</p>
      </div>
      <div>
        <f.Form method="POST">
          <LabeledInput
            name="name"
            title="Full Name"
            type="text"
            placeholder="Brenda Gutenberg"
            onBlur={()=>{}}
          />
          <LabeledInput
            name="email"
            title="Email"
            type="email"
            placeholder="correo@service.com"
            onBlur={()=>{}}
            required
          />
          <LabeledInput
            name="password"
            title="Password"
            type="password"
            placeholder="********"
            onBlur={()=>{}}
            required
          />
          <Button
            text="Sign up"
            disabled ={f.state !== "idle"}
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
