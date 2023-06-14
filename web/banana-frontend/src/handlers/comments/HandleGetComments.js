import axios from "axios";

async function GetComments(post_id) {
    const token = "Bearer " + localStorage.getItem("token");

    const options = {
        url: "http://localhost:3000/comment/" + post_id,
        config: {
            headers: {
                'Content-Type': "application/json",
                'Authorization': token
            }
        }
    }

    return axios.get(options.url, options.config);
}

export default GetComments;