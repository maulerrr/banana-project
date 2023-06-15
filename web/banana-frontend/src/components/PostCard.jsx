import React, {useEffect, useState} from 'react';
import classes from "./classes/components.module.css"
import Like from "../handlers/likes/HandleLike";
import DeletePost from "../handlers/posts/HandleDelete";
import GetPosts from "../handlers/posts/HandleGetPosts";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faTrash, faHeart} from "@fortawesome/free-solid-svg-icons"
import GetLike from "../handlers/likes/HandleGetLike";
import GetLikesCount from "../handlers/likes/HandleCountLikes";
import {Alert, AlertTitle} from "@mui/material";

function PostCard(props) {
    const [color, setColor] = useState("black")
    const [alert, setAlert] = useState(<></>)
    const [likesCount, setLikesCount] = useState(0)
    const [likes, setLikes] = useState(<></>)

    const user_id = JSON.parse(localStorage.getItem("user"))?.id

    function closeAlert(e){
        e.preventDefault()
        setAlert(<></>)
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
        }).catch((error)=> {
            setAlert(
                <Alert severity={"error"}
                       variant={"filled"}
                       onClose={(e)=>{closeAlert(e)}}
                >
                    <AlertTitle>Error occurred</AlertTitle>
                    {error.response.data.message}
                </Alert>
            )
            setTimeout(()=>{
                setAlert(<></>)
            }, 3000)
        })
    }

    useEffect(()=>{
        GetLike(props.post.post_id)
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
        <div className={classes.PostContainer}>
            <div onClick={()=>{
                window.location.href=`/post/${props.post.post_id}`
            }
            }>
                <h1 className={classes.PostHeader}>
                    {props.post.header}
                </h1>
                <p className={classes.PostBody}>
                    {props.post.body}
                </p>
            </div>

            <div className={classes.ButtonGroup}>
                <button
                    className={classes.LikeButton}
                    onClick={HandleLike}>
                    <img className={classes.Banana} src={`/img/${color}.png`} alt={"like"}/>
                    {/*<FontAwesomeIcon icon="fa-solid fa-banana" color={color}/>*/}
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


            <div className={classes.Alerts}>
                {alert}
            </div>
        </div>
    );
}

export default PostCard;