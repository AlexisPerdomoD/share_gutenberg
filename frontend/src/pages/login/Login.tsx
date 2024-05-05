import {Form} from "react-router-dom"
import LabeledInput from "../../components/LabeledInput"
import Button from "../../components/Button"
import {Link} from "react-router-dom"

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
              <Link to={"/register"}>
                <h4>
                  don't have an account? Sign Up
                </h4>
              </Link>
            </div>    
        </div>
        
    )
}
