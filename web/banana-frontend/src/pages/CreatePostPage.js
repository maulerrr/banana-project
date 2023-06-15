import React from 'react';
import CreatePostForm from "../components/CreatePostForm";

function CreatePostPage() {
    if (!localStorage.getItem("token"))
        window.location.href = "/login"

    return (
        <div className={"CreatePostPage"}>
            <CreatePostForm/>
        </div>
    );
}

export default CreatePostPage;