export default new class user {
    username;
    password;
    constructor() {

    }

    setUsername(username) {
        this.username = username
    }

    setPassword(password) {
        this.password = password
    }

    getUsername() {
       return this.username 
    }

    getPassword() {
        return this.password
    }
}
