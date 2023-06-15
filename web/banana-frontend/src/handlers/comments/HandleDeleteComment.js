import axios from "axios";

async function DeleteComment(comment_id){
    const token = "Bearer " + localStorage.getItem("token")

    if (!token) return

    const user = JSON.parse(localStorage.getItem("user"))

    if (!user) return

    const options = {
        url: "http://localhost:3000/comment/" + comment_id,
        config: {
            headers: {
                "Content-Type": "application/json",
                "Authorization": token
            }
        }
    }

    return (await axios.delete(
        options.url,
        options.config,
    )).data;
}

export default DeleteComment;