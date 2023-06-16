import React, {useState} from 'react';
import classes from "./classes/components.module.css"
import handleRegistration from "../handlers/auth/HandleRegistration";
import {Alert, AlertTitle} from "@mui/material";

function SignupForm(props) {
    const [username, setUsername] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [error, setError] = useState(false)
    const [alert, setAlert] = useState(<></>)

    function HandleRegistration(e){
        e.preventDefault()

        let validUsername = username.replace(/\s/g, '')

        if (validUsername === ""){
            setAlert(
                <Alert severity={"error"}
                       variant={"filled"}
                       onClose={(e)=>{closeAlert(e)}}
                >
                    <AlertTitle>Error occurred</AlertTitle>
                    Specify Username!
                </Alert>
            )
            setTimeout(()=>{
                setAlert(<></>)
            }, 3000)
            return
        }
        handleRegistration(username, email, password)
            .then((response)=>{
                console.log("Trying to u sign up..")
                if (response) window.location.href = "/"
            })
            .catch((err)=>{
                console.log(err)
                setError(true)
                setAlert(
                    <Alert severity={"error"}
                           variant={"filled"}
                           onClose={(e)=>{closeAlert(e)}}
                    >
                        <AlertTitle>Error occurred</AlertTitle>
                        {err.response.data.error}
                    </Alert>
                )
                setTimeout(()=>{
                    setAlert(<></>)
                }, 3000)
            })
    }

    function closeAlert(e){
        e.preventDefault()
        setAlert(<></>)
    }

    return (
        <div className={classes.FormWrapper}>
            <form onSubmit={HandleRegistration}>
                <label htmlFor="username">Username:</label>
                <input type="text" name="username"
                       onChange={(e)=>{setUsername(e.target.value)}}/>

                <label htmlFor="email">Email: </label>
                <input type="text" name="email"
                       onChange={(e)=>{setEmail(e.target.value)}}/>

                <label htmlFor="password">Password: </label>
                <input type="password" name="password"
                       onChange={(e)=>{setPassword(e.target.value)}}/>

                <label> Already have an account? <a href={"/login"}>Sign in</a>
                </label>

                <input type="submit" value="Sign Up"/>
            </form>

            <div className={classes.Alerts}>
                {alert}
            </div>
        </div>
    );
}

export default SignupForm;