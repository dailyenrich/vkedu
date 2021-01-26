import app from "../config/app"
import token from "../lib/token"
import axios from "../lib/ajax"
import backend_uri from "../api/backend-api"
import c from "../lib/const"

export default new class auth {
    token;
    local_storage_key="auth_token";

    constructor() {
        if (app.auth_token_key) {
            this.local_storage_key = app.auth_token_key
        }
    }

    login(user, callback) {
        if (typeof callback != "function") {
            throw "参数二必须为函数"
        }
        
        axios.post(backend_uri.ADMIN_LOGIN_URI, user).then((response) => {
            if (response.data.code == c.success) {
                this.token = response.data.data.token
                token.set(this.local_storage_key, this.token)
                token.set(app.user_info_key, JSON.stringify(response.data.data))
            }
            callback(this.check())
        })
    }

    logout() {
        this.token = ""
        token.set(this.local_storage_key, this.token)
        token.set(app.user_info_key, "")
        return !this.check()
    }

    check() {
        this.token = token.get(this.local_storage_key)
        if(!this.token) {
            return false
        }

        return true;
    }

    user() {
        let user = token.get(app.user_info_key)
        if (!user) {
            this.logout()
        }

        return JSON.parse(user)
    }
}