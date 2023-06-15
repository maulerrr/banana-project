import axios from "axios";

async function GetLikesCount(post_id){
    const token = "Bearer " + localStorage.getItem("token");

    if (!token) return

    const options = {
        url: "https://localhost:3000/api/v1/post/like/count/" + post_id,
        config: {
            headers: {
                "Content-Type": "application/json",
                "Authorization": token
            }
        }
    }

    return (await axios.get(options.url, options.config)).data
}

export default GetLikesCount;