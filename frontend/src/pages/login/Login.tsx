import {Form} from "react-router-dom"
import LabeledInput from "../../components/LabeledInput"
import Button from "../../components/Button"

export default function Login() {
    return (
        <div>
            <div>
              {/* header */}
              <h2>Sign in</h2>
              <p>some other small information</p>
            </div>
            <div>
              {/* body */}
              <Form>
                <LabeledInput
                  name="username"
                  title="Username"
                  type="text"
                  placeholder="guntenberg"
                  onBlur={()=>{}}
                  required
                />
                <LabeledInput
                  name="password"
                  title="Password"
                  placeholder="********"
                  type="password"
                  required
                />
                <Button
                  text="Log in"
                 />
              </Form>
            </div>    
        </div>
        
    )
}
