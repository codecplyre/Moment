export default function CheckCookie() {
    //Send a get request to check if the user is valid
    let valid
    fetch("/validate").then(async resp => {
        valid = resp.json().valid
    })
    return valid
}