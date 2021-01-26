class Token {
    constructor() {

    }

    set(key, value) {
        localStorage.setItem(key, value)
    }

    get(key) {
       return localStorage.getItem(key)
    }
}

export default new Token