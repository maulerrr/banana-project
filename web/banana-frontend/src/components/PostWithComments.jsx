import React, {useState, useEffect} from 'react';
import classes from "./classes/components.module.css";
import Comment from "./Comment";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faHeart, faTrash} from "@fortawesome/free-solid-svg-icons";
import Like from "../handlers/likes/HandleLike";
import DeletePost from "../handlers/posts/HandleDelete";
import GetPosts from "../handlers/posts/HandleGetPosts";
import DeleteComment from "../handlers/comments/HandleDeleteComment";
import CreateComment from "../handlers/comments/HandleComment";
import GetLike from "../handlers/likes/HandleGetLike";
import {Alert, AlertTitle} from "@mui/material";

function PostWithComments(props) {
    const [text, setText] = useState("")
    const [color, setColor] = useState("black")
    const [alert, setAlert] = useState(<></>)

    const user_id = JSON.parse(localStorage.getItem("user"))?.id

    function closeAlert(e){
        e.preventDefault()
        setAlert(<></>)
    }

    function HandleCommentCreation(e){
        e.preventDefault()

        let validText = text.replace(/\s/g, '')

        if (validText === ""){
            setAlert(
                <Alert severity={"error"}
                       variant={"filled"}
                       onClose={(e)=>{closeAlert(e)}}
                >
                    <AlertTitle>Error occurred</AlertTitle>
                    You cannot create empty comment!
                </Alert>
            )
            setTimeout(()=>{
                setAlert(<></>)
            }, 3000)
            return
        }

        CreateComment(props.post.post_id, text)
            .then((response) => {
                console.log(response.data)
            })
            .catch((error) => {
                console.log(error)
            })
            .finally(()=> {
                window.location.reload()
            })
    }

    function HandleLike(e){
        e.preventDefault()

        Like(props.post.post_id).then((response)=> {
            console.log("Liked post with ID="+props.post.post_id)
            if (response.liked === true){
                setColor("red")
            } else setColor("black")
        })

    }

    function HandleDelete(e) {
        e.preventDefault()
        DeletePost(props.post.post_id).then(()=>{
            props.onPostDelete(props.post.post_id);
            GetPosts()
                .then((response)=>response.data)
                .then(()=> {
                    console.log("Updated posts after delete..")
                })
            console.log("Deleted post with id=" + props.post.post_id)
        })
    }

    useEffect(()=>{
        GetLike(props?.post.post_id)
            .then((response)=> {
                if (response.liked === true) {
                    setColor("red")
                }
            })
            .catch(()=>{
                console.log("error on getting like data")
                setColor("black")
            })
    }, [])

    return (
        <div className={classes.PostPageMain}>
            <section className={classes.PostDetails}>
                <div>
                    <h1 className={classes.PostHeader}>
                        {props.post?.header}
                    </h1>
                    <p className={classes.PostBodyFull}>
                        {props.post?.body}
                    </p>
                    <div className={classes.ButtonGroup}>
                        <button
                            className={classes.LikeFullPost}
                            onClick={HandleLike}>
                            <img className={classes.Banana} src={`/img/${color}.png`} alt={"like"}/>
                            {/*<FontAwesomeIcon icon={faHeart} size={"2x"} color={color}/>*/}
                        </button>
                        {
                            props.post.user_id === user_id ?
                                <button
                                    className={classes.DeleteButton}
                                    onClick={HandleDelete}>
                                    <FontAwesomeIcon icon={faTrash}/>
                                </button> : <></>
                        }
                    </div>

                </div>
            </section>

            <h1 className={classes.CommentSectionHeader}>Comments:</h1>
            <section className={classes.Comments}>
                <form onSubmit={HandleCommentCreation}>
                    <textarea placeholder="Enter comment"
                              onChange={(e)=>{setText(e.target.value)}}/>
                    <input type={"submit"} value={"Comment"}/>
                </form>
                <div>
                    {props.comments && props.comments.map((comment)=>
                        <Comment comment={comment} key={comment.comment_id} onCommentDelete={props.onCommentDelete}/>
                    )}
                </div>
            </section>

            <div className={classes.Alerts}>
                {alert}
            </div>
        </div>
    );
}

export default PostWithComments;