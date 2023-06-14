import React from 'react';
import {useEffect, useState} from "react";
import HandleGetComments from "../handlers/comments/HandleGetComments";
import {useParams} from "react-router-dom";
import PostWithComments from "../components/PostWithComments";
import HandleGetPostByID from "../handlers/posts/HandleGetPostByID";

function PostPage(props) {
    const [post, setPost] = useState(null);
    const [comments, setComments] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(false);

    const {post_id} = useParams()

    if (!localStorage.getItem("token"))
        window.location.href = "/login"

    useEffect(() => {
        setError(false);

        HandleGetPostByID(post_id)
            .then((response) => {
                return response.data
            })
            .then((post) => {
                setPost(post)
                setLoading(false)
            })
            .catch(() => setError(true))
            .finally(() => setLoading(false))

        HandleGetComments(post_id)
            .then((response) => {
                return response.data.comments
            })
            .then((comments) => {
                setComments(comments)
                setLoading(false)
            })
            .catch(() => {
                setError(true);
            })
            .finally(() => {
                setLoading(false);
            });
    }, []);

    if (loading) {
        return <p className="Indicator">Loading...</p>;
    }

    if (error) {
        return <p className="Indicator">Fetching Error..</p>;
    }

    //test
    function handleCommentDelete(deletedCommentId){
        setComments((prevComments) => prevComments.filter((comment) => comment.comment_id !== deletedCommentId));
    }

    return (
        <div className="PostPage">
            <PostWithComments post={post} comments={comments} onCommentDelete={handleCommentDelete}/>
        </div>
    );
}

export default PostPage;