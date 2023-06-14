import axios from "axios";

async function Like(post_id){
    const token = "Bearer " + localStorage.getItem("token")
    const userID = JSON.parse(localStorage.getItem("user")).id

    if (!token) return

    const options = {
        url: "http://localhost:3000/like/",
        config: {
            headers: {
                "Content-Type": "application/json",
                "Authorization": token
            }
        },
        body: {
            user_id: userID,
            post_id: post_id,
        }
    }

    const response = await axios.post(
        options.url,
        options.body,
        options.config,
        )

    return response.data
}

export default Like;