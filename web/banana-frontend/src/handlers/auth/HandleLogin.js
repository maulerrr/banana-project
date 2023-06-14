import axios from "axios"
import {useCookies} from "react-cookie";

async function Login(email, password){
    const options = {
        url: "http://localhost:3000/auth/login",
        headers: "",
        body: {
            email: email,
            password: password
        }
    }

    const response = (await axios.post(
        options.url,
        options.body,
        options.headers
    )).data

    localStorage.setItem("token", response.token)

    const user = {
        id : response?.user_id,
        username : response?.username,
        email : response?.email,
    }

    localStorage.setItem("user", JSON.stringify(user))

    return response
}

export default Login;