import { ReactElement, useContext } from "react"
import { userContext } from "../contexts/user.context"
import { ActionFunction, ActionFunctionArgs, Navigate } from "react-router-dom"
import validate, { userSchema } from "../ts_models/validations"
import { models as m } from "../../wailsjs/go/models"
import { CreateUser, GetUserByEmail } from "../../wailsjs/go/services/UMT"
// export const isLoggedLoader:LoaderFunction<LoaderFunctionArgs>
//     = async ({request, params}) =>{

// }

export function IsLogged({ children }: { children: ReactElement }) {
    const { current } = useContext(userContext)
    if (current === null) {
        return <Navigate to="/login" />
    }
    return children
}
export const signUp: ActionFunction<ActionFunctionArgs> = async ({
    request,
}) => {
    try {
        const data = await request.formData()
        const userInfo = validate<m.UserInfo>(
            userSchema,
            Object.fromEntries(data)
        )
        if (!userInfo)
            return {
                error: "CastError",
                status: 400,
                message: "field's validations went wrong",
            }
        await CreateUser(userInfo)
        const user = await GetUserByEmail(userInfo.email)
        if ("error" in user)
            return {
                error: "CastError",
                message: user.message,
                status: user.status,
            }

        return {
            name: user.name,
            username: user.id,
            role: user.role,
        }
    } catch (err) {
        if (err instanceof Error)
            return {
                error: err.name,
                message: err.message,
                status: 0,
            }
        return {
            error: "internal error server",
            status: 500,
        }
    }
}
