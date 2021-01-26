import axios from "../lib/ajax"
import backend_uri from "../api/backend-api"
import c from "../lib/const"

class Course {
    constructor() {

    }

    getCourceByID(id) {
        return new Promise(function(resolve, reject) {
            axios.get(backend_uri.COURSE_ONE_BY_ID_URI.replace(":id", id)).then((res) => {
                if (res.status == 200) {
                    resolve(res.data)
                } else {
                    reject(res.data)
                }
            }).catch((res) => {
                reject(res)
            })
        })
    }

    edit(cource) {
        return new Promise(function(resolve, reject) {
            axios.put(backend_uri.COURSE_UPDATE_URI.replace(":id", cource.id), cource).then((res) => {
                if (res.status == 200) {
                    resolve(res.data)
                } else {
                    reject(res.data)
                }
            }).catch((res) => {
                reject(res)
            })
        })
    }

    delete(id) {
        return new Promise(function(resolve, reject) {
            axios.delete(backend_uri.COURSE_DELETE_URI.replace(":id", id)).then((res) => {
                if (res.status == 200) {
                    resolve(res.data)
                } else {
                    reject(res.data)
                }
            }).catch((res) => {
                reject(res)
            })
        })
    }

    store(courese, successhandle, errorhandle) {
        axios.post(backend_uri.COURSE_STORE_URI, courese).then((res) => {
            if (res.data.code == c.success && res.data.code == 200 && typeof successhandle == "function") {
                successhandle(res.data)
            } else if (typeof errorhandle == "function") {
                errorhandle(res.data)
            }
        }).catch((res) => {
            if (typeof errorhandle == "function") {
                errorhandle(res)
            }
        })
    }

    list(condition, successhandle, errorhandle) {
        axios.get(backend_uri.COURSE_LIST_URI, condition).then((res) => {
            if (res.data.code == c.success && res.status == 200 && typeof successhandle == "function") {
                successhandle(res.data)
            } else if (typeof errorhandle == "function") {
                errorhandle(res.data)
            }
        }).catch((res) => {
            if (typeof errorhandle == "function") {
                errorhandle(res)
            }
        })
    }
}

export default new Course