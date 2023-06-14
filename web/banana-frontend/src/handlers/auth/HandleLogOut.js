
async function LogOut(){
    if(localStorage.getItem("token")){
        localStorage.clear()
    }
}

export default LogOut;