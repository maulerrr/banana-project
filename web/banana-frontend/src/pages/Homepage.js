import React, {useEffect, useState} from 'react';
import PostCard from "../components/PostCard";
import handleGetPosts from "../handlers/posts/HandleGetPosts";
import DeletePost from "../handlers/posts/HandleDelete";

function Homepage() {
    const [posts, setPosts] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(false);

    if (!localStorage.getItem("token"))
        window.location.href = "/login"

    const username = JSON.parse(localStorage.getItem("user")).username

    useEffect(() => {
        setError(false);

        handleGetPosts()
            .then((response) => {
                return response.data.posts
            })
            .then((posts) => {
                console.log(posts)
                setPosts(posts)
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
        return <p className="Indicator">Error fetching posts</p>;
    }

    function handlePostDelete(deletedPostId) {
        setPosts((prevPosts) => prevPosts.filter((post) => post.post_id !== deletedPostId));
    }

    return (
        <div className="Home">
            {posts && posts.map((post) => <PostCard
                post={post}
                key={post.post_id}
                onPostDelete={handlePostDelete}
            />)}
        </div>
    );
}

export default Homepage;