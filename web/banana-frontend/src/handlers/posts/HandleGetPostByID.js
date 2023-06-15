import axios from "axios"

async function HandleGetPostByID(post_id){
    const token = "Bearer " + localStorage.getItem("token")

    if (!token) return

    const options = {
        url: "http://localhost:3000/post/" + post_id,
        config: {
            headers: {
                'Content-Type': "application/json",
                'Authorization': token
            }
        }
    }

    return axios.get(options.url, options.config)
}

export default HandleGetPostByID;