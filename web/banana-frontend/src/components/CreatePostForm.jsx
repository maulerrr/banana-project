import React, {useState} from 'react';
import classes from "./classes/components.module.css";
import handleCreatePost from "../handlers/posts/HandleCreatePost";
import handleGetPosts from "../handlers/posts/HandleGetPosts";
import {Alert, AlertTitle} from "@mui/material";

function CreatePostForm(props) {
    const [header, setHeader] = useState("")
    const [body, setBody] = useState("")
    const [error, setError] = useState(false)
    const [alert, setAlert] = useState(<></>)

    function closeAlert(e){
        e.preventDefault()
        setAlert(<></>)
    }

    function HandleCreatePost(e){
        e.preventDefault()

        let validHeader = header.replace(/\s/g, '')
        let validBody = body.replace(/\s/g, '')

        if (validHeader === "" || validBody === "") {
            setError(true)
            setAlert(
                <Alert severity={"error"}
                       variant={"filled"}
                       onClose={(e)=>{closeAlert(e)}}
                >
                    <AlertTitle>Error occurred</AlertTitle>
                    You cannot create empty post!
                </Alert>
            )
            setTimeout(()=>{
                setAlert(<></>)
            }, 3000)
        } else {
            handleCreatePost(header, body)
                .then(() => {
                    handleGetPosts()
                        .then(response => response.data)
                        .then(() => {
                            console.log("Updated posts!")
                            window.location.href = "/"
                        })
                        .catch(() => {
                            setError(true)
                        })
                })
            setHeader("")
            setBody("")
        }
    }

    return (
        <div className={classes.FormWrapper}>
            <form onSubmit={HandleCreatePost}>
                <label htmlFor="header">Header: </label>
                <input type="text" name="header"
                       onChange={(e)=>{setHeader(e.target.value)}}/>

                <label htmlFor="body">Body: </label>
                <textarea type="text" name="body"
                       onChange={(e)=>{setBody(e.target.value)}}/>

                <input type="submit" value="Create Post"/>
            </form>

            <div className={classes.Alerts}>
                {alert}
            </div>
        </div>
    );
}

export default CreatePostForm;