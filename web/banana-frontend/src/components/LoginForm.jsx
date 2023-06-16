import React from 'react';
import classes from "./classes/components.module.css"
import handleLogin from "../handlers/auth/HandleLogin";
import {useState} from "react";
import {Alert, AlertTitle, Button, Fade} from "@mui/material";


function LoginForm(props) {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [error, setError] = useState(false)
    const [alert, setAlert] = useState(<></>)

    function closeAlert(e){
        e.preventDefault()
        setAlert(<></>)
    }

    function HandleLogin(e){
        e.preventDefault()
        handleLogin(email, password)
            .then((response)=>{
                console.log("Trying to log in..")
                if (response) window.location.href = "/"
            })
            .catch((err)=>{
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

    return (
        <div className={classes.FormWrapper}>
            <form onSubmit={HandleLogin}>
                <label htmlFor="email">Email: </label>
                <input type="text" name="email"
                       onChange={(e)=>{setEmail(e.target.value)}}/>

                <label htmlFor="password">Password: </label>
                <input type="password" name="password"
                       onChange={(e)=>{setPassword(e.target.value)}}/>

                <label> Don't have an account? <a href={"/signup"}>Sign Up</a>
                </label>

                <input type="submit" value="Login"/>
            </form>

            <div className={classes.Alerts}>
                {alert}
            </div>

        </div>
    );
}

export default LoginForm;