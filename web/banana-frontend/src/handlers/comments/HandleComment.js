import axios from "axios";

async function CreateComment(post_id, text) {
    const token = "Bearer " + localStorage.getItem("token")

    if (!token) return

    const user = JSON.parse(localStorage.getItem("user"))

    if (!user) return

    const options = {
        url: "http://localhost:3000/comment/",
        config: {
            headers: {
                "Content-Type": "application/json",
                "Authorization": token
            }
        },
        body: {
            user_id: user.id,
            post_id: post_id,
            text: text,
        }
    }

    return (await axios.post(
        options.url,
        options.body,
        options.config
    )).data
}

export default CreateComment;